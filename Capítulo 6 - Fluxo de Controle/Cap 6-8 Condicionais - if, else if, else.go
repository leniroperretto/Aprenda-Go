/*
- If, else.
- If, else if, else.
- If, else if, else if, ..., else.

Solução:https://go.dev/play/p/vKJWOkdnfRQ
*/

package main

import (
	"fmt"
)

func main() {

	if x := 500; x > 100 {
		fmt.Println("x é maior do que 100")
	} else if x < 10 {
		fmt.Println("x não é maior que 100")
	} else {
		fmt.Println("x não é menor que 10 e nem maior que 100")
	}
}
