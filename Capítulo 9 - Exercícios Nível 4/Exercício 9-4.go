/*
Cap. 9 – Exercícios: Nível #4 – 4
- Começando com a seguinte slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Anexe a ela o valor 52;
- Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
- Demonstre a slice;
- Anexe a ela a seguinte slice:
    - y := []int{56, 57, 58, 59, 60}
- Demonstre a slice x.

Solução:https://go.dev/play/p/klS5X8uUO1Y
*/

package main

import (
	"fmt"
)

func main() {

	sliceX := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(sliceX)

	sliceX = append(sliceX, 52)
	fmt.Println(sliceX)

	sliceX = append(sliceX, 53, 54, 55)
	fmt.Println(sliceX)

	sliceY := []int{56, 57, 58, 59, 60}

	sliceX = append(sliceX, sliceY...)
	fmt.Println(sliceX)
}
