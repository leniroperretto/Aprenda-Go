/*
Cap. 22 – Exercícios: Nível #10 – 5
- Utilizando este código: https://play.golang.org/p/YHOMV9NYKK
- ...demonstre o comma ok idiom.
- Solução: https://go.dev/play/p/WHGdN_x9cvi

*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
