/*
 - Overflow

Go playground: https://go.dev/play/p/hltOmK7Xri5
*/

package main

import (
	"fmt"
)

func main() {

	var i uint16
	i = 65535
	fmt.Println(i)
	i++
	fmt.Println(i)
	i++
	fmt.Println(i)

}
