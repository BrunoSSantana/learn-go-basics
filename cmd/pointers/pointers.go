package pointers

import (
	"fmt"
)

// `demonstratePointer` ilustra o conceito de ponteiros no Go.
func DemonstratePointer() {
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
