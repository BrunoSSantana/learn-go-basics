package utils

import "fmt"

// `show` aceita qualquer tipo de dado e imprime no console.
func Show(data interface{}) {
	fmt.Println("Valor recebido:", data)
}

// `updateValueUsingPointer` altera o valor de um inteiro usando um ponteiro.
func UpdateValueUsingPointer(data *int) int {
	*data = 20 // Atualiza o valor armazenado no endereço de memória
	return *data
}
