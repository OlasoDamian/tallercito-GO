package main

import (
	"fmt"
	"suma"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Printf("La suma de los elementos del arreglo es %d.\n", suma.SumaElementos(numeros))
}
