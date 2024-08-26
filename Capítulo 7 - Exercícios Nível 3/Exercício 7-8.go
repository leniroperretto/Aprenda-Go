/*
- Crie um programa que utilize a declaração switch, sem switch statement especificado.

Solução:https://go.dev/play/p/KiQrNooQONi
*/

package main

import (
	"fmt"
)

func main() {

	x := 11
	switch {

	case x == 10:
		fmt.Println("x é igual a 10")
	case x == 12:
		fmt.Println("x é 12")
	case x != 10, x != 12:
		fmt.Println("x é diferente de 10 e de 12")
	}
}
