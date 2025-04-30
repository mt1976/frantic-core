package dateHelpers

import (
	"strings"
	"time"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/logHandler"
)

var name = "DATE"
var Format DateFormat
var cfg *commonConfig.Settings

type DateFormat struct {
	External     string
	DMY          string
	Internal     string
	Detail       string
	YMD          string
	Calendar     string
	BackupDate   string
	BackupFolder string
}

func init() {
	cfg = commonConfig.Get()

	Format.External = cfg.GetDateFormat_Human()
	Format.DMY = cfg.GetDateFormat_DMY2()
	Format.Internal = cfg.GetDateFormat_Internal()
	Format.Detail = cfg.GetDateFormat_DateTime()
	Format.YMD = cfg.GetDateFormat_YMD()
	Format.Calendar = "2006-01-02T15:04:05"
	Format.BackupDate = cfg.GetDateFormat_Backup()
	Format.BackupFolder = cfg.GetDateFormat_BackupDirectory()
}

// func FormatYMD(in time.Time) string {
// 	return in.Format(Format.YMD)
// }

func FormatAudit(in time.Time) string {
	return in.Format(Format.Detail)
}

func FormatDMY(in time.Time) string {
	if in.IsZero() {
		return ""
	}
	return in.Format(Format.DMY)
}

func FormatYMD(in time.Time) string {
	if in.IsZero() {
		return ""
	}
	return in.Format(Format.YMD)
}

func FormatCalendar(in time.Time) string {
	if in.IsZero() {
		return ""
	}
	return in.Format(Format.Calendar)
}

func FormatHumanFromString(in string) (time.Time, error) {
	if in == "" {
		return time.Time{}, nil
	}
	return time.Parse(Format.External, in)
}

func FormatHuman(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(Format.External)
}

func StartOfDay(t time.Time) time.Time {
	// Purpose: To remove the time from a date
	w := t.Format(Format.DMY)
	r, err := time.Parse(Format.DMY, w)
	if err != nil {
		logHandler.WarningLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
		return t
	}
	if cfg.IsApplicationMode(commonConfig.MODE_DEVELOPMENT) {
		logHandler.InfoLogger.Printf("[%v] [DateStartOfDay] Date=[%v] Result=[%v]", strings.ToUpper(name), t, r)
	}
	return r
}

func EndOfDay(t time.Time) time.Time {
	// Purpose: To remove the time from a date
	w := t.Format(Format.DMY)
	r, err := time.Parse(Format.DMY, w)
	if err != nil {
		logHandler.WarningLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
		return t
	}
	r = r.AddDate(0, 0, 1)
	r = r.Add(-time.Second)
	if cfg.IsApplicationMode(commonConfig.MODE_DEVELOPMENT) {
		logHandler.InfoLogger.Printf("[%v] [DateEndOfDay] Date=[%v] Result=[%v]", strings.ToUpper(name), t, r)
	}
	return r
}

func IsBeforeOrEqualTo(t1, t2 time.Time) bool {
	// Purpose: check if a time is before or equal to a time
	if cfg.IsApplicationMode(commonConfig.MODE_DEVELOPMENT) {
		//	logger.InfoLogger.Printf("HLP: [HELPER] Date=[%v] Check=[%v]", DateStartOfDay(t1), DateStartOfDay(t2))
	}
	check := StartOfDay(t1)
	if check.Before(StartOfDay(t2)) || check.Equal(StartOfDay(t2)) {
		//	logger.InfoLogger.Printf("HLP: [HELPER] Date=[%v] Check=[%v] Result=[%v]", DateStartOfDay(t1), DateStartOfDay(t2), true)
		return true
	}
	//logger.InfoLogger.Printf("HLP: [HELPER] Date=[%v] Check=[%v] Result=[%v]", DateStartOfDay(t1), DateStartOfDay(t2), false)
	return false
}

func IsAfterOrEqualTo(t1, t2 time.Time) bool {
	// Purpose: check if a time is after or equal to a time
	if t1.After(StartOfDay(t2)) || t1.Equal(StartOfDay(t2)) {
		return true
	}
	return false
}

// AddWorkingDays adds n working days (excluding weekends) to a given date
func AddWorkingDays(startDate time.Time, daysToAdd int) time.Time {
	currentDate := startDate
	addedDays := 0

	for addedDays < daysToAdd {
		currentDate = currentDate.AddDate(0, 0, 1) // Add one day

		// Skip weekends
		if IsWorkingDay(currentDate) {
			addedDays++
		}
	}
	return currentDate
}

func IsWorkingDay(t time.Time) bool {
	// Purpose: check if a date is a working day
	if t.Weekday() == time.Saturday || t.Weekday() == time.Sunday {
		return false
	}
	return true
}

// AdjustToNextWorkingday moves the date forward to the next weekday if it's a weekend
func AdjustToNextWorkingday(date time.Time) time.Time {
	switch date.Weekday() {
	case time.Saturday:
		return date.AddDate(0, 0, 2) // Move to Monday
	case time.Sunday:
		return date.AddDate(0, 0, 1) // Move to Monday
	default:
		return date
	}
}

// AddMonthAndAdjust adds 1 calendar month and adjusts to the next weekday if needed
func AddMonthAndAdjust(startDate time.Time, noMonths int) time.Time {
	datePlusMonth := startDate.AddDate(0, noMonths, 0)
	return AdjustToNextWorkingday(datePlusMonth)
}
