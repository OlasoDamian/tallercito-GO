package main

import (
	"conjuntos"
	"fmt"
)

func main() {
	A := []int{1, 2, 3, 4, 5}
	B := []int{4, 5, 6, 7, 8}

	unión, intersección := conjuntos.UniónIntersección(A, B)

	fmt.Println("Conjunto A:", A)
	fmt.Println("Conjunto B:", B)
	fmt.Println("Unión:", unión)
	fmt.Println("Intersección:", intersección)
}
