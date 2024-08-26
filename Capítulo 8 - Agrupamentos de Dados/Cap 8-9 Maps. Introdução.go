/*
- MAPS
	- Utiliza o formato key:value.
	- E.g. nome e telefone
	- Performance excelente para lookups.
	- map[key]value{ key: value }
	- Acesso: m[key]
	- Key sem value retorna zero. Isso pode trazer problemas.
	- Para verificar: comma ok idiom.
	    - v, ok := m[key]
	    - ok é um boolean, true/false
- Na prática: if v, ok := m[key]; ok { }
	- Para adicionar um item: m[v] = value
	- Maps não tem ordem.

Solução: https://go.dev/play/p/E_jB-bWM71E
*/

package main

import (
	"fmt"
)

func main() {

	amigos := map[string]int{
		"alfredo": 5551234,
		"joana":   9996674,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["joana"])

	amigos["gopher"] = 444444

	fmt.Println(amigos)
	fmt.Println(amigos["gopher"], "\n\n")

	// comma ok idiom
	if amigoFantasma, ok := amigos["fantasma"]; !ok {
		fmt.Println("não tem!")
	} else {
		fmt.Println(amigoFantasma)
	}

}
