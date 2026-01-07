package banking

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/logHandler"
)

type UTI struct {
	uti string
}

// NewISO23897UTI generates an ISO 23897:2020 compliant Unique Transaction Identifier (UTI).
func NewISO23897UTI(generatingEntity string) (UTI, error) {

	if len(generatingEntity) != 20 {
		return UTI{}, commonErrors.ErrValidationWrapper(fmt.Errorf("generatingEntity is not 20 chars [%v][%s]", len(generatingEntity), generatingEntity))
	}

	// Generate a random number using the current time as the seed
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(1000000) // 6-digit random number

	// Get the current date in YYYYMMDD format
	date := time.Now().Format("20060102")

	// Construct the UTI using the generating entity, date, and random number
	uti := fmt.Sprintf("%s%s%d", generatingEntity, date, randomNumber)

	// Pad the UTI to 52 characters with zeros
	uti = fmt.Sprintf("%-52s", uti)

	nu := UTI{uti: uti}
	//fmt.Printf("UTI: %s\n", nu.Formatted())
	logHandler.InfoLogger.Printf("UTI: %s\n", nu.Formatted())

	return nu, nil
}

func (U *UTI) String() string {
	return U.Get()
}

func (U *UTI) Set(in string) error {
	U.uti = in
	val, err := U.IsValid()
	if err != nil || !val {
		return commonErrors.ErrValidationWrapper(fmt.Errorf("invalid UTI [%s]", in))
	}
	if !val {
		return commonErrors.ErrValidationWrapper(fmt.Errorf("invalid UTI [%s]", in))
	}
	return nil
}

func (U *UTI) Get() string {
	return U.uti
}

func (U *UTI) IsValid() (bool, error) {
	// Max length is 52
	if len(U.uti) > 52 {
		return false, commonErrors.ErrValidationWrapper(fmt.Errorf("invalid UTI length [%s]", U.uti))
	}
	// US Min length is 42
	if len(U.uti) < 42 {
		return false, commonErrors.ErrValidationWrapper(fmt.Errorf("invalid UTI length [%s]", U.uti))
	}

	//fmt.Printf("UTI: %s\n", U.Formatted())
	logHandler.InfoLogger.Printf("UTI: %s\n", U.Formatted())

	return true, nil
}

func (U *UTI) IsEmpty() bool {
	return U.uti == ""
}

func (U *UTI) Formatted() string {
	// Format as 1 group of 20 chars, a space, the remaining chars
	return fmt.Sprintf("%s %s %s", U.uti[:20], U.uti[20:28], U.uti[28:])
}
