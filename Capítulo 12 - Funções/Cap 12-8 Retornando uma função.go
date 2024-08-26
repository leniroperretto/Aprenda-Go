/*
RETORNANDO UMA FUNÇÃO
- Pode-se usar uma função como retorno de uma função
- Declaração: func f() return
- Exemplo: func f() func() int { [...]; return func() int{ return [int] } }
    - ????: fmt.Println(f()())

Solução: https://go.dev/play/p/QqDM3WIAXxD

*/

package main

import (
	"fmt"
)

func main() {

	x := retornaumafuncao()
	y := x(3)
	fmt.Println(y)
}

func retornaumafuncao() func(int) int {

	return func(i int) int {
		return i * 10
	}

}
