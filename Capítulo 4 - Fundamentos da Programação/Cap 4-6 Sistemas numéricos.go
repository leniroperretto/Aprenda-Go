/*
 - Sistemas numéricos
	- Base-10: decimal, 0-9 usa-se o %d para demonstrar na saída o valor em decimal
	- Base-2: binário, 0-1 usa-se o %b para demonstrar na saída o valor em binário
	- Base-16: hexadecimal, 0-f  usa-se o %#x para demonstrar na saída o valor em hexadecimal

Go playground: https://go.dev/play/p/hfjVxqAS8B3
*/

package main

import (
	"fmt"
)

func main() {

	a := 31337

	fmt.Printf("%d\t%b\t%#x", a, a, a)
}
