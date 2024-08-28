/*
Cap. 22 – Exercícios: Nível #10 – 4
- Utilizando este código: https://play.golang.org/p/MvL6uamrJP
- ...use um select statement para demonstrar os valores do canal.
- Solução: https://go.dev/play/p/RlckhDLMra0

*/

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 0
	}()
	return c
}

func receive(c <-chan int, q chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
