// CALCULADORA
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// SOMA, SUBTRACAO, MULTIPLICACAO, DIVISAO
	// + , - , *, /
	// 3 + 3 => 6
	// 3-3 => 0
	// 3*3 => 9

	// Criar um scanner para ler do terminal
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite um numero: ")
	scanner.Scan()
	numero1 := scanner.Text()
	fmt.Print("Digite outro numero: ")
	scanner.Scan()
	numero2 := scanner.Text()

	numero1Inteiro, err := strconv.Atoi(numero1)
	if err != nil {
		panic(err)
	}
	numero2Inteiro, err := strconv.Atoi(numero2)
	if err != nil {
		panic(err)
	}
	if numero1Inteiro == numero2Inteiro {
		fmt.Println("Verdade")
	} else {
		fmt.Println("Mentiraa")
	}
	fmt.Println("Soma:", soma(numero1Inteiro, numero2Inteiro))
	fmt.Println("Subtracao:", subtracao(numero1Inteiro, numero2Inteiro))
}

// Funcao
func soma(numero1 int, numero2 int) int {
	return numero1 + numero2
}

func subtracao(numero1 int, numero2 int) *int {
	subtracao := numero1 - numero2 // usando ponteiros
	return &subtracao
}
