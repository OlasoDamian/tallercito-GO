package main

import (
	"fmt"
)

func swap(px, py *int) {
	*px, *py = *py, *px
}

func main() {
	x := 5
	y := 10
	fmt.Printf("Antes del intercambio: x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("Después del intercambio: x = %d, y = %d\n", x, y)
}
