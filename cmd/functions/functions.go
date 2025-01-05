package functions

import "fmt"

// DemonstrateFunctions exemplifica o uso de diferentes tipos de funções em Go:
// - Função anônima: uma função sem nome é definida e executada imediatamente.
// - Função nomeada: uma função tradicional com nome que aceita parâmetros e retorna um valor.
// - Função com retorno nomeado: permite especificar nomes para valores de retorno.
// - Função com retorno múltiplo: demonstra como retornar múltiplos valores de uma função.
// - Função variádica: aceita um número variável de argumentos.
// - Função dentro de função: uma função que retorna outra função.

func DemonstrateFunctions() {
	// Função anônima
	func() {
		fmt.Println("Função anônima")
	}()

	// Função nomeada
	fmt.Println(funcNamed(2, 3))

	// Função com retorno nomeado
	fmt.Println(funcNamedReturn(2, 3))

	// Função com retorno múltiplo
	fmt.Println(multipleReturn())

	// Função variádica
	fmt.Println("Soma: ", variadicFunction(1, 2, 3, 4, 5))

	// Função dentro de função
	fmt.Println("Resposta da OuterFunction", outerFunction()())
}

// Função nomeada
func funcNamed(firstInput, secondInput int) int {
	return firstInput * secondInput
}

// Função com retorno nomeado
func funcNamedReturn(firstInput, secondInput int) (result int) {
	result = firstInput * secondInput
	return
}

// Função com múltiplos retornos
func multipleReturn() (int, int) {
	return 1, 2
}

// variadicFunction aceita um número variável de argumentos
func variadicFunction(data ...int) int {
	// sum values
	sum := 0
	for _, value := range data {
		sum += value
	}

	return sum
}

// fução dentro de função
func outerFunction() func() int {
	x := 10

	innerFunction := func() int {
		return x * x
	}

	return innerFunction
}
