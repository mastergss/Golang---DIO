// If é uma estrutura de controle de fluxo que permite executar um bloco de código apenas se uma condição for verdadeira. Em Go, a sintaxe do if é simples e direta, e pode ser usada para criar condições complexas usando operadores lógicos.
package main

import "fmt"

func main() {
	//utilizar o for para imprimir os números de 1 a 5
	for i := 1; i <= 5; i++ {
		// Verificar o valor de i e imprimir o número correspondente
		if i == 1 {
			fmt.Println("Número: ", i)
		} else if i == 2 {
			fmt.Println("Número: ", i)
		} else if i == 3 {
			fmt.Println("Número: ", i)
		} else if i == 4 {
			fmt.Println("Número: ", i)
		} else if i == 5 {
			fmt.Println("Número: ", i)
		}
	}
}