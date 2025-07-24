package types

import (
	"fmt"
)

func ReturnSlice() {
	numeros := []int{10, 20, 30}
	numeros = append(numeros, 30)
	numeros = append(numeros, 40)
	numeros = append(numeros, 50)

	fmt.Println(numeros)
}
