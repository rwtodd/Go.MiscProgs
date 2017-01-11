package main

import (
	"fmt"
	"math/rand"
	"time"
)

var LINES = [2]string{"*   *", "  *  "}

// randFig generates a random geomantic figure.
func randFig() []string {
	var fig = make([]string, 4)
	for idx := range fig {
		fig[idx] = LINES[rand.Int31n(2)]
	}
	return fig
}

// transposeFigs takes a slice of figures and returns
// their transpose.
func transposeFigs(in [][]string) [][]string {
	var ans = make([][]string, len(in[0]))
	for idx := range ans {
		ans[idx] = make([]string, len(in))
		for idy := range ans[idx] {
			ans[idx][idy] = in[idy][idx]
		}
	}
	return ans
}

// combineFigs takes adjaces figures in a slice, and combines
// them.  The resulting slice is half the length of the input.
func combineFigs(in [][]string) [][]string {
	var ans = make([][]string, len(in)/2)
	for idx := range ans {
		a, b := in[idx*2], in[idx*2+1]
		ans[idx] = make([]string, 4)
		for i, aline := range a {
			if b[i][0] == aline[0] {
				ans[idx][i] = LINES[0]
			} else {
				ans[idx][i] = LINES[1]
			}
		}
	}
	return ans
}

// spaces is a helper function to generate a given amount of spaces
func spaces(amt int) string {
	ans := make([]byte, amt)
	for idx := range ans {
		ans[idx] = ' '
	}
	return string(ans)
}

// display outputs a row of figures, with 'ispace' initial space,
// and 'mspace' spaces between each figure.
func display(ispace int, mspace int, figs [][]string) {
	isp, msp := spaces(ispace), spaces(mspace)
	strs := transposeFigs(figs)
	for _, l := range strs {
		fmt.Print(isp)
		for _, s := range l {
			fmt.Print(s, msp)
		}
		fmt.Println()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// generate the top line (moms and daughters)
	line1 := make([][]string, 8)
	for idx := 0; idx < 4; idx++ {
		line1[idx] = randFig()
	}
	copy(line1[4:], transposeFigs(line1[:4]))

	// reverse the figures
	for idx := 3; idx >= 0; idx-- {
		opp := 7 - idx
		line1[idx], line1[opp] = line1[opp], line1[idx]
	}

	// generate the remaining lines via combination
	nieces := combineFigs(line1)
	witnesses := combineFigs(nieces)
	judge := combineFigs(witnesses)

	// display the shield
	display(2, 5, line1)
	fmt.Println()
	display(7, 15, nieces)
	fmt.Println()
	display(17, 35, witnesses)
	fmt.Println()
	display(37, 0, judge)
}
