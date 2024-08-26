/*
- && = quer dizer E
- || = quer dizer OU
- ! = quer dizer NÃO
- Qual o resultado de fmt.Println...
    - true && true
    - true && false
    - true || true
    - true || false
    - !true

Solução:https://go.dev/play/p/T7FXzH-yeB3
*/

package main

import (
	"fmt"
)

func main() {

	x := 6

	if x == 2 || x == 3 {
		fmt.Println("x é 2 ou 3")
	}

	if x%2 == 0 && x%3 == 0 {
		fmt.Println("x é multiplo de 2 e 3")
	}
	fmt.Println("x não é 2 ou 3")
}
