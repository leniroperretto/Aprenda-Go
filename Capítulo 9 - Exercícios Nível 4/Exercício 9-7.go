/*
Cap. 9 – Exercícios: Nível #4 – 7
- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

Solução:https://go.dev/play/p/g3ucAuTwwjz
*/

package main

import (
	"fmt"
)

func main() {

	slice := [][]string{
		[]string{
			"João",
			"Silva",
			"Pescar",
		},
		[]string{
			"Maria",
			"Pereira",
			"Costurar",
		},
		[]string{
			"Pedro",
			"Santos",
			"Caçar",
		},
	}

	for _, v := range slice {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}
	}
}
