package main

import (
	"fmt"
)

func imprimiresultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimiresultado(7.3)
	imprimiresultado(5.1)
}
