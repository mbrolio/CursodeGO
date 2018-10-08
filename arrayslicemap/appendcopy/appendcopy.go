package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) //copiei os 2 primeiros elementos do slice 1 para o slice2 - o copy nao faz o slice crescer
	fmt.Println(slice2)  //ele só copiou 2 elementos pois criei o slice2 com tamanho 2
}
