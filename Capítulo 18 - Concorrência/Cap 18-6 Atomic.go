/*
ATOMIC
- Agora vamos fazer a mesma coisa, mas com atomic ao invés de mutex.
    - atomic.AddInt64
    - atomic.LoadInt64

Solução: https://go.dev/play/p/DgeM5bXP9SM
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())

	var contador int64
	contador = 0
	totalGoRoutines := 1000

	var wg sync.WaitGroup
	wg.Add(totalGoRoutines)

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			v := contador
			runtime.Gosched()
			v++
			contador = v
			fmt.Println("Contador:\t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final", contador)
}
