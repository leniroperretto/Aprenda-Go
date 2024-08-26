/*
- Escreva um programa que mostre um número em decimal, binário e hexadecimal.

Solução: https://go.dev/play/p/L1IHxXz098h?v=gotip
*/

package main

import "fmt"

func main() {

	x := 10

	fmt.Printf("%d\n%b\n%#x", x, x, x)
}
