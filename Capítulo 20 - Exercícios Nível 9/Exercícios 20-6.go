/*
Cap. 20 – Exercícios: Nível #9 – 6
- Crie um programa que demonstra seu OS e ARCH.
- Rode-o com os seguintes comandos:
    - go run
    - go build
    - go install
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Seu OS é:\t", runtime.GOOS)
	fmt.Println("Seu ARCH é:\t", runtime.GOARCH)
}
