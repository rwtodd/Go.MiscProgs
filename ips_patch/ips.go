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

// A patch will be represented by the following struct:
type patch struct {
	values   []byte
	location int64
}

// readIPS reads an entire patch file, pushing Patchers into
// a channel.
func readIPS(ips *bufio.Reader, out chan *patch, errs chan error) {
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
	send := func(p *patch) {
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
			plen = read2()
			buf = bytes.Repeat([]byte{byte(read1())}, plen)
		default:
			buf = make([]byte, plen)
			_, err = io.ReadFull(ips, buf)
		}
		send(&patch{buf, int64(offs)})
		if err != nil {
			if (offs == EOF_BYTES) && (err == io.EOF) {
				out <- nil
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
	pchan, echan := make(chan *patch, 100), make(chan error, 1)
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
		switch p {
		case nil:
			fmt.Printf("%d: END-OF-PATCHES MARKER\n", idx)
		default:
			fmt.Printf("%d: Patch of length %d at 0x%X\n", idx, len(p.values), p.location)
			if _, err := outfile.WriteAt(p.values, p.location); err != nil {
				return fmt.Errorf("Applying patches: %v\n", err)
			}
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
