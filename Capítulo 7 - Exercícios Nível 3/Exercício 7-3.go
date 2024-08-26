/*
- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.

Solução:https://go.dev/play/p/nh85KfVJnzx
*/

package main

import (
	"fmt"
)

func main() {

	anodenascimento := 1986
	anoatual := 2024

	for anodenascimento <= anoatual {
		fmt.Println(anodenascimento)
		anodenascimento++
	}
}
