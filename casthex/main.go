// A program to cast an I Ching hexagram.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func doCasting(input string) {
	// STEP ONE: determine the command type
	var lines string
	switch input {
	case "coins":
		lines = casting(coinsMtd)
	case "stalks":
		lines = casting(stalksMtd)
	case "static":
		lines = casting(randomMtd)
	default:
		lines = input
	}

	if len(lines) != 6 {
		fmt.Fprintf(os.Stderr, "Bad input <%v>\n", lines)
		os.Exit(1)
	}

	// STEP TWO: parse the input, displaying the hexagram
	fmt.Printf("Casting for <%s>:\n\n", lines)
	var h1, h2 int
	var output = make([]string, 0, 6)

	for idx := 5; idx >= 0; idx-- {
		h1, h2 = (h1 << 1), (h2 << 1)
		switch lines[idx] {
		case '6':
			h2 |= 1
			output = append(output, "  ---   ---   =>   ---------")
		case '7':
			h1 |= 1
			h2 |= 1
			output = append(output, "  ---------        ---------")
		case '8':
			output = append(output, "  ---   ---        ---   ---")
		case '9':
			h1 |= 1
			output = append(output, "  ---------   =>   ---   ---")
		default:
			fmt.Fprintf(os.Stderr, "Bad input <%v>\n", lines[idx])
			os.Exit(1)
		}
	}

	var changed = h1 != h2
	for _, l := range output {
		if !changed {
			l = l[:11] // get the first 11 bytes
		}
		fmt.Println(l)
	}
	fmt.Println()

	// STEP THREE:  display the hexagram names
	fmt.Println(hexname[h1])
	if changed {
		fmt.Printf(" == Changing To =>\n%s\n", hexname[h2])
	}
	fmt.Println("\n")
}

func main() {
	if len(os.Args) > 1 {
		for _, casting := range os.Args[1:] {
			doCasting(casting)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			doCasting(strings.TrimSpace(scanner.Text()))
		}
	}

}
