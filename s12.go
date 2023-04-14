package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("digite uma string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	qntd := len(str)
	for i := 0; i < qntd/2; i++ {
		if str[i] != str[qntd-i-1] {
			fmt.Print("a string nao é um palindromo")
			return
		}
	}
	fmt.Print("a string é um palindromo")
}
