// Pacote error fornece funcionalidades para manipulação de erros em Go. Ele inclui funções para criar e manipular erros, bem como para verificar o tipo de erro e extrair informações úteis dele.
// O pacote error é amplamente utilizado para lidar com erros em Go, permitindo que os desenvolvedores criem mensagens de erro personalizadas e tratem erros de forma eficiente em seus programas.

// Exemplo de uso do pacote error:
package main

import (
	"fmt"
	"time"
)

// MyError é uma estrutura personalizada que implementa a interface de erro do Go. Ela contém um campo When para armazenar a hora do erro e um campo What para armazenar a mensagem de erro.
type MyError struct {
	When time.Time
	What string
}

// O método Error é implementado para MyError, permitindo que ele satisfaça a interface de erro do Go. Ele retorna uma string formatada que inclui a hora e a mensagem de erro.
func (e MyError) Error() string {
	// A função fmt.Sprintf é usada para formatar a string de erro, combinando o valor do campo When e o valor do campo What em uma única string legível.
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

// A função oops retorna um erro do tipo MyError, preenchendo os campos When e What com informações relevantes sobre o erro.
func oops() error {
	// A função time.Date é usada para criar um valor de tempo específico, representando a data e hora do erro. O campo What é preenchido com uma mensagem de erro personalizada.
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"O arquivo não foi encontrado no sistema",
	}
}

func main() {
	// Chama a função oops e verifica se ela retorna um erro. Se um erro for retornado, ele é impresso na saída padrão.
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}