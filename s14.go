package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Print("digite uma sequencia numerica: ")
	fmt.Scan(&str)

	pos := -1

	for _, c := range str {
		n, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Print("nao tem numeros")
			return
		}
		if n >= pos && pos != -1 {
			fmt.Print("a sequencia nao é decrescente")
			return
		}
		pos = n
	}
	fmt.Print("a sequencia é decrescente")
}
