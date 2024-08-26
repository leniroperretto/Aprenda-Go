/*
Cap. 11 – Exercícios: Nível #5 – 4
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.

Solução: https://go.dev/play/p/k0F9wIzmS_-
*/

package main

import "fmt"

func main() {

	carros := struct {
		nome    string
		ano     int
		cores   []string
		modelos map[string]string
	}{
		nome:  "Volkswagem",
		ano:   1956,
		cores: []string{"Azul", "Amarelo", "Branco", "Preto"},
		modelos: map[string]string{
			"Fusca":    "Itamar",
			"Brasilia": "dos Mamonas",
			"Gol":      "Chaleira",
		},
	}

	fmt.Println(carros)

}
