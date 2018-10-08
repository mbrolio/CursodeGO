package main

import "fmt"

func main() {
	//homogenea (mesmo tipo) e estática (fixo) - exemplo?: um array de inteiro só vai ter inteiro
	//indice começa sempre no 0
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ { //o len() serve para pegar o tamanho do array, assim sabe quantas vezes tem que fazer o for
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)
}
