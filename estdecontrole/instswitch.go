// Switch é uma estrutura de controle de fluxo que permite executar diferentes blocos de código com base no valor de uma expressão. Em Go, o switch é mais flexível do que em outras linguagens, pois pode ser usado para comparar valores de diferentes tipos e não se limita apenas a casos constantes.
package main

import "fmt"

func main() {
	//utilizar o for para imprimir os números de 1 a 5
	for i := 1; i <= 5; i++ {
		// Verificar o valor de i e imprimir o número correspondente
		switch i {
			// Caso para cada número de 1 a 5
		case 1:
			fmt.Println("Número: ", i)
		case 2:
			fmt.Println("Número: ", i)
		case 3:
			fmt.Println("Número: ", i)
		case 4:
			fmt.Println("Número: ", i)
		case 5:
			fmt.Println("Número: ", i)
			// Caso para o número 6, que não está no intervalo de 1 a 5 - apenas para demonstrar o uso do switch
		case 6:
			fmt.Println("Número: ", i)
			// Caso padrão para valores que não correspondem a nenhum dos casos anteriores
		default:
			fmt.Println("Número fora do intervalo de 1 a 5")
		}
	}
}