package main

import "fmt"

func contarVogais(texto string) int {
	vogais := "aeiouAEIOU"
	contador := 0

	for _, char := range texto {
		for _, vogal := range vogais {
			if char == vogal {
				contador++
				break
			}
		}
	}

	return contador
}

func main() {
	textoExemplo := "Hello, World!"
	quantidadeDeVogais := contarVogais(textoExemplo)
	fmt.Println("Quantidade de vogais:", quantidadeDeVogais)
}
