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
	s1 := scanner.Text()

	fmt.Print("digite outra string: ")
	scanner.Scan()
	s2 := scanner.Text()

	if s1 == s2 {
		fmt.Print("as duas strings são iguais")
	} else {
		fmt.Print("as duas strings são diferentes")
	}
}
