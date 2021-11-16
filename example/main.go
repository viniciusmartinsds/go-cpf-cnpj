package main

import (
	"fmt"

	"github.com/viniciusmartinsds/go-cpf-cnpj/cpf"
)

func main() {
	cpf := cpf.Generate()
	fmt.Println(cpf)
}
