package main

import (
	"fmt"
	"github.com/brunopaixao87/cursoGO/funcoes_basico/matematica"
)

func main()  {
	valorResultadoOperacao := matematica.Calcular(matematica.Somar, 10,2)
	fmt.Printf("O resultado da soma é: %d\r\n", valorResultadoOperacao)

	valorResultadoOperacao = matematica.Calcular(matematica.Multiplicar, 10,2)
	fmt.Printf("O resultado da multiplicação é: %d\r\n", valorResultadoOperacao)

	valorResultadoOperacao = matematica.Calcular(matematica.Dividir, 10,2)
	fmt.Printf("O resultado da divisão é: %d\r\n", valorResultadoOperacao)

	resultado, resto := matematica.DividirComResto(9,2)
	fmt.Printf("O resultado da divisão é: %d e o resto é: %d\r\n", resultado, resto)

}
