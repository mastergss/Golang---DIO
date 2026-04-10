//Função é um bloco de código que pode ser reutilizado, ou seja, você pode chamar a função quantas vezes quiser, e ela irá executar o mesmo código.
//Função é uma forma de organizar o código, tornando-o mais legível e fácil de entender.
//também chamada de procedimento ou sub-rotina.
//recebe um nome, parâmetros de entrada (opcional) e um corpo de código que é executado quando a função é chamada. e pode retornar um valor (opcional).

package main

import "fmt"
//Função para calcular a média de uma lista de notas
func media(notas []float64) float64 {
	total := 0.0//variável para armazenar o total das notas
	//loop para percorrer as notas e somá-las
	for _, valor := range notas { //loop para percorrer as notas e somá-las
		total += valor
	}
	return total/float64(len(notas))//retorna a média das notas
}

func main() {
	//Progama que cálcula a média de notas de uma sala
	notas := []float64{98, 93, 77, 82, 83}//lista de notas
	fmt.Println("Média: ", media(notas))//chama a função media e imprime o resultado
	/*
	total := 0.0//variável para armazenar o total das notas

	for _, valor := range notas { //loop para percorrer as notas e somá-las
		total += valor
		//fmt.Println("Total: ", total)
	}
	fmt.Println("Média: ", total/float64(len(notas)))
	*/
}
