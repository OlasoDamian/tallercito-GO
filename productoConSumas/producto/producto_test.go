package main

import "testing"

func TestProducto(t *testing.T) {
	casos := []struct {
		a, b, esperado int
	}{
		{3, 4, 12},
		{5, 5, 25},
		{0, 5, 0},
		{5, 0, 0},
		{-3, 4, -12},
		{3, -4, -12},
		{-3, -4, 12},
	}

	for _, caso := range casos {
		resultado := Producto(caso.a, caso.b)
		if resultado != caso.esperado {
			t.Errorf("Producto(%d, %d) = %d, esperado %d", caso.a, caso.b, resultado, caso.esperado)
		}
	}
}
