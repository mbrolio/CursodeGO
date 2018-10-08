package main

import (
	"fmt"
)

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriele Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Jr": 1200.00,
		},
	}
	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
