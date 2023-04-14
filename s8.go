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
	s := scanner.Text()

	slice := []string{}
	for _, i := range s {
		slice = append(slice, string(i))
	}

	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	resultado := ""
	for _, i := range slice {
		resultado += i
	}
	fmt.Print(resultado)
}
