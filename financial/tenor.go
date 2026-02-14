package financial

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/mt1976/frantic-core/logHandler"
)

type Tenor struct {
	term string
}

// The function NewTenor takes a string as input and returns a Tenor object.
func NewTenor(term string) (Tenor, error) {
	newTenor := Tenor{}
	_, err := newTenor.Set(term)
	if err != nil {
		return Tenor{}, err
	}
	return newTenor, nil
}

// The function String returns the term of a Tenor object.
func (t *Tenor) String() string {
	return t.term
}

// The function Set takes a string as input and sets the term of a Tenor object.
func (t *Tenor) Set(term string) (*Tenor, error) {
	newTenor, err := validateAndFormatTenor(term)
	if err != nil {
		logHandler.Event.Printf("invalid tenor [%s] [%v]", term, err.Error())
		return nil, err
	}
	t.term = newTenor
	return t, nil
}

func validateAndFormatTenor(tenor string) (string, error) {
	//Validates that the term string is valid
	// Validation is that the string is at least 2 characters long, and the last character is a valid unit
	// i.e. D, W, M, Y
	if len(tenor) < 2 {
		logHandler.Error.Printf("invalid tenor [%s] must be at least 2 characters long", tenor)
		return "", fmt.Errorf("invalid tenor [%s] must be at least 2 characters long", tenor)
	}
	unit := tenor[len(tenor)-1]
	unit = byte(unicode.ToUpper(rune(unit)))
	factor := tenor[:len(tenor)-1]

	// Deal with special cases of SP and TD
	uTerm := strings.ToUpper(tenor)
	// Special cases SP, TD, ON, TN
	if uTerm == "SP" || uTerm == "TD" || uTerm == "ON" || uTerm == "TN" || uTerm == "SN" {
		return uTerm, nil
	}

	_, err := strconv.Atoi(factor)
	if err != nil {
		logHandler.Error.Printf("supplied value [%s] is not a number %v", factor, err.Error())
		return "", fmt.Errorf("supplied value [%s] is not a number", factor)
	}

	clean := fmt.Sprintf("%s%c", factor, unit)

	switch unit {
	case 'D':
		return clean, nil
	case 'W':
		return clean, nil
	case 'M':
		return clean, nil
	case 'Y':
		return clean, nil
	default:
		logHandler.Error.Printf("invalid tenor mnemonic [%c]", unit)
		return "", fmt.Errorf("invalid tenor mnemonic")
	}
}
