package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("digite uma string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	str = strings.ToLower(str)
	s2 := ""
	for _, i := range str {
		if i != 'a' && i != 'e' && i != 'i' && i != 'o' && i != 'u' {
			s2 += string(i)
		}
	}
	fmt.Print(s2)
}
