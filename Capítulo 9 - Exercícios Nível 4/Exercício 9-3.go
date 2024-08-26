/*
Cap. 9 – Exercícios: Nível #4 – 3
- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
    - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
    - Do quinto ao último item do slice (incluindo o último item!)
    - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
    - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
    - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

Solução: https://go.dev/play/p/cEY8Ic3EPt0
*/

package main

import (
	"fmt"
)

func main() {
	// índice       0     1    2    3    4    5    6    7    8     9
	slice := []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:9])
	fmt.Println(slice[2 : len(slice)-1])
}
