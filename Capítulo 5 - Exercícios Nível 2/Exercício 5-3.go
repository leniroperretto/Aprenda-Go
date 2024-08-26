/*
- Crie constantes tipadas e não-tipadas
- Demonstre estes valores.

Solução:https://go.dev/play/p/b-Z9IHLX6od?v=gotip
*/

package main

import (
	"fmt"
)

const (
	a int     = 10
	b string  = "Tipada"
	c float64 = 10.4
	d bool    = true
	e         = 15
	f         = "não tipada"
)

func main() {
	fmt.Printf("%v\n%T\n%v\n%T\n%v\n%T\n%v\n%T\n%v\n%T\n%v\n%T\n", a, a, b, b, c, c, d, d, e, e, f, f)
}
