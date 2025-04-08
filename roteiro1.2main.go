package main
import "fmt"

func ClassificarNota(nota int) string{
	switch {
		
	case nota >= 9:
		return "Exelente"
	case nota >= 7:
		return "Bom"
	case nota >= 5:
		return "Regular"
	case nota >= 3: 
		return "Insuficiente"
	default:
		return "Muito Ruim"
		
	}
}

func main(){

	var n int
	fmt.Println("Digite uma nota:" )
	fmt.Scan(&n)
	fmt.Println(ClassificarNota(n))
}