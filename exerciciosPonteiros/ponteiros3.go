package main

import "fmt"

func dobrar(a *int, b *int){

	*a = *a * 2
	*b = *b * 2

}

func main(){

	a := 10
	b := 20
	fmt.Println("Antes da dobra", a, b)
	dobrar(&a,&b)
	fmt.Println("Depois da dobra", a, b)
}