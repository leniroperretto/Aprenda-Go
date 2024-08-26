/*
Cap. 18 – Concorrência – 2. Goroutines & WaitGroups

- O código abaixo é linear. Como fazer as duas funções rodarem concorrentemente?
    - https://play.golang.org/p/XP-ZMeHUk4
- Goroutines!
- O que são goroutines? São "threads."
- O que são threads? [WP](https://pt.wikipedia.org/wiki/Thread_...)
- Na prática: go func.
- Exemplo: código termina antes da go func executar.
- Ou seja, precisamos de uma maneira pra "sincronizar" isso.
- Ah, mas então... não.
- Qualé então? sync.WaitGroup:
- Um WaitGroup serve para esperar que uma coleção de goroutines termine sua execução.
    - func Add: "Quantas goroutines?"
    - func Done: "Deu!"
    - func Wait: "Espera todo mundo terminar."
- Ah, mas então... sim!
- Só pra ver: runtime.NumCPU() & runtime.NumGoroutine()

Solução: https://go.dev/play/p/M9i-si2RtMk
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(2)

	go func1()
	go func2()

	fmt.Println(runtime.NumGoroutine())

	wg.Wait()
}

func func1() {
	for i := 0; i < 100; i++ {
		fmt.Println("func1:", i)
		time.Sleep(20)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2:", i)
	}
	wg.Done()
}
