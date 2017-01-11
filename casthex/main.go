// A program to cast an I Ching hexagram.
package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: %s [opts|casting]
  ... where casting is 6 digits from the set: {6,7,8,9} for I Ching lines
  ... and options are below:
       -coins  use the 3-Coins method
       -stalks use the yarrow stalks method
       -static generate a random hexagram with no moving lines
`, os.Args[0])
	os.Exit(1)
}

func validate(input string) error {
	if len(input) != 6 {
		return fmt.Errorf("Wrong-sized input <%s>!", input)
	}

	for _, ch := range input {
		if ch < '6' || ch > '9' {
			return fmt.Errorf("Bad character <%c> in input!", ch)
		}
	}
	return nil
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
		case "-static":
			lines = casting(randomMtd)
		default:
			lines = os.Args[1]
		}
	} else {
		lines = casting(coinsMtd) // default to coins
	}

	// STEP TWO: validate the input, give usage on bad input
	err := validate(lines)
	if (len(os.Args) > 2) || (err != nil) {
		usage()
	}

	// STEP THREE: parse the input, displaying the hexagram
	fmt.Printf("Casting for <%s>:\n\n", lines)
	var h1, h2 int
	var output = make([]string, 0, 6)

	for idx := 5; idx >= 0; idx-- {
		h1, h2 = (h1 << 1), (h2 << 1)
		switch lines[idx] {
		case '6':
			h2 |= 1
			output = append(output, "  --   --    -->    -------")
		case '7':
			h1 |= 1
			h2 |= 1
			output = append(output, "  -------           -------")
		case '8':
			output = append(output, "  --   --           --   --")
		case '9':
			h1 |= 1
			output = append(output, "  -------    -->    --   --")
		}
	}
	var changed = h1 != h2
	for _, l := range output {
		if !changed {
			l = l[:9]
		}
		fmt.Println(l)
	}
	fmt.Println()

	// STEP FOUR:  display the hexagram names
	fmt.Println(hexname[h1])
	if changed {
		fmt.Printf(" --Changing To-->\n%s\n", hexname[h2])
	}
	fmt.Println()
}
