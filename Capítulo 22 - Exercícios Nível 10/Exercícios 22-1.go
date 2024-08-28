/*
Cap. 22 – Exercícios: Nível #10 – 1
- Nível 10?! Êita! Parabéns!
- Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
    - Usando uma função anônima auto-executável
    - Usando buffer
- Solução:
    - 1. https://go.dev/play/p/cxXvCA0j7th
    - 2. https://go.dev/play/p/CPpipyvp_Nr

*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}
