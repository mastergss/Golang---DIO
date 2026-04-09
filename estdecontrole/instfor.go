// For é a única estrutura de controle de repetição em Go, mas é muito flexível e pode ser usada para criar loops de diferentes tipos, como loops while e loops infinitos. O exemplo abaixo mostra um loop for que imprime os números de 1 a 10 e verifica se cada número é par ou ímpar.
package main

import "fmt"

func main() {
	// Loop for para imprimir os números de 1 a 10
	for i := 1; i <= 10; i++ {
		// Verifica se o número é par ou ímpar
		if i%2 == 0 {
			fmt.Println("Par: ", i)
		} else {
			fmt.Println("Ímpar: ", i)
		}
	}
}