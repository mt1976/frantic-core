package stringHelpers

import (
	"crypto/sha3"
	"fmt"
	"strings"

	"github.com/mt1976/frantic-core/htmlHelpers"
	"github.com/mt1976/frantic-core/logHandler"
)

// Lowers the first character of a string
func LowerFirst(s string) string {
	return lowerFirst(s)
}

// ArrToString converts an array of strings to a printable string
func ArrToString(strArray []string) string {
	return strings.Join(strArray, "\n")
}

// StrArrayToString converst a string array into a string
func StrArrayToString(inArray []string) string {
	return strArrayToStringWithSep(inArray, "\n")
}

// StrArrayToStringWithSep converts a string array to a string using a given separator
func StrArrayToStringWithSep(inArray []string, inSep string) string {
	return strArrayToStringWithSep(inArray, inSep)
}

// RemoveSpecialChars removes special characters from a string and replaces them with a dash
func RemoveSpecialChars(in string) string {
	return removeSpecialChars(in)
}

// ReplaceWildcard replaces a wildcard {{wildcard}} with a given string
func ReplaceWildcard(orig string, replaceThis string, withThis string) string {
	return replaceWildcard(orig, replaceThis, withThis)
}

// PadRight pads a string on the right
func PadRight(s string, p string, l int) string {
	return padRight(s, p, l)
}

// PadLeft pads a string on the left
func PadLeft(s string, p string, l int) string {
	return padLeft(s, p, l)
}

// Encode encodes a string to base64
func Encode(rawStr string) string {
	// Encode a string to base64
	out := rawStr
	out = strings.Replace(out, " ", "", -1)
	out = strings.Trim(out, " ")
	out, err := htmlHelpers.ToPathSafe(out)
	if err != nil {
		logHandler.Error.Printf("error encoding string: %v", err.Error())
		return ""
	}

	z := sha3.Sum256([]byte(out))
	out = fmt.Sprintf("%x", z)

	return out
}

// DQuote wraps the string in double quotes - ""
func DQuote(s string) string {
	return dQuote(s)
}

// SQuote wraps the string in single quotes - ”
func SQuote(s string) string {
	return sQuote(s)
}

// DBracket wraps the string in double brackets - [[]]
func DBracket(s string) string {
	return dBracket(s)
}

// SBracket wraps the string in square brackets - []
func SBracket(s string) string {
	return sBracket(s)
}

// DChevrons wraps the string in double chevrons - <<>>
func DChevrons(s string) string {
	return dChevrons(s)
}

// SChevrons wraps the string in single chevrons - <>
func SChevrons(s string) string {
	return sChevrons(s)
}

// DCurlies wraps the string in double curlies - {{}}
func DCurlies(s string) string {
	return dCurlies(s)
}

// SCurlies wraps the string in single curlies - {}
func SCurlies(s string) string {
	return sCurlies(s)
}

// DParentheses wraps the string in double Parentheses - (())
func DParentheses(s string) string {
	return dParentheses(s)
}

// SParentheses wraps the string in single Parentheses - ()
func SParentheses(s string) string {
	return sParentheses(s)
}

// The function "MakeStringStorable" takes a string as input and returns a storable version of the
// string. Replaces ascii cr/lf with {{cr}}/{{lf}}.
func MakeStringStorable(in string) string {
	return makeStorable(in)
}

// The function "MakeStringDisplayable" takes a string as input and returns a modified version of the
// string that is displayable. Replaces {{cr}}/{{lf}} with ascii cr/lf.
func MakeStringDisplayable(in string) string {
	return makeDisplayable(in)
}
