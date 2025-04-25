package main

import "fmt"


func main()  {
	numero := 10

	fmt.Println("Antes da modificação", numero)

	modificarValor(&numero)

	fmt.Println("depois da modificação", numero)


}

func modificarValor(numero *int){

	*numero = 100
}