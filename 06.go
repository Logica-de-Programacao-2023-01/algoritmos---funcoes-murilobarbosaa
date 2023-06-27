package main

import (
	"errors"
	"fmt"
)

func concatenarComVirgulas(strings []string) (string, error) {
	if len(strings) == 0 {
		return "", errors.New("O slice está vazio")
	}

	concatenacao := strings[0]
	for i := 1; i < len(strings); i++ {
		concatenacao += ", " + strings[i]
	}

	return concatenacao, nil
}

func main() {
	valores := []string{"valor1", "valor2", "valor3"}
	resultado, err := concatenarComVirgulas(valores)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Concatenação:", resultado)
	}

	valoresVazios := []string{}
	resultadoVazio, errVazio := concatenarComVirgulas(valoresVazios)
	if errVazio != nil {
		fmt.Println("Erro", errVazio)
	} else {
		fmt.Println("Concatenação:", resultadoVazio)
	}
}
