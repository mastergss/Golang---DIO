//Ler JSON: "usuarios.json"

package main

import (
	"fmt"
	"os"
)

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