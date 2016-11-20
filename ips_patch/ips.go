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

// There are 2 kinds of patch, which follow the following
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
	return fmt.Sprintf("Patch (length %d) at location: %08X",
		len(p.values),
		p.location)
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

// readIPS reads an entire patch file, pushing Patchers into
// a channel.
func readIPS(ips *bufio.Reader, out chan Patcher) {
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

	defer close(out)

	header := make([]byte, 5)
	_, err = io.ReadFull(ips, header)
	if string(header) != "PATCH" {
		fmt.Fprintln(os.Stderr, "Not a valid IPS file!")
		return
	}

	var buf []byte
	for {
		offs := read3()
		if offs == EOF_BYTES {
			out <- EOFMarker{}
			return
		}

		plen := read2()
		switch plen {
		case 0:
			plen = read2()
			val := byte(read1())
			buf = bytes.Repeat([]byte{val}, plen)
		default:
			buf = make([]byte, plen)
			_, err = io.ReadFull(ips, buf)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"Error reading ips file: %v\n",
				err)
			return
		}
		out <- &BytePatch{buf, int64(offs)}
	}
}

func main() {
	// check the arguments
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr,
			"Usage: %s patchfile orig newfile\n",
			os.Args[0])
		os.Exit(1)
	}

	// copy the source to the new name
	if err := copyFileContents(os.Args[2], os.Args[3]); err != nil {
		fmt.Fprintf(os.Stderr, "File copy: %v\n", err)
		os.Exit(1)
	}

	// open the IPS file and start the reader
	pchan := make(chan Patcher, 100)
	infile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Opening IPS file: %v\n", err)
		os.Exit(1)
	}
	br := bufio.NewReader(infile)
	go readIPS(br, pchan)

	// open the target for patching
	outfile, err := os.OpenFile(os.Args[3], os.O_WRONLY, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Opening output file: %v\n", err)
		os.Exit(1)
	}

	// drain the channel, applying patches
	idx := 0
	for p := range pchan {
		idx++
		fmt.Printf("%d: %v\n", idx, p)
		if err := p.patch(outfile); err != nil {
			fmt.Fprintf(os.Stderr, "Applying patches: %v\n", err)
			os.Exit(1)
		}
	}
}
