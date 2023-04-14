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

	possui := false
	for _, i := range s {
		if i >= '0' && i <= '9' {
			possui = true
			break
		}
	}
	if possui {
		fmt.Print("a string possui numero")
	} else {
		fmt.Print("a string nao possui numero")
	}
}
