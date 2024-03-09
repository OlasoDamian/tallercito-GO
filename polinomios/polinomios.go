package polinomios

import (
	"fmt"
)

// Pre: El parámetro coeficientes no debe ser nulo.
// Post: La cadena devuelta representa el polinomio formado por los coeficientes.
func ImprimirPolinomio(coeficientes []float64) string {
	if len(coeficientes) == 0 {
		return "El array está vacío"
	} else {
		grado := len(coeficientes) - 1
		var termino string
		fmt.Print("Polinomio: ")
		for i := 0; i <= grado; i++ {
			if i == 0 {
				termino = fmt.Sprintf("%.1f", coeficientes[i])
			} else if i == 1 {
				termino = termino + fmt.Sprintf("%.1f", coeficientes[i]) + " x"
			} else if i > 1 {
				termino = termino + fmt.Sprintf("%.1f", coeficientes[i]) + " x**" + fmt.Sprintf("%d", i)
			}
			if i != grado {
				termino = termino + " + "
			}
		}
		return termino
	}

}
