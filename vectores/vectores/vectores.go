package vectores

import "errors"

type Vector []int

func SumaYProductoEscalar(v1, v2 Vector) (Vector, int, error) {
	if len(v1) != len(v2) {
		return nil, 0, errors.New("los vectores deben tener el mismo tamaño")
	}

	suma := make(Vector, len(v1))
	productoEscalar := 0

	for i := range v1 {
		suma[i] = v1[i] + v2[i]
		productoEscalar += v1[i] * v2[i]
	}

	return suma, productoEscalar, nil
}
