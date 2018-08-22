package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516
	y := 12345
	xs := fmt.Sprint(x)
	ys := fmt.Sprint(y)
	fmt.Println("O valor de x é: " + xs)
	fmt.Println("O valor de y é: " + ys)

	//outro modo
	fmt.Printf("O valor de x é %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "Opa"

	fmt.Printf("\n%d \n%f \n%.1f \n%t \n%s", a, b, b, c, d)
}
