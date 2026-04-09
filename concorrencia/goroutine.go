//Concorrência é a capacidade de um programa lidar com múltiplas tarefas ao mesmo tempo. Em Go, isso é facilitado por meio de goroutines, que são funções leves que podem ser executadas concorrentemente.
//Goroutines são criadas usando a palavra-chave "go" seguida da chamada da função. Elas permitem que o programa execute várias tarefas simultaneamente, sem bloquear a execução principal.

package main
import (
	"fmt"
	"time"
)

func main() {
	// Iniciando uma goroutine para executar a função "sayHello"
	go sayHello()
	// A função principal continua a execução enquanto a goroutine está rodando
	fmt.Println("Main function is running...")
	// Iniciando outra goroutine para executar a função "concorrentFunction"
	go concorrentFunction(5)
	// Aguardando um pouco para garantir que a goroutine tenha tempo de executar
	time.Sleep(1 * time.Second)

	var escrever string
	fmt.Print("Digite algo: ")
	// Scanln é usado para ler a entrada do usuário e armazená-la na variável "escrever"
	fmt.Scanln(&escrever)
	fmt.Printf("Você digitou: %s\n", escrever)
}

func sayHello() {
	fmt.Println("Hello from the goroutine!")
}

func concorrentFunction(n int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Concurrent function iteration %d\n", n, i)
	}
}