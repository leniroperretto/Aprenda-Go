/*
DIVERGÊNCIA
- Divergência é o contrário de convergência :)
- Na prática, exemplos:
    - 2. Com throttling! Ou seja, com um número máximo de go funcs.
        - Ídem acima, mas a func que lança go funcs é assim:
        - Cria X go funcs, cada uma com um range no primeiro canal que, para cada item, poe o retorno de trabalho() no canal dois.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	funções := 5
	go manda(100, canal1)
	go outra(funções, canal1, canal2)
	for v := range canal2 {
		fmt.Println(v)
	}
}

func manda(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func outra(funções int, canal1, canal2 chan int) {
	var wg sync.WaitGroup
	for i := 0; i < funções; i++ {
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalho(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * 1000) //time.Duration(rand.Intn(1e3)))
	return n
}
