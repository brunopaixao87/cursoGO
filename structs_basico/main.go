package main

import "fmt"

// Imovel é uma struct que respresenta um imóvel
type Imovel struct {
	X int
	Y int
	Nome string
	valor int
}

func main()  {

	casa := Imovel{}
	fmt.Printf("A casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu AP", 20000}
	fmt.Printf("O apartamento é: %+v\r\n", apartamento)

	chacara := Imovel{
		Y: 85,
		Nome: "Minha chácara",
		X: 50,
		valor:500000,
	}
    fmt.Printf("A chácara é: %+v\r\n", chacara)

	casa.X = 70
	casa.Y = 90
	casa.Nome = "Minha casa"
	casa.valor = 230000
	fmt.Printf("A casa é: %+v\r\n", casa)

}
