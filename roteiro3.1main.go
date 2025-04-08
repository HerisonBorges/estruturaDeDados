package main

import "fmt"

func main(){
	var b float64	 = 33.6
	var p2 *float64 = &b

	fmt.Println("Valor antes da modificação", b)
	*p2 = 25.0
	fmt.Println("Valor depois da modificação", b)
}