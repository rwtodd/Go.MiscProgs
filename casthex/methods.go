package main

import (
	"math/rand"
	"time"
)

// a file of iching divination methods

// coinsMtd calculates an i-ching reading via the 3-coins
// method.  The result is a string of characters from the set
// { 6,7,8,9 }.
func coinsMtd() string {
	var casting [6]byte
	for idx := range casting {
		casting[idx] = '6' + byte(rand.Int31n(2)+rand.Int31n(2)+rand.Int31n(2))
	}
	return string(casting[:])
}

// stalksMtd calculates an i-ching reading via the yarrow stalks method.
// The result is a string of characters from the set { 6,7,8,9 }.
func stalksMtd() string {
	var casting [6]byte
	for idx := range casting {
		rnd := rand.Int31n(16)
		switch rnd {
		case 0:
			casting[idx] = '6'
		case 1, 2, 3, 4, 5:
			casting[idx] = '7'
		case 6, 7, 8, 9, 10, 11, 12:
			casting[idx] = '8'
		case 13, 14, 15:
			casting[idx] = '9'
		}
	}
	return string(casting[:])
}

// randomMtd calculates an i-ching reading that is a single
// hexagram with no moving lines. The result is a string of characters
// from the set { 6,7,8,9 }.
func randomMtd() string {
	var casting [6]byte
	for idx := range casting {
		casting[idx] = '7' + byte(rand.Int31n(2))
	}
	return string(casting[:])
}

// we need to set up a random number generator...
func init() {
	rand.Seed(time.Now().UnixNano())
}
