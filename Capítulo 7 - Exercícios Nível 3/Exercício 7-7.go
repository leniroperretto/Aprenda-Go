/*
- Utilizando a solução anterior, adicione as opções else if e else.

Solução:https://go.dev/play/p/gVjk43e5PkS
*/

package main

import (
	"fmt"
)

func main() {

	x := 12

	if x == 10 {
		fmt.Println("x é igual a 10")
	} else if x == 12 {
		fmt.Println("x é 12")
	} else {
		fmt.Println("x é diferente de 10")
	}
}
