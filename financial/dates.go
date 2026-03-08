package financial

import (
	"fmt"
	"strconv"
	"time"

	ce "github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/mockData"
)

// The "FinDate" type represents a date with various properties and formats.
// @property {string} Code - The "Code" property is a string that represents a code associated with the
// date.
// @property {string} Name - The "Name" property is a string that represents the name of the date.
// @property FinDate - The "FinDate" property is of type "time.Time" and represents a specific date and time.
// @property {int} Sort - The "Sort" property is an integer that is used to determine the order in
// which the dates should be sorted. It can be used to sort the dates in ascending or descending order
// based on this value.
// @property {string} Simple - The "Simple" property is a string that represents the date in a
// simplified format. It could be a formatted string that only includes the day, month, and year of the
// date.
// @property {string} External - The "External" property in the FinDate struct is a string that represents
// an external reference or identifier related to the date.
// @property {string} Human - The "Human" property in the "FinDate" struct represents the date in a
// human-readable format. It is likely used to display the date to users in a more understandable way,
// such as "January 1, 2022".
type FinDate struct {
	Code     string
	Name     string
	Date     time.Time
	Index    int
	Simple   string
	External string
	Human    string
}

// The function `GetDateFromTenor` calculates the settlement date based on the given tenor, trade date, and
// currency.
func GetDateFromTenor(tenor Tenor, tradeDate time.Time, ccy ...string) (time.Time, error) {

	if len(ccy) == 0 {
		//xlogs.Warn("no currency provided")
		logHandler.Error.Printf("no currency provided")
		return time.Now(), ce.ErrNoCurrencyProvided
	}

	if tenor.term == "" {
		//xlogs.Warn("no tenor provided")
		logHandler.Error.Printf("no tenor provided")
		return time.Now(), ce.ErrNoTenorProvided
	}

	//fmt.Printf("GetDateFromTenor Tenor [%v] Trade Date [%v] Ccys %v [%v]\n", tenor.String(), tradeDate.Format("2006-01-02"), ccy, len(ccy))

	// Calculate the settlement days, and adjust the date based on the term string provided i.e. 1D, 1W, 1M, 1Y
	// loop thgouth currencies

	if !mockData.IsValidPeriod(tenor.String()) {
		logHandler.Error.Printf("invalid tenor [%s]", tenor.String())
		return time.Now(), ce.ErrInvalidTenorWrapper(tenor.String())
	}

	spotDays := 0

	for _, c := range ccy {
		//fmt.Printf("c: %v\n", c)
		ccySpot, spotError := getSettlementDaysCCY(c)
		if spotError != nil {
			logHandler.Error.Printf("error getting settlement days for currency [%s]", c)
			return time.Now(), spotError
		}
		if ccySpot > spotDays {
			spotDays = ccySpot
		}

	}

	switch tenor.term {
	case "SP":
		//	fmt.Printf("SP: %v\n", spotDays)
		return tradeDate.AddDate(0, 0, spotDays), nil
	case "ON":
		return tradeDate.AddDate(0, 0, spotDays-1), nil
	case "TN":
		return tradeDate.AddDate(0, 0, 1), nil
	case "TD":
		return tradeDate, nil
	}

	tenorPeriod, err := tenorToDuration(tenor)
	if err != nil {
		logHandler.Error.Printf("error converting tenor to duration [%s] [%v]", tenor.String(), err.Error())
		return time.Now(), err
	}
	//fmt.Printf("tenorPeriod: %v\n", tenorPeriod)

	rtn := tradeDate.AddDate(0, 0, spotDays).Add(tenorPeriod)
	//fmt.Printf("GetDateFromTenor rtn: %v\n", rtn)
	return rtn, nil
}

// The function `tenorToDuration` converts a financial tenor (e.g., "1D", "2W", "3M", "4Y") into a
// corresponding time duration.
func tenorToDuration(tenor Tenor) (time.Duration, error) {
	term := tenor.String()
	if len(term) < 2 {
		logHandler.Error.Printf("invalid term length [%s] must be 2 characters", term)
		return 0, fmt.Errorf("invalid term length [%s] must be 2 characters", term)
	}

	valueStr := term[:len(term)-1]

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		logHandler.Error.Printf("invalid term prefix [%s]", term)
		return 0, fmt.Errorf("invalid term prefix [%s]", term)
	}

	switch unit := term[len(term)-1]; unit {
	case 'D':
		return time.Duration(value) * 24 * time.Hour, nil
	case 'W':
		return time.Duration(value) * 7 * 24 * time.Hour, nil
	case 'M':
		return time.Duration(value) * 30 * 24 * time.Hour, nil // Assuming 30 days per month
	case 'Y':
		return time.Duration(value) * 365 * 24 * time.Hour, nil // Assuming 365 days per year
	default:
		logHandler.Error.Printf("invalid term unit: %c", unit)
		return 0, fmt.Errorf("invalid term unit: %c", unit)
	}
}

// The function `GetLadder` takes a pivot date and a list of currency codes as input, and returns a
// list of dates based on a rate ladder.
func GetLadder(pivotDate time.Time, ccy ...string) ([]FinDate, int, error) {
	//fmt.Printf("GetLadder pivotDate: [%v] [%v]\n", pivotDate, ccy)
	if len(ccy) == 0 {
		logHandler.Error.Printf("no currency provided")
		return []FinDate{}, 0, ce.ErrNoCurrencyProvided
	}
	var DateList []FinDate
	//rateLadder := xmock.Ladder
	//xmock.LadderToString(rateLadder)
	//fmt.Printf("rateLadder: %v\n", rateLadder)
	//fmt.Printf("ccy: %v\n", ccy)
	//fmt.Printf("pivotDate: %v\n", pivotDate.Format("2006-01-02"))
	// range over the ladder

	for i := 1; i < mockData.LadderSize; i++ {
		ladder := mockData.GetRateLadderByIndex(i)
		//fmt.Printf("ladder[%v]: %v\n", i, ladder)
		thisTenor, err := NewTenor(ladder.Code)
		if err != nil {
			logHandler.Error.Printf("error [%v]\n", err.Error())
			return DateList, 0, ce.ErrFunctionalWrapper(err, "tenor validation")
		}
		date, err := GetDateFromTenor(thisTenor, pivotDate, ccy...)
		if err != nil {
			logHandler.Error.Printf("error [%v]\n", err.Error())
			return DateList, 0, ce.ErrFunctionalWrapper(err, "tenor -> date calculation")
		}
		//fmt.Printf("thisTenor: [%v] [%v] -> [%v]\n", ladder.Code, thisTenor.String(), date.Format("2006-01-02"))
		di := FinDate{}
		di.Code = ladder.Code
		di.Name = ladder.Name
		di.Date = date
		di.Index = ladder.Index
		di.Simple = date.Format("01/02/2006")
		di.External = date.Format("2006-01-02")
		di.Human = date.Format("Mon 02 Jan 2006")
		DateList = append(DateList, di)
	}
	return DateList, len(DateList), nil
}

func getLadderCCY(pivotdate time.Time, ccy ...string) ([]FinDate, int, error) {
	// TODO - this is a stub, please write logic
	return GetLadder(pivotdate, ccy...)
}

// The function `GetTenorFromDate` takes a date and optional currency as input and returns the
// corresponding tenor or an error.
func GetTenorFromDate(inDate, baseDate time.Time, ccy ...string) (Tenor, error) {
	// TODO - TEST
	// TEST
	//fmt.Printf("GetTenorFromDate inDate: [%v] base: [%v] ccy [%v]\n", dFormat(inDate), dFormat(baseDate), ccy)

	if len(ccy) == 0 {
		logHandler.Error.Printf("no currency provided")
		return Tenor{}, ce.ErrNoCurrencyProvided
	}

	if inDate.Before(baseDate) {
		logHandler.Warning.Printf("date before base date [%v] [%v]", inDate, baseDate)
		return Tenor{}, fmt.Errorf("date before base date")
	}

	tenorList, ladderSize, err := GetLadder(baseDate, ccy...)
	if err != nil {
		logHandler.Error.Printf("Error [%v]\n", err.Error())
		return Tenor{}, err
	}

	var lastTenor FinDate
	lastTenor = FinDate{}
	// loop through tenors in sequence
	for i := 1; i < ladderSize; i++ {
		tenor := tenorList[i]
		//fmt.Printf("GetTenorFromDate loop [%v] tenor.Code: %v tenor.Date: %v inDate: %v d0: %v\n", i, tenor.Code, tenor.Date, inDate, d0)
		if tenor.Date.Equal(inDate) {
			//fmt.Printf("loop tenor.Date.Equal(inDate) %v %v", tenor.Date, inDate)
			rtn, err := NewTenor(tenor.Code)
			if err != nil {
				logHandler.Error.Printf("Error [%v]\n", err.Error())
				return Tenor{}, err
			}
			return rtn, nil
		}
		//fmt.Printf("\"phase2\": %v\n", "phase2")
		// Now check if the tenor date is greater than the input date but less than the next tenor date
		// if so, return the tenor
		//fmt.Printf("GetTenorFromDate P2 inDate: [%v] tenor.Date: [%v] tenor.Code [%v] d0: [%v]\n", inDate, tenor.Date, tenor.Code, d0)
		//fmt.Printf("GetTenorFromDate P2 if [%v].After([%v]) in[%v] cd[%v] td[%v] d0[%v] lt[%v]\n", dFormat(inDate), dFormat(tenor.Date), dFormat(inDate), tenor.Code, dFormat(tenor.Date), dFormat(d0), lastTenor.Code)

		if inDate.After(tenor.Date) {
			lastTenor = tenor
		}

		if inDate.Before(tenor.Date) && inDate.After(lastTenor.Date) {
			//	fmt.Printf("this is the one [%v]\n", lastTenor.Code)
			return NewTenor(lastTenor.Code)
		}

		// 	// get the previous tenor
		// 	previousTenor := tenorList[tenor.Sort-1]
		// 	approximatedTenor, err := NewTenor(previousTenor.Code)
		// 	if err != nil {
		// 		return Tenor{}, err
		// 	}
		// 	fmt.Printf("rtn: %v\n", approximatedTenor)
		// 	return approximatedTenor, nil
		// }
	}
	logHandler.Error.Printf("no tenor found [%v] [%v]", inDate, baseDate)
	return Tenor{}, ce.ErrFunctionalWrapper(fmt.Errorf("no tenor found"), fmt.Sprintf("no tenor found for [%v][%v]", inDate, baseDate))
}

func dFormat(d time.Time) string {
	return d.Format("2006-01-02")
}

func getTenorFromDateCCY(inDate, pivotDate time.Time, ccy string) (Tenor, error) {
	// TODO - this is a stub, please write logic
	logHandler.Trace.Printf("getTenorFromDateCCY [%v] [%v] [%v]", inDate, pivotDate, ccy)
	return Tenor{}, nil
}

func getTenorFromDateCCYPAIR(inDate, pivotDate time.Time, ccy1 string, ccy2 string) (Tenor, error) {
	// TODO - this is a stub, please write logic
	logHandler.Trace.Printf("getTenorFromDateCCYPAIR [%v] [%v] [%v] [%v]", inDate, pivotDate, ccy1, ccy2)
	return Tenor{}, nil
}

func getTenorFromDateCCYCROSS(inDate, pivotDate time.Time, ccy1 string, via string, ccy2 string) (Tenor, error) {
	// TODO - this is a stub, please write logic
	logHandler.Trace.Printf("getTenorFromDateCCYCROSS [%v] [%v] [%v] [%v] [%v]", inDate, pivotDate, ccy1, via, ccy2)
	return Tenor{}, nil
}
