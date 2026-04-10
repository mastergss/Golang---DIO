//Panic e Recover são funções em Go que permitem lidar com erros de tempo de execução e interromper a execução normal de um programa. Panic é usado para gerar um erro de tempo de execução, enquanto Recover é usado para recuperar do erro e continuar a execução do programa.

package main

import "fmt"

func main() {
	/*
	panic("Pânico - Erro tempo de execução")
	fmt.Println("Este código não será executado") //este código não será executado devido ao panic
	x := recover()
	fmt.Println("Valor recuperado:", x) */ //o código acima gera um panic e tenta recuperar o valor, mas como o panic já interrompeu a execução, o recover não é chamado

	//o código abaixo mostra como usar defer e recover para lidar com o panic e recuperar o valor do erro
	defer func() {
		x := recover() //a função defer é usada para recuperar do panic, caso ocorra
		fmt.Println("Valor recuperado:", x) //imprime o valor recuperado do panic
	} ()
	//o código abaixo gera um panic, mas como a função defer está presente, o recover é chamado e o valor do panic é recuperado
	panic("Pânico - Erro tempo de execução")
}