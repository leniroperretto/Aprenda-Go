/*
SORT
- Sort serve para ordenar slices.
    - Como faz?
    - golang.org/pkg/ → sort
    - godoc.org/sort → examples
    - Sort altera o valor original!
- Exemplo: Ints, Strings.

Solução:
    - sort.Strings: https://go.dev/play/p/nhaaXlZzxa9
    - sort.Ints:

*/

package main

import (
	"fmt"
	"sort"
)

func main() {

	ss := []string{"Morango", "Uva", "Kiwi", "Laranja", "Pêssego", "Damasco", "Amora", "Côco"}

	fmt.Println(ss)

	sort.Strings(ss)

	fmt.Println(ss)
}
