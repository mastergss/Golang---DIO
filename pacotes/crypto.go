//Pacote crypto fornece funções para criptografia e segurança em Go. Ele inclui implementações de algoritmos de criptografia simétrica, como AES, e algoritmos de criptografia assimétrica, como RSA. O pacote crypto é amplamente utilizado para proteger dados sensíveis, criar conexões seguras e implementar autenticação em Go.

//Exemplo de uso do pacote crypto para criptografia SHA-1:

package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	// Criando um novo hash SHA-1
	h := sha1.New()
	// Escrevendo dados no hash
	h.Write([]byte("código com pacote crypto"))
	// Calculando o hash e imprimindo o resultado
	res := h.Sum([]byte{})
	// Imprimindo o hash em formato hexadecimal
	fmt.Printf("SHA-1 Hash: %x\n", res)
}