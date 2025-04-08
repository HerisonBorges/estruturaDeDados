package main

import(
	"fmt"
)
type produtos struct{
	nome string
	valor float64
	quantidade int
	
}

func main(){

	p := produtos{
		nome: "feijão",
		valor: 5.99,
		quantidade: 10,	
	}
	h := produtos{
		nome: "arroz",
		valor: 4.99,
		quantidade: 10,	
	}
	
	fmt.Println(p.nome,p.valor,p.quantidade)
	fmt.Println(h.nome,h.valor,h.quantidade)
}