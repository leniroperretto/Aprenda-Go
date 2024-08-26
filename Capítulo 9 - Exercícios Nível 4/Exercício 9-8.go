/*
Cap. 9 – Exercícios: Nível #4 – 8
- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.

Solução: https://go.dev/play/p/1OtH76liFBf
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

	for key, value := range pessoas {
		fmt.Println(key)
		for i, h := range value {
			fmt.Println("\t", i, " - ", h)
		}
	}
}
