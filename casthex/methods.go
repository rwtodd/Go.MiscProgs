package main

import (
	"math/rand"
	"time"
)

// a file of iching divination methods

// casting calls the given method (e.g., coinsMtd) six times
// and builds a string of the results.
func casting(proc func() byte) string {
	var results [6]byte
	for idx := range results {
		results[idx] = proc()
	}
	return string(results[:])
}

// coinsMtd calculates an i-ching reading via the 3-coins
// method.  The result is a string of characters from the set
// { 6,7,8,9 }.
func coinsMtd() byte {
	return '6' + byte(rand.Int31n(2)+rand.Int31n(2)+rand.Int31n(2))
}

// stalksMtd calculates an i-ching reading via the yarrow stalks method.
// The result is a string of characters from the set { 6,7,8,9 }.
func stalksMtd() byte {
	var result byte
	rnd := rand.Int31n(16)
	switch rnd {
	case 0:
		result = '6'
	case 1, 2, 3, 4, 5:
		result = '7'
	case 6, 7, 8, 9, 10, 11, 12:
		result = '8'
	case 13, 14, 15:
		result = '9'
	}
	return result
}

// randomMtd calculates an i-ching reading that is a single
// hexagram with no moving lines. The result is a string of characters
// from the set { 6,7,8,9 }.
func randomMtd() byte {
	return '7' + byte(rand.Int31n(2))
}

// we need to set up a random number generator...
func init() {
	rand.Seed(time.Now().UnixNano())
}
