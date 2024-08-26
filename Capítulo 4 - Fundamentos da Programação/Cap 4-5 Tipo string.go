/*
 - Tipo String

Go playground: https://go.dev/play/p/h2L6bKmgMKP
*/

package main

import (
	"fmt"
)

func main() {

	s := "ascii éºâ \u9999"

	for _, v := range s {
		fmt.Printf("%v- %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println(" ")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v- %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}

}
