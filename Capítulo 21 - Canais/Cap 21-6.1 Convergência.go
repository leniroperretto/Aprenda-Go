/*
CANAIS
Convergência
- Observamos convergência quando informação de vários canais é enviada a um número menor de canais.
- Interessante:
- Na prática, exemplos:
    - 2. Rob Pike (palestra Go Concurrency Patterns):
        - Func trabalho cria um canal, cria uma go func que manda dados pra esse canal, e retorna o canal. Interessante: time.Duration(rand.Intn(1e3))
        - Func converge toma dois canais, cria um canal novo, e cria duas go funcs com for infinito que passa tudo para o canal novo. Retorna o canal novo.
        - Por fim chamamos canal := converge(trabalho(nome1), trabalho(nome2)) e usamos um for para receber dados do canal var.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	canal := converge(trabalho("Primeiro"), trabalho("Segundo"))
	for x := 0; x < 16; x++ {
		fmt.Println(<-canal)
	}
}

func trabalho(s string) chan string {
	canal := make(chan string)
	go func(s string, c chan string) {
		for i := 1; ; i++ {
			c <- fmt.Sprintf("Função %v diz: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(s, canal)
	return canal
}

func converge(x, y chan string) chan string {
	novo := make(chan string)
	go func() {
		for {
			novo <- <-x
		}
	}()
	go func() {
		for {
			novo <- <-y
		}
	}()
	return novo
}
