package mockData

import (
	"fmt"
	"strings"

	"github.com/mt1976/frantic-core/logHandler"
)

type Rung struct {
	Code        string
	Name        string
	Alternative string
	Index       int
}

type Indexer struct {
	Index int
}

func (i *Indexer) inc() int {
	i.Index++
	return i.Index
}

var Ladder map[string]Rung
var LadderSize int

func init() {

	report("Tenor Ladder")

	Ladder = make(map[string]Rung)

	var idx Indexer

	Ladder["TD"] = Rung{Code: "TD", Name: "Today", Index: idx.inc()}

	Ladder["ON"] = Rung{Code: "ON", Name: "Overnight", Index: idx.inc()}
	//Ladder["TN"] = Rung{Code: "TN", Name: "Tom/Next", Index: 3}
	Ladder["SP"] = Rung{Code: "SP", Name: "Spot", Index: idx.inc()}
	Ladder["1W"] = Rung{Code: "1W", Name: "1 Week", Alternative: "7D", Index: idx.inc()}
	Ladder["2W"] = Rung{Code: "2W", Name: "2 Weeks", Index: idx.inc()}
	Ladder["3W"] = Rung{Code: "3W", Name: "3 Weeks", Index: idx.inc()}
	Ladder["1M"] = Rung{Code: "1M", Name: "1 Month", Alternative: "30D", Index: idx.inc()}
	Ladder["2M"] = Rung{Code: "2M", Name: "2 Months", Index: idx.inc()}
	Ladder["3M"] = Rung{Code: "3M", Name: "3 Months", Alternative: "90D", Index: idx.inc()}
	Ladder["4M"] = Rung{Code: "4M", Name: "4 Months", Index: idx.inc()}
	Ladder["5M"] = Rung{Code: "5M", Name: "5 Months", Index: idx.inc()}
	Ladder["6M"] = Rung{Code: "6M", Name: "6 Months", Alternative: "180D", Index: idx.inc()}
	Ladder["7M"] = Rung{Code: "7M", Name: "7 Months", Index: idx.inc()}
	Ladder["8M"] = Rung{Code: "8M", Name: "8 Months", Index: idx.inc()}
	Ladder["9M"] = Rung{Code: "9M", Name: "9 Months", Index: idx.inc()}
	Ladder["10M"] = Rung{Code: "10M", Name: "10 Months", Index: idx.inc()}
	Ladder["11M"] = Rung{Code: "11M", Name: "11 Months", Index: idx.inc()}
	Ladder["1Y"] = Rung{Code: "1Y", Name: "1 Year", Alternative: "12M", Index: idx.inc()}
	Ladder["15M"] = Rung{Code: "15M", Name: "13 Months", Index: idx.inc()}
	Ladder["18M"] = Rung{Code: "18M", Name: "18 Months", Index: idx.inc()}
	Ladder["21M"] = Rung{Code: "21M", Name: "21 Months", Index: idx.inc()}
	Ladder["2Y"] = Rung{Code: "2Y", Name: "2 Years", Index: idx.inc()}
	Ladder["3Y"] = Rung{Code: "3Y", Name: "3 Years", Index: idx.inc()}
	Ladder["4Y"] = Rung{Code: "4Y", Name: "4 Years", Index: idx.inc()}
	Ladder["5Y"] = Rung{Code: "5Y", Name: "5 Years", Index: idx.inc()}
	Ladder["6Y"] = Rung{Code: "6Y", Name: "6 Years", Index: idx.inc()}
	Ladder["7Y"] = Rung{Code: "7Y", Name: "7 Years", Index: idx.inc()}
	Ladder["8Y"] = Rung{Code: "8Y", Name: "8 Years", Index: idx.inc()}
	Ladder["9Y"] = Rung{Code: "9Y", Name: "9 Years", Index: idx.inc()}
	Ladder["10Y"] = Rung{Code: "10Y", Name: "10 Years", Index: idx.inc()}
	Ladder["11Y"] = Rung{Code: "11Y", Name: "11 Years", Index: idx.inc()}
	Ladder["12Y"] = Rung{Code: "12Y", Name: "12 Years", Index: idx.inc()}
	Ladder["99Y"] = Rung{Code: "99Y", Name: "99 Years", Index: idx.inc()}
	LadderSize = idx.inc()
}

func GetRateLadderList() []string {
	rtn := []string{}
	for k := range Ladder {
		rtn = append(rtn, k)
	}
	//rtn.sort()
	return rtn
}

func IsValidPeriod(in string) bool {
	_, ok := Ladder[in]
	return ok
}

func GetRateLadderByIndex(index int) Rung {
	for _, v := range Ladder {
		if v.Index == index {
			return v
		}
	}
	return Rung{}
}

func test() bool {
	noitems := len(Ladder)
	for i := 1; i <= noitems; i++ {
		rli := GetRateLadderByIndex(i)
		logHandler.Info.Printf("rate ladder info: rli=[%v] i=[%v]\n", rli, i)
	}
	return true
}

func LadderToString(R map[string]Rung) string {
	var output strings.Builder
	//	noItems := len(R)
	noitems := len(Ladder)
	for i := 1; i <= noitems; i++ {
		rli := GetRateLadderByIndex(i)
		// add to output
		output.WriteString(fmt.Sprintf("%v:%v,", rli.Code, rli.Index))
	}

	return output.String()
}

func GetTenorInfo(tenor string) (Rung, error) {
	return Ladder[tenor], nil
}
