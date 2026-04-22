package main

import "testing"

func TestSoma(t *testing.T) {
	teste := soma(3 ,2 ,1)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func TestSoma2(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado :=7

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}