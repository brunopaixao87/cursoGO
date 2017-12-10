package main

import "fmt"

var (
	//Endereco uma váriavel pública exposta fora do pacote, var pública primeira caracter UpperCase
	Endereco string
	//Telefone uma váriavel pública exposta fora do pacote, var pública primeira caracter UpperCase
	Telefone            string
	quantidade, estoque int
	comprou             bool
	valor               float64
	palavras            rune
)

func main() {
	//Criando váriael inferindo seu tipo pelo valor atribuido
	teste := "Valor de teste"
	fmt.Printf("endereco: %s\r\n", Endereco)
	fmt.Printf("telefone: %s\r\n", Telefone)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("valor: %v\r\n", valor)
	fmt.Printf("palavras: %v\r\n", palavras)
	fmt.Printf("O valor de teste é: %v\r\n", teste)
}
