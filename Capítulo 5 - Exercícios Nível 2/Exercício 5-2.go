/*
- Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis.
	- ==
	- !=
	- <=
	- <
	- >=
	- >
- Demonstre estes valores.

Solução: https://go.dev/play/p/CCuhp78rUyh?v=gotip
*/

package main

import (
	"fmt"
)

func main() {
	a := (10 == 100)
	b := (10 != 100)
	c := (10 <= 100)
	d := (10 < 100)
	e := (10 >= 100)
	f := (10 > 100)
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}
