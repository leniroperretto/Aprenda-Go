/*
- Exercício:
   	- Crie um tipo struct "pessoa" que contenha os campos:
	    - nome
	    - sobrenome
	    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.

Solução: https://go.dev/play/p/v79mnywO0YY

*/

package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) demonstrar() {
	fmt.Println("O nome completo dessa pessoa é:", p.nome, p.sobrenome, "\nEssa pessoa tem", p.idade, "anos")
}

func main() {

	p := pessoa{
		nome:      "João",
		sobrenome: "Silva",
		idade:     14,
	}

	fmt.Println(p)
	p.demonstrar()
}
