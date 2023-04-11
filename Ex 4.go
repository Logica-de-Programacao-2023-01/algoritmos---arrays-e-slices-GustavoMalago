package main

import "fmt"

func main() {
	var tamanho int
	fmt.Println(" Informe o tamanho da lista: ")
	fmt.Scan(&tamanho)
	slice := make([]int, tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Println("informe os numeros da lista: ")
		fmt.Scan(&slice[i])
	}
	fmt.Print(slice)
	soma := 0
	for i := 0; i < len(slice); i++ {
		soma += slice[i]
	}
	fmt.Println("A soma dos números da lista é: ", soma)
}
