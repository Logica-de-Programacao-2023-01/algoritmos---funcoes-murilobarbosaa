package main

import (
	"fmt"
	"sort"
)

func segundoMaiorValor(valores []int) int {
	sort.Ints(valores)

	unicos := make([]int, 0, len(valores))
	ultimoValor := valores[0]
	unicos = append(unicos, ultimoValor)

	for i := 1; i < len(valores); i++ {
		valorAtual := valores[i]
		if valorAtual != ultimoValor {
			unicos = append(unicos, valorAtual)
			ultimoValor = valorAtual
		}
	}

	if len(unicos) < 2 {
		panic("Não há segundo maior valor")
	}

	return unicos[len(unicos)-2]
}

func main() {
	numeros := []int{10, 5, 8, 15, 3, 20}
	segundoMaior := segundoMaiorValor(numeros)
	fmt.Println("Segundo maior valor:", segundoMaior)
}
