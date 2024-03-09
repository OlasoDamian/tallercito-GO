package menu

import (
	"fmt"
)

// Post: Se muestra el menú con las opciones disponibles.
func mostrarMenu() {
	fmt.Println("Menú:")
	fmt.Println("* Opción 1")
	fmt.Println("* Opción 2")
	fmt.Println("* Opción 3")
	fmt.Println("* Opción 4")
}

// Post: Se muestra un mensaje indicando la opción seleccionada o "Opción incorrecta" si la opción no es válida.
func elegirOpcion(opcion int) int {
	switch opcion {
	case 1:
		fmt.Println("Se eligió la Opción 1")
		return 1
	case 2:
		fmt.Println("Se eligió la Opción 2")
		return 2
	case 3:
		fmt.Println("Se eligió la Opción 3")
		return 3
	case 4:
		fmt.Println("Se eligió la Opción 4")
		return 4
	default:
		fmt.Println("Opción incorrecta")
		return 0
	}
}
