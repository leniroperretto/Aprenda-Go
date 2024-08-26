/*
Cap. 20 – Exercícios: Nível #9 – 1
 Alem da goroutine principal, crie duas outras goroutines.
- Cada goroutine adicional devem fazer um print separado.
- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	novaGoRoutine(100)
	wg.Wait()
}

func novaGoRoutine(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		x := j
		go func(i int) {
			fmt.Println("Eu sou a GoRoutine número:", i+1)
			wg.Done()
		}(x)
	}
}
