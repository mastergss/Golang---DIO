//Exibir números entre 1 e 100 que sejam múltiplos de 3
package main

import "fmt"

func main() {
	// Percorre os números de 1 a 100
	for num := 1; num <= 100; num++ {
		// Verifica se o número é múltiplo de 3
		if num%3 == 0 {
			// Exibe o número no console
			fmt.Print(num, " ")
		}
	}
}