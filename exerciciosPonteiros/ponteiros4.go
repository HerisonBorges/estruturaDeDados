package main

import "fmt"

func somarSubtrair(a *int, b *int){

	*a = *a + *b
	*b = *a - *b 
}

func main(){

	a := 10
	b := 5

	fmt.Println("Antes da soma e subtração", a, b)
	somarSubtrair(&a, &b)
	fmt.Println("Depois da soma e subtração", a, b)
}