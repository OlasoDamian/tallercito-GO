package producto

func Producto(a, b int) int {
	resultado := 0
	for i := 0; i < b; i++ {
		resultado += a
	}
	return resultado
}
