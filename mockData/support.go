package mockData

import (
	"github.com/mt1976/frantic-core/logHandler"
)

func report(in string) {
	logHandler.Info.Printf("Mocking - %s\n", in)
}
