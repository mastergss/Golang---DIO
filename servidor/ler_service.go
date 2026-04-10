//Exemplo de um cliente simples em Go usando o pacote "net/http" e "bufio" para ler a resposta do servidor. O cliente faz uma requisição GET para a URL "http://localhost:8090/hello" e imprime a resposta no console.

package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// Fazendo uma requisição GET para a URL "http://localhost:8090/hello"
	// A função "Get" é usada para enviar uma requisição HTTP GET para o servidor e obter a resposta
	// A variável "resp" armazena a resposta do servidor, e a variável "err" armazena qualquer erro que possa ocorrer durante a requisição
	// O endereço "http://localhost:8090/hello" é o endpoint do servidor que o cliente está tentando acessar para obter a resposta
	resp, err := http.Get("http://localhost:8090/hello")
	// Verificando se houve um erro ao fazer a requisição
	if err != nil {
		// Se houver um erro, ele será impresso no console e a função "main" será encerrada
		fmt.Println("Erro ao fazer requisição:", err)
		return
	}
	// A função "defer" é usada para garantir que o corpo da resposta seja fechado após a leitura, mesmo que ocorra um erro durante a leitura
	defer resp.Body.Close()

	// Lendo a resposta do servidor
	// A função "Scanner" do pacote "bufio" é usada para ler a resposta linha por linha
	scanner := bufio.NewScanner(resp.Body)
	// O loop "for" é usado para iterar sobre as linhas da resposta, onde "scanner.Scan()" lê a próxima linha e retorna "true" se houver mais linhas para ler
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}