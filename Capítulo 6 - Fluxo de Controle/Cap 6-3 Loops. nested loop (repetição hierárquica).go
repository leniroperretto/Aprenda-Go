/*
- For
  - Repetição hierárquica
  - Exemplos: relógio, calendário

Solução:https://go.dev/play/p/Ax2ER1ZLoDm?v=gotip
*/
package main

import (
	"fmt"
)

func main() {

	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Hora:", horas)
		for x := 0; x < 60; x++ {
			fmt.Print(" ", x)
		}
		fmt.Println("\n")
	}
}
