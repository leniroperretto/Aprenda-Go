/*
RECURSIVIDADE
- WP: "The most common application of recursion is in mathematics and computer science, where a function being defined is applied within its own definition."
- Exemplos de recursividade: Fractais, matrioscas, efeito Droste (o efeito produzido por uma imagem que aparece dentro dela própria), GNU (“GNU is Not Unix”), etc.
- No estudo de funções: é uma função que chama a ela própria.
- Exemplo: fatoriais.
    - 4! = 4 * 3 * 2 * 1 (e no zero, deu.)
    - Com recursividade. Go Playground: https://go.dev/play/p/MUDd3eJQjPy
    - Com loops. Go Playground: https://go.dev/play/p/LMCD-KUTOna

Solução: https://go.dev/play/p/LMCD-KUTOna

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(fatorial(4))
	// fatorial de 4! = 4 * 3 * 2 * 1
	fmt.Println(loops(4))
}

func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}

func loops(x int) int {
	total := 1
	for x > 1 {
		total *= x // total = total * x
		x--
	}
	return total
}
