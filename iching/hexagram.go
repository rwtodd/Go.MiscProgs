// A package for dealing with I Ching
// hexagrams and trigrams.  It has basic
// facilities for using the King Wen sequence,
// as well as decomposing and recomposing hexagrams.
//
// In this package, lines are represented as bools,
// where false = yin and true = yang.
package iching

// Represents an I Ching hexagram.
type Hexagram uint8

// Deconstructs the hexagram into its lower and upper
// trigrams.
func (h Hexagram) Trigrams() (Trigram, Trigram) {
	return Trigram(h & 0x07), Trigram(h >> 3)
}

// Deconstructs a hexagram into a slice of
// lines (true = yang, false = yin).
func (h Hexagram) Lines() [6]bool {
	var ans [6]bool
	ans[0] = (h & 1) > 0
	ans[1] = (h & 2) > 0
	ans[2] = (h & 4) > 0
	ans[3] = (h & 8) > 0
	ans[4] = (h & 16) > 0
	ans[5] = (h & 32) > 0
	return ans
}

// translates King Wen hexagram numbers (1 to 64)
// to the hexagram.
func FromKingWen(idx int) Hexagram {
	return num2hex[idx-1]
}

// translates the given slice of lines into a hexagram.
func FromLines(l [6]bool) Hexagram {
	bint := func(b bool) int {
		if b {
			return 1
		} else {
			return 0
		}
	}
	return Hexagram(
		bint(l[0]) + (bint(l[1]) << 1) + (bint(l[2]) << 2) +
			(bint(l[3]) << 3) + (bint(l[4]) << 4) + (bint(l[5]) << 5))
}

// translates from Trigrams to a hexagram.
func FromTrigrams(lower, upper Trigram) Hexagram {
	return Hexagram(lower + (upper << 3))
}

// Gives the name of a hexagram.
func (h Hexagram) String() string {
	return hexnames[h]
}

// Gives the number of the hexagram in the King Wen sequence.
func (h Hexagram) KingWen() int {
	return hexnums[h]
}

// Returns the hexagram resulting from changing the
// given list of lines (1-based indices, 1-6).
func ChangeLines(h Hexagram, which ...int) Hexagram {
	lines := h.Lines()
	for _, idx := range which {
		lines[idx-1] = !lines[idx-1]
	}
	return FromLines(lines)
}

// Gives the inverse of the provided Hexagram.
func Invert(h Hexagram) Hexagram {
	return ChangeLines(h, 1, 2, 3, 4, 5, 6)
}

// Gives the next hexagram in the King Wen sequence.
func Next(h Hexagram) Hexagram {
	n := h.KingWen() + 1
	if n == 65 {
		n = 1
	}
	return FromKingWen(n)
}

// Gives the previous hexagram in the King Wen sequence.
func Prev(h Hexagram) Hexagram {
	n := h.KingWen() - 1
	if n == 0 {
		n = 64
	}
	return FromKingWen(n)
}
