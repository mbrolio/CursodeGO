package main

import (
	"fmt"
)

//func main() {
//	t := time.Now()
//	switch { //switch true
//	case t.Hour() < 12:
//		fmt.Println("Bom dia.")
//	case t.Hour() < 18:
//		fmt.Println("Boa tarde.")
//	default:
//		fmt.Println("Boa noite.")
//	}
//}

func notaTeste(n float64) string {
	var nota = int(n)
	switch nota {
	case 9, 10:
		return "Nota A"
	case 8:
		return "Nota B"
	case 5, 6, 7:
		return "Nota C"
	case 3, 4:
		return "Nota D"
	case 0, 1, 2:
		return "Nota E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaTeste(4.7))
	fmt.Println(notaTeste(2.1))
	fmt.Println(notaTeste(10))
}
