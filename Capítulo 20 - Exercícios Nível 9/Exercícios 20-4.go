/*
Cap. 20 – Exercícios: Nível #9 – 4
- Utilize mutex para consertar a condição de corrida do exercício anterior.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador int
var mu sync.Mutex

const quantidadeDeGoroutines = 500

func main() {

	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de Goroutines:", quantidadeDeGoroutines, "\nTotal do contador:\t", contador)

}

func criarGoroutines(i int) {
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
		}()
	}
}
