// Desafio: Ping Pong
// Descrição: baseado nas aulas de concorrência, crie um programa que simule um jogo de ping pong entre duas goroutines. Uma goroutine deve representar o jogador "Ping" e a outra o jogador "Pong". As goroutines devem enviar mensagens alternadamente para um canal, indicando quem está jogando. O programa deve imprimir as mensagens no console, mostrando o andamento do jogo.

package main

import (
	"fmt"
	"time"
)
// A função "ping" envia a string "Ping" para o canal "c" em um loop infinito
func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "Ping"
	}
}
// A função "pong" envia a string "Pong" para o canal "c" em um loop infinito
func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "Pong"
	}
}
// A função "imprimir" recebe mensagens do canal "c" e as imprime
func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Criando um canal de string chamado "c"
	c := make(chan string)
	// Iniciando as goroutines "ping", "pong" e "imprimir"
	go ping(c)
	go pong(c)
	go imprimir(c)
	// A função "Scanln" é usada para ler a entrada do usuário e armazená-la na variável "entrada"
	var entrada string
	// A função "Scanln" é usada para ler a entrada do usuário e armazená-la na variável "entrada"
	// Neste caso, a função "Scanln" é usada para manter o programa em execução, permitindo que as goroutines continuem enviando mensagens para o canal "c" e imprimindo-as no console. O programa só será encerrado quando o usuário digitar algo e pressionar Enter.
	fmt.Scanln(&entrada)
}