/*
CANAIS
- v, ok := ←chan
- Se receber valor: v, true
- Canal fechado, nada, etc.: zero v, false
- Agora vamos resolver o problema do exercício anterior usando comma ok.

Solução: https://go.dev/play/p/-w2DqWFmAzJ

*/

package main

import "fmt"

func main() {
	canal := make(chan int)

	go func() {
		canal <- 42
		close(canal)
	}()

	v, ok := <-canal

	fmt.Println(v, ok)

	v, ok = <-canal

	fmt.Println(v, ok)
}
