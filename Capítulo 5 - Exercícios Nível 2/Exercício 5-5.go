/*
- Crie uma variável do tipo string utilizando um raw string literal e demonstre-a.

Solução:https://go.dev/play/p/o_YhMGHQP3G?v=gotip
*/
package main

import (
	"fmt"
)

func main() {

	a := `criando uma 
			variável do 
				tipo 
					string`

	fmt.Println(a)

}
