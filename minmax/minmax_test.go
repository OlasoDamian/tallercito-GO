package minmax

import (
	"testing"
)

func TestEncontrarMinimoMaximo(t *testing.T) {
	lista := []int{4, 7, 2, 9, 1}

	minimo, maximo := EncontrarMinimoMaximo(lista)
	if minimo != 1 || maximo != 9 {
		t.Errorf("Error en EncontrarMinimoMaximo. Se esperaba (min: 1, max: 9) pero se obtuvo (min: %d, max: %d)", minimo, maximo)
	}
}
