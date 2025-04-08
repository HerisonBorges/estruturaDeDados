package main

import "fmt"

func increment (val int){
	val++
	fmt.Println("Dentro da função increment:", val)
}

func main(){
x := 9
increment(x)
fmt.Println("fora da função increment:", x)
}