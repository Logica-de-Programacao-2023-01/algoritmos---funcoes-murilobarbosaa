package main

import (
	"errors"
	"fmt"
	"strings"
)

func extraisPalavras(str string) ([]string, error) {
	if str == "" {
		return nil, errors.New("A string está vazia")
	}

	palavras := strings.Fields(str)
	return palavras, nil
}

func main() {
	frase := "Olá, como vai você hoje?"
	palavras, err := extraisPalavras(frase)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Palavras:", palavras)
	}

	fraseVazia := ""
	palavrasVazias, errVazias := extraisPalavras(fraseVazia)
	if errVazias != nil {
		fmt.Println("Erro:", errVazias)
	} else {
		fmt.Println("Palavras:", palavrasVazias)
	}
}
