// Criar servidor estático para acessar "./estatico.html"
// Utilizar os pacotes "log" e "net/http"
// Servir arquivos ".html" de um local específico


package main

import (
	"log"
	"net/http"
)

func main() {
	//FileServer retorna um handler que atende a solicitações HTTP com o conteúdo do sistema de arquivos enraizado na raiz.
	//Para usar a implementação do sistema de arquivos do sistema operacional, use "http.Dir"
	fs := http.FileServer(http.Dir("./estatico"))
	http.Handle("/", fs)
	log.Print("Listening on :3000...")
	//ListenAndServe inicia um servidor HTTP com um endereço e handler específicos.
	//Determina a porta de entrada "localhost:3000"
	err := http.ListenAndServe(":3000", nil)
	//Em caso de erro vamos chamar a função "Fatal"
	if err != nil {
		//Função "Fatal" chama "os.Exit" após escrever a mensagem do log
		log.Fatal(err)
	}
}