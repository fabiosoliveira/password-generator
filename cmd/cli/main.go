package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fabiosoliveira/password-generator/pkg/password"
)

func main() {
	// gere senhas aleatórias com critérios como comprimento e tipos de caracteres
	args := os.Args[1:]
	if len(args) != 5 {
		fmt.Println("Usage: gen-pass <length:number> <uppercase:bool> <lowercase:bool> <numbers:bool> <symbols:bool>")
		fmt.Println("Example: gen-pass 16 true true true true")
		os.Exit(1)
	}

	// converte os argumentos para os tipos corretos
	length, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid length:", args[0])
		os.Exit(1)
	}

	uppercase, err := strconv.ParseBool(args[1])
	if err != nil {
		fmt.Println("Invalid uppercase:", args[1])
		os.Exit(1)
	}

	lowercase, err := strconv.ParseBool(args[2])
	if err != nil {
		fmt.Println("Invalid lowercase:", args[2])
		os.Exit(1)
	}

	numbers, err := strconv.ParseBool(args[3])
	if err != nil {
		fmt.Println("Invalid numbers:", args[3])
		os.Exit(1)
	}

	symbols, err := strconv.ParseBool(args[4])
	if err != nil {
		fmt.Println("Invalid symbols:", args[4])
		os.Exit(1)
	}

	password := password.GeneratePassword(length, uppercase, lowercase, numbers, symbols)
	fmt.Println(password)
}
