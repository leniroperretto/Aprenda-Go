/*
CALLBACK
- Primeiro veja se você entende isso: https://play.golang.org/p/QkAtwMZU-z
- Callback é passar uma função como argumento.
- Exemplo:
    - Criando uma função que toma uma função e um []int, e usa somente os números pares como argumentos para a função.
    - Go Playground:
- Desafio: Crie uma função no programa acima que utilize somente os números ímpares.

Solução: https://go.dev/play/p/0A4fIvM0MWb

*/

package main

import (
	"fmt"
)

func main() {
	t := somenteImpares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println(t)
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func somenteImpares(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 != 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}
