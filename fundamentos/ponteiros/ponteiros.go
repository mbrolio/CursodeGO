package main

import "fmt"

func main() {
	i := 1

	//ponteiro é uma referencia de memoria
	var p *int = nil

	p = &i //pegando o endereco da variavel e atribuir ao ponteiro
	*p++   // desreferenciando (pegando o valor)
	i++

	//Go nao tem aritimética de ponteiros
	// p++ Nao pode fazer
	fmt.Println(p, *p, i, &i)
}
