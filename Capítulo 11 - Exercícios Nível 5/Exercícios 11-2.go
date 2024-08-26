/*
Cap. 11 – Exercícios: Nível #5 – 2
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior

Solução: https://go.dev/play/p/4x-P_VA27V8
*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	mapa := make(map[string]pessoa)

	mapa["Silva"] = pessoa{
		nome:      "João",
		sobrenome: "Silva",
		sabores:   []string{"Abacaxi", "Morango", "Baunilha"},
	}

	mapa["Santos"] = pessoa{
		nome:      "Maria",
		sobrenome: "Santos",
		sabores:   []string{"Uva", "Melancia", "Pistache"},
	}

	for _, value := range mapa {
		fmt.Println("Meu nome é", value.nome, "e meus sabores favoritos são:")
		for _, value := range value.sabores {
			fmt.Println("-", value)
		}
	}

}
