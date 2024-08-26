/*
 - Zero value
 - Atribuindo um valor
 - Bool como resultado de operadores relacionais

Go playground: https://go.dev/play/p/hLjt6VKAtJx
*/

package main

import (
	"fmt"
)

var x bool

func main() {

	fmt.Println(x) // zero value
	x = true
	fmt.Println(x)   // valor atribuído
	x = (10 < 100)   // operadores relacionais
	y := (10 == 100) // operadores relacionais
	z := (10 > 100)  // operadores relacionais
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
