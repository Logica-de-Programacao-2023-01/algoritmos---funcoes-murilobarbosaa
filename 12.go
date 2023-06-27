package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func filtrarMaiusculas(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	resultado := ""
	for _, palavra := range slice {
		if len(palavra) > 0 && unicode.IsUpper(rune(palavra[0])) {
			resultado += palavra + ""
		}
	}

	return strings.TrimSpace(resultado), nil
}

func main() {
	palavras := []string{"Gato", "Cachorro", "Casa", "Janela", "Porta"}
	resultado, err := filtrarMaiusculas(palavras)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Palavras filtradas:", resultado)
	}

	palavrasVazias := []string{}
	resultadoVazio, errVazio := filtrarMaiusculas(palavrasVazias)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Palavras filtradas:", resultadoVazio)
	}
}
