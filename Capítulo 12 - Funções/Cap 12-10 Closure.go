/*
CLOSURE
- Closure é cercar ou capturar um scope para que possamos utilizá-lo em outro contexto. Já vimos:
    - Package-level scope
    - Function-level scope
    - Code-block-in-code-block scope
- Exemplo de closure:
    - func i() func() int { x := 0; return func() int { x++; return x } }
    - Quando fizermos a := i() teremos um scope, um valor para x.
    - Quando fizermos b := i() teremos outro scope, e x terá um valor independente do x acima.
- Closures nos permitem salvar dados entre function calls e ao mesmo tempo isolar estes dados do resto do código.

Solução: https://go.dev/play/p/ehL3Tt_Fl3u

*/

package main

import (
	"fmt"
)

func main() {
	a := i()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := i()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
