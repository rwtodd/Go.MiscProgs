package iching

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestCycle(t *testing.T) {
	for n := 0 ; n < 20 ; n++ {
		wen := rand.Intn(64)+1
		hx := FromKingWen(wen)
		for o := 0 ; o < 64 ; o++ {
			hx = Next(hx)
		}
		if hx.KingWen() != wen {
			t.Fatalf("After cycle of 64, %v != %v", wen, hx.KingWen())
		}
	}
}
