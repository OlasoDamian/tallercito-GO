package conjuntos

func UniónIntersección(A, B []int) (unión, intersección []int) {
	for _, a := range A {
		if !contiene(unión, a) {
			unión = append(unión, a)
		}
		if contiene(B, a) && !contiene(intersección, a) {
			intersección = append(intersección, a)
		}
	}

	for _, b := range B {
		if !contiene(unión, b) {
			unión = append(unión, b)
		}
	}

	return unión, intersección
}

func contiene(arreglo []int, número int) bool {
	for _, a := range arreglo {
		if a == número {
			return true
		}
	}
	return false
}
