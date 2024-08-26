/*
- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.

Solução:https://go.dev/play/p/F4hGDdMct1M
*/

package main

import (
	"fmt"
)

func main() {

	anoemqueeunasci := 1988
	anoateoqualeuquerocontar := 2088

	for {
		if anoemqueeunasci > anoateoqualeuquerocontar {
			break
		}
		fmt.Println(anoemqueeunasci)
		anoemqueeunasci++
	}
}
