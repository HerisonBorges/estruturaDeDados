package main

import (
	"fmt"
)

// Definição da pilha
type Stack struct {
	items []rune
}

// Push adiciona um caractere na pilha
func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

// Pop remove e retorna o caractere do topo da pilha
func (s *Stack) Pop() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

// Função para verificar se uma palavra é um palíndromo usando pilha
func isPalindrome(word string) bool {
	stack := Stack{}
	// Empilhar os caracteres da palavra
	for _, char := range word {
		stack.Push(char)
	}

	// Comparar com os caracteres da palavra original
	for _, char := range word {
		top, _ := stack.Pop()
		if top != char {
			return false
		}
	}

	return true
}

func main() {
	// Testes
	words := []string{"radar", "arara", "golang", "level", "hello"}
	for _, word := range words {
		if isPalindrome(word) {
			fmt.Printf("\"%s\" é um palíndromo\n", word)
		} else {
			fmt.Printf("\"%s\" NÃO é um palíndromo\n", word)
		}
	}
}