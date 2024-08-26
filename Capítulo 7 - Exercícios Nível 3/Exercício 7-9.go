/*
- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".

Solução:https://go.dev/play/p/OfdL8w64I54
*/

package main

import (
	"fmt"
)

func main() {

	esporteFavorito := "simulador"
	switch esporteFavorito {

	case "futebol":
		fmt.Println("seu esporte favorito é futebol")
	case "volei":
		fmt.Println("seu esporte favorito é volei")
	case "basquete":
		fmt.Println("seu esporte favorito é basquete")
	case "simulador":
		fmt.Println("seu esporte favorito é simulador")
	}
}
