package menu

import (
	"testing"
)

func TestElegirOpcionValida(t *testing.T) {
	// Probamos cada opción válida
	for i := 1; i <= 4; i++ {
		resultado := elegirOpcion(i)
		if resultado != i {
			t.Errorf("Para la opción %d, se esperaba %d pero se obtuvo %d", i, i, resultado)
		}
	}
}
func TestElegirOpcionInvalida(t *testing.T) {
	// Probamos una opción inválida
	resultado := elegirOpcion(5)
	if resultado != 0 {
		t.Errorf("Para la opción inválida, se esperaba 0 pero se obtuvo %d", resultado)
	}
}
