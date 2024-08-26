/*
Cap. 9 – Exercícios: Nível #4 – 9
- Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.

Solução: https://go.dev/play/p/cm8lPnkKbuc
*/

package main

import (
	"fmt"
)

func main() {

	pessoas := map[string][]string{
		"Silva_João": []string{
			"Pescar", "Jogar dominó", "Caçar",
		},
		"Senna_Ayrton": []string{
			"Pilotar", "Esportes radicais", "Esquiar",
		},
		"Nazário_Ronaldo": []string{
			"Jogar futebol", "Comprar Ilhas", "Comprar Clubes",
		},
	}

	pessoas["Santos_Silvio"] = []string{"Lançar aviões de dinheiro", "MAhhhhh Oi", "Vem pra cá, vem pra cá"}

	for key, value := range pessoas {
		fmt.Println(key)
		for i, h := range value {
			fmt.Println("\t", i, " - ", h)
		}
	}
}
