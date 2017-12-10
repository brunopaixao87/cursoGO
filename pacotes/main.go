package main

import (
	"fmt"
	"github.com/brunopaixao87/cursoGO/pacotes/prefixo"
	"github.com/brunopaixao87/cursoGO/pacotes/operadora"
)

//NomeDoUsuario Logado na aplicação
var NomeDoUsuario = "Bruno Paixão"

func main() {
	fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)
}
