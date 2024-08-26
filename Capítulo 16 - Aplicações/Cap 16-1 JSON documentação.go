/*
DOCUMENTAÇÃO JSON
- Já entendemos ponteiros, já entendemos métodos. Já temos o conhecimento necessário para começar a utilizar a standard library.
- Nesse vídeo faremos uma orientação sobre como abordar a documentação.
- Essa aula não foi preparada. Vai ser tudo ao vivo no improviso pra vocês verem como funciona o processo.
    - golang.org → Documents → Package Documentation
    - godoc.org → encoding/json
        - files
        - examples
        - funcs
        - types
        - methods

Solução: https://go.dev/play/p/og4K6VxRJaG
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
