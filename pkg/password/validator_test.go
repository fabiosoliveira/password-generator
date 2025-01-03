package password

import (
	"testing"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		password  string
		uppercase bool
		lowercase bool
		numbers   bool
		symbols   bool
		expected  bool
	}{
		{"Password123!", true, true, true, true, true},
		{"Password123", true, true, true, false, true},
		{"Password!", true, true, false, true, true},
		{"password123!", false, true, true, true, true},
		{"PASSWORD123!", true, false, true, true, true},
		{"Password", true, true, false, false, true},
		{"password", false, true, false, false, true},
		{"PASSWORD", true, false, false, false, true},
		{"123456", false, false, true, false, true},
		{"!@#$%", false, false, false, true, true},
		{"", false, false, false, false, true},
	}

	for _, test := range tests {
		result := validatePassword(test.password, test.uppercase, test.lowercase, test.numbers, test.symbols)
		if result != test.expected {
			t.Errorf("validatePassword(%q, %v, %v, %v, %v) = %v; want %v", test.password, test.uppercase, test.lowercase, test.numbers, test.symbols, result, test.expected)
		}
	}
}
