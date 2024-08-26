/*
- MAPS
	- Range: for k, v := range map { }
	- Reiterando: maps não tem ordem e um range usará uma ordem aleatória.
	- delete(map, key)
	- Deletar uma key não-existente não retorna erros!

Solução: https://go.dev/play/p/xkaxakoDkB-
*/

package main

import (
	"fmt"
)

func main() {

	mapRange := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é bem legal",
		18:  "idade para ir para a festa",
	}

	fmt.Println(mapRange)

	for key, value := range mapRange {
		fmt.Println(key, value)
	}

	fmt.Println(mapRange)

	total := 0

	for key, _ := range mapRange {
		total += key

	}

	fmt.Println(total)

	delete(mapRange, 123) // deletando um item do Map

	fmt.Println(mapRange) // verificando a saída com o item 123 deletado
}
