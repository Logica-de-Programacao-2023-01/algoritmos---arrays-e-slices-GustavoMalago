package main

import "fmt"

func main() {
	var numeros = [3]int{1, 2, 3}
	soma := 0
	for i := 0; i < len(numeros); i++ {
		soma += numeros[i]
	}
	fmt.Print("A soma dos valores é: ", soma)
}
