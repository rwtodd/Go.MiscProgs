// a program to apply IPS patches
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

const (
	EOF_BYTES = ('E' << 16) + ('O' << 8) + 'F'
)

// There are 3 kinds of patch, which follow the following
// simple interface:
type Patcher interface {
	fmt.Stringer
	patch(fl io.WriterAt) error
}

// The EOFMarker should be the last patch the user sees:
type EOFMarker struct{}

func (p EOFMarker) patch(fl io.WriterAt) error { return nil }
func (p EOFMarker) String() string             { return "END-OF-PATCHES MARKER" }

// A BytePatch has []byte and an offset to write to the destination.
type BytePatch struct {
	values   []byte
	location int64
}

func (p *BytePatch) patch(fl io.WriterAt) error {
	_, err := fl.WriteAt(p.values, p.location)
	return err
}
func (p *BytePatch) String() string {
	return fmt.Sprintf("Patch (length %d) at location: %X", len(p.values), p.location)
}

// An RLEPatch fills a section of bytes with a single value.
type RLEPatch struct {
	location int64
	length   int
	value    byte
}

func (p *RLEPatch) patch(fl io.WriterAt) error {
	buf := bytes.Repeat([]byte{p.value}, p.length)
	_, err := fl.WriteAt(buf, p.location)
	return err
}
func (p *RLEPatch) String() string {
	return fmt.Sprintf("RLEPatch (length %d, value %02X) at location: %X", p.length, p.value, p.location)
}

// readIPS reads an entire patch file, pushing Patchers into
// a channel.
func readIPS(ips *bufio.Reader, out chan Patcher, errs chan error) {
	var err error
	read1 := func() int {
		var v byte
		if err == nil {
			v, err = ips.ReadByte()
		}
		return int(v)
	}
	read2 := func() int { return (read1() << 8) + read1() }
	read3 := func() int { return (read2() << 8) + read1() }

	send := func(p Patcher) {
		if err == nil {
			out <- p
		}
	}

	defer close(out)
	defer close(errs)

	header := make([]byte, 5)
	_, err = io.ReadFull(ips, header)
	if string(header) != "PATCH" {
		errs <- fmt.Errorf("Not a valid IPS file!")
		return
	}

	var buf []byte
	for {
		offs := read3()
		plen := read2()
		switch plen {
		case 0:
			send(&RLEPatch{int64(offs), read2(), byte(read1())})
		default:
			buf = make([]byte, plen)
			_, err = io.ReadFull(ips, buf)
			send(&BytePatch{buf, int64(offs)})
		}
		if err != nil {
			if (offs == EOF_BYTES) && (err == io.EOF) {
				out <- EOFMarker{}
			} else {
				errs <- fmt.Errorf("Error reading ips file: %v\n", err)
			}
			return
		}
	}
}

func process(ipsf, srcf, tgtf string) error {
	// open the IPS file and start the reader
	infile, err := os.Open(ipsf)
	if err != nil {
		return fmt.Errorf("Opening IPS file: %v\n", err)
	}
	defer infile.Close()
	br := bufio.NewReader(infile)
	pchan, echan := make(chan Patcher, 100), make(chan error, 1)
	go readIPS(br, pchan, echan)

	// copy the source to the new name
	if err = copyFileContents(srcf, tgtf); err != nil {
		return fmt.Errorf("File copy: %v\n", err)
	}

	// open the target for patching
	outfile, err := os.OpenFile(tgtf, os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("Opening output file: %v\n", err)
	}
	defer outfile.Close()

	// drain the channel, applying patches
	idx := 0
	for p := range pchan {
		idx++
		fmt.Printf("%d: %v\n", idx, p)
		if err := p.patch(outfile); err != nil {
			return fmt.Errorf("Applying patches: %v\n", err)
		}
	}

	return <-echan
}

func main() {
	// check the arguments
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr,
			"Usage: %s patchfile orig newfile\n",
			os.Args[0])
		os.Exit(1)
	}

	if err := process(os.Args[1], os.Args[2], os.Args[3]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
