/*
- For: inicialização, condição, pós
- For: condição ("while")
- For: ...ever? (http servers)
- For: break
- golang.org/ref/spec#For_statements, Effective Go
- (Range vem mais pra frente.)

Solução: https://go.dev/play/p/XFOsWFJglBH
*/

package main

import (
	"fmt"
)

func main() {

	x := 0
	for {
		if x < 10 {
			x++
			fmt.Println("X é menor que 10")
		} else {
			fmt.Println("X é maior ou igual a 10")
			break
		}
	}
	fmt.Println("o break está funcionando!")

}
