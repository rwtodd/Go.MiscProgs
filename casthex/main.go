// A program to cast an I Ching hexagram.
package main

import (
	"fmt"
	"os"
)

const (
	yang = "\u2584\u2584\u2584\u2584\u2584\u2584\u2584\u2584"
	yin  = "\u2584\u2584\u2584  \u2584\u2584\u2584"
)

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: %s [opts|casting]
  ... where casting is 6 digits from the set: {6,7,8,9} for I Ching lines
  ... and options are below:
       -coins  use the 3-Coins method
       -stalks use the yarrow stalks method
       -random generate a random hexagram with no moving lines
`, os.Args[0])
	os.Exit(1)
}

func main() {
	// STEP ONE: get the input
	var lines string

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-coins":
			lines = coinsMtd()
		case "-stalks":
			lines = stalksMtd()
		case "-random":
			lines = randomMtd()
		default:
			lines = os.Args[1]
		}
	} else {
		lines = coinsMtd() // default to coins
	}

	// STEP TWO: validate the input, give usage on bad input
	if (len(os.Args) > 2) || (len(lines) != 6) {
		usage()
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
	fmt.Println(hexname[h1])
	if h1 != h2 {
		fmt.Printf(" --Changing To-->\n%s\n", hexname[h2])
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
