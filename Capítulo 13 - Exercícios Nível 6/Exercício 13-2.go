/*
- Exercício:
   	- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
	- Passe um valor do tipo slice of int como argumento para a função;
	- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
	- Passe um valor do tipo slice of int como argumento para a função.

Solução: https://go.dev/play/p/Hpz2Opfi4uO

*/

package main

import (
	"fmt"
)

func main() {
	si := []int{1, 2, 3, 4, 5, 6, 7}
	si2 := []int{10, 22, 33, 44, 55, 66, 77}

	fmt.Println(variadicoInt(si...))
	fmt.Println(sliceOfInt(si2))
}

func variadicoInt(x ...int) int {
	total := 0
	for _, valor := range x {
		total += valor
	}
	return total
}

func sliceOfInt(x []int) int {
	total := 0
	for _, valor := range x {
		total += valor
	}
	return total
}
