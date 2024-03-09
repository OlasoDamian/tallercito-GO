package figuras

import "testing"

func TestCirculo(t *testing.T) {
	c := Circulo{Centro: Punto{X: 0, Y: 0}, Radio: 5}

	if c.Perimetro() != 31 {
		t.Errorf("Perimetro incorrecto. Esperado: %v, Obtenido: %v", 31, c.Perimetro())
	}

	if c.Area() != 78 {
		t.Errorf("Area incorrecta. Esperado: %v, Obtenido: %v", 78, c.Area())
	}

	expectedString := "Círculo: Centro (0, 0), Radio 5"
	if c.ToString() != expectedString {
		t.Errorf("ToString incorrecto. Esperado: %v, Obtenido: %v", expectedString, c.ToString())
	}
}
