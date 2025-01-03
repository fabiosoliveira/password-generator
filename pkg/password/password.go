package password

import "math/rand"

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
