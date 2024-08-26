/*
- Exercício:
   - Crie um tipo "quadrado"
   - Crie um tipo "círculo"
   - Crie um método "área" para cada tipo que calcule e retorne a área da figura
	- Área do círculo: 2 * π * raio
	- Área do quadrado: lado * lado
   - Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
   - Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
   - Crie um valor de tipo "quadrado"
   - Crie um valor de tipo "círculo"
   - Use a função "info" para demonstrar a área do "quadrado"
   - Use a função "info" para demonstrar a área do "círculo"

Solução: https://go.dev/play/p/7E7YpzJ82_j

*/

package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

type info interface {
	area()
}

func medida(i info) {
	i.area()
}

func (q quadrado) area() {
	resultado := q.lado * q.lado
	fmt.Println("Área do quadrado:", resultado)
}

func (c circulo) area() {
	resultado := math.Pi * 2 * c.raio
	fmt.Println("Área do círculo:", resultado)
}

func main() {

	x := quadrado{
		lado: 15.0,
	}

	y := circulo{
		raio: 13.0,
	}

	medida(x)
	medida(y)
}
