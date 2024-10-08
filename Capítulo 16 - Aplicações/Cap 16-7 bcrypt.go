/*
CRYPT
 É uma maneira de encriptar senhas utilizando hashes.
- x/crypto/bcrypt
    - GenerateFromPassword
    - CompareHashAndPassword
    - go get golang.org/x/crypto/bcrypt

Solução: https://go.dev/play/p/Yv2hhXCkrEg
*/

package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "20julho1980"
	senhaerrada := "20julho1990"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))

	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada)))

}
