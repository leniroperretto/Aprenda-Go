/*
- O que são tipos de dados compostos?
    - Wikipedia: Composite_data_type
    - Effective Go: Composite literals
    - ref/spec: Composite literals
- Uma slice agrupa valores de um único tipo.
- Criando uma slice: literal composta → x := []type{values}

Solução:https://go.dev/play/p/UZ4F3QL3qHd
*/

package main

import (
	"fmt"
)

func main() {

	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice2 := append(slice, 6) // adicionando mais um campo usando o append no slice original
	fmt.Println(slice2)

	fmt.Println(slice[3]) // alterando o índice 3 para o valor de 348756
	slice[3] = 348756
	fmt.Println(slice[3])

	fmt.Println(slice) // verificando o output com a alteração do índice 3 para 348756
}
