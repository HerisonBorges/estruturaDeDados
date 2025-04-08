package main

import(
	"fmt"
) 

func main(){
	var arr [6] int
arr [0] = 1
arr [1] = 2
arr [2] = 4
arr [3] = 1
arr [4] = 4
arr [5] = 9


fmt.Println(arr [0])
fmt.Println(arr [1])
fmt.Println(arr [2])
fmt.Println(arr [3])
fmt.Println(arr [4])
fmt.Println(arr [5])


for i := 0; i < 6; i++{
	fmt.Println(arr[i])	
}


for _ , v := range arr{
	fmt.Println(v)
}
}