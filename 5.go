package marcelo

import "fmt"

func cinco(a float64, b string, c string) (float64, error) {
	validas := "CFK"
	if b != validas || c != validas {
		return 0, fmt.Errorf("escala inválida")
	}
	if b == "C" && c == "F" {
		a = (a * 9 / 5) + 32
	} else if b == "C" && c == "K" {
		a = a + 273.15
	} else if b == "F" && c == "C" {
		a = (a - 32) * 5 / 9
	} else if b == "F" && c == "K" {
		a = (a-32)*5/9 + 273.15
	} else if b == "K" && c == "C" {
		a = a - 273.15
	} else if b == "K" && c == "F" {
		a = (a-273.15)*9/5 + 32
	}
	return a, nil
}
