package main

import "fmt"

func main() {

	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go mandaNumeros(par, impar, quit)

	receive(par, impar, quit)

}

func mandaNumeros(par, impar chan int, quit chan bool) {
	total := 100
	for i := 0; i < total; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
	quit <- true
}

func receive(par, impar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("O número", v, "é par.")
		case v := <-impar:
			fmt.Println("O número", v, "é impar.")
		case v, ok := <-quit:
			if !ok {
				fmt.Println("Ainda está errado", v)
			} else {
				fmt.Println("Encerrando. Recebemos:", v)
			}
			return
		}
	}
}
