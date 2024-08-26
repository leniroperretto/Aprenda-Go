// utilizando a palavra var

package main

import (
	"fmt"
)

var x int = 10 // aqui estamos declarando especificamente que o tipo da nossa variável ser int

func main() {

	x = 20.5 // isso não pode ser feito, pois quando setamos a var, colocamos que x seria int (número inteiro) e aqui estamos usando um float
	fmt.Println(x)
	fmt.Printf("%v, %T", x, x)

}

// Tipos de dados primitivos: disponíveis na linguagem nativamente como blocos básicos de construção
//	- int, string, bool
// Tipos de dados compostos: são tipos compostos de tipos primitivos e criados pelo usuário
//	- slice, array, struct, map
