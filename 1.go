package marcelo

import "fmt"

func um(a float64, b []float64) (float64, error) {
	if a <= 0 {
		return 0, fmt.Errorf("valor da compra inválido")
	}
	disconto := 0.0
	media := 0.0
	valortotal := 0.0
	for _, i := range b {
		valortotal += i + a
		media = valortotal / float64(len(b)+1)
	}
	if media > 1000 {
		disconto = 0.2
	} else if len(b) == 0 {
		disconto = 0.1
	} else if valortotal > 1000 {
		disconto = 0.1
	} else if valortotal <= 500 {
		disconto = 0.02
	} else if valortotal <= 1000 {
		disconto = 0.05
	}
	valorfinal := a * disconto
	return valorfinal, nil
}
