package password

import (
	"testing"
	"unicode"
)

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		length    int
		uppercase bool
		lowercase bool
		numbers   bool
		symbols   bool
	}{
		{10, true, true, true, true},
		{8, true, true, true, false},
		{12, true, true, false, true},
		{6, false, true, true, true},
		{15, true, false, true, true},
		{5, true, true, false, false},
		{7, false, true, false, false},
		{9, true, false, false, false},
		{10, false, false, true, false},
		{11, false, false, false, true},
	}

	for _, test := range tests {
		password := GeneratePassword(test.length, test.uppercase, test.lowercase, test.numbers, test.symbols)
		if len(password) != test.length {
			t.Errorf("GeneratePassword(%d, %v, %v, %v, %v) = %q; length = %d; want length = %d", test.length, test.uppercase, test.lowercase, test.numbers, test.symbols, password, len(password), test.length)
		}

		if test.uppercase && !containsUppercase(password) {
			t.Errorf("GeneratePassword(%d, %v, %v, %v, %v) = %q; want at least one uppercase letter", test.length, test.uppercase, test.lowercase, test.numbers, test.symbols, password)
		}

		if test.lowercase && !containsLowercase(password) {
			t.Errorf("GeneratePassword(%d, %v, %v, %v, %v) = %q; want at least one lowercase letter", test.length, test.uppercase, test.lowercase, test.numbers, test.symbols, password)
		}

		if test.numbers && !containsNumber(password) {
			t.Errorf("GeneratePassword(%d, %v, %v, %v, %v) = %q; want at least one number", test.length, test.uppercase, test.lowercase, test.numbers, test.symbols, password)
		}

		if test.symbols && !containsSymbol(password) {
			t.Errorf("GeneratePassword(%d, %v, %v, %v, %v) = %q; want at least one symbol", test.length, test.uppercase, test.lowercase, test.numbers, test.symbols, password)
		}
	}
}

func containsUppercase(s string) bool {
	for _, char := range s {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func containsLowercase(s string) bool {
	for _, char := range s {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}

func containsNumber(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func containsSymbol(s string) bool {
	symbols := "!@#$%&*()_+{}:<>?|"
	for _, char := range s {
		if containsRune(symbols, char) {
			return true
		}
	}
	return false
}

func containsRune(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}
