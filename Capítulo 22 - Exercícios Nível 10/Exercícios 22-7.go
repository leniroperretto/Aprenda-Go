/*
Cap. 22 – Exercícios: Nível #10 – 7
- Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
- Tire estes números do canal e demonstre-os.
- Solução: https://go.dev/play/p/MqahkC4F_NJ
*/

package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				canal <- j
			}
		}()
	}
	for k := 0; k < 100; k++ {
		fmt.Println(k, "\t", <-canal)
	}

}
