/*
- Effective Go: append (package builtin)
- x = append(slice, ...values)
- x = append(slice, slice...)
- Todd: unfurl → desdobrar, desenrolar
- Nome oficial: enumeration

Solução:https://go.dev/play/p/pxjGrH76V7Z
*/

package main

import (
	"fmt"
)

func main() {

	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{9, 10, 11, 12}

	fmt.Println(umaslice)

	umaslice = append(umaslice, 5, 6, 7, 8)
	fmt.Println(umaslice)

	umaslice = append(umaslice, outraslice...) // é necessário utilizar o "..." para o append entender que você quer adicionar o conteúdo do slice para adicionar ao append, ou seja, estamos renumerando.
	fmt.Println(umaslice)

}
