package marcelo

import (
	"fmt"
	"strings"
)

func dois(str string) (float64, error) {
	if len(str) == 0 {
		return 0, fmt.Errorf("texto vazio")
	}
	novo := strings.ReplaceAll(str, ",", " ")
	palavras := 1.0
	letras := 0.0
	for i := 0; i < len(novo); i++ {
		if novo[i] == ' ' {
			palavras++
		} else {
			letras++
		}
	}
	media := float64(letras) / float64(palavras)
	return media, nil
}
