package main

import (
	"errors"
	"fmt"
)

func stringsMaisDe5Caracteres(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	resultado := []string{}

	for _, str := range slice {
		if len(str) > 5 {
			resultado = append(resultado, str)
		}
	}

	return resultado, nil
}

func main() {
	slice := []string{"banana", "abacaxi", "laranja", "uva", "mamão", "limão"}
	maisDe5Caracteres, err := stringsMaisDe5Caracteres(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Slice original: %v\n", slice)
		fmt.Printf("Strings com mais de 5 caracteres: %v\n", maisDe5Caracteres)
	}

	sliceVazio := []string{}
	maisDe5CaracteresVazio, errVazio := stringsMaisDe5Caracteres(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Printf("Slice original: %v\n", sliceVazio)
		fmt.Printf("Strings com mais de 5 caracteres: %v\n", maisDe5CaracteresVazio)
	}
}
