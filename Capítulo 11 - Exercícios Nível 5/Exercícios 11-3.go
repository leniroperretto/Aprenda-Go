/*
Cap. 11 – Exercícios: Nível #5 – 3
- Crie um novo tipo: veículo
    - O tipo subjacente deve ser struct
    - Deve conter os campos: portas, cor
- Crie dois novos tipos: caminhonete e sedan
    - Os tipos subjacentes devem ser struct
    - Ambos devem conter "veículo" como struct embutido
    - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
    - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veículo, caminhonete e sedan:
    - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
    - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
- Demonstre estes valores.
- Demonstre um único campo de cada um dos dois.

Solução: https://go.dev/play/p/6m0tmSQn2pi
*/

package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	traçãoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {

	caminhonete1 := caminhonete{
		veiculo: veiculo{portas: 4,
			cor: "Preta",
		},
		traçãoNasQuatro: true,
	}

	sedan1 := sedan{
		veiculo: veiculo{portas: 4,
			cor: "Azul",
		},
		modeloLuxo: true,
	}

	fmt.Println(caminhonete1, sedan1)
	fmt.Println(caminhonete1.cor)
	fmt.Println(sedan1.modeloLuxo)

}
