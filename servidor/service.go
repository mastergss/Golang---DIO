//Servidor Service é um termo genérico que pode se referir a um serviço ou aplicação que é executada em um servidor. Ele pode ser um serviço web, um serviço de banco de dados, um serviço de autenticação, entre outros. O servidor service é responsável por fornecer funcionalidades específicas para os clientes que se conectam a ele, como processar solicitações, armazenar dados, autenticar usuários, etc. Ele pode ser implementado usando diferentes tecnologias e linguagens de programação, dependendo das necessidades do projeto.
// Exemplo de um servidor service simples em Go usando o pacote "net/http"

package main

import (
	"fmt"
	"net/http"
)

// A função "helloHandler" é um manipulador de rota que responde a solicitações HTTP com uma mensagem de saudação
// Ela recebe um "http.ResponseWriter" para escrever a resposta e um "http.Request" para acessar os detalhes da solicitação
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
	fmt.Fprintf(w, "Seja bem-vindo ao servidor service em Go!")
}

func cabecalhoHandler(w http.ResponseWriter, r *http.Request) {
	// A função "Header" é usada para acessar os cabeçalhos da solicitação HTTP
	// O método "Get" é usado para obter o valor do cabeçalho "User-Agent"
	userAgent := r.Header.Get("User-Agent")
	fmt.Fprintf(w, "User-Agent: %s", userAgent)
	// Iterando sobre todos os cabeçalhos da solicitação e imprimindo seus nomes e valores
	// O loop "for" é usado para percorrer os cabeçalhos da solicitação, onde "nome" é o nome do cabeçalho e "cabecalhoHandler" é o valor do cabeçalho
	// A função "Fprintf" é usada para formatar a saída e escrever os nomes e valores dos cabeçalhos na resposta HTTP
	// O resultado será uma lista de todos os cabeçalhos da solicitação, cada um em uma linha, no formato "Nome: Valor"
	for nome, cabecalhoHandler := range r.Header {
		fmt.Fprintf(w, "%s: %s\n", nome, cabecalhoHandler)
	}
}

func main() {
	// Registrando o manipulador de rota para a URL "/hello"
	// Quando um cliente acessar a URL "/hello", a função "helloHandler" será chamada para processar a solicitação
	// A função "HandleFunc" é usada para associar a URL "/hello" ao manipulador de rota "helloHandler"
	http.HandleFunc("/hello", helloHandler)
	// Registrando o manipulador de rota para a URL "/cabecalho"
	// Quando um cliente acessar a URL "/cabecalho", a função "cabecalhoHandler" será chamada para processar a solicitação
	// A função "HandleFunc" é usada para associar a URL "/cabecalho" ao manipulador de rota "cabecalhoHandler"
	http.HandleFunc("/cabecalho", cabecalhoHandler)
	// Iniciando o servidor na porta 8090
	http.ListenAndServe(":8090", nil)
}