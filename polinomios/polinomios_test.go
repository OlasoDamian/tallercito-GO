package polinomios

import (
	"testing"
)

func TestImprimirPolinomio(t *testing.T) {
	coeficientes := []float64{3.0, 2.0, 1.0}
	resultadoEsperado := "3.0 + 2.0 x + 1.0 x**2"
	resultado := ImprimirPolinomio(coeficientes)
	println(resultadoEsperado)
	println(resultado)
	if resultado != resultadoEsperado {
		t.Errorf("Error: resultado incorrecto. Se esperaba '%s' pero se obtuvo '%s'", resultadoEsperado, resultado)
	}
}
func TestImprimirPolinomioVacio(t *testing.T) {
	coeficientes := []float64{}
	resultadoEsperado := "El array está vacío"
	resultado := ImprimirPolinomio(coeficientes)
	println(resultadoEsperado)
	println(resultado)
	if resultado != resultadoEsperado {
		t.Errorf("Error: resultado incorrecto. Se esperaba '%s' pero se obtuvo '%s'", resultadoEsperado, resultado)
	}
}
