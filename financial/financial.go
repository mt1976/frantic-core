package financial

import (
	"time"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/mathHelpers"
	"github.com/mt1976/frantic-core/mockData"
)

// The function adjustSettlementForWeekends adjusts the input date to the next weekday if it falls on a Saturday or
// Sunday.
func adjustSettlementForWeekends(inDate time.Time) time.Time {
	if inDate.Weekday() == time.Saturday {
		inDate = inDate.AddDate(0, 0, 2)
	}
	if inDate.Weekday() == time.Sunday {
		inDate = inDate.AddDate(0, 0, 1)
	}
	return inDate
}

func getSettlementDaysCCY(ccy1 string) (int, error) {

	// Validate the two currencues using the mock package
	days1, err := mockData.GetCurrency(ccy1)
	if err != nil {
		logHandler.Error.Printf("error getting settlement days for currency [%s] [%v]", ccy1, err.Error())
		return -1, err
	}

	// Calculate the settlement days
	return days1.SpotDays, nil
}

func getSettlementDaysPAIR(ccy1 string, ccy2 string) (int, error) {

	// Validate the two currencues using the mock package
	days1, err := getSettlementDaysCCY(ccy1)
	if err != nil {
		logHandler.Error.Printf("error getting settlement days for currency [%s] [%v]", ccy1, err.Error())
		return -1, err
	}
	days2, err := getSettlementDaysCCY(ccy2)
	if err != nil {
		logHandler.Error.Printf("error getting settlement days for currency [%s] [%v]", ccy2, err.Error())
		return -1, err
	}

	// Calculate the settlement days
	return mathHelpers.Max(days1, days2), nil
}

func getSettlementDaysCROSS(ccy1 string, via string, ccy2 string) (int, error) {
	days1, err := getSettlementDaysPAIR(ccy1, via)
	if err != nil {
		logHandler.Error.Printf("error getting settlement days for currency [%s] [%v]", ccy1, err.Error())
		return -1, err
	}
	days2, err := getSettlementDaysPAIR(via, ccy2)
	if err != nil {
		logHandler.Error.Printf("error getting settlement days for currency [%s] [%v]", ccy2, err.Error())
		return -1, err
	}

	// Calculate the settlement days
	return mathHelpers.Max(days1, days2), nil
}

func getSettlementDateCCY(ccy1 string, tradeDate time.Time) (time.Time, error) {
	// Calculate the settlement days

	days, err := getSettlementDaysCCY(ccy1)
	if err != nil {
		logHandler.Error.Printf("error getting settlement days for currency [%s] [%v]", ccy1, err.Error())
		return time.Now(), err
	}

	// Adjust the date
	return adjustSettlementForWeekends(tradeDate.AddDate(0, 0, days)), nil
}

func getSettlementDatePAIR(ccy1 string, ccy2 string, tradeDate time.Time) (time.Time, error) {
	// Calculate the settlement days
	days, err := getSettlementDaysPAIR(ccy1, ccy2)
	if err != nil {
		logHandler.Error.Printf("error getting settlement days for currency [%s] [%v]", ccy1, err.Error())
		return time.Now(), err
	}

	// Adjust the date
	return adjustSettlementForWeekends(tradeDate.AddDate(0, 0, days)), nil
}

func getSettlementDateCROSS(ccy1 string, via string, ccy2 string, tradeDate time.Time) (time.Time, error) {
	// Calculate the settlement days
	days, err := getSettlementDaysCROSS(ccy1, via, ccy2)
	if err != nil {
		logHandler.Error.Printf("error getting settlement days for currency [%s] [%v]", ccy1, err.Error())
		return time.Now(), err
	}

	// Adjust the date
	return adjustSettlementForWeekends(tradeDate.AddDate(0, 0, days)), nil
}
