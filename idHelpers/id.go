package idHelpers

import (
	"crypto/sha3"
	"fmt"
	"strings"
	"time"

	"github.com/mt1976/frantic-core/commonConfig"
	ce "github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/htmlHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/stringHelpers"
	"github.com/segmentio/ksuid"
)

var cfg *commonConfig.Settings
var DELIMITER string
var SEP string = "⋮"

func init() {
	cfg = commonConfig.Get()
	DELIMITER = cfg.Delimiter()
	SEP = DELIMITER
}

// Encode returns a SHA3-256 hex digest of a path-safe version of the input.
// It trims spaces, removes internal spaces, converts to a path-safe string,
// and then hashes the result. On error, it logs and returns an empty string.
func Encode(in string) string {

	out := in
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

// Decode reverses a path-safe transformation applied by `htmlHelpers.ToPathSafe`.
// If decoding fails, it logs the error and returns an empty string.
func Decode(in string) string {
	in, err := htmlHelpers.FromPathSafe(in)
	if err != nil {
		logHandler.Error.Printf("error decoding string: %v", err.Error())
		return ""
	}
	return in
}

// GetUUID returns a legacy time-based identifier string comprising the current
// date/time, process UID (if available), and a random component. It is intended
// for human-readable tracing rather than cryptographic uniqueness.
func GetUUID() string {
	return GetUUIDv2()
	// Get a new UUID
	// Get TODAY and convert to string
	// today := time.Now().Format("060102-150405.000000")
	// today = today + ""
	// today = strings.Replace(today, ".", "-", -1)
	// //xx := shortuuid.New()
	// uid := 000000
	// if os.Getuid() > 0 {
	// 	uid = os.Getuid()
	// }

	// //ip, _ := get_IP()
	// //ip = strings.Replace(ip, ".", "", -1)
	// xx := rand.Intn(100000)
	// yy := fmt.Sprintf("%s-%04d-%06d", today, uid, xx)
	// yy = strings.Replace(yy, ".", "", -1)
	// yy = strings.Replace(yy, "-", "", -1)
	//yy = base64Encode(yy)

	//	logger.Info.Printf("[UUID] %v %v", yy, UUID2String(yy))

	// return yy
}

// UUID2String formats a legacy `GetUUID` value into a readable string showing
// segmented components and a human-friendly date/time breakdown.
// DEPRECATED: Use KSUID-based functions instead.
// func UUID2String(uuid string) string {
// 	// Convert UUID to string
// 	// 2407032122304271385011014720229731 convert to 240703\212230\427138\501\1014720229\731
// 	// 2407032122304271385011014720229731 convert to 240703.212230.427138.501.1014720229.731
// 	// 2407032122304271385011014720229731 convert to 240703-212230-427138-501-1014720229-731
// 	//logger.Info.Println("UID: UUID: ", uuid, len(uuid))

// 	fmtr := "%s" + SEP + "%s" + SEP + "%s" + SEP + "%s" + SEP + "%s"
// 	op := fmt.Sprintf(fmtr, uuid[0:6], uuid[6:12], uuid[12:18], uuid[18:24], uuid[24:])
// 	day, _ := time.Parse("060102150405", uuid[0:12])
// 	fmtr2 := "(Date=[%s]" + " " + "Time=[%s]" + " " + "ms=[%sms]" + " " + "uid=[%s]" + " " + "rnd=[%s])"
// 	op2 := fmt.Sprintf(fmtr2, dateHelpers.FormatHuman(day), day.Format("15:04:05"), uuid[12:18], strings.TrimLeft(uuid[18:24], "0"), uuid[24:])
// 	//logger.Info.Println("UID: String:", op, len(op))
// 	return op + ", " + op2
// }

// GetUUIDv2 generates a KSUID (K-Sortable Unique ID) and returns it as a string.
// KSUIDs are globally unique and time-sortable identifiers.
func GetUUIDv2() string {
	return ksuid.New().String()
}

// GetUUIDv2WithPayload generates a KSUID containing a fixed-size payload.
// The payload is right-padded to 16 bytes if shorter and must not exceed 16 bytes.
// Returns the KSUID string or an error if generation fails.
func GetUUIDv2WithPayload(payload string) (string, error) {
	// Ensure payload is 16 bytes
	length := 16
	if len(payload) > length {
		return "", ce.ErrIDGenerationWrapper(fmt.Errorf("payload must be %d bytes or less", length))
	}
	if len(payload) < 16 {
		payload = fmt.Sprintf("%-16s", payload)
	}
	ksuid, err := ksuid.FromParts(time.Now(), []byte(payload))
	if err != nil {
		logHandler.Error.Printf("Error generating KSUID: [%v]", err.Error())
		return "", ce.ErrIDGenerationWrapper(err)
	}
	return ksuid.String(), nil
}

// GetUUIDv2Payload extracts and returns the payload from a KSUID string.
// If parsing fails, it logs the error and returns an empty string.
func GetUUIDv2Payload(uuid string) string {
	ksuid, err := ksuid.Parse(uuid)
	if err != nil {
		logHandler.Error.Printf("Error generating KSUID: [%v]", err.Error())
		return ""
	}
	val := string(ksuid.Payload())
	val = strings.TrimLeft(strings.TrimRight(strings.Trim(val, " "), " "), " ")
	return val
}

// InspectUUIDv2 returns a human-readable summary of a KSUID including its time
// and payload contents. Returns an empty string if parsing fails.
func InspectUUIDv2(uuid string) string {
	ksuid, err := ksuid.Parse(uuid)
	if err != nil {
		logHandler.Error.Println("Error parsing KSUID:", err, " got:", len(uuid), " uuid", uuid)
		return ""
	}
	payload := ksuid.Payload()
	return fmt.Sprintf("Time: %v, Payload: %v", ksuid.Time(), string(payload))
}

// BuildCompositeID builds a composite identifier by sanitizing each part and
// joining them with the configured delimiter. It also returns the encoded
// (hashed) form of the composite ID.
func BuildCompositeID(part ...any) (compositeID string, encodedCompositeID string, err error) {

	logHandler.Trace.Printf("BuildCompositeID called with parts: %v", part)

	// Based on parts, build a composite ID, e.g., "part1::part2::part3"
	// parts must be convertible to string

	for i, p := range part {
		if i > 0 {
			compositeID += DELIMITER
		}
		compositeID += SanitizeID(fmt.Sprintf("%v", p))
	}

	encodedCompositeID = Encode(compositeID)
	logHandler.Trace.Printf("Built CompositeID: %v, Encoded: %v", compositeID, encodedCompositeID)
	return compositeID, encodedCompositeID, nil
}

// ParseCompositeID is a placeholder that will parse a composite ID back into
// its constituent parts. Currently returns an empty string.
func ParseCompositeID(compositeID string) string {
	// Based on compositeID, parse and return the parts
	// This is a placeholder implementation

	return ""
}

const exampleID = "5907763e7e4043e610d88bf2feae302a0ceb2578f923c4eed94004dbfdfba723"

// SanitizeID normalizes an identifier by camel-casing, removing special
// characters, trimming whitespace, and stripping underscores/hyphens.
// If the input appears to be a hash (example length), it is lowercased.
func SanitizeID(id string) string {
	orig := id
	logHandler.Trace.Printf("Sanitizing ID: '%v'", id)
	sanitizedID := stringHelpers.CamelCase(id)
	logHandler.Trace.Printf("CamelCased ID: '%v'", sanitizedID)
	sanitizedID = stringHelpers.RemoveSpecialChars(sanitizedID)
	logHandler.Trace.Printf("Removed Special Chars ID: '%v'", sanitizedID)
	sanitizedID = strings.TrimSpace(sanitizedID)
	logHandler.Trace.Printf("Trimmed ID: '%v'", sanitizedID)
	// cleanID = strings.ToLower(cleanID)
	// logHandler.Trace.Printf("Lowercased ID: %v", cleanID)
	sanitizedID = strings.ReplaceAll(sanitizedID, "_", "")
	logHandler.Trace.Printf("Final Sanitized ID: '%v'", sanitizedID)
	sanitizedID = strings.ReplaceAll(sanitizedID, "-", "")

	sanitizedID = strings.ReplaceAll(sanitizedID, ".local", "")
	logHandler.Trace.Printf("Removed .local ID: '%v'", sanitizedID)
	r := strings.NewReplacer("\n", "", "\r", "", "\t", "")
	sanitizedID = r.Replace(sanitizedID)
	logHandler.Trace.Println("Removed \n \r \t")

	sanitizedID = stringHelpers.RemoveSpecialCharacters(sanitizedID)
	logHandler.Trace.Printf("Stripped Special ID: '%v'", sanitizedID)

	if len(orig) == len(exampleID) {
		// Looks like a hash - make lowercase
		sanitizedID = strings.ToLower(sanitizedID)
		logHandler.Trace.Printf("Sanitized ID looks like a hash, lowercased to: '%v'", sanitizedID)
	}
	logHandler.Trace.Printf("Final Sanitized ID after removing hyphens: '%v'", sanitizedID)
	if sanitizedID == "" {
		logHandler.Trace.Printf("Sanitized ID is empty for original ID: '%v'", orig)
	}
	if orig == sanitizedID {
		logHandler.Trace.Printf("Sanitized ID is unchanged: '%v'", sanitizedID)
	} else {
		logHandler.Trace.Printf("Sanitized ID changed from '%v' to '%v'", orig, sanitizedID)
		if len(orig) == len(exampleID) {
			stringHelpers.CompareStrings(orig, sanitizedID)
		}
	}
	return sanitizedID
}
