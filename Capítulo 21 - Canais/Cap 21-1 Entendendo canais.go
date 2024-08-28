/*
CANAIS
- Canais são o Jeito Certo® de fazer sincronização e código concorrente.
- Eles nos permitem trasmitir valores entre goroutines.
- Servem pra coordenar, sincronizar, orquestrar, e buffering.
- Na prática:
    - make(chan type, b)
- Canais bloqueiam:
    - Eles são como corredores em uma corrida de revezamento
    - Eles tem que "passar o bastão" de maneira sincronizada
    - Se um corredor tentar passar o bastão pro próximo, mas o próximo corredor não estiver lá...
    - Ou se um corredor ficar esperando receber o bastão, mas ninguem entregar...
    - ...não dá certo.
- Exemplos:
    - Poe um valor num canal e faz um print. Block.
        - Código acima com goroutine.
        - Ou com buffer. Via de regra: má idéia; é legal em certas situações, mas em geral é melhor sempre passar o bastão de maneira sincronizada.
- Interessante: ref/spec → types
- Código:
    - Block: https://go.dev/play/p/TRO4w9Q-ceZ (não roda!)
    - Go routine: https://go.dev/play/p/95RMxyDHVKz
    - Buffer: https://go.dev/play/p/5KVzCb9GTHH
    - Buffer block: https://go.dev/play/p/Mo5hEWSvRxZ
    - Mais buffer: https://go.dev/play/p/w4vmIhw4zF4
*/

package main

import (
	"fmt"
)

func main() {
	canal := make(chan int, 2)
	canal <- 42
	canal <- 43
	fmt.Println(<-canal)
	fmt.Println(<-canal)
}
