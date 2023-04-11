package main

import "fmt"

func main() {
	array := [4]float64{2.1, 3.7, 9.2, 1.3}
	produto := array[0] * array[1] * array[2] * array[3]
	fmt.Print("O prouduto é igual a: ", produto)
}
