/*
Cap. 18 – Concorrência – 3. Discussão: Condição de corrida

- Aqui vamos replicar a race condition mencionada no artigo anterior.
    - time.Sleep(time.Second) vs. runtime.Gosched()
- go help → go help build → go run -race main.go
- Como resolver? Mutex.

Solução:https://go.dev/play/p/2RJFCJsotIz
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	contador := 0
	totalGoRoutines := 1000

	var wg sync.WaitGroup
	wg.Add(totalGoRoutines)

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final", contador)
}
