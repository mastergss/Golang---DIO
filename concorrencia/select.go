//Select é um comando utilizado para realizar operações de leitura em múltiplos canais. Ele permite que um programa aguarde por múltiplas operações de leitura ou escrita em canais, e execute uma ação quando uma delas estiver pronta.
//Select funciona como um switch, mas para canais. Ele bloqueia a execução até que uma das operações de leitura ou escrita esteja pronta, e então executa o caso correspondente.
//Combinar goroutines e canais com select é uma maneira poderosa de criar programas concorrentes e eficientes em Go.
//No exemplo abaixo, temos duas goroutines que enviam mensagens para dois canais diferentes. A função "select" é usada para aguardar por mensagens em ambos os canais e imprimir a mensagem recebida.

package main

import (
	"fmt"
	"time"
)

func main() {
	// Criando dois canais de string
	canal1 := make(chan string)
	canal2 := make(chan string)
	// Iniciando a goroutine para enviar mensagens para o canal1
	go func() {
		time.Sleep(1 * time.Second) // Simulando um atraso para mostrar a concorrência
		canal1 <- "mensagem do canal 1"
	}()
	// Iniciando a goroutine para enviar mensagens para o canal2
	go func() {
		// Simulando um atraso para mostrar a concorrência
		time.Sleep(2 * time.Second)
		canal2 <- "mensagem do canal 2"
	}()
	// Utilizando select para aguardar por mensagens em ambos os canais
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-canal1:
			fmt.Println("recebida:", msg1)
		case msg2 := <-canal2:
			fmt.Println("recebida:", msg2)
		}
	}
}