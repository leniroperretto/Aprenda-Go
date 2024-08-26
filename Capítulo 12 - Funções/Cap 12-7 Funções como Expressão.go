/*
FUNÇÕES COMO EXPRESSÃO
- f := func(p params){ ... }
- f()

Solução: https://go.dev/play/p/Jejvjbw3VLf

*/

package main

import (
	"fmt"
)

func main() {

	x := 12
	y := func(x int) int {
		return x * 873648
	}

	fmt.Println(x, "vezes 873648 é:", y(x))
}
