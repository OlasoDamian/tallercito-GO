package factoriales

// Factorial calcula el factorial de un número entero no negativo.
// Precondición: n >= 0.
// Postcondición: Devuelve el factorial de n.
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}
