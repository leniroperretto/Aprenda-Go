/*
Cap. 20 – Exercícios: Nível #9 – 5
- Utilize atomic para consertar a condição de corrida do exercício #3.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var contador int32

const quantidadeDeGoroutines = 500

func main() {

	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de Goroutines:", quantidadeDeGoroutines, "\nTotal do contador:\t", contador)

}

func criarGoroutines(i int) {
	for j := 0; j < i; j++ {
		go func() {
			atomic.AddInt32(&contador, 1)
			runtime.Gosched()
			v := atomic.LoadInt32(&contador)
			fmt.Println(v)
		}()
	}
}
