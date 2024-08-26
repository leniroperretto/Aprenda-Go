/*
Cap. 9 – Exercícios: Nível #4 – 10
- Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.

Solução: https://go.dev/play/p/8ZhUa1q3wM1
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

	delete(pessoas, "Silva_João")

	for key, value := range pessoas {
		fmt.Println(key)
		for i, h := range value {
			fmt.Println("\t", i, " - ", h)
		}

	}
}
