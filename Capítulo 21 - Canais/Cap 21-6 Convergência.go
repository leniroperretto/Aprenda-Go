/*
CANAIS
Convergência
- Observamos convergência quando informação de vários canais é enviada a um número menor de canais.
- Interessante:
- Na prática, exemplos:
    - 1. Todd:
        - Canais par, ímpar, e converge.
        - Func send manda pares pra um, ímpares pro outro, depois fecha.
        - Func receive cria duas go funcs, cada uma com um for range, enviando dados dos canais par e ímpar pro canal converge. Não esquecer de WGs!
        - Por fim um range retira todas as informações do canal converge.
Solução: https://go.dev/play/p/HQpxHGcSiJj

    - 2. Rob Pike (palestra Go Concurrency Patterns):
        - Func trabalho cria um canal, cria uma go func que manda dados pra esse canal, e retorna o canal. Interessante: time.Duration(rand.Intn(1e3))
        - Func converge toma dois canais, cria um canal novo, e cria duas go funcs com for infinito que passa tudo para o canal novo. Retorna o canal novo.
        - Por fim chamamos canal := converge(trabalho(nome1), trabalho(nome2)) e usamos um for para receber dados do canal var.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

	go envia(par, impar)
	go recebe(par, impar, converge)

	for v := range converge {
		fmt.Println("Valor recebido:", v)
	}
}

func envia(p, i chan int) {
	x := 100
	for n := 0; n < x; n++ {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	close(p)
	close(i)
}

func recebe(p, i, c chan int) {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for v := range p {
			c <- v
		}
		wg.Done()

	}()

	go func() {
		wg.Add(1)
		for v := range i {
			c <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(c)
}
