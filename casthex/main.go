// A program to cast an I Ching hexagram.
package main

import (
	"fmt"
	"os"
)

var YANGYIN = [...]string{
	"\u2584\u2584\u2584  \u2584\u2584\u2584",
	"\u2584\u2584\u2584\u2584\u2584\u2584\u2584\u2584",
}

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

func validate(input string) (bool, error) {
	changed := false
	if len(input) != 6 {
		return false, fmt.Errorf("Wrong-sized input <%s>!", input)
	}

	for _, ch := range input {
		switch ch {
		case '6', '9':
			changed = true
		case '7', '8': // do nothing
		default:
			return false, fmt.Errorf("Bad character <%c> in input!", ch)
		}
	}
	return changed, nil
}

func main() {
	// STEP ONE: get the input
	var lines string

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-coins":
			lines = casting(coinsMtd)
		case "-stalks":
			lines = casting(stalksMtd)
		case "-random":
			lines = casting(randomMtd)
		default:
			lines = os.Args[1]
		}
	} else {
		lines = casting(coinsMtd) // default to coins
	}

	// STEP TWO: validate the input, give usage on bad input
	changed, err := validate(lines)
	if (len(os.Args) > 2) || (err != nil) {
		usage()
	}

	// STEP THREE: parse the input, displaying the hexagram
	fmt.Printf("Casting for <%s>:\n\n", lines)
	var h1, h2 int
	var h1line, h2line int

	for idx := 5; idx >= 0; idx-- {
		h1line = int(lines[idx]) & 1
		h2line = h1line
		if lines[idx] == '6' || lines[idx] == '9' {
			h2line = 1 - h2line
		}
		h1 = (h1 << 1) | h1line
		h2 = (h2 << 1) | h2line

		if changed {
			middle := "   "
			if h1line != h2line {
				middle = "-->"
			}
			fmt.Printf("  %s %s %s\n", YANGYIN[h1line], middle, YANGYIN[h2line])
		} else {
			fmt.Printf("  %s\n", YANGYIN[h1line])
		}
	}
	fmt.Println()

	// STEP FOUR:  display the hexagram names 
	fmt.Println(hexname[h1])
	if changed {
		fmt.Printf(" --Changing To-->\n%s\n", hexname[h2])
	}
	fmt.Println()

}
