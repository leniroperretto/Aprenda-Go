/*
- Exercício:
	- Crie uma função que retorna uma função.
	- Atribua a função retornada a uma variável.
	- Chame a função retornada.

Solução: https://go.dev/play/p/afTPSeLGzfN

*/

package main

import (
	"fmt"
)

func main() {

	x := retornaUmaFuncao()
	x()
}

func retornaUmaFuncao() func() {
	return func() {
		fmt.Println("Retorna uma função")
	}
}
