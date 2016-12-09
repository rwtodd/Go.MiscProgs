// A program to cast an I Ching hexagram.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	yang = "\u2584\u2584\u2584\u2584\u2584\u2584\u2584\u2584"
	yin  = "\u2584\u2584\u2584  \u2584\u2584\u2584"
)

func main() {
	// STEP ONE: get the input
	var lines string
	if len(os.Args) > 1 {
		lines = os.Args[1]
	} else {
		inp := bufio.NewReader(os.Stdin)
		lines, _ = inp.ReadString('\n')
	}
	lines = strings.TrimSpace(lines)

	// STEP TWO: validate the input, give usage on bad input
	if len(lines) != 6 {
		fmt.Fprintf(os.Stderr, `Usage: %s [casting]
  ... where casting is 6 digits (6,7,8,9) for I Ching lines
  If casting is not given, it is read from stdin. 
`, os.Args[0])
		os.Exit(1)
	}

	// STEP THREE: parse the input
	var h1, h2 int
	var h1lines, h2lines [6]string

	for idx := 5; idx >= 0; idx-- {
		h1 = h1 << 1
		h2 = h2 << 1

		switch lines[idx] {
		case '6':
			h1lines[idx] = yin
			h2lines[idx] = yang
			h2 |= 1
		case '7':
			h1lines[idx] = yang
			h2lines[idx] = yang
			h1 |= 1
			h2 |= 1
		case '8':
			h1lines[idx] = yin
			h2lines[idx] = yin
		case '9':
			h1lines[idx] = yang
			h2lines[idx] = yin
			h1 |= 1
		default:
			fmt.Fprintf(os.Stderr, "Bad input <%c>\n", lines[idx])
			h1lines[idx] = yin
			h2lines[idx] = yin
		}
	}

	// STEP FOUR:  display the output
	fmt.Printf("Casting for <%s>:\n\n", lines)
	fmt.Printf("%d %v\n", hex2wen[h1], hexname[h1])
	if h1 != h2 {
		fmt.Printf(" --Changing To-->\n%d %v\n", hex2wen[h2], hexname[h2])
	}
	fmt.Println()

	for idx := 5; idx >= 0; idx-- {
		l1, middle, l2 := h1lines[idx], "   ", ""
		if h1 != h2 {
			l2 = h2lines[idx]
			if l1 != l2 {
				middle = "-->"
			}
		}
		fmt.Printf("  %s %s %s\n", l1, middle, l2)
	}
	fmt.Println()
}