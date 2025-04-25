package main

import "fmt"

func main (){

	a := 10
	b := 20

	fmt.Println("Antes da troca", a, b)
	trocar(&a,&b)
	fmt.Println("Depois da troca", a, b)
}

func trocar(a *int, b *int){

	temp := *a
	*a = *b
	*b = temp
}