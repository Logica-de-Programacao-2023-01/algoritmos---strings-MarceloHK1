package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Print("digite uma string com numeros: ")
	fmt.Scan(&str)

	ant := -1

	for _, i := range str {
		n, err := strconv.Atoi(string(i))
		if err != nil {
			fmt.Println("nao tem numeros")
			return
		}
		if n <= ant {
			fmt.Println("a sequencia nao é crescente")
			return
		}
		ant = n
	}
	fmt.Print("a sequencia é crescente")
}
