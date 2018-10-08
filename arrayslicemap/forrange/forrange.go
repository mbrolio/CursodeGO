package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador que conta os elemento

	for i, numero := range numeros {
		fmt.Printf("%d)  %d\n", i, numero) // o /n é para quebrar a linha
	}

	for _, num := range numeros { // o _ é para ignorar o valor do indice, como nao usamos a variavel i usamos o _
		fmt.Println(num) // se tirar o _, eu conto o indice e nao os valores
	}
}
