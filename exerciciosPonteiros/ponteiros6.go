package main

import "fmt"

type Pessoa struct{


	nome string
	idade int
}

func main(){

	p := Pessoa{"Herison", 33}

	fmt.Println("Antes da alteração", p)

	alteracao(&p)

	fmt.Println("Apos a alteração", p)

}

func alteracao(p *Pessoa){

	p.idade = 25
	p.nome = "Thoor"
}