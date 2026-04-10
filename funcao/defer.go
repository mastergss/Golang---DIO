//Defer é uma instrução em Go que permite adiar a execução de uma função até que a função que contém o defer termine sua execução. Isso é útil para liberar recursos, fechar arquivos, entre outros. Exemplo: defer fmt.Println("Fim do programa") irá imprimir "Fim do programa" somente após a função main terminar sua execução.

package main

import "fmt"

func dia1() {
	fmt.Println("Dia 1: Aprendendo Go")
}

func dia2() {
	fmt.Println("Dia 2: Aprendendo Go")
}


func main() {
	//a instrução defer adia a execução da função fmt.Println até o final da função main, ou seja, "Fim do programa" será impresso somente após a execução de todas as outras instruções na função main.
	defer fmt.Println("Fim do programa")	
	fmt.Println("Início do programa")	//imprime "Início do programa" imediatamente

	defer dia1()	//a instrução defer adia a execução da função dia1 até o final da função main
	dia2()	//imprime "Dia 2: Aprendendo Go" imediatamente
}