package banking

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/mockData"
)

type ISIN struct {
	value string
}

// IsValid checks if the given ISIN is valid.
func (I *ISIN) IsValid() bool {

	//fmt.Printf("Validate: %s\n", I.value)
	// Remove spaces and convert to uppercase
	val := strings.ToUpper(strings.ReplaceAll(I.String(), " ", ""))

	// Check if the ISIN length is valid (it should be 12 characters)
	if len(val) != 12 {
		//log.Printf("[WARN] ISIN != 12 characters (%v)\n", len(val))
		logHandler.WarningLogger.Printf("ISIN != 12 characters (%v)\n", len(val))
		return false
	}

	//Validate the ISIN prefix, first two characters are a valud countrycode
	countryCode := val[:2]
	countryInfo, err := mockData.GetCountryInfo(countryCode)
	if err != nil || countryCode != countryInfo.ISOCode {
		//log.Printf("[WARN] ISIN prefix not a valid country code (%v)\n", countryCode)
		logHandler.WarningLogger.Printf("ISIN prefix not a valid country code (%v)\n", countryCode)
		return false
	}
	checksum, _ := strconv.Atoi(val[11:])
	rtnVal := I.calculateChecksum() == checksum

	return rtnVal
}

func (I *ISIN) String() string {
	return I.Get()
}

func (I *ISIN) Set(in string) error {
	I.value = in
	if !I.IsValid() {
		return commonErrors.ErrValidationWrapper(fmt.Errorf("invalid ISIN [%s]", in))
	}
	return nil
}

func (I *ISIN) Get() string {
	return I.value
}

// calculateChecksum calculates the ISIN checksum digit using Luhn algorithm.
func (I *ISIN) calculateChecksum() int {
	// Convert characters to numbers (A = 10, B = 11, ..., Z = 35)
	var numericISIN string
	for _, char := range I.value[:11] {
		if '0' <= char && char <= '9' {
			numericISIN += string(char)
		} else {
			numericISIN += fmt.Sprintf("%d", int(char-'A'+10))
		}
	}

	//fmt.Printf("Numeric ISIN: %s\n", numericISIN)

	// Calculate the checksum digit using Luhn algorithm
	// Start from the rightmost digit, double every other digit and sum the digits
	// If the result is greater than 9, subtract 9 from it
	sum := 0
	for i := 0; i < len(numericISIN); i++ {
		digit := int(numericISIN[i] - '0')
		//	fmt.Printf("Start %v %v", digit, (i%2 == 0))
		if i%2 == 0 {
			//fmt.Printf("*")
			digit *= 2
			if digit > 9 {
				rmd, _ := strconv.Atoi(strconv.Itoa(digit)[1:])
				lmd, _ := strconv.Atoi(strconv.Itoa(digit)[:1])
				//		fmt.Printf(" -> %v %v %v", rmd, lmd, digit)
				digit = rmd + lmd
				//digit -= 9
			}
		}
		//fmt.Printf(" -> %v \n", digit)
		sum += digit
	}
	//fmt.Printf("\n")

	//fmt.Printf("Sum: %v\n", sum)
	//Find the smallest number ending with a zero that is greater than or equal to sum, and call it val
	val := sum
	for val%10 != 0 {
		val++
	}

	//fmt.Printf("Val: %v\n", val)

	// Calculate the checksum value (next highest multiple of 10 - sum) modulo 10
	checksumValue := val - sum

	//fmt.Printf("Checksum value: %v\n", checksumValue)
	// Convert the checksum value back to the character representation
	checksum := fmt.Sprintf("%v", checksumValue)
	//TODO  Handle Errors Better
	rtn, _ := strconv.Atoi(string(checksum[0]))
	return rtn
}

func (I *ISIN) Printable() string {
	working := I.Get()
	country := working[:2]
	nsin := working[2:11]
	checksum := working[11:]
	return fmt.Sprintf("%s %s %s", country, nsin, checksum)
}
