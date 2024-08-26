/*
- Crie um programa que:
	- Atribua um valor int a uma variável
	- Demonstre este valor em decimal, binário e hexadecimal
	- Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
	- Demonstre estra outra variável em decimal, binário e hexadecimal

Solução:https://go.dev/play/p/KAHrO0ejMeU?v=gotip
*/

package main

import "fmt"

func main() {

	a := 200
	fmt.Printf("%b\t%d\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%b\t%d\t%#x", b, b, b)
}
