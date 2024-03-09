package primos

import (
	"fmt"
)

func esPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var num int
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&num)
	if esPrimo(num) {
		fmt.Println("Verdadero")
	} else {
		fmt.Println("Falso")
	}
}
