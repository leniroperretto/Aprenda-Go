/*
- Operação módulo: %
- For: break
- For: continue

Solução:https://go.dev/play/p/__0JIJuOQo0
*/

package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 20; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
