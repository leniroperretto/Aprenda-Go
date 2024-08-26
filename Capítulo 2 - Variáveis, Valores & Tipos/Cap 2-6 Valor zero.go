// Valor zero

package main

import (
	"fmt"
)

var x int

func main() {

	x = 10
	fmt.Printf("%v, %T", x, x)
	x = 20
	fmt.Printf("%v, %T", x, x)

}

// podemos remover o var e atribuir com o operador curto :=

// Valor zero - Trabalhando com tipos primitivos

package main

import (
	"fmt"
)

var a int
var b float64
var c string
var d bool

func main() {

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

}