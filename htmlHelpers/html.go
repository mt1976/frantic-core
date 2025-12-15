package htmlHelpers

import (
	b64 "encoding/base64"
	"net/url"
	"strconv"
	"strings"

	"github.com/mt1976/frantic-core/logHandler"
)

var name = "HTML"

func ValueToInt(s string) int {
	if s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		logHandler.WarningLogger.Printf("[%v] [Math] Error=[%v]", strings.ToUpper(name), err.Error())
		return 999999999
	}
	return i
}

func ValueToBool(s string) bool {
	return s == "on"
}

func BoolToValue(b bool) string {
	if b {
		return "on"
	}
	return "off"
}

func ToPathSafe(s string) (string, error) {
	r := url.PathEscape(s)
	sEnc := b64.StdEncoding.EncodeToString([]byte(r))
	return sEnc, nil
}

func FromPathSafe(s string) (string, error) {

	uDec, err := b64.URLEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	r, err := url.PathUnescape(string(uDec))
	if err != nil {
		return "", err
	}

	return r, nil
}
