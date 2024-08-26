/*
- Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

Solução:https://go.dev/play/p/9tnfH4-c2Jv
*/

package main

import (
	"fmt"
)

func main() {

	for x := 10; x <= 100; x++ {
		fmt.Println(x, x%4)
	}
}
