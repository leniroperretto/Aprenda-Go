/*
RANGE
- Slices:
    - Tamanho: len(x)
    - Índice específico: x[i] (0-based)
- Para ver todos os itens de uma slice utilizamos o loop for com range.
- Range significa alcance, faixa, extensão.
- For range: for i, v := range x {}

Solução:https://go.dev/play/p/V4mfLAV7R5Q
*/

package main

import (
	"fmt"
)

func main() {

	slice := []string{"banana", "maça", "abacaxi"}

	// range roda uma função para cada elemento do slice
	for índice, valor := range slice {
		fmt.Println("No índice", índice, "temos o valor:", valor)
	}

	slice = append(slice, "melancia") // só irá adicionar valor usando o append

	for índice, valor := range slice {
		fmt.Println("No índice", índice, "temos o valor:", valor)
	}
}

// outra opção para somar os campos do slice

func main() {
	slice := []int{20, 21, 22, 23}

	total := 0

	for _, valor := range slice {
		// mesma coisa que total = total + valor
		total += valor
	}

	fmt.Println("O valor total é:", total)
}
