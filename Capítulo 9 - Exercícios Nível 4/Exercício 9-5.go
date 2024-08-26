/*
Cap. 9 – Exercícios: Nível #4 – 5
- Comece com essa slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Utilizando slicing e append, crie uma slice y que contenha os valores:
    - [42, 43, 44, 48, 49, 50, 51]

Solução:https://go.dev/play/p/QWRijgWfyfR
*/

package main

import (
	"fmt"
)

func main() {

	sliceX := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(sliceX)

	slice1 := (sliceX[:3])
	fmt.Println(slice1)

	slice2 := (sliceX[6:])
	fmt.Println(slice2)

	sliceY := append(slice1, slice2...)
	fmt.Println(sliceY)
}

// ou de forma mais resumida conforme a vídeo aula, podemos fazer assim

	sliceX := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(sliceX)
	
	sliceY = append(sliceX[:3], sliceX[len(sliceX)-4]...)
	fmt.Println(sliceY)