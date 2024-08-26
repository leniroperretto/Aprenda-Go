/*
- Exercício:
	- Callback: passe uma função como argumento a outra função.
Solução: https://go.dev/play/p/uYGpzUsLCIK

*/

package main

import (
	"fmt"
)

func main() {
	recebeOutraFuncao(argumento)
}

func argumento() {
	fmt.Println("Isso é um argumento")
}

func recebeOutraFuncao(x func()) {
	fmt.Println("Recebe outra função:")
	x()
}
