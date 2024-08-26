/*
 - Na prática:
	- Defaults com :=
	- Tipagem com var
	- Dá para colocar número com vírgula em tipo int?
	- Overflow

Go playground: https://go.dev/play/p/_eFzqY-Ebk1
*/

package main

import (
	"fmt"
)

func main() {

	a := "e"
	b := "é"
	c := "\u9999"
	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("%v, %v, %v\n", d, e, f)

	g := 10
	h := 10.1

	fmt.Printf("%v, %T\n", g, g)
	fmt.Printf("%v, %T\n", h, h)

}
