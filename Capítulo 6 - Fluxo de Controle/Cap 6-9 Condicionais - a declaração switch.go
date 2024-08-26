/*
- Switch:
    - pode avaliar uma expressão
        - switch statement == case (value)
        - default switch statement == true (bool)
    - não há fall-through por padrão
    - criando fall-through
    - default
    - cases compostos

Solução:https://go.dev/play/p/XPWtmpEGWp6
*/

package main

import (
	"fmt"
)

func main() {

	quemfezoswitch := ""

	switch quemfezoswitch {
	case "jose":
		fmt.Println("quem fez o switch foi o jose")
		fallthrough // posso usar esse fallthrough para que ele verificar a outra condição se é verdadeira ou não
	case "maria":
		fmt.Println("sempre que o jose faz um switch a maria faz também")
	case "joao":
		fmt.Println("quem fez o switch foi o joao")
	case "joana":
		fmt.Println("quem fez o switch foi a joana")
	default:
		fmt.Println("ninguém da lista fez o switch")
	}
}

// podemos utilizar várias maneiras do switch como por exemplo

func main() {

	quemfezoswitch := "jose"

	switch quemfezoswitch {
	case "jose", "maria":
		fmt.Println("quem fez o switch foi o jose e a maria")
	case "joao", "joana":
		fmt.Println("quem fez o switch foi o joao e a joana")
	default:
		fmt.Println("ninguém da lista fez o switch")
	}
}
