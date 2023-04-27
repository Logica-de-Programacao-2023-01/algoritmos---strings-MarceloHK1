package marcelo

import "fmt"

func tres(a []int) (int, int, float64, error) {
	if len(a) == 0 {
		return 0, 0, 0.0, fmt.Errorf("lista vazia")
	}
	maior := a[0]
	menor := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > maior {
			maior = a[i]
		} else if a[i] < menor {
			menor = a[i]
		}
	}
	soma := 0
	for _, j := range a {
		soma += j
	}
	media := float64(soma-maior-menor) / float64(len(a)-2)
	return maior, menor, media, nil
}
