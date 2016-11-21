package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/rwtodd/apputil/cmdline"
)

// counts all the bytes in the input buffer which aren't ASCII spaces or tabs.
func count(buf []byte) int64 {
	var c int64
	for _, v := range buf {
		if v != ' ' && v != '\t' {
			c++
		}
	}
	return c
}

// summer sums all the integers from the 'in' channel, and
// then leaves the result int he 'out' channel.
func summer(in chan int64, out chan int64) {
	var sum int64
	for v := range in {
		sum += v
	}
	out <- sum
}

func main() {
	cmdline.GlobArgs()

	// initialization phase... set up the needed channels and WaitGroup,
	// and start the summer goroutine
	var wg sync.WaitGroup
	var err error

	limiter := make(chan struct{}, 8)
	counts, sum := make(chan int64, 8), make(chan int64)
	go summer(counts, sum)

	// calculation phase... read all input files, counting their characters in parallel
	var fl io.ReadCloser
	for _, fname := range os.Args[1:] {
		fl, err = os.Open(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", fname, err)
		}

		// keep up to 8 buffers in memory (via 'limiter'), reading the whole file:
		for err == nil {
			var n int
			limiter <- struct{}{}
			buf := make([]byte, 4096)
			n, err = fl.Read(buf)
			if n > 0 {
				wg.Add(1)
				go func(b []byte) {
					counts <- count(b)
					wg.Done()
					<-limiter // read the next batch
				}(buf[:n])
			}
		}
		if err != io.EOF {
			fmt.Fprintf(os.Stderr, "%s: %v\n", fname, err)
		}
		fl.Close()
	}

	// cleanup phase... wait for all work to be done, then close the
	// 'counts' channel so the summer will finish
	wg.Wait()
	close(counts)
	fmt.Printf("Total: %d\n", <-sum)
}
