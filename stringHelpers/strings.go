package stringHelpers

import (
	"encoding/base64"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/galsondor/go-ascii"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/rivo/uniseg"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	// specialChars is a list of special characters that are not allowed
	specialChars = "[^A-Za-z0-9]+"
	// wildcardOpen is the open wildcard
	wildcardOpen = "{{"
	// wildcardClose is the close wildcard
	wildcardClose = "}}"
)

// LowerFirst lowers the first character of a string
func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

/*
* leftPad and rightPad just repoeat the padStr the indicated
* number of times
*
 */
// leftPad pads a string on the left - OBSOLETE
// Deprecated: Use padLeft instead
func leftPad(s string, padStr string, pLen int) string {
	return padLeft(s, padStr, pLen)
}

// rightPad pads a string on the right - OBSOLETE
// Deprecated: Use padRight instead
func rightPad(s string, padStr string, pLen int) string {
	return padRight(s, padStr, pLen)
}

/* the Pad2Len functions are generally assumed to be padded with short sequences of strings
* in many cases with a single character sequence
*
* so we assume we can build the string out as if the char seq is 1 char and then
* just substr the string if it is longer than needed
*
* this means we are wasting some cpu and memory work
* but this always get us to want we want it to be
*
* in short not optimized to for massive string work
*
* If the overallLen is shorter than the original string length
* the string will be shortened to this length (substr)
*
 */
// rightPad2Len pads a string on the right to a given length - if the string is longer than the length it is shortened
func rightPad2Len(s string, padStr string, overallLen int) string {
	padCountInt := 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

// leftPad2Len pads a string on the left to a given length - if the string is longer than the length it is shortened
func leftPad2Len(s string, padStr string, overallLen int) string {
	padCountInt := 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

// PadRight pads a string on the right
func padRight(s string, p string, l int) string {
	return rightPad2Len(s, p, l)
}

// PadLeft pads a string on the left
func padLeft(s string, p string, l int) string {
	return leftPad2Len(s, p, l)
}

// dQuote adds double quotes to a string
func dQuote(s string) string {
	return "\"" + s + "\""
}

// sQuote adds single quotes to a string
func sQuote(s string) string {
	return "'" + s + "'"
}

// sBracket adds square brackets to a string
func sBracket(s string) string {
	return "[" + s + "]"
}

func dBracket(s string) string {
	return "[[" + s + "]]"
}

// dChevrons wraps the string in double chevrons
func dChevrons(s string) string {
	return "<<" + s + ">>"
}

// sChevrons wraps the string in single chevrons
func sChevrons(s string) string {
	return "<" + s + ">"
}

// DCurlies wraps the string in double curlies
func dCurlies(s string) string {
	return "{{" + s + "}}"
}

// SCurlies wraps the string in single curlies
func sCurlies(s string) string {
	return "{" + s + "}"
}

// dParentheses wraps the string in double Parentheses
func dParentheses(s string) string {
	return "((" + s + "))"
}

// sParentheses wraps the string in single Parentheses
func sParentheses(s string) string {
	return "(" + s + ")"
}

// strArrayToStringWithSep converts a string array to a string using a given separator
func strArrayToStringWithSep(inArray []string, inSep string) string {

	var outString strings.Builder
	noRows := len(inArray)
	for ii := range noRows {
		outString.WriteString(inArray[ii] + inSep)
	}
	return outString.String()
}

// removeSpecialChars removes special characters from a string and replaces them with a dash
func removeSpecialChars(in string) string {
	reg, err := regexp.Compile(specialChars)
	if err != nil {
		logHandler.Error.Printf("Error=[%v]", err.Error())
	}
	newStr := reg.ReplaceAllString(in, "-")
	return newStr
}

// replaceWildcard replaces a wildcard {{wildcard}} with a given string
func replaceWildcard(orig string, replaceThis string, withThis string) string {
	wrkThis := wildcardOpen + replaceThis + wildcardClose
	//log.Printf("Replace %s with %q", wrkThis, withThis)
	return strings.ReplaceAll(orig, wrkThis, withThis)
}

// Encode encodes a string to base64
func encode(rawStr string) string {

	// base64.StdEncoding: Standard encoding with padding
	// It requires a byte slice so we cast the string to []byte
	encodedStr := base64.URLEncoding.EncodeToString([]byte(rawStr))

	return encodedStr
}

// Decode decodes a base64 encoded string
func decode(encodedStr string) string {
	decodedStr, err := base64.URLEncoding.DecodeString(encodedStr)
	if err != nil {
		logHandler.Warning.Printf("Error=[%v]", err.Error())
	}

	return string(decodedStr)
}

var cr = "{cr}"
var lf = "{lf}"

func makeStorable(in string) string {
	s := strings.ReplaceAll(in, string(rune(ascii.CR)), cr)
	s = strings.ReplaceAll(s, string(rune(ascii.LF)), lf)
	return s
}

func makeDisplayable(in string) string {
	s := strings.ReplaceAll(in, cr, string(rune(ascii.CR)))
	s = strings.ReplaceAll(s, lf, string(rune(ascii.LF)))
	return s
}

func CamelCase(s string) string {

	// Remove all characters that are not alphanumeric or spaces or underscores
	s = regexp.MustCompile("[^a-zA-Z0-9_ ]+").ReplaceAllString(s, "")

	// Replace all underscores with spaces
	s = strings.ReplaceAll(s, "_", " ")

	// Title case s
	s = cases.Title(language.BritishEnglish, cases.NoLower).String(s)

	// Remove all spaces
	s = strings.ReplaceAll(s, " ", "")

	// Lowercase the first letter
	if len(s) > 0 {
		s = strings.ToLower(s[:1]) + s[1:]
	}

	return s
}

func CompareStrings(a, b string) {
	linesA := strings.Split(a, "\n")
	linesB := strings.Split(b, "\n")

	max := len(linesA)
	if len(linesB) > max {
		max = len(linesB)
	}

	for i := 0; i < max; i++ {
		var lineA, lineB string

		if i < len(linesA) {
			lineA = linesA[i]
		}
		if i < len(linesB) {
			lineB = linesB[i]
		}

		// If identical, skip
		if lineA == lineB {
			continue
		}

		// If one line is missing
		switch {
		case lineA == "" && lineB != "":
			logHandler.Warning.Printf("[Line %d] Added:    %q", i+1, lineB)
		case lineA != "" && lineB == "":
			logHandler.Warning.Printf("[Line %d] Removed:  %q", i+1, lineA)
		default:
			logHandler.Warning.Printf("[Line %d] Changed:", i+1)
			logHandler.Warning.Printf("   - %q", lineA)
			logHandler.Warning.Printf("   + %q", lineB)
		}
	}
}

func RemoveSpecialCharacters(s string) string {
	var b strings.Builder

	for gr := uniseg.NewGraphemes(s); gr.Next(); {

		if r := gr.Runes()[0]; unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteString(gr.Str())
		}
	}
	return b.String()
}
