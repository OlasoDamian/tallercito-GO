package suma

import "testing"

func TestSumaElementos(t *testing.T) {
	casos := []struct {
		numeros  []int
		esperado int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{-1, -2, -3, -4, -5}, -15},
		{[]int{0, 0, 0, 0, 0}, 0},
		{[]int{}, 0},
	}

	for _, caso := range casos {
		resultado := SumaElementos(caso.numeros)
		if resultado != caso.esperado {
			t.Errorf("SumaElementos(%v) = %d, esperado %d", caso.numeros, resultado, caso.esperado)
		}
	}
}
