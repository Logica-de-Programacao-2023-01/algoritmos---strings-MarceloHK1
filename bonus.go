package marcelo

import "fmt"

func bonus(a, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return 0, fmt.Errorf("nível inválido")
	}
	dano := a
	if a > 100 {
		dano = a * 20
	} else if b > 100 {
		dano = a * 2
	}
	if a > b {
		dano *= 10
	} else if a < b {
		dano *= 5
	} else if a == b {
		dano *= 7
	}
	if a-b > 20 ||b-a>20{
		dano*=5
	}else if a-b<20 || b-a<20{
		dano*=2
	}
	return dano,nil
}
