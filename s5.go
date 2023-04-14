package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Print("digite uma string: ")
	fmt.Scan(&s)

	_, err := strconv.ParseFloat(s, 64)
	if err == nil {
		fmt.Print("a string é um numero valido em ponto flutuante")
	} else {
		fmt.Print("a string nao é um numero valido em ponto flutuante")
	}
}
