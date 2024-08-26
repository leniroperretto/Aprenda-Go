## Compilação cruzada
- GOOS
- GOARCH
- `GOOS=darwin GOARCH=amd64 go build test.go`
- https://godoc.org/runtime#pkg-constants
- git push
- git clone
- go get

Segue aqui o exemplo do exercício de compilação cruzada.

package main

import ( "fmt"
        "runtime"
        )

func main() {
    fmt.Println("Esse é um programa do exercício de compilação cruzada. Foi compilado em um Windows 11 - 64, e deverá rodar no seu sistema operacional:",
    runtime.GOARCH, runtime.GOOS)

}

## -----------------------------------------------------------------------------------------------------------------------------------------------------------------
Se o arquivo for compilado no Linux, usaremos os comandos abaixo para executar.

-> No terminal, usar esses comandos para usá-lo em WINDOWS
GOOS=windows GOARCH=amd64 go build main.go

-> No terminal, usar esses comandos para usá-lo em MAC
GOOS=darwin GOARCH=amd64 go build main.go

