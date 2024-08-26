/*
- Exercício:
   	- Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.

Solução: https://go.dev/play/p/KnBTRYmWk8M

*/

package main

import (
	"fmt"
)

func main() {

	defer fmt.Println("Primeiro que vira último")
	defer fmt.Println("Segundo que vira penúltimo")
	defer fmt.Println("Terceiro que vira antepenúltimo")
	fmt.Println("Quarto que vira primeiro")
	fmt.Println("Quinto que vira segundo")
	fmt.Println("Sexto que vira terceiro")
}
