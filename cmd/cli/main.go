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
		PrintError()
		os.Exit(1)
	}

	length := Must(strconv.Atoi(args[0]))
	uppercase := Must(strconv.ParseBool(args[1]))
	lowercase := Must(strconv.ParseBool(args[2]))
	numbers := Must(strconv.ParseBool(args[3]))
	symbols := Must(strconv.ParseBool(args[4]))

	password := password.GeneratePassword(length, uppercase, lowercase, numbers, symbols)
	fmt.Println(password)
}

func Must[T bool | int](option T, err error) T {
	if err != nil {
		fmt.Println("Invalid symbols:", option)
		PrintError()
		os.Exit(1)
	}
	return option
}

func PrintError() {
	fmt.Println("Usage: gen-pass <length:number> <uppercase:bool> <lowercase:bool> <numbers:bool> <symbols:bool>")
	fmt.Println("Example: gen-pass 16 true true true true")
}
