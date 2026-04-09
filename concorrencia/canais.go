//Canal modo de comunicação e sincronização entre goroutines
//O canal é um tipo de dado que permite a comunicação entre goroutines. Ele é usado para enviar e receber valores entre goroutines, permitindo a sincronização e a coordenação entre elas. Os canais são criados usando a função "make" e podem ser usados para enviar e receber valores de qualquer tipo.
package main

import (
	"fmt"
	"time"
)
// A função "pingar" envia a string "ping" para o canal "c" em um loop infinito
func pingar(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
		// Intervalos de tempo para evitar que a função "pingar" consuma muita CPU
		time.Sleep(1 * time.Second)
	}
}
// A função "imprimir" recebe mensagens do canal "c" e as imprime
func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func main() {
	// Criando um canal de string chamado "c"
	c := make(chan string)
	// Iniciando a goroutine "pingar" para enviar mensagens para o canal "c"
	go pingar(c)
	// Iniciando a goroutine "imprimir" para receber e imprimir mensagens do canal "c"
	go imprimir(c)
	
	var entrada string
	// A função "Scanln" é usada para ler a entrada do usuário e armazená-la na variável "entrada"
	fmt.Scanln(&entrada)
}