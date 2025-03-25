package timehelpers

import (
	"fmt"
	"strings"
	"time"
)

var regionToTimezone = map[string]string{
	"US": "America/New_York",
	"GB": "Europe/London",
	"DE": "Europe/Berlin",
	"FR": "Europe/Paris",
	"IN": "Asia/Kolkata",
	"JP": "Asia/Tokyo",
	"CN": "Asia/Shanghai",
	"AU": "Australia/Sydney",
	"CA": "America/Toronto",
	"BR": "America/Sao_Paulo",
	"RU": "Europe/Moscow",
	"KR": "Asia/Seoul",
	"IT": "Europe/Rome",
	"ES": "Europe/Madrid",
	"MX": "America/Mexico_City",
	"ID": "Asia/Jakarta",
	"TR": "Europe/Istanbul",
	"NL": "Europe/Amsterdam",
	"CH": "Europe/Zurich",
	"SE": "Europe/Stockholm",
	"SA": "Asia/Riyadh",
	"AE": "Asia/Dubai",
	"SG": "Asia/Singapore",
	"PL": "Europe/Warsaw",
	"TH": "Asia/Bangkok",
	"MY": "Asia/Kuala_Lumpur",

	// Add more as needed
}

func InferTimezoneFromLocale(locale string) (string, error) {
	locale = strings.ToUpper(locale)
	locale = strings.ReplaceAll(locale, "-", "_")
	parts := strings.Split(locale, "_")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid locale format")
	}
	region := parts[1]
	tzName, ok := regionToTimezone[region]
	if !ok {
		return "", fmt.Errorf("unsupported region code: %s", region)
	}
	loc, err := time.LoadLocation(tzName)
	if err != nil {
		return "", fmt.Errorf("failed to load timezone: %v", err)
	}
	return loc.String(), nil
}
