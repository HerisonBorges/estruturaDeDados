package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
	prox  *Pessoa // ponteiro para o próximo nó
}

func main() {
	// criando os nós
	p1 := &Pessoa{"Herison", 33, nil}
	p2 := &Pessoa{"Thoor", 25, nil}
	p3 := &Pessoa{"Adriana", 57, nil}

	// encadeando os nós
	p1.prox = p2
	p2.prox = p3

	// Exibindo a lista
	imprimirLista(p1)
}

// função para percorrer e imprimir a lista encadeada
func imprimirLista(inicio *Pessoa) {
	atual := inicio
	for atual != nil {
		fmt.Printf("Nome: %s, Idade: %d\n", atual.nome, atual.idade)
		atual = atual.prox
	}
}
