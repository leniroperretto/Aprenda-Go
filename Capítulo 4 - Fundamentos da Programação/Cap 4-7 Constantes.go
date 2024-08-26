/*
 - Constantes
	- São valores imutáveis.
	- Podem ser tipadas ou não:
		- const oi = "Bom dia"
		- const oi string = "Bom dia"
	- As não tipadas só terão um tipo atribuído a elas quando forem usadas.
		- exemplo: qual tipo de 42? int? uint? float64?
		- ou seja, é uma flexibilidade conveniente.
	- Na prática: int, float, string.
		- const x = y
		- const (x = y)

Go playground: https://go.dev/play/p/jAhxeRG1lgF
*/

package main

import (
	"fmt"
)

const (
	x = 10
	y = 20
	z = 30
)

func main() {

	fmt.Println(x, y, z)
}
