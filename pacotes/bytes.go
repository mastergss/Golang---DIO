//O pacote bytes implementa funções para manipulação de fatias de bytes. É análogo às funcionalidades do pacote strings. Ele fornece funções para criar, comparar, dividir, juntar e modificar fatias de bytes. O pacote bytes é útil para trabalhar com dados binários ou texto em formato de bytes, oferecendo uma maneira eficiente de manipular esses dados em Go.

//Exemplo de uso do pacote bytes:

package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Criando uma fatia de bytes a partir de uma string
	data := []byte("Hello, World!")
	fmt.Println("Data:", data)
	// Criando um buffer de bytes
	var buffer bytes.Buffer
	// Escrevendo dados no buffer
	buffer.Write(data)
	// Convertendo o buffer de volta para uma string
	result := buffer.String()
	fmt.Println("Result:", result)
	// Comparando duas fatias de bytes
	compare := bytes.Equal([]byte("Hello, World!"), []byte(result))
	fmt.Println("Are they equal?", compare)
}