/*
Cap. 9 – Exercícios: Nível #4 – 2
- Usando uma literal composta:
    - Crie uma slice de tipo int
    - Atribua 10 valores a ela
- Utilize range para demonstrar todos estes valores.
- E utilize format printing para demonstrar seu tipo.

Solução:https://go.dev/play/p/EGLwAH9gL67
*/

package main

import (
	"fmt"
)

func main() {

	slice := []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}

	for i, v := range slice {
		fmt.Println("Índice", i, "temos o valor", v)
	}

	fmt.Printf("%T", slice)
}
