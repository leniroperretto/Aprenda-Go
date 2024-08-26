## - $GOPATH/
    bin/
    pkg/
    src/
        github.com/
            [Nome do usuário (github.com)]/
                [Nome do projeto ou repo]/
                [Nome do projeto ou repo]/
                [Nome do projeto ou repo]/
                [Nome do projeto ou repo]/
                ...
                [Nome do projeto ou repo]/
## - GOROOT: onde os binários da instalação do Go foram instalados
    - GOROOT="/usr/lib/go"
## - GOPATH: onde seus arquivos de trabalho, seu workspace, fica localizado
    - GOPATH="/home/ellen/go"
    - export GOPATH=$HOME/go (.bashrc)
## - Package management? go get.
    - Na prática → e.g. gouuid