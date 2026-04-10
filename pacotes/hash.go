//Pacote hash fornece funções para calcular hashes de dados. Ele inclui implementações de algoritmos de hash comuns, como MD5, SHA-1 e SHA-256. O pacote hash é amplamente utilizado para verificar a integridade dos dados, criar assinaturas digitais e armazenar senhas de forma segura em Go.
//aceita um conjunto de dados e o reduz a um tamanho fixo menor
//são usadas para buscar dados e detectar alterações em dados, como arquivos ou mensagens
//são divididas em CRIPTOGRAFADAS e NÃO CRIPTOGRAFADAS

//Exemplo de uso do pacote hash CRC32:

package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	// Criando um novo hash CRC32 usando a tabela IEEE
	crc32q := crc32.NewIEEE()
	// Escrevendo dados no hash
	crc32q.Write([]byte("código com pacote hash"))
	// Calculando o hash e imprimindo o resultado
	res := crc32q.Sum32()
	// Imprimindo o hash em formato hexadecimal
	fmt.Printf("CRC32 Hash: %08x\n", res)
}