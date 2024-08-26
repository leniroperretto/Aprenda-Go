/*
INTERFACES E POLIMORFISMO
- Em Go, valores podem ter mais que um tipo.
- Uma interface permite que um valor tenha mais que um tipo.
- Declaração: keyword identifier type → type x interface
- Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface.
- Se um tipo possuir todos os métodos necessários (que, no caso da interface{}, pode ser nenhum) então esse tipo implicitamente implementa a interface.
- Esse tipo será o seu tipo e também o tipo da interface.

- Exemplos:
  - Os tipos profissão1 e profissão2 contem o tipo pessoa
  - Cada um tem seu método oibomdia()*, e podem dar oi utilizando *pessoa.oibomdia()
  - Implementam a interface gente
  - Ambos podem acessar o função serhumano() que chama o método oibomdia() de cada gente
  - Tambem podemos no método serhumano() tomar ações diferentes dependendo do tipo:

Solução antes do switch: https://go.dev/play/p/El0YJVES-xq

    switch pessoa.(type) { case profissão1: fmt.Println(h.(profissão1).valorquesóexisteemprofissão1) [...] }*

- Onde se utiliza?
  - Área de formas geométricas (gobyexample.com)
  - Sort
  - DB
  - Writer interface: arquivos locais, http request/response

- Se isso estiver complicado, não se desespere. É foda mesmo. Com tempo e prática a fluência vem.

Solução: https://go.dev/play/p/CsGqF0ac_t9

*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesarrancados int
	salario          float64
}

type arquiteto struct {
	pessoa
	tipodeconstrução string
	tamanhodaloucura string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é", x.nome, x.sobrenome, "bom dia! E eu já arranquei", x.dentesarrancados, "dentes.")
}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é", x.nome, x.sobrenome, "bom dia!")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
	switch g.(type) {
	case dentista:
		fmt.Println("Eu ganho R$", g.(dentista).salario, "reais.")
	case arquiteto:
		fmt.Println("Eu construo", g.(arquiteto).tipodeconstrução, ".")
	}

}

func main() {

	mrdente := dentista{
		pessoa: pessoa{
			nome:      "Dr. Dentista",
			sobrenome: "Obturação",
			idade:     45,
		},
		dentesarrancados: 8973,
		salario:          15486.10,
	}

	mrpredio := arquiteto{
		pessoa: pessoa{
			nome:      "Ar. Arquiteto",
			sobrenome: "Arranha Céu",
			idade:     28,
		},
		tipodeconstrução: "só predio grandão",
		tamanhodaloucura: "acima de quase tudo",
	}

	serhumano(mrdente)
	fmt.Println("")
	serhumano(mrpredio)

}
