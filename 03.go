package main

import "fmt"

func concatenarStrings(strings []string) string {
	concatenacao := ""

	for _, str := range strings {
		concatenacao += str
	}

	return concatenacao
}

func main() {
	stringsExemplo := []string{"Olá", " ", "mundo", "!"}
	resultado := concatenarStrings(stringsExemplo)
	fmt.Println("Concatenação: ", resultado)
}
