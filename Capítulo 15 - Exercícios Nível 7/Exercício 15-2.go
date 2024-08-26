/*
Cap. 15 – Exercícios: Nível #7 – 2
- Crie um struct "pessoa"
- Crie uma função chamada mudeMe que tenha *pessoa como parâmetro. Essa função deve mudar um valor armazenado no endereço *pessoa.
    - Dica: a maneira "correta" para fazer dereference de um valor em um struct seria (*valor).campo
    - Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors
    - "As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method),
    - → x.f is shorthand for (*x).f." ←
    - Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
- Crie um valor do tipo pessoa;
- Use a função mudeMe e demonstre o resultado.

Solução: https://go.dev/play/p/8wXIr13w3Az
*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {

	alguem := pessoa{
		nome:      "João",
		sobrenome: "Zinho",
		idade:     21,
	}

	fmt.Println(alguem)
	mudeMe(&alguem)
	fmt.Println(alguem)
}

func mudeMe(p *pessoa) {
	(*p).nome = "José"
	p.sobrenome = "Zão"
}
