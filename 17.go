package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func ordenarStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	// Ordenar o slice em ordem alfabética
	sort.Strings(slice)

	// Juntar as strings em uma nova string separadas por vírgula
	novaString := strings.Join(slice, ", ")

	return novaString, nil
}

func main() {
	slice := []string{"banana", "maçã", "laranja", "abacaxi"}
	novaString, err := ordenarStrings(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Strings originais: %v\n", slice)
		fmt.Printf("Nova string ordenada: %s\n", novaString)
	}

	sliceVazio := []string{}
	novaStringVazia, errVazia := ordenarStrings(sliceVazio)
	if errVazia != nil {
		fmt.Println("Erro:", errVazia)
	} else {
		fmt.Printf("Strings originais: %v\n", sliceVazio)
		fmt.Printf("Nova string ordenada: %s\n", novaStringVazia)
	}
}
