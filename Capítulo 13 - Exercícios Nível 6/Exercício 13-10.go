/*
- Exercício:
	- Demonstre o funcionamento de um closure.
	- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.

Solução: https://go.dev/play/p/6bSUYrtSlLn

*/

package main

import (
	"fmt"
)

func main() {
	a := i()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

}

func i() func() int {
	x := 2
	return func() int {
		x++
		return x
	}
}
