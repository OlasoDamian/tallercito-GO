package conjuntos

import (
	"reflect"
	"testing"
)

func TestUniónIntersección(t *testing.T) {
	casos := []struct {
		A, B                 []int
		esperadoUnión        []int
		esperadoIntersección []int
	}{
		{[]int{1, 2, 3}, []int{2, 3, 4}, []int{1, 2, 3, 4}, []int{2, 3}},
		{[]int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}, []int{}},
		{[]int{}, []int{}, []int{}, []int{}},
	}

	for _, caso := range casos {
		unión, intersección := UniónIntersección(caso.A, caso.B)
		if !reflect.DeepEqual(unión, caso.esperadoUnión) || !reflect.DeepEqual(intersección, caso.esperadoIntersección) {
			t.Errorf("UniónIntersección(%v, %v) = %v, %v, se esperaba %v, %v", caso.A, caso.B, unión, intersección, caso.esperadoUnión, caso.esperadoIntersección)
		}
	}
}
