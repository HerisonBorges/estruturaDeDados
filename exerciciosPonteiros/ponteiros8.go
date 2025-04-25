package main

import "fmt"


type Pessoa struct{

	nome string
	idade int

}

func main(){

	p:=[]Pessoa{
		{"herison", 33},
		{"thoor", 25},
		{"Adriana", 57},
	}

	fmt.Println("Antes", p)

	alterarIdade(&p[0], 10)

	alterarNome(&p[2], "Cidão")

	exibirPessoa(&p[1])

	fmt.Println("depois", p)
}

func alterarIdade(p *Pessoa, numero int){

	p.idade += numero
}

func alterarNome(p *Pessoa, novoNome string){

	p.nome = novoNome
}

func exibirPessoa(p *Pessoa){

	fmt.Println("nome", p.nome)
	fmt.Println("idade", p.idade)
}