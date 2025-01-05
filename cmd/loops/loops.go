package loops

import "fmt"

// DemonstrateLoops - Mostra como funciona o loop for em Go.
//
// Demonstra 4 formas de se usar o loop for em Go:
// 1. Loop for comum: `for i := 0; i < 5; i++`.
// 2. Loop for com range: `for index, value := range data`.
// 3. Loop infinito: `for { ... }`.
// 4. Loop for while like: `for i < 5 { ... i++ }`.
func DemonstrateLoops() {
	// Loop for
	demonstrateSampleLoop()

	// Loop for com range
	demonstrateRangeLoop()

	// Loop infinito
	demonstrateInfiniteLoop()

	// Loop for while like
	demonstrateWhileLikeLoop()
}

// demonstrateSampleLoop - Exibe o loop for comum, que é
// responsável por executar um bloco de código
// enquanto uma condição for verdadeira.
func demonstrateSampleLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

// demonstrateRangeLoop - Exibe o loop for com range, que é
// responsável por executar um bloco de código para cada
// elemento de uma slice, acessando o índice e valor de cada
// elemento.
func demonstrateRangeLoop() {
	data := []int{1, 2, 3, 4, 5}
	for index, value := range data {
		fmt.Printf("Índice: %v, Valor: %v\n", index, value)
	}
}

// demonstrateInfiniteLoop - Exibe um loop infinito que imprime uma mensagem no console.
// O loop é interrompido imediatamente após a primeira execução usando a instrução "break".

func demonstrateInfiniteLoop() {
	for {
		fmt.Println("Loop infinito")
		break
	}
}

// demonstrateWhileLikeLoop - Exibe o loop for estilo while, que é
// responsável por executar um bloco de código enquanto uma
// condição for verdadeira.
func demonstrateWhileLikeLoop() {
	fmt.Println("Loop for while like")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}
