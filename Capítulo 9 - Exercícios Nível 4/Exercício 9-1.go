/*
Cap. 9 – Exercícios: Nível #4 – 1
- Usando uma literal composta:
     - Crie um array que suporte 5 valores to tipo int
     - Atribua valores aos seus índices
- Utilize range e demonstre os valores do array.
- Utilizando format printing, demonstre o tipo do array.

Solução: https://go.dev/play/p/XE-MnHo6Vk0
*/

package main

import (
	"fmt"
)

func main() {

	array := [5]int{1, 2, 3, 4, 5}

	for i, v := range array {
		fmt.Println("Índice", i, "temos o valor", v)
	}

	fmt.Printf("%T", array)
}
