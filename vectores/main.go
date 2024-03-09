package main

import (
	"fmt"
	"vectores"
)

func main() {
	v1 := vectores.Vector{1, 2, 3}
	v2 := vectores.Vector{4, 5, 6}

	suma, productoEscalar, err := vectores.SumaYProductoEscalar(v1, v2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("La suma de los vectores es %v.\n", suma)
	fmt.Printf("El producto escalar de los vectores es %d.\n", productoEscalar)
}
