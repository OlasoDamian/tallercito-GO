package factoriales

import "testing"

func TestFactorialde0(t *testing.T) {
	resultado := Factorial(0)
	if resultado != 1 {
		t.Errorf("Para el caso de factorial de cero, se esperaba 1 pero se obtuvo %d", resultado)
	}
}

func TestFactorialDe5(t *testing.T) {
	resultado := Factorial(5)
	if resultado != 120 {
		t.Errorf("Para el caso de factorial de 5, se esperaba 120 pero se obtuvo %d", resultado)
	}
}
