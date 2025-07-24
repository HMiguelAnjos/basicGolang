package types

import (
	"fmt"
)

func ReturnArray() {
	var numeros [5]int
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	fmt.Println(numeros)
}
