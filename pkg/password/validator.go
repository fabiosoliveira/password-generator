package password

import "unicode"

func validatePassword(password string, uppercase, lowercase, numbers, symbols bool) bool {

	uppercaseValid := false
	lowercaseValid := false
	numbersValid := false
	symbolsValid := false

	options := []bool{}

	for _, char := range password {
		if !uppercaseValid && uppercase {
			uppercaseValid = unicode.IsUpper(char)
		}

		if !lowercaseValid && lowercase {
			lowercaseValid = unicode.IsLower(char)
		}

		if !numbersValid && numbers {
			numbersValid = unicode.IsDigit(char)
		}

		if !symbolsValid && symbols {
			symbolsValid = isSymbol(char)
		}
	}

	if uppercase {
		options = append(options, uppercaseValid)
	}

	if lowercase {
		options = append(options, lowercaseValid)
	}

	if numbers {
		options = append(options, numbersValid)
	}

	if symbols {
		options = append(options, symbolsValid)
	}

	for _, option := range options {
		if !option {
			return false
		}
	}

	return true
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
