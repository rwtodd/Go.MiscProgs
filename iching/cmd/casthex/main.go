// A program to cast an I Ching hexagram.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/rwtodd/misc-go/iching"
)

func line2str(l bool) string {
	if l {
		return "\u2584\u2584\u2584\u2584\u2584\u2584\u2584\u2584"
	} else {
		return "\u2584\u2584\u2584  \u2584\u2584\u2584"
	}

}

func printOneHex(h iching.Hexagram) {
	fmt.Printf("%d %v\n\n", h.KingWen(), h)

	lines := h.Lines()
	for idx := 5; idx >= 0; idx-- {
		fmt.Print("  ")
		fmt.Println(line2str(lines[idx]))
	}
}

func printHex(h1, h2 iching.Hexagram) {
	if h1 == h2 {
		printOneHex(h1)
		return
	}

	fmt.Printf("%d %v\n --Changing To-->\n%d %v\n\n", h1.KingWen(), h1, h2.KingWen(), h2)

	lines1, lines2 := h1.Lines(), h2.Lines()

	// print the hexagrams side-by-side...
	for idx := 5; idx >= 0; idx-- {
		l1, l2 := lines1[idx], lines2[idx]
		middle := "   "
		if l1 != l2 {
			middle = "-->"
		}
		fmt.Printf("  %s %s %s\n", line2str(l1), middle, line2str(l2))
	}
}

//  three coins...                combinations   range [0,256)
//  6 changing yin    all tails   1 way          [0,32)
//  7 yang            1 head      3 ways         [32,128)
//  8 yin             2 heads     3 ways         [128,224)
//  9 changing yang   all heads   1 way          [224,256)
func castLine() rune {
	coins := rand.Int31n(256)
	switch {
	case coins < 32:
		return '6'
	case coins < 128:
		return '7'
	case coins < 224:
		return '8'
	default:
		return '9'
	}
}

// returns a string suitable as input to this program
// only generated at random...
func castHex() string {
	rand.Seed(time.Now().UTC().UnixNano())
	lines := make([]rune, 6)
	for idx := range lines {
		lines[idx] = castLine()
	}

	return string(lines)
}

func parse(lines string) (iching.Hexagram, iching.Hexagram) {
	var h1, h2 [6]bool

	if len(lines) != 6 {
		fmt.Fprintf(os.Stderr, "Bad input <%s>!\n", lines)
		return iching.FromKingWen(1), iching.FromKingWen(2)
	}

	for idx, val := range lines {
		switch val {
		case '6':
			h1[idx], h2[idx] = false, true
		case '7':
			h1[idx], h2[idx] = true, true
		case '8':
			h1[idx], h2[idx] = false, false
		case '9':
			h1[idx], h2[idx] = true, false
		default:
			fmt.Fprintf(os.Stderr, "Bad input <%c>\n", val)
		}
	}

	return iching.FromLines(h1), iching.FromLines(h2)
}

func main() {
	var input string
	if len(os.Args) < 2 {
		input = castHex()
		fmt.Printf("Random I-Ching Cast: %s\n", input)
	} else {
		input = os.Args[1]
	}

	fmt.Println()
	oldH, newH := parse(input)
	printHex(oldH, newH)
}
