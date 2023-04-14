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
	str1 := scanner.Text()

	fmt.Print("digite outra string: ")
	scanner.Scan()
	str2 := scanner.Text()

	if strings.Contains(str1, str2) == true {
		fmt.Print("a string 1 contem a string 2")
	} else {
		fmt.Print("a string 1 nao contem a string 2")
	}
}
