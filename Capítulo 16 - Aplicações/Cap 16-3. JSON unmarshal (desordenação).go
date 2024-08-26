/*
JSON
- E agora o contrário.
- Link: https://cdn.jsdelivr.net/gh/GoesToEleven/golang-web-dev@17e3852d/040_json/README.html
- JSON-to-Go
- Tags
- Marshal/unmarshal vs. encoder/decoder
    - Marshal vai pra uma variável
    - Encoder "vai direto"
- Go Playground: https://go.dev/play/p/7ez-uZ4cP76
- Com Encoder:

Solução: https://go.dev/play/p/7ez-uZ4cP76
*/

package main

import (
	"encoding/json"
	"fmt"
)

type informacoes struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Profissao"`
	Contabancaria float64 `json:"Contabancaria"`
}

func main() {

	sb := []byte(`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente secreto","Contabancaria":1000000.5}`)

	var jamesbond informacoes
	err := json.Unmarshal(sb, &jamesbond)
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println(jamesbond)
	fmt.Println(jamesbond.Profissao)
}
