package mockData

import (
	"fmt"
	"log"

	ce "github.com/mt1976/frantic-core/commonErrors"
)

// Country represents information about the IBAN format for a specific country.
type Country struct {
	IBANLength int    // The expected length of the IBAN for the country.
	Currency   string // The currency code for the country.
	ISOCode    string // The ISO country code
	ISOCode3   string // The ISO country code
}

var Countries map[string]Country

func init() {
	report("Countries")
	Countries = make(map[string]Country)
	Countries["DE"] = Country{IBANLength: 22, Currency: "EUR", ISOCode: "DE", ISOCode3: "DEU"}
	Countries["US"] = Country{IBANLength: 18, Currency: "USD", ISOCode: "US", ISOCode3: "USA"}
	Countries["GB"] = Country{IBANLength: 22, Currency: "GBP", ISOCode: "GB", ISOCode3: "GBR"}
	Countries["FR"] = Country{IBANLength: 27, Currency: "EUR", ISOCode: "FR", ISOCode3: "FRA"}
	Countries["ES"] = Country{IBANLength: 24, Currency: "EUR", ISOCode: "ES", ISOCode3: "ESP"}
	Countries["IT"] = Country{IBANLength: 27, Currency: "EUR", ISOCode: "IT", ISOCode3: "ITA"}
	Countries["NL"] = Country{IBANLength: 18, Currency: "EUR", ISOCode: "NL", ISOCode3: "NLD"}
	Countries["ZA"] = Country{IBANLength: 20, Currency: "ZAR", ISOCode: "ZA", ISOCode3: "ZAF"}
}

func GetCountryInfo(countryCode string) (Country, error) {

	rtn := Country{}
	if len(countryCode) == 2 {
		rtn = Countries[countryCode]
	}

	if len(countryCode) == 3 {
		for _, v := range Countries {
			if v.ISOCode3 == countryCode {
				rtn = v
			}
		}
	}

	if rtn.IBANLength == 0 {
		log.Printf("[WARN] Invalid country code: [%s]", countryCode)
		return Country{}, ce.ErrMockingWrapper(fmt.Errorf("invalid country code: [%s]", countryCode))
	}
	return rtn, nil
}
