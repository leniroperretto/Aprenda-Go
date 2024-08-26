// Sprint formats using the default formats for its operands and returns the resulting string. Spaces are added between operands when neither is a string.
// Print -> string, pode ser usado como variável
// - func Sprint(a ...interface{}) string
// - func Sprintf(format string, a ...interface{}) string
// - func Sprintln(a ...interface{}) string
package main

import (
	"fmt"
)

func main() {

	x := "oi"
	y := "bom dia"

	z := fmt.Sprint(x, y)
	z = fmt.Sprint(x, " ", y) // caso eu queira adicionar um espaço entre as strings

	fmt.Println(z)
}
