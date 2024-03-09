package minmax

// MinMax devuelve el menor y el mayor número de una lista de enteros.
// Pre: La lista no debe estar vacía.
// Post: Devuelve el menor y el mayor número de la lista.
func EncontrarMinimoMaximo(lista []int) (int, int) {
	if len(lista) == 0 {
		panic("La lista no debe estar vacía")
	}
	min, max := lista[0], lista[0]

	for _, num := range lista {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}
