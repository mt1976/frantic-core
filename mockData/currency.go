package mockData

import (
	"fmt"
	"log"
	"time"

	ce "github.com/mt1976/frantic-core/commonErrors"
)

// Currency represents information about the IBAN format for a specific country.
type Currency struct {
	Code               string // The ISO currency code
	SpotDays           int    // The number of spot days for the currency
	Name               string // The name of the currency
	Character          string // The character of the currency
	DPS                int    // The number of decimal places for the currency
	QuoteDPS           int    // The number of decimal places for the currency when quoting
	Type               string // The type of currency
	MajorUnit          string // The major unit of the currency
	MinorUnit          string // The minor unit of the currency
	MinorCharacter     string // The minor character of the currency
	ISONumericCode     string // The ISO numeric code of the currency
	KnownAs            string // The known as name of the currency
	YearOfIntroduction int    // The year the currency was introduced
}

const (
	Fiat    = "Fiat"
	Crypto  = "Crypto"
	Metals  = "Psuedo"
	Testing = "Testing"
)

var Currencies map[string]Currency

func init() {

	report("Currencies")
	Currencies = make(map[string]Currency)
	Currencies["EUR"] = Currency{Code: "EUR", SpotDays: 2, Name: "Euro", Character: "€", DPS: 2, QuoteDPS: 4, Type: Fiat, MajorUnit: "Euro", MinorUnit: "Cent", ISONumericCode: "978", KnownAs: "Euros", MinorCharacter: "c", YearOfIntroduction: 1999}
	Currencies["USD"] = Currency{Code: "USD", SpotDays: 1, Name: "US Dollar", Character: "$", DPS: 2, QuoteDPS: 4, Type: Fiat, MajorUnit: "Dollar", MinorUnit: "Cent", ISONumericCode: "840", KnownAs: "Bucks", MinorCharacter: "c", YearOfIntroduction: 1792}
	Currencies["GBP"] = Currency{Code: "GBP", SpotDays: 2, Name: "Pound Sterling", Character: "£", DPS: 2, QuoteDPS: 4, Type: Fiat, MajorUnit: "Pound", MinorUnit: "Pence", ISONumericCode: "826", KnownAs: "Quids", MinorCharacter: "p", YearOfIntroduction: 800}
	Currencies["ZAR"] = Currency{Code: "ZAR", SpotDays: 2, Name: "South African Rand", Character: "R", DPS: 2, QuoteDPS: 4, Type: Fiat, MajorUnit: "Rand", MinorUnit: "Cent", ISONumericCode: "710", KnownAs: "Bucks", MinorCharacter: "c", YearOfIntroduction: 1961}
	Currencies["MXN"] = Currency{Code: "MXN", SpotDays: 3, Name: "Mexican Peso", Character: "$", DPS: 2, QuoteDPS: 4, Type: Fiat, MajorUnit: "Peso", MinorUnit: "Centavo", ISONumericCode: "484", KnownAs: "Pesos", MinorCharacter: "c", YearOfIntroduction: 1993}
	Currencies["CAD"] = Currency{Code: "CAD", SpotDays: 1, Name: "Canadian Dollar", Character: "$", DPS: 2, QuoteDPS: 4, Type: Fiat, MajorUnit: "Dollar", MinorUnit: "Cent", ISONumericCode: "124", KnownAs: "Loonies", MinorCharacter: "c", YearOfIntroduction: 1858}
	Currencies["JPY"] = Currency{Code: "JPY", SpotDays: 2, Name: "Japanese Yen", Character: "¥", DPS: 0, QuoteDPS: 2, Type: Fiat, MajorUnit: "Yen", MinorUnit: "Sen", ISONumericCode: "392", KnownAs: "Yen", MinorCharacter: "s", YearOfIntroduction: 1871}
	Currencies["CHF"] = Currency{Code: "CHF", SpotDays: 2, Name: "Swiss Franc", Character: "Fr", DPS: 2, QuoteDPS: 4, Type: Fiat, MajorUnit: "Franc", MinorUnit: "Rappen", ISONumericCode: "756", KnownAs: "Swissies", MinorCharacter: "rp", YearOfIntroduction: 1850}
	Currencies["AUD"] = Currency{Code: "AUD", SpotDays: 2, Name: "Australian Dollar", Character: "$", DPS: 2, QuoteDPS: 4, Type: Fiat, MajorUnit: "Dollar", MinorUnit: "Cent", ISONumericCode: "036", KnownAs: "Aussie", MinorCharacter: "c", YearOfIntroduction: 1966}
	Currencies["INR"] = Currency{Code: "INR", SpotDays: 2, Name: "Indian Rupee", Character: "₹", DPS: 2, QuoteDPS: 4, Type: Fiat, MajorUnit: "Rupee", MinorUnit: "Paisa", ISONumericCode: "356", KnownAs: "Rupayya", MinorCharacter: "p", YearOfIntroduction: 1540}
	Currencies["CLF"] = Currency{Code: "CLF", SpotDays: 2, Name: "Chilean Unidad de Fomento", Character: "UF", DPS: 4, QuoteDPS: 4, Type: Fiat, MajorUnit: "Unidad de Fomento", MinorUnit: "Peso", ISONumericCode: "990", YearOfIntroduction: 1967}
	Currencies["CNY"] = Currency{Code: "CNY", SpotDays: 2, Name: "Chinese Yuan Renminbi", Character: "¥", DPS: 2, QuoteDPS: 4, Type: Fiat, MajorUnit: "Yuan", MinorUnit: "Fen", ISONumericCode: "156", KnownAs: "Yuan", MinorCharacter: "f", YearOfIntroduction: 1949}
	Currencies["IQD"] = Currency{Code: "IQD", SpotDays: 2, Name: "Iraqi Dinar", Character: "ع.د", DPS: 3, QuoteDPS: 4, Type: Fiat, MajorUnit: "Dinar", MinorUnit: "Fils", ISONumericCode: "368", KnownAs: "Dinar", MinorCharacter: "f", YearOfIntroduction: 1932}
	Currencies["XAG"] = Currency{Code: "XAG", SpotDays: 2, Name: "Silver", Character: "Ag", DPS: 0, QuoteDPS: 0, Type: Metals, MajorUnit: "Ounce", MinorUnit: "Ounce", ISONumericCode: "961", KnownAs: "Silver"}
	Currencies["XAU"] = Currency{Code: "XAU", SpotDays: 2, Name: "Gold", Character: "Au", DPS: 0, QuoteDPS: 0, Type: Metals, MajorUnit: "Ounce", MinorUnit: "Ounce", ISONumericCode: "959", KnownAs: "Gold"}
	Currencies["XTS"] = Currency{Code: "XTS", SpotDays: 2, Name: "Testing Currency Code", Character: "¤", DPS: 4, QuoteDPS: 4, Type: Testing, MajorUnit: "Unit", MinorUnit: "Unit", ISONumericCode: "999", KnownAs: "Testing Currency Code", MinorCharacter: "u", YearOfIntroduction: 1970}
	Currencies["BTC"] = Currency{Code: "BTC", SpotDays: 0, Name: "Bitcoin", Character: "₿", DPS: 8, QuoteDPS: 8, Type: Crypto, MajorUnit: "Bitcoin", MinorUnit: "Satoshi", ISONumericCode: "1001", KnownAs: "Bitcoin", MinorCharacter: "s", YearOfIntroduction: 2009}
	Currencies["ETH"] = Currency{Code: "ETH", SpotDays: 0, Name: "Ethereum", Character: "Ξ", DPS: 8, QuoteDPS: 8, Type: Crypto, MajorUnit: "Ether", MinorUnit: "Wei", ISONumericCode: "1002", KnownAs: "Ether", MinorCharacter: "w", YearOfIntroduction: 2015}
	Currencies["LTC"] = Currency{Code: "LTC", SpotDays: 0, Name: "Litecoin", Character: "Ł", DPS: 8, QuoteDPS: 8, Type: Crypto, MajorUnit: "Litecoin", MinorUnit: "Litetoshi", ISONumericCode: "1003", KnownAs: "Litecoin", MinorCharacter: "l", YearOfIntroduction: 2011}
	Currencies["XRP"] = Currency{Code: "XRP", SpotDays: 0, Name: "Ripple", Character: "Ʀ", DPS: 8, QuoteDPS: 8, Type: Crypto, MajorUnit: "Ripple", MinorUnit: "Drops", ISONumericCode: "1004", KnownAs: "XRup", MinorCharacter: "d", YearOfIntroduction: 2012}
	Currencies["DASH"] = Currency{Code: "DASH", SpotDays: 0, Name: "Dash", Character: "Đ", DPS: 8, QuoteDPS: 8, Type: Crypto, MajorUnit: "Dash", MinorUnit: "Duffs", ISONumericCode: "1005", KnownAs: "Dash", MinorCharacter: "d", YearOfIntroduction: 2014}
	Currencies["DOGE"] = Currency{Code: "DOGE", SpotDays: 0, Name: "Dogecoin", Character: "Ð", DPS: 8, QuoteDPS: 8, Type: Crypto, MajorUnit: "Dogecoin", MinorUnit: "Shibes", ISONumericCode: "1006", KnownAs: "Much Wow", MinorCharacter: "s", YearOfIntroduction: 2013}
}

func GetCurrency(code string) (Currency, error) {

	rtn := Currency{}
	if len(code) > 1 && len(code) <= 4 {
		rtn = Currencies[code]
	}

	if len(rtn.Code) == 0 {
		log.Printf("[WARN] Invalid currency code: [%s]", code)
		return Currency{}, ce.ErrMockingWrapper(fmt.Errorf("invalid currency code: [%s]", code))
	}
	//fmt.Printf("CurrencyInfo: %v Age %v years\n", rtn, rtn.Age())
	return rtn, nil
}

func (C *Currency) Age() int {
	return time.Now().Year() - C.YearOfIntroduction
	//return C.SpotDays
}
