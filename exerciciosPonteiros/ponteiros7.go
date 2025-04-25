package main

import "fmt"

type Pessoa struct{
	nome string
	idade int
}

func main(){

	p:=Pessoa{"Herison", 33}

	fmt.Println("Antes", p)

	aumentarIdade(&p, 5)
	renomear(&p, "Thoor")

	fmt.Println("Depois", p)

}

func aumentarIdade(p *Pessoa, numero int){

	p.idade += numero
}

func renomear(p *Pessoa, novoNome string){

	p.nome = novoNome
}