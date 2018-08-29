package main

import (
	"fmt"
)

//Não tem operador ternario
func obterResultados(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultados(5.2))
}
