/*
PONTEIROS
- Ponteiros permitem compartilhar endereços de memória. Isso é útil quando:
    - Não queremos passar grandes volumes de dados pra lá e pra cá
    - Queremos mudar um valor em sua localização original (tudo em Go é pass by value!)
- Exemplos:
    - x := 0; funçãoquemudaovalordoargumentopra1(x); Print(x)
    - x := 0; funçãoquemudaovalordo*argumentopra1(&x); Print(x)

Solução: https://go.dev/play/p/XR8iMRAsVfD
*/

package main

import (
	"fmt"
)

func main() {

	x := 11

	estarecebeovalor(x)
	estarecebeumponteiro(&x)
}

func estarecebeovalor(x int) {
	x++
	fmt.Println("Na função:", x)

}

func estarecebeumponteiro(x *int) {
	*x++
	fmt.Println("Na função:", *x)
}
