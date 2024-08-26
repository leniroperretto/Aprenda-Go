/*
MUTEX
- Agora vamos resolver a race condition do programa anterior utilizando mutex.
- Mutex é mutual exclusion, exclusão mútua.
- Utilizando mutex somente uma thread poderá utilizar a variável contador de cada vez, e as outras deve aguardar sua vez "na fila."
- Na prática:
    - type Mutex
        - func (m *Mutex) Lock()
        - func (m *Mutex) Unlock()
- RWMutex

Solução: https://go.dev/play/p/Uwtbp7PGisV
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

	var mu sync.Mutex

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final", contador)
}
