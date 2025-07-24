package types

import (
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func ReturnStruct() {
	p := Pessoa{
		Nome:  "Henrique",
		Idade: 29,
	}
	fmt.Printf("Nome: %s, Idade: %d\n", p.Nome, p.Idade)
}
