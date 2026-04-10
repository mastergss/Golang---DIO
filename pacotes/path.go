//Pacote path fornece funcionalidades para manipulação de caminhos de arquivos e diretórios. Ele inclui funções para construir, analisar e manipular caminhos de arquivos de forma independente do sistema operacional. O pacote path é útil para lidar com caminhos de arquivos em diferentes plataformas, como Windows, Linux e macOS.

// Exemplo de uso do pacote path.filepath para percorrer um diretório e imprimir os caminhos dos arquivos encontrados:
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Percorre o diretório atual e imprime os caminhos dos arquivos encontrados
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		// Retorna nil para continuar a caminhada
		return nil
	})
}