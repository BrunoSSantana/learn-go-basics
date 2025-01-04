// Pacote principal e importações
package main

import (
	"fmt"
)

// Constantes - a nível de pacote, são acessíveis por qualquer função dentro deste pacote
const PackageLevelScope = "Package Level Scope"

func main() {
	// Variáveis e operadores
	variablesAndOperators()
	fmt.Println("---")

	// Demonstra o uso de ponteiros
	demonstratePointer()
	fmt.Println("---")

	// Condicionais
	conditionalsExample()
	fmt.Println("---")

	// Loops
	loopsExample()
	fmt.Println("---")

	// Switch
	switchExample()
	fmt.Println("---")

	// Funções
	demonstrateFunctions()
	fmt.Println("---")
}

func variablesAndOperators() {
	// Declaração de variável usando `var`. Aqui, o tipo é especificado porque nenhum valor inicial foi atribuído.
	var oldValue string

	// Operador de atribuição curta ":=" declara e atribui valor ao mesmo tempo.
	// Só pode ser usado dentro de blocos de código, nunca no nível de pacote.
	shortDeclaration := "Olá, Mundo!"

	fmt.Printf("O valor inicial é %v e o tipo é %T\n", shortDeclaration, oldValue)

	// Operador de atribuição "=" reatribui valores a variáveis já declaradas.
	shortDeclaration = "Novo Valor"

	// Chama a função `show` com diferentes tipos de dados.
	show(42)       // Inteiro
	show(oldValue) // String vazia
	oldValue = "Atualizado"
	show(oldValue) // String com valor atualizado
}

// `show` aceita qualquer tipo de dado e imprime no console.
func show(data interface{}) {
	fmt.Println("Valor recebido:", data)
}

// `demonstratePointer` ilustra o conceito de ponteiros no Go.
func demonstratePointer() {
	// Declaração de variável normal e ponteiro
	data := 10
	pointer := &data // Operador "&" pega o endereço de memória da variável `data`

	fmt.Println("Valor inicial de data:", data)
	fmt.Println("Endereço de memória de data:", pointer)
	fmt.Println("Valor no endereço de memória:", *pointer) // Operador "*" acessa o valor no endereço de memória

	// Alterando o valor usando o ponteiro
	*pointer = 20
	fmt.Println("Novo valor de data:", data)

	// Criando um segundo ponteiro para a mesma variável
	var secondPointer *int = &data
	fmt.Println("Endereço de memória através de secondPointer:", secondPointer)
	fmt.Println("Valor acessado através de secondPointer:", *secondPointer)
}

// `updateValueUsingPointer` altera o valor de um inteiro usando um ponteiro.
func updateValueUsingPointer(data *int) int {
	*data = 20 // Atualiza o valor armazenado no endereço de memória
	return *data
}

func conditionalsExample() {
	fmt.Println("Exemplo de condicionais:")
	conditional := true

	if data := 10; conditional {
		fmt.Println(data)
	}
}

func loopsExample() {
	// Loop for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Loop for com range
	data := []int{1, 2, 3, 4, 5}
	for index, value := range data {
		fmt.Printf("Índice: %v, Valor: %v\n", index, value)
	}

	// Loop infinito
	infiniteLoopExample()

	// Loop for while like
	whileLikeLoopExample()
}

func infiniteLoopExample() {
	for {
		fmt.Println("Loop infinito")
		break
	}
}

func whileLikeLoopExample() {
	fmt.Println("Loop for while like")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func switchExample() {
	// Exemplo de switch
	switch data := 10; data {
	case 1:
		fmt.Println("Caso 1")
	case 2:
		fmt.Println("Caso 2")
	default:
		fmt.Println("Caso padrão")
	}
}

func demonstrateFunctions() {
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
