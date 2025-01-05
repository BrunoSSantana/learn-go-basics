package variables

import (
	"fmt"

	"github.com/BrunoSSantana/learn-go-basics/cmd/utils"
)

func VariablesAndOperators() {
	// Declaração de variável usando `var`. Aqui, o tipo é especificado porque nenhum valor inicial foi atribuído.
	var oldValue string

	// Operador de atribuição curta ":=" declara e atribui valor ao mesmo tempo.
	// Só pode ser usado dentro de blocos de código, nunca no nível de pacote.
	shortDeclaration := "Olá, Mundo!"

	fmt.Printf("O valor inicial é %v e o tipo é %T\n", shortDeclaration, oldValue)

	// Operador de atribuição "=" reatribui valores a variáveis já declaradas.
	shortDeclaration = "Novo Valor"

	// Chama a função `show` com diferentes tipos de dados.
	utils.Show(42)       // Inteiro
	utils.Show(oldValue) // String vazia
	oldValue = "Atualizado"
	utils.Show(oldValue) // String com valor atualizado
}
