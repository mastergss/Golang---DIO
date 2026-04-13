//Ler JSON: "usuarios.json"

package main

import (
	"fmt"
	"os"
)

//Criar estruturas em JSON para receber os dados de "usuarios.json"

//estrutura "Usuarios" para listar "usuarios.json"
type Usuarios struct {
	Usuarios []Usuario `json: "usuarios"`
}

//estrutura para receber os dados em Json que estão antes de "Social"
type Usuario struct {
	Nome string `json: "Nome"`
	Tipo string `json: "Tipo"`
	Idade int `json: "Idade"`
	Social string `json: "Social"`
}

//estrutura para receber os dados em Json que estão em "Social"
type Social struct {
	Facebook string `json: "Facebook"`
	Twitter  string `json: "Twitter"`
}

func main() {
	//Abre o arquivo nomeado com a flag especificada
	//Se bem-sucedido, o arquivo retornado pode ser usado para leitura
	jsonFile, err := os.Open("usuarios.json")
	//Se ocorrer um erro, imprimir "err"
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Arquivo JSON aberto com sucesso!")
	//Após executar a função principal, fechar o arquivo
	defer jsonFile.Close()
	//fmt.Println("Hello World")
}