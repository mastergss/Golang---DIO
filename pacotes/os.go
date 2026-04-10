//Pacote os fornece uma plataforma independente de interface para funcionalidades do sistema operacional. Ele inclui funções para manipulação de arquivos, diretórios, processos e variáveis de ambiente. O pacote os é amplamente utilizado para realizar operações relacionadas ao sistema operacional em Go, como leitura e escrita de arquivos, criação de diretórios, execução de comandos e manipulação de variáveis de ambiente.

//Exemplo de uso do pacote os.Create para criar um arquivo e escrever dados nele. O código cria um arquivo chamado "example.txt", escreve a string "Hello, World!" no arquivo e fecha o arquivo após a escrita. Se ocorrer algum erro durante a criação ou escrita do arquivo, ele será impresso no console.:
package main

import (
	"fmt"
	"os"
)

func main() {
	// Criando um novo arquivo chamado "example.txt"
	file, err := os.Create("example.txt")
	// Verificando se houve algum erro na criação do arquivo
	if err != nil {
		// Imprimindo o erro e retornando
		fmt.Println("Error creating file:", err)
		return
	}
	// Garantindo que o arquivo seja fechado após a função main terminar
	defer file.Close()

	// Escrevendo dados no arquivo
	_, err = file.WriteString("Hello, World!")
	// Verificando se houve algum erro na escrita do arquivo
	if err != nil {
		// Imprimindo o erro e retornando
		fmt.Println("Error writing to file:", err)
		return
	}
	// Imprimindo uma mensagem de sucesso
	fmt.Println("File created and written successfully.")
}