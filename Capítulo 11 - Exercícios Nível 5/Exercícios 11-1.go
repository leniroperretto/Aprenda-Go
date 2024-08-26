/*
Cap. 11 – Exercícios: Nível #5 – 1
- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.

Solução:https://go.dev/play/p/doMvrnyGDuR
*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	pessoa1 := pessoa{
		nome:      "João",
		sobrenome: "Silva",
		sabores:   []string{"Abacaxi", "Morango", "Baunilha"}}

	pessoa2 := pessoa{
		nome:      "Maria",
		sobrenome: "Santos",
		sabores:   []string{"Uva", "Melancia", "Pistache"}}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.sabores {
		fmt.Println("\t", v)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, v := range pessoa2.sabores {
		fmt.Println("\t", v)
	}
}
