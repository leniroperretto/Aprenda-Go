/*
 - Deslocamento de Bits
	- Deslocamento de bits é quando deslocamos digitos binários para a esquerda ou direita.
	- Na prática:
		- %d %b
		- x << y
		- iota * 10 << 10 = kb, mb, gb

Go playground: https://go.dev/play/p/nZZ115tzTSP?v=gotip
*/

package main

import (
	"fmt"
)

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {

	fmt.Println("bynary\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
}
