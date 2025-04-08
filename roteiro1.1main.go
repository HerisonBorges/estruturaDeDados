package main

import "fmt"

func SomaAte(n int) int{
	var soma int
	for i := 1; i <= n; i++{
		soma += i
	}
	return soma
	
}
func main(){
	var n int
	fmt.Println("Digite um número:" )
	fmt.Scan(&n)
	fmt.Println(SomaAte(n))
}