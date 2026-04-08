// Escopo em Go se refere à visibilidade e ao tempo de vida das variáveis. Existem dois tipos principais de escopo: escopo global e escopo local. Variáveis declaradas fora de qualquer função têm escopo global e podem ser acessadas por qualquer função dentro do mesmo pacote. Variáveis declaradas dentro de uma função têm escopo local e só podem ser acessadas dentro dessa função.
package main

import "fmt"
// Variável global
var globalVar = "I am a global variable"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(globalVar)
    // Variável local
    localVar := "I am a local variable"
    fmt.Println(localVar)
}