//Ler JSON: "usuarios.json"

package main

import (
	"io"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

//Criar estruturas em JSON para receber os dados de "usuarios.json"

//estrutura "Usuarios" para listar "usuarios.json"
type Usuarios struct {
	Usuarios []Usuario `json:"usuarios"`
}

//estrutura para receber os dados em Json que estão antes de "Social"
type Usuario struct {
	Nome string `json:"Nome"`
	Tipo string `json:"Tipo"`
	Idade int `json:"Idade"`
	Social string `json:"Social"`
}

//estrutura para receber os dados em Json que estão em "Social" por ser um objeto dentro de outro objeto
type Social struct {
	Facebook string `json:"Facebook"`
	Twitter  string `json:"Twitter"`
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
	
	//Ler o conteúdo do arquivo JSON e armazenar os dados numa variável
	byteValue, _ := io.ReadAll(jsonFile)

	//Criar uma variável do tipo "Usuarios" para armazenar os dados do arquivo JSON
	var usuarios Usuarios
	
	//Desempacotar os dados JSON
	//Unmarshal é a função que converte os dados JSON em uma estrutura GO
	json.Unmarshal(byteValue, &usuarios)

	//Imprimir todos os dados desempacotados
	for i := 0; i < len(usuarios.Usuarios); i ++ {
		fmt.Println("Nome do usuário: ", usuarios.Usuarios[i].Nome)
		fmt.Println("Tipo de Usuário: ", usuarios.Usuarios[i].Tipo)
		fmt.Println("Idade do Usuário: ", strconv.Itoa(usuarios.Usuarios[i].Idade))
	}
}