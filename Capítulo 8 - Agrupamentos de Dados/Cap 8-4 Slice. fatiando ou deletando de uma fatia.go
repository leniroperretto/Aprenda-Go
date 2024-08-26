/*
- x[:], x[a:], x[:b], x[a:b]
- "a" é incluso;
- "b" não é.
- Exemplo: cabeça magnética de um disco rígido (relógio, fita).
    - Off-by-one error.


- É fatiando que se deleta um item de uma slice. Na prática:
    - x := append(x[:i], x[:i]...)
    - Go Playground: https://play.golang.org/p/xK2HwCqvwd
- Exercício: tente acessar todos os itens de uma slice sem utilizar range.

Solução: https://go.dev/play/p/C6BE_QCotXS
*/

package main

import (
	"fmt"
)

func main() {

	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatro queijos", "calabresa"}

	fatia := sabores[0:2]
	fatia1 := sabores[2:len(sabores)]
	fatia2 := sabores[3:4]
	fatia3 := sabores[0:5]
	fatia4 := sabores[:] // nesse exemplo estamos percorrendo todo o slice sem especificar o início e o fim.

	fmt.Println(fatia)
	fmt.Println(fatia1)
	fmt.Println(fatia2)
	fmt.Println(fatia3)
	fmt.Println(fatia4)
}
