package primos

import "testing"

func TestEsPrimo(t *testing.T) {
	casos := []struct {
		n        int
		esperado bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{12, false},
		{13, true},
		{14, false},
		{15, false},
		{16, false},
		{17, true},
		{18, false},
		{19, true},
		{20, false},
	}

	for _, caso := range casos {
		resultado := esPrimo(caso.n)
		if resultado != caso.esperado {
			t.Errorf("esPrimo(%d) = %v, se esperaba %v", caso.n, resultado, caso.esperado)
		}
	}
}
