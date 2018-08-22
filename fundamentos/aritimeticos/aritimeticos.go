package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma = ", a+b)
	fmt.Println("Subtração = ", a-b)
	fmt.Println("Multiplicação = ", a*b)
	fmt.Println("Divisão = ", a/b)

	// bitwise
	fmt.Println("E =", a&b)   // 11 & 10 = 10
	fmt.Println("Ou = ", a|b) // 11 | 10 = 11

	//Outras usando math

	c := 3.0
	d := 2.0

	fmt.Println("Maior = ", math.Max(float64(c), float64(d)))
	fmt.Println("Menor = ", math.Min(c, d))
}
