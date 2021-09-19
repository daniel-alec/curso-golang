package matematica

import (
	"fmt"
	"strconv"
)

// Média é responsável por calcular o que vc já sabe... :)
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArrendondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArrendondada
}
