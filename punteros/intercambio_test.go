package main

import (
	"testing"
)

func TestSwap(t *testing.T) {
	casos := []struct {
		x         int
		y         int
		esperadoX int
		esperadoY int
	}{
		{1, 2, 2, 1},
		{3, 4, 4, 3},
		{5, 6, 6, 5},
		{-7, 8, 8, -7},
		{9, -10, -10, 9},
	}

	for _, caso := range casos {
		swap(&caso.x, &caso.y)
		if caso.x != caso.esperadoX || caso.y != caso.esperadoY {
			t.Errorf("swap(%d, %d) = (%d, %d), se esperaba (%d, %d)", caso.esperadoX, caso.esperadoY, caso.x, caso.y, caso.esperadoX, caso.esperadoY)
		}
	}
}
