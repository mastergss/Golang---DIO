//Servidor Client o que é?
//O cliente é a parte do sistema que faz requisições para o servidor. Ele pode ser um aplicativo, um navegador ou qualquer outro tipo de software que se comunica com o servidor para obter dados ou realizar ações.
//No contexto de um servidor web, o cliente geralmente é um navegador que faz requisições HTTP (Hypertext Transfer Protocol) para o servidor para obter páginas da web, imagens, vídeos ou outros recursos. O cliente pode enviar dados para o servidor por meio de formulários ou APIs, e o servidor processa essas requisições e retorna as respostas apropriadas.
//Exemplo de um cliente simples em Go usando o pacote "net/http" e "bufio" para ler a resposta do servidor. O cliente faz uma requisição GET para a URL "https://gobyexample.com" e imprime a resposta no console.

package main

import (
    "bufio"
    "fmt"
    "net/http"
)

func main() {
	// Fazendo uma requisição GET para a URL "https://gobyexample.com"
	// A função "Get" é usada para enviar uma requisição HTTP GET para o servidor e obter a resposta
	// A variável "resp" armazena a resposta do servidor, e a variável "err" armazena qualquer erro que possa ocorrer durante a requisição
	// O endereço "https://gobyexample.com" é o endpoint do servidor que o cliente está tentando acessar para obter a resposta
    resp, err := http.Get("https://gobyexample.com")
	// Verificando se houve um erro ao fazer a requisição
    if err != nil {
		// Se houver um erro, ele será impresso no console e a função "main" será encerrada
        panic(err)
    }
	// A função "defer" é usada para garantir que o corpo da resposta seja fechado após a leitura, mesmo que ocorra um erro durante a leitura
    defer resp.Body.Close()

	// Imprimindo o status da resposta do servidor
    fmt.Println("Response status:", resp.Status)

	// Lendo a resposta do servidor
	// A função "Scanner" do pacote "bufio" é usada para ler a resposta linha por linha
    scanner := bufio.NewScanner(resp.Body)
	// O loop "for" é usado para iterar sobre as linhas da resposta, onde "scanner.Scan()" lê a próxima linha e retorna "true" se houver mais linhas para ler
    for i := 0; scanner.Scan() && i < 15; i++ {
		// Imprimindo as primeiras 15 linhas da resposta do servidor
		// A função "Text" do scanner é usada para obter o texto da linha atual
        fmt.Println(scanner.Text())
    }
	// Verificando se houve um erro durante a leitura da resposta
    if err := scanner.Err(); err != nil {
		// Se houver um erro, ele será impresso no console e a função "main" será encerrada
        panic(err)
    }
}