package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var letraa, letrab string

	fmt.Print("digite uma string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	fmt.Print("digite a letra a ser substituida: ")
	fmt.Scan(&letraa)
	fmt.Print("digite a letra a ser colocada: ")
	fmt.Scan(&letrab)

	novo := strings.ReplaceAll(s, letraa, letrab)
	fmt.Print(novo)
}
