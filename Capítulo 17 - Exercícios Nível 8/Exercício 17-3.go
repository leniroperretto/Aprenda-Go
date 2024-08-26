/*
Cap. 17 – Exercícios: Nível #8 – 3

- Partindo do código abaixo, utilize NewEncoder() e Encode() para enviar as informações como JSON para Stdout.
    - https://play.golang.org/p/BVRZTdlUZ_
- Desafio: descubra o que é, e utilize, method chaining para conectar os dois métodos acima.

Solução: https://go.dev/play/p/OyQ0UV2HcKF
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(users)

	// usando o Method Chaining

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("Erro", err)
	}

}
