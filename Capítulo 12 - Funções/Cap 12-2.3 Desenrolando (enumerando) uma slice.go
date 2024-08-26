/*
- Quando temos uma slice, podemos passar os elementos individuais através "deste..." operador.
- Exemplos:
    - Desenrolando uma slice de ints com como argumento para a função "soma" anterior
        - Go Playground: https://go.dev/play/p/RUKlJGZqGUt
    - Pode-se passar zero ou mais valores
        - Go Playground: https://go.dev/play/p/D6lhGTJ0zv0
    - O parâmetro variádico deve ser o parâmetro final → ref/spec#Passing_arguments_to_..._parameters
        - Go Playground: https://go.dev/play/p/mh-iHZhjhto
        - Não roda: https://go.dev/play/p/mh-iHZhjhto

Solução:
*/

package main

import (
	"fmt"
)

func main() {
	total, quantos, oi := soma("tarde", 10, 10, 1, 2, 3, 5)

	fmt.Println(total, quantos, oi)
}

// só é possível usar o ... como último parâmetro da lista
func soma(s string, x ...int) (int, int, string) {
	oi := ""
	if s == "manhã" {
		oi = "Oi, bom dia!"
	} else if s == "tarde" {
		oi = "Oi, boa tarde!"
	} else {
		oi = "Oi, boa noite!"
	}

	soma := 0
	for _, v := range x {
		soma += v
	}

	return soma, len(x), oi
}
