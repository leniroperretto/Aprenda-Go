/*
- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos

Solução:https://go.dev/play/p/s4_4nth52Eh?v=gotip
*/
package main

import (
	"fmt"
)

const (
	_ = 2024 + iota
	b
	c
	d
	e
)

func main() {

	fmt.Println(b, c, d, e)

}
