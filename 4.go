package marcelo

import "fmt"

func quatro(a float64, b string, c int) (float64, error) {
	if c == 0 {
		return 0, fmt.Errorf("imposto não encontrado")
	} else if a <= 0 {
		return 0, fmt.Errorf("preço base inválido")
	}
	imposto := 0.0
	if c == 1 {
		imposto = 0.05
	} else if c == 2 {
		imposto = 0.1
	} else if c == 3 {
		imposto = 0.15
	}
	frete := 0.0
	if b == "SP" {
		frete = 0.1
	} else if b == "RJ" {
		frete = 0.15
	} else if b == "MG" {
		frete = 0.2
	} else if b == "ES" {
		frete = 0.25
	} else {
		frete = 0.3
	}
	precofinal := a + a*imposto + a*frete
	return precofinal, nil

}