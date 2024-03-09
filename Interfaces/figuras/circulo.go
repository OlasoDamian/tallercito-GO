package figuras

import (
	"fmt"
	"math"
)

type Circulo struct {
	Centro Punto
	Radio  int
}

func (c *Circulo) Mover(incX, incY int) {
	c.Centro.Mover(incX, incY)
}

func (c Circulo) Perimetro() int {
	return int(2 * math.Pi * float64(c.Radio))
}

func (c Circulo) Area() int {
	return int(math.Pi * float64(c.Radio) * float64(c.Radio))
}

func (c Circulo) ToString() string {
	return fmt.Sprintf("Círculo: Centro %s, Radio %d", c.Centro.ToString(), c.Radio)
}
