package banking

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/mockData"
)

type IBAN struct {
	iban string
}

// isValidIBAN checks if the given IBAN is valid.
// The function `isValidIBAN` checks if a given IBAN (International Bank Account Number) is valid by
// performing various checks and calculations.
func (I *IBAN) isValid(iban string) bool {
	// Remove spaces and convert to uppercase
	iban = strings.ToUpper(strings.ReplaceAll(iban, " ", ""))

	// Check if the IBAN length is valid for the country code
	countryCode := iban[:2]
	countryInfo, err := mockData.GetCountryInfo(countryCode)
	if err != nil || len(iban) != countryInfo.IBANLength {
		return false
	}

	// Move the first 4 characters to the end
	iban = iban[4:] + iban[:4]

	// Convert characters to numbers (A = 10, B = 11, ..., Z = 35)
	var numericIBAN string
	for _, char := range iban {
		if '0' <= char && char <= '9' {
			numericIBAN += string(char)
		} else {
			numericIBAN += fmt.Sprintf("%d", int(char-'A'+10))
		}
	}

	// Convert numeric IBAN to a big.Int for modulo calculation
	bigIntIBAN, _ := new(big.Int).SetString(numericIBAN, 10)

	// Check if the modulo of the numeric IBAN with 97 is equal to 1
	return new(big.Int).Mod(bigIntIBAN, big.NewInt(97)).Int64() == 1
}

func NewIBAN(iban string) (IBAN, error) {
	i := IBAN{}
	i.iban = iban
	if !i.isValid(iban) {
		return IBAN{}, commonErrors.ErrValidationWrapper(fmt.Errorf("invalid IBAN: %s", iban))
	}
	return i, nil
}

func (I *IBAN) String() string {
	return I.iban
}
