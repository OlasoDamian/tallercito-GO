package main

import (
	"figuras"
	"fmt"
)

func main() {
	var figuras [5]figuras.Figura

	for i := 0; i < 5; i++ {
		fmt.Println("Seleccione la figura a crear:")
		fmt.Println("a. Rectangulo")
		fmt.Println("b. Cuadrado")
		fmt.Println("c. Circulo")

		var opcion string
		fmt.Scanln(&opcion)

		switch opcion {
		case "a":
			rectangulo := figuras.Rectangulo{
				P1: figuras.Punto{X: 0, Y: 0},
				P2: figuras.Punto{X: 5, Y: 4},
			}
			figuras[i] = &rectangulo
		case "b":
			cuadrado := figuras.Cuadrado{
				Pto:  figuras.Punto{X: 0, Y: 0},
				Lado: 5,
			}
			figuras[i] = &cuadrado
		case "c":
			circulo := figuras.Circulo{
				Centro: figuras.Punto{X: 0, Y: 0},
				Radio:  5,
			}
			figuras[i] = &circulo
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
			i--
		}
	}

	for _, figura := range figuras {
		fmt.Println(figura.ToString())
	}
}
