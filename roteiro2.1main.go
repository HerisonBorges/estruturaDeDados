package main

import(
	"fmt"
) 

func main(){

var matriz [3][3] string

for i := 0; i < 3; i++{
	for j := 0; j < 3; j++{
		fmt.Println("Digite o valor", i, j)
		fmt.Scan(&matriz[i][j])
	}
}
for i := 0; i < 3; i++{
	for j := 0; j < 3; j++{
		fmt.Print(matriz[i][j], " ")
	}
	fmt.Println()
}

}


