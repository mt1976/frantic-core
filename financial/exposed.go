package financial

import (
	"strconv"
	"strings"
	"time"

	"github.com/leekchan/accounting"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/mockData"
)

// The function AbbrToInt converts a string representation of a number with abbreviations (such as "1M"
// for 1 million) into its corresponding integer value.
func AbbrToInt(str string) int {
	str = strings.ToUpper(str)
	//number := ""
	number := strings.ReplaceAll(str, "$", "")
	number = strings.ReplaceAll(number, "€", "")
	number = strings.ReplaceAll(number, "£", "")
	fact := strings.ToUpper(number[len(number)-1:])
	number = strings.ReplaceAll(number, "M", "")
	number = strings.ReplaceAll(number, "K", "")
	number = strings.ReplaceAll(number, "T", "")
	number = strings.ReplaceAll(number, "B", "")

	intNum, err := strconv.Atoi(number)
	if err != nil {
		logHandler.Error.Printf("error converting string to int [%s]", number)
	}

	var retNum int
	switch fact {
	case "T":
		retNum = intNum * 1000
	case "K":
		retNum = intNum * 1000
	case "M":
		retNum = intNum * 1000000
	case "B":
		retNum = intNum * 1000000000
	default:
		retNum = intNum
	}

	return retNum
}

// GetSpotDate(inTime invalid type)
// The function "GetSpotDate" takes a time input and returns a modified time value.
func GetSpotDate(inTime time.Time) time.Time {
	spot := inTime.AddDate(0, 0, 2)
	return adjustSettlementForWeekends(spot)
}

// CalculateSpotDate(inTime invalid type)
// The function `GetTenorDate` takes a time and a month as input, adds the specified number of months
// to the time, and returns the resulting date.
func GetTenorDate(inTime time.Time, inMonth string) time.Time {
	month, _ := strconv.Atoi(inMonth)
	spot := inTime.AddDate(0, month, 0)
	return adjustSettlementForWeekends(spot)
}

// The function "GetFirstDayOfYear" returns the first day of the year, assuming that January 1st is a
// holiday and the first day is shifted to January 2nd.
func GetFirstDayOfYear(inTime time.Time) time.Time {
	// Assuking 1st Jan is a holiday therefore first day is 2, then wibble the date.
	tempDate := time.Date(inTime.Year(), 1, 2, 0, 0, 0, inTime.Nanosecond(), inTime.Location())
	return adjustSettlementForWeekends(tempDate)
}

// FormatAmount returns a formated string version of a CCY amount
// The function takes an amount and currency code as input, formats the amount with the specified
// currency symbol and precision, and returns the formatted amount as a string.
//func FormatAmount(inAmount string, inCCY string) string {
//	ac := accounting.Accounting{Symbol: inCCY, Precision: 2, Format: "%v", FormatNegative: "-%v", FormatZero: "\u2013 ;\u2013"}
//	bum, _ := strconv.ParseFloat(inAmount, 64)
//	return ac.FormatMoney(bum)
//}

// FormatAmountFullDPS returns a formated string version of a CCY amount to 7dps
// The function takes an amount and currency code as input, formats the amount with the specified
// precision and currency symbol, and returns the formatted amount as a string.
func FormatAmountFullDPS(inAmount string, inCCY string) string {
	prec, _ := strconv.Atoi("7")
	ac := accounting.Accounting{Symbol: inCCY, Precision: prec, Format: "%v", FormatNegative: "-%v", FormatZero: "\u2013 \u2013"}
	bum, _ := strconv.ParseFloat(inAmount, 64)
	return ac.FormatMoney(bum)
}

// FormatAmountToDPS returns a formated string version of a CCY amount to a given DPS
// The function takes an amount, currency, and precision as input, and formats the amount to the
// specified decimal places with the currency symbol.
func FormatAmountToDPS(inAmount string, inCCY string, inPrec string) string {
	prec, _ := strconv.Atoi(inPrec)
	ac := accounting.Accounting{Symbol: inCCY, Precision: prec, Format: "%v", FormatNegative: "-%v", FormatZero: "\u2013 \u2013"}
	bum, _ := strconv.ParseFloat(inAmount, 64)
	return ac.FormatMoney(bum)
}

func FormatAmount(inAmount float64, inCCY string) string {
	ccyInfo, err := mockData.GetCurrency(inCCY)
	if err != nil {
		logHandler.Error.Printf("Accounting Currenty Error=[%v]", err.Error())
	}
	ac := accounting.Accounting{Symbol: ccyInfo.Character, Precision: ccyInfo.DPS, Format: "%v", FormatNegative: "-%v", FormatZero: "\u2013 \u2013"}
	return ac.FormatMoney(inAmount)
}

func SettlementDate(major string, minor string, pivotDate time.Time) (time.Time, error) {
	rtn, err := getSettlementDatePAIR(major, minor, pivotDate)
	return rtn, err
}

func SettlementDateVia(major string, minor string, pivotDate time.Time, via string) (time.Time, error) {
	rtn, err := getSettlementDateCROSS(major, minor, via, pivotDate)
	return rtn, err
}
