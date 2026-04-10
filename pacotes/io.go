//O pacote io em Go é uma coleção de interfaces e funções para realizar operações de entrada e saída (I/O).
//função io.WriteString escreve uma string em um writer
//ela recebe dois parâmetros: o writer e a string que queremos escrever
//ex: io.WriteString(w, "Hello, World!")

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Escreve "Hello World" na saída padrão
	// A função io.WriteString retorna o número de bytes escritos e um erro, se houver
	// Se ocorrer um erro ao escrever, o programa será encerrado com log.Fatal
	if _, err := io.WriteString(os.Stdout, "Hello World\n"); err != nil {
		//log.Fatal é usado para registrar um erro e encerrar o programa
		log.Fatal(err)
	}
	// Escreve "Teste package io!" na saída padrão
	_, err := io.WriteString(os.Stdout, "Teste package io!\n")
	// Verifica se houve um erro ao escrever na saída padrão
	if err != nil {
		// Se houver um erro, imprime uma mensagem de erro na saída padrão
		fmt.Println("Erro ao escrever na saída padrão:", err)
	}

}