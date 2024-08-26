/*
Cap. 20 – Exercícios: Nível #9 – 2
 Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
    - Crie um tipo para um struct chamado "pessoa"
    - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
    - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
    - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
    - Demonstre no seu código:
        - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
        - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
- Se precisar de dicas, veja: https://gobyexample.com/interfaces

*/

package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Println(p.nome, "ele fala")
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {

	p1 := pessoa{"Fulano Ciclano", 60}

	p1.falar() // É um atalho para (&p1).falar()

	(&p1).falar() // É a maneira "mais correta".

	dizerAlgumaCoisa(&p1) // Funciona

	// dizerAlgumaCoisa(p1) // Não funciona.
}
