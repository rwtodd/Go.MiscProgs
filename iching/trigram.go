package iching

// A type to represent an I Ching trigram, the
// three-line building block of a hexagram.
type Trigram uint8

const (
	KUN Trigram = iota
	CHEN
	KAN
	TUI
	KEN
	LI
	SUN
	CHIEN
)

var trinames = [...]string{
	"K'UN / RECEPTIVE / EARTH",
	"CHEN / AROUSING / THUNDER",
	"K'AN / ABYSMAL / WATER",
	"TUI / JOYOUS / LAKE",
	"KEN / KEEPING STILL / MOUNTAIN",
	"LI / CLINGING / FIRE",
	"SUN / GENTLE / WIND",
	"CH'IEN / CREATIVE / HEAVEN",
}

// Gives the name of a trigram.
func (t Trigram) String() string {
	return trinames[t]
}
