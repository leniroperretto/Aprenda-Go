/*
Cap. 17 – Exercícios: Nível #8 – 5

- Partindo do código abaixo, ordene os []user por idade e sobrenome.
    - https://play.golang.org/p/BVRZTdlUZ_
- Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.

Solução: https://go.dev/play/p/vAKykHB2Cmq
*/

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ordenarPorIdade []user

func (x ordenarPorIdade) Len() int           { return len(x) }
func (x ordenarPorIdade) Less(i, j int) bool { return x[i].Age < x[j].Age }
func (a ordenarPorIdade) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ordenarPorSobrenome []user

func (x ordenarPorSobrenome) Len() int           { return len(x) }
func (x ordenarPorSobrenome) Less(i, j int) bool { return x[i].Last < x[j].Last }
func (a ordenarPorSobrenome) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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

	// your code goes here

	fmt.Println(users)
	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	sort.Sort(ordenarPorIdade(users))
	fmt.Println("Ordenando por Idade:")
	harmoniosa(users)

	sort.Sort(ordenarPorSobrenome(users))
	fmt.Println("Ordenando por Sobrenome:")
	harmoniosa(users)

}

func harmoniosa(x []user) {
	for i, v := range x {
		fmt.Println(i, "\tIdade:", v.Age, "\tNome Completo:", v.First, v.Last)
		for _, c := range v.Sayings {
			fmt.Println("\t", c)
		}
	}
}
