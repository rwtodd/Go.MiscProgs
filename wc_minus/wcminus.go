package main

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/rwtodd/apputil/cmdline"
)

// counts all the bytes in the input buffer which aren't ASCII spaces or tabs.
func counter(in chan []byte, out chan int64) {
	var c int64

	for buf := range in {
		for _, v := range buf {
			if v != ' ' && v != '\t' {
				c++
			}
		}
	}

	out <- c
}

func main() {
	cmdline.GlobArgs()

	var err error
	ncpu := runtime.NumCPU()

	bufs, sum := make(chan []byte, ncpu), make(chan int64, ncpu)
	for idx := 0; idx < ncpu; idx++ {
		go counter(bufs, sum)
	}

	// calculation phase... read all input files, counting their characters in parallel
	var fl io.ReadCloser
	for _, fname := range os.Args[1:] {
		fl, err = os.Open(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", fname, err)
		}

		for err == nil {
			var n int
			buf := make([]byte, 4096)
			n, err = fl.Read(buf)
			if n > 0 {
				bufs <- buf[:n]
			}
		}
		if err != io.EOF {
			fmt.Fprintf(os.Stderr, "%s: %v\n", fname, err)
		}
		fl.Close()
	}

	// cleanup phase... close the bufs channel and add all the results
	close(bufs)
	var answer int64
	for idx := 0; idx < ncpu; idx++ {
		answer += <-sum
	}

	fmt.Printf("Total: %d\n", answer)
}
