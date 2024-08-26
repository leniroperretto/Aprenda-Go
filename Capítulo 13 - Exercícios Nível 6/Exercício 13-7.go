/*
- Exercício:
	- Atribua uma função a uma variável.
	- Chame essa função.

Solução: https://go.dev/play/p/Fq7j7yl-n6D

*/

package main

import (
	"fmt"
)

func main() {

	x := func() {
		fmt.Println("Uma função a uma variável")
	}
	x()
}
