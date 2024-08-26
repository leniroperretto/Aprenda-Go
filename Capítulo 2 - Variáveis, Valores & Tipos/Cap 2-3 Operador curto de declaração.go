// para declarar uma variável usa-se o operador := (o famoso gopher ou marmota)

package main

import (
	"fmt"
)

func main() {

	x := 10
	y := "olá bom dia"
	z := 10.2

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n", z, z)

	x = 20 // o = nesse caso enra como operador de atribuição, colocando agora que o valor de x será 20
	fmt.Printf("x: %v, %T\n", x, x)
}
