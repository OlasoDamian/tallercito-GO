package vectores

import "testing"

func TestSumaYProductoEscalar(t *testing.T) {
	casos := []struct {
		v1, v2          Vector
		suma            Vector
		productoEscalar int
	}{
		{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{5, 7, 9}, 32},
		{Vector{-1, -2, -3}, Vector{-4, -5, -6}, Vector{-5, -7, -9}, 32},
		{Vector{0, 0, 0}, Vector{0, 0, 0}, Vector{0, 0, 0}, 0},
	}

	for _, caso := range casos {
		suma, productoEscalar, err := SumaYProductoEscalar(caso.v1, caso.v2)
		if err != nil {
			t.Errorf("SumaYProductoEscalar(%v, %v) devolvió un error: %v", caso.v1, caso.v2, err)
		} else if !equal(suma, caso.suma) || productoEscalar != caso.productoEscalar {
			t.Errorf("SumaYProductoEscalar(%v, %v) = %v, %d, se esperaba %v, %d", caso.v1, caso.v2, suma, productoEscalar, caso.suma, caso.productoEscalar)
		}
	}
}

func equal(v1, v2 Vector) bool {
	if len(v1) != len(v2) {
		return false
	}
	for i := range v1 {
		if v1[i] != v2[i] {
			return false
		}
	}
	return true
}
