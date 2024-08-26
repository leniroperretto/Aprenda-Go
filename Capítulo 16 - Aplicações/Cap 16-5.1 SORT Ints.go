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
    - sort.Ints: https://go.dev/play/p/ycRIVv62z8X

*/

package main

import (
	"fmt"
	"sort"
)

func main() {

	si := []int{10, 3, 5, 8, 9, 55, 22, 1, 4, 7}

	fmt.Println(si)

	sort.Ints(si)

	fmt.Println(si)
}
