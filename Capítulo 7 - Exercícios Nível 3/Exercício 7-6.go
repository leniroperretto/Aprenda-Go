/*
- Crie um programa que demonstre o funcionamento da declaração if.

Solução:https://go.dev/play/p/ZTFuLtRoZCU
*/

package main

import (
	"fmt"
)

func main() {

	x := 11

	if x == 10 {
		fmt.Println("x é igual a 10")
	} else {
		fmt.Println("x é diferente de 10")
	}

}
