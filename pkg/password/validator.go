package password

func validatePassword(password string, uppercase, lowercase, numbers, symbols bool) bool {

	uppercaseValid := false
	lowercaseValid := false
	numbersValid := false
	symbolsValid := false

	for _, char := range password {
		if !uppercaseValid && uppercase {
			uppercaseValid = isUppercase(char)
		}

		if !lowercaseValid && lowercase {
			lowercaseValid = isLowercase(char)
		}

		if !numbersValid && numbers {
			numbersValid = isNumber(char)
		}

		if !symbolsValid && symbols {
			symbolsValid = isSymbol(char)
		}
	}

	return uppercaseValid && lowercaseValid && numbersValid && symbolsValid
}

func isSymbol(char rune) bool {
	strSymbols := "!@#$%&*()_+{}:<>?|"
	for _, symbol := range strSymbols {
		if char == symbol {
			return true
		}
	}
	return false
}

func isNumber(char rune) bool {
	return char >= '0' && char <= '9'
}

func isLowercase(char rune) bool {
	return char >= 'a' && char <= 'z'
}

func isUppercase(char rune) bool {
	return char >= 'A' && char <= 'Z'
}
