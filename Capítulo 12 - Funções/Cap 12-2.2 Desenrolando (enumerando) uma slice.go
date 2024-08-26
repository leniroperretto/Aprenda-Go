/*
- Quando temos uma slice, podemos passar os elementos individuais através "deste..." operador.
- Exemplos:
    - Desenrolando uma slice de ints com como argumento para a função "soma" anterior
        - Go Playground: https://go.dev/play/p/RUKlJGZqGUt
    - Pode-se passar zero ou mais valores
        - Go Playground: https://go.dev/play/p/D6lhGTJ0zv0
    - O parâmetro variádico deve ser o parâmetro final → ref/spec#Passing_arguments_to_..._parameters
        - Go Playground:
        - Não roda:

Solução:
*/

package main

import (
	"fmt"
)

func main() {

	total := soma()
	fmt.Println(total)
}

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}
