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
		switch i {
		case 'a', 'e', 'i', 'o', 'u':
			s2 += "*"
		default:
			s2 += string(i)
		}
	}
	fmt.Print(s2)
}
