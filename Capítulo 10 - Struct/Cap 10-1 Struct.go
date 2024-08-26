/*
- Struct é um tipo de dados composto que nos permite armazenar valores de tipos diferentes.
- Seu nome vem de "structure," ou estrutura.
- Declaração: type x struct { y: z }
- Acesso: x.y
- Exemplo: nome, idade, fumante.

Solução: https://go.dev/play/p/rn9nfg707Ar
*/

package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {

	cliente1 := cliente{
		nome:      "João",
		sobrenome: "Silva",
		fumante:   false,
	}

	cliente2 := cliente{"Joana", "Pereira", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
}
