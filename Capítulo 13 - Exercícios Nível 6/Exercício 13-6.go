/*
- Exercício:
	- Crie e utilize uma função anônima.


Solução: https://go.dev/play/p/y2k9ZBj6Mo_6

*/

package main

import (
	"fmt"
)

func main() {

	x := 12
	func(x int) {
		fmt.Println(x, "é uma dúzia")
		fmt.Println(x * 1)
	}(x)
}
