/*
- Structs dentro de structs dentro de structs.
- Exemplo: um corredor de fórmula 1 é uma pessoa (nome, sobrenome, idade) e tambem um competidor (nome, equipe, pontos).

Solução: https://go.dev/play/p/0T8zJjo2Sw9
*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type competidor struct {
	pessoa
	equipe string
	pontos int
}

func main() {

	pessoa1 := pessoa{
		nome:      "Ayrton",
		sobrenome: "Senna",
		idade:     30,
	}

	pessoa := competidor{
		pessoa: pessoa{nome: "Rubens",
			sobrenome: "Barrichelo",
			idade:     45},
		equipe: "Ferrari",
		pontos: 150,
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa)
	fmt.Println(pessoa1.idade)
	fmt.Println(pessoa.equipe)
}
