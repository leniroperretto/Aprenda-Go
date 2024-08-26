/*
- Numa declaração de constantes, o identificador iota representa números sequenciais.
- Na prática.
    - iota, iota + 1, a = iota b c, reinicia em cada const, _

Go playground: https://go.dev/play/p/ZuIzsBNV_xE
*/

package main

import (
	"fmt"
)

const (
	a = iota
	_ = iota // usando o _ no lugar do b estou descartando esse valor do output
	c = iota
	x = iota
	_ = iota // usando o _ no lugar do y estou descartando esse valor do output
	z = iota
) // o Println deveria ser mostrado como 0 1 2 3 4 5, mas como trocamos o 1 e o 4 por _ ele some

func main() {

	fmt.Println(a, c, x, z)
}
