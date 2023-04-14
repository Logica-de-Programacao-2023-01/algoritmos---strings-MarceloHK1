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

	fmt.Print("digite outra string: ")
	scanner.Scan()
	s2 := scanner.Text()

	for i := 0; i < len(s); i++ {
		caracterachado := false
		for j := 0; j < len(s2); j++ {
			if s[i] == s2[j] {
				s2 = s2[:j] + s2[j+1:]
				caracterachado = true
				break
			}
		}
		if caracterachado == false {
			fmt.Print("eles nao sao anagramas")
			return
		}
	}
	fmt.Print("eles sao anagramas")
}
