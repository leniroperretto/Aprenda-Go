/*
FUNÇÕES ANÔNIMOS
- Anonymous self-executing functions → Funções anônimas auto-executáveis.
- func(p params) { ... }()
- Vamos ver bastante quando falarmos de goroutines.

Solução:https://go.dev/play/p/-dJNzLPB-ED

*/

package main

import (
	"fmt"
)

func main() {

	x := 387
	func(x int) {
		fmt.Println(x, "vezes 873648 é:")
		fmt.Println(x * 873648)
	}(x)
}
