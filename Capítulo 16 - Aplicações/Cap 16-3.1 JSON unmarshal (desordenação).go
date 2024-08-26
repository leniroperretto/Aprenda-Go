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
- Com Encoder: https://go.dev/play/p/isLNZoX_9ip

Solução: https://go.dev/play/p/isLNZoX_9ip
*/

package main

import (
	"encoding/json"
	"os"
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
		Contabancaria: 1000000.5,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(jamesbond)
}
