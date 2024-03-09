package suma

func SumaElementos(numeros []int) int {
	suma := 0
	for _, numero := range numeros {
		suma += numero
	}
	return suma
}
