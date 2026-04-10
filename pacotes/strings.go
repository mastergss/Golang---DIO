//strings é um pacote da biblioteca padrão do Go que fornece funções para manipulação de strings
//exemplo: contains verifica se uma string contém outra string
//ela recebe dois parâmetros: a string principal e a string que queremos verificar
//ex: Terra - 'rr' - true

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Terra", "rr"))
	texto := "O rato roeu a roupa do rei de Roma"
	fmt.Println(strings.Contains(texto, "roeu"))
	fmt.Println(strings.Contains(texto, "roeu a roupa do rei de Roma"))
	fmt.Println(strings.Contains(texto, "roeu a roupa do rei de Roma e foi embora"))
}