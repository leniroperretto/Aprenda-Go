/*
Cap. 15 – Exercícios: Nível #7 – 1
- Crie um valor e atribua-o a uma variável.
- Demonstre o endereço deste valor na memória.

Solução: https://go.dev/play/p/H5cC2kw0lg8
*/

package main

import (
	"fmt"
)

func main() {

	x := 12
	fmt.Println(x)
	fmt.Println(&x)
}
