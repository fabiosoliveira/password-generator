package password

import "math/rand"

func GeneratePassword(length int, uppercase, lowercase, numbers, symbols bool) string {

	strPassword := make([]rune, length)
	functions := []func() rune{}

	if uppercase {
		functions = append(functions, randleUppercase)
	}

	if lowercase {
		functions = append(functions, randleLowercase)
	}

	if numbers {
		functions = append(functions, randleNumbers)
	}

	if symbols {
		functions = append(functions, randleSymbols)
	}

	for i := 0; i < length; i++ {
		strPassword[i] = functions[rand.Intn(len(functions))]()
	}

	if !validatePassword(string(strPassword), uppercase, lowercase, numbers, symbols) {
		return GeneratePassword(length, uppercase, lowercase, numbers, symbols)
	}

	return string(strPassword)
}

func randRune(str string) rune {
	return rune(str[rand.Intn(len(str))])
}

func randleUppercase() rune {
	strUppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return randRune(strUppercase)
}

func randleLowercase() rune {
	strLowercase := "abcdefghijklmnopqrstuvwxyz"
	return randRune(strLowercase)
}

func randleNumbers() rune {
	strNumbers := "0123456789"
	return randRune(strNumbers)
}

func randleSymbols() rune {
	strSymbols := "!@#$%&*()_+{}:<>?|"
	return randRune(strSymbols)
}
