package main

import "fmt"



func main()	{
	var a int = 42
	var p *int = &a

	fmt.Println("valor de a:", a)
	fmt.Println("endereço de a:", &a)
	fmt.Println("valor de p (endereço de a):", p)
	fmt.Println("valor apontado por p", *p)
}

 