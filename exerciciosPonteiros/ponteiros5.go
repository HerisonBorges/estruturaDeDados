package main

import "fmt"

func main(){

	numeros := [5]int{1,2,3,4,5}

	fmt.Println("antes da inversão", numeros)
	inversão(&numeros)
	fmt.Println("depois da inversão", numeros)
}

func inversão(arr *[5]int){
	arr[0], arr[4] = arr[4], arr[0]
	arr[1], arr[3] = arr[3], arr[1]
}