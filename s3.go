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

	fmt.Print("digite o caractere a ser substituido: ")
	scanner.Scan()
	old := scanner.Text()

	fmt.Print("digite o caractere a ser acrescentado: ")
	scanner.Scan()
	novo := scanner.Text()

	s1 := ""
	for _, i := range s {
		if string(i) == old {
			s1 += novo
		} else {
			s1 += string(i)
		}
	}
	fmt.Println(s1)
}
