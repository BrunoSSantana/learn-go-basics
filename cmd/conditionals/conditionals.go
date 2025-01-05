package conditionals

import "fmt"

// ConditionalsExample - Exibe um exemplo de como condicionais funcionam no Go.
//
// A variável "conditional" é usada como condição para o bloco de código.
// A variável "data" é criada e atribuída o valor 10, e imprimida no console.
// O bloco de código dentro do "if" é executado somente se "conditional" for verdadeiro.
func ConditionalsExample() {
	fmt.Println("Exemplo de condicionais:")
	conditional := true

	if data := 10; conditional {
		fmt.Println(data)
	}
}

// DemonstrateSwitch - Exibe um exemplo de como o switch funciona no Go.
//
// O switch tem uma declaração curta para a variável "data" e faz uma
// compara o com os valores 1 e 2. Caso o valor seja diferente de 1 e 2,
// o bloco de código dentro do "default"   executado.
func DemonstrateSwitch() {
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
