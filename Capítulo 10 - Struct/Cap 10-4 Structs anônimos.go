/*
- Structs dentro de structs dentro de structs.
- Exemplo: um corredor de fórmula 1 é uma pessoa (nome, sobrenome, idade) e tambem um competidor (nome, equipe, pontos).

Solução: https://go.dev/play/p/9DUto8539ZN
*/

package main

import "fmt"

func main() {

	pessoa := struct {
		nome  string
		idade int
	}{

		nome:  "José",
		idade: 45,
	}

	fmt.Println(pessoa)

}
