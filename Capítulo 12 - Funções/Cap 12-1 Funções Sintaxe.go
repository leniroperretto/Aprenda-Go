/*
- Qual a utilidade de funções?
    - Abstrair funcionalidade
    - Reutilização de código
- func (receiver) identifier(parameters) (returns) { code }
- A diferença entre parâmetros e argumentos:
    - Funções são definidas com parâmetros
    - Funções são chamadas com argumentos
- Tudo em Go é pass by value.
    - Pass by reference, pass by copy, ... não.
- Parâmetro pode ser ...variádico.
- Exemplos:
    - Função básica.
        - Go Playground: https://go.dev/play/p/ebkwS1a3XJr
    - Função que aceita um argumento.
        - Go Playground: https://go.dev/play/p/Xc0jaOdbDKg
    - Função com retorno.
        - Go Playground: https://go.dev/play/p/LwxIYwZpPch
    - Função com múltiplos retornos e parâmetro variádico.
        - Go Playground: https://go.dev/play/p/aWJD9aADn3y

Solução:
*/

// Básica

package main

import (
	"fmt"
)

func main() {
	basica()
}

func basica() {
	fmt.Println("Oi, bom dia")
}

// Função que aceita um argumento.

package main

import (
	"fmt"
)

func main() {
	basica()
	argumento("tarde")
}

func basica() {
	fmt.Println("Oi, bom dia")
}

func argumento(s string) {
	if s == "manhã" {
		fmt.Println("Oi, bom dia!")
	} else if s == "tarde" {
		fmt.Println("Oi, boa tarde!")
	} else {
		fmt.Println("Oi, boa noite!")
	}
}

// Função com retorno.

package main

import (
	"fmt"
)

func main() {

	total, quantos := soma(10, 10, 1, 2, 3, 5)
	fmt.Println(total, quantos)
}

func soma(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)
}

// Função com múltiplos retornos e parâmetro variádico.
package main

import (
	"fmt"
)

func main() {

	total, quantos, oi := soma(10, 10, 1, 2, 3, 5)
	fmt.Println(total, quantos, oi)
}

func soma(x ...int) (int, int, string) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), "Bom dia!"
}
