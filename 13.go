package main

import (
	"errors"
	"fmt"
	"strconv"
)

func somaDigitos(numero int) (int, error) {
	if numero < 0 {
		return 0, errors.New("O número é negativo")
	}

	digitos := strconv.Itoa(numero)
	soma := 0
	for _, d := range digitos {
		digito, _ := strconv.Atoi(string(d))
		soma += digito
	}

	return soma, nil
}

func main() {
	numero := 12345
	soma, err := somaDigitos(numero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("A soma dos dígitos de %d é %d \n", numero, soma)
	}

	numeroNegativo := -1234
	somaNegativa, errNegativa := somaDigitos(numeroNegativo)
	if errNegativa != nil {
		fmt.Println("Erro:", errNegativa)
	} else {
		fmt.Printf("A soma dos digitos de %d é %d \n", numeroNegativo, somaNegativa)
	}
}
