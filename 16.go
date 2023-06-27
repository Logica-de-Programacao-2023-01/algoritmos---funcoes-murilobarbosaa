package main

import (
	"errors"
	"fmt"
	"strings"
)

func substituirVogais(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("A string está vazia")
	}

	str = strings.ToLower(str)
	vogais := []string{"a", "e", "i", "o", "u"}

	for _, vogal := range vogais {
		str = strings.ReplaceAll(str, vogal, "*")
	}

	return str, nil
}

func main() {
	str := "Olá, mundo"
	novaStr, err := substituirVogais(str)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("String original: %s \n", str)
		fmt.Printf("Nova string: %s \n", novaStr)
	}

	strVazia := ""
	novaStrVazia, errVazia := substituirVogais(strVazia)
	if errVazia != nil {
		fmt.Println("Erro:", errVazia)
	} else {
		fmt.Printf("String original: %s \n", strVazia)
		fmt.Printf("Nova string: %s \n", novaStrVazia)
	}
}
