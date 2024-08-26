/*
JSON
- Exemplo: transformando structs em Go em código JSON.
- No improviso tambem.

Solução: https://go.dev/play/p/HN7E6CLA5hH
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

func main() {

	jamesbond := Pessoa{
		Nome:          "James",
		Sobrenome:     "Bond",
		Idade:         40,
		Profissao:     "Agente secreto",
		Contabancaria: 1000000.50,
	}

	darthvader := Pessoa{
		Nome:          "Anakin",
		Sobrenome:     "Skywalker",
		Idade:         50,
		Profissao:     "Vilão Intergalático",
		Contabancaria: 5002500.80,
	}

	j, err := json.Marshal(jamesbond)
	if err != nil {
		fmt.Println(err)
	}
	d, err := json.Marshal(darthvader)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))
	fmt.Println(string(d))

}
