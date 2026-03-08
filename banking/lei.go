package banking

import (
	"fmt"

	"github.com/mt1976/frantic-core/commonErrors"
	zlei "github.com/trisacrypto/lei"
)

type LEI struct {
	lei string
}

// isValidLEI checks if the given LEI is valid.
// The function `isValidLEI` checks if a given LEI (Legal Entity Identifier) is valid by removing
// spaces and converting it to uppercase.
func isValidLEI(inLei string) bool {
	// Remove spaces and convert to uppercase

	if err := zlei.LEI(inLei).Check(); err != nil {
		return false
	}
	return true
}

// The function `Formatted` takes a Legal Entity Identifier (LEI) as input and returns a formatted
// string with the LEI divided into its components.
func (l *LEI) Formatted() string {
	lou := l.lei[:4]
	reserved := l.lei[4:6]
	entity := l.lei[6:18]
	checksum := l.lei[18:]
	return fmt.Sprintf("lou=%s res=%s entity=%s checksum=%s", lou, reserved, entity, checksum)
}

func NewLEI(lei string) (LEI, error) {
	l := LEI{}
	l.lei = lei
	if !isValidLEI(lei) {
		return LEI{}, commonErrors.ErrValidationWrapper(fmt.Errorf("invalid LEI: %s", lei))
	}
	return l, nil
}

func (L *LEI) String() string {
	return L.lei
}
