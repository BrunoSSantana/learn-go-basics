package main

import (
	"fmt"

	"github.com/BrunoSSantana/learn-go-basics/cmd/conditionals"
	"github.com/BrunoSSantana/learn-go-basics/cmd/functions"
	"github.com/BrunoSSantana/learn-go-basics/cmd/loops"
	"github.com/BrunoSSantana/learn-go-basics/cmd/pointers"
	"github.com/BrunoSSantana/learn-go-basics/cmd/variables"
)

// Constantes - a nível de pacote, são acessíveis por qualquer função dentro deste pacote
const PackageLevelScope = "Package Level Scope"

func main() {
	// Variáveis e operadores
	variables.VariablesAndOperators()
	fmt.Println("---")

	// Demonstra o uso de ponteiros
	pointers.DemonstratePointer()
	fmt.Println("---")

	// Condicionais
	conditionals.ConditionalsExample()
	fmt.Println("---")

	// Switch
	conditionals.DemonstrateSwitch()
	fmt.Println("---")

	// Loops
	loops.DemonstrateLoops()
	fmt.Println("---")

	// Funções
	functions.DemonstrateFunctions()
	fmt.Println("---")
}
