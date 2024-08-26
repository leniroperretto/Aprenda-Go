/*
DEFER
- Funções são ótimas pois tornam nosso código modular. Podemos alterar partes do nosso programa sem afetar o resto!
- Uma declaração defer chama uma função cuja execução ocorrerá no momento em que a função da qual ela faz parte finalizar.
- Essa finalização pode ocorrer devido a um return, ao fim do code block da função, ou no caso de pânico em uma goroutine correspondente.
- "Deixa pra última hora!"
- ref/spec
- Sempre usamos para fechar um arquivo após abri-lo.

Solução: https://go.dev/play/p/PPk7MXtzw15
*/
package main

import (
	"fmt"
)

func main() {
	//	defer fmt.Println("com defer, veio primeiro") // quando usamos o DEFER, ele automáticamente joga a instrução para a última "Hora"
	//	fmt.Println("sem defer, veio depois")

	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
}
