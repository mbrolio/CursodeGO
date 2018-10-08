package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  //slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	//slice nao é um array. Ele é um pedaço de um array.
	s2 := a2[1:3] //exemplo é o indice 1 até o indice 3 nao incluindo o indice 3
	fmt.Println(a2, s2)

	s3 := a2[:2] //nao colocar nada é definido o inicio
	fmt.Println(a2, s3)

	s4 := s2[:1]
	fmt.Println(s2, s4) //exemplo de um slice de um slice
}
