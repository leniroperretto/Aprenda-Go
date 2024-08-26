/*
- If: bool
- If: o operador não → "!"
- If: declaração de inicialização

Solução:https://go.dev/play/p/QL1AhHichnT
*/

package main

import (
	"fmt"
)

func main() {

	x := 10
	if !(x < 100) {
		fmt.Println("Executei")
	}
	fmt.Println("Não Executei")
}
