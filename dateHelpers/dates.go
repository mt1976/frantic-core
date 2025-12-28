// Package dateHelpers provides utilities for date/time formatting, parsing,
// arithmetic, and business-day adjustments used across the project.
package dateHelpers

import (
	"fmt"
	"strings"
	"time"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/logHandler"
)

var name = "DATE"

// Format contains the active set of date/time layout strings loaded from
// application configuration. These layouts are used by the helper functions
// for formatting and parsing dates in consistent forms.
var Format DateFormat
var cfg *commonConfig.Settings

// DateFormat defines the collection of Go time layouts used by this package
// for formatting/parsing dates in various contexts.
type DateFormat struct {
	External     string
	DMY          string
	Internal     string
	Detail       string
	YMD          string
	Calendar     string
	BackupDate   string
	BackupFolder string
	HTMLInput    string
	DMY4         string
	MYH          string
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
	Format.HTMLInput = "02/01/2006" // 14/12/2025
	Format.DMY4 = "02/01/2006"
	Format.MYH = "Jan 2006"
}

// func FormatYMD(in time.Time) string {
// 	return in.Format(Format.YMD)
// }

// FormatAudit returns the formatted date/time using the audit/detail layout.
func FormatAudit(in time.Time) string {
	return in.Format(Format.Detail)
}

// ParseAudit parses a date/time string using the audit/detail layout.
func ParseAudit(in string) (time.Time, error) {
	return time.Parse(Format.Detail, in)
}

// FormatDMY formats a time as day/month/year using the configured DMY layout.
func FormatDMY(in time.Time) string {
	if in.IsZero() {
		return ""
	}
	return in.Format(Format.DMY)
}

// ParseDMY parses a day/month/year string using the configured DMY layout.
func ParseDMY(in string) (time.Time, error) {
	if in == "" {
		return time.Time{}, nil
	}
	return time.Parse(Format.DMY, in)
}

// FormatYMD formats a time as year-month-day using the configured YMD layout.
func FormatYMD(in time.Time) string {
	if in.IsZero() {
		return ""
	}
	return in.Format(Format.YMD)
}

// ParseYMD parses a year-month-day string using the configured YMD layout.
func ParseYMD(in string) (time.Time, error) {
	if in == "" {
		return time.Time{}, nil
	}
	return time.Parse(Format.YMD, in)
}

// FormatCalendar formats the time using the calendar (ISO-like) layout.
func FormatCalendar(in time.Time) string {
	if in.IsZero() {
		return ""
	}
	return in.Format(Format.Calendar)
}

// ParseCalendar parses a date/time string using the calendar layout.
func ParseCalendar(in string) (time.Time, error) {
	if in == "" {
		return time.Time{}, nil
	}
	return time.Parse(Format.Calendar, in)
}

// FormatHumanFromString parses a human-readable date/time string using the
// External layout. This is functionally equivalent to ParseHumanFromString.
func FormatHumanFromString(in string) (time.Time, error) {
	if in == "" {
		return time.Time{}, nil
	}
	return time.Parse(Format.External, in)
}

// ParseHumanFromString parses a human-readable date/time string using the
// External layout.
func ParseHumanFromString(in string) (time.Time, error) {
	if in == "" {
		return time.Time{}, nil
	}
	return time.Parse(Format.External, in)
}

// FormatHuman returns the human-friendly formatted date/time string using the
// External layout.
func FormatHuman(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(Format.External)
}

// ParseHuman parses a human-friendly date/time string using the External
// layout.
func ParseHuman(in string) (time.Time, error) {
	if in == "" {
		return time.Time{}, nil
	}
	return time.Parse(Format.External, in)
}

// FormatInternal returns the internal/system formatted date/time string.
func FormatInternal(in time.Time) string {
	if in.IsZero() {
		return ""
	}
	return in.Format(Format.Internal)
}

// ParseInternal parses a date/time string using the internal/system layout.
func ParseInternal(in string) (time.Time, error) {
	if in == "" {
		return time.Time{}, nil
	}
	return time.Parse(Format.Internal, in)
}

// FormatHTMLInput formats a time for HTML input controls (e.g. dd/mm/yyyy).
func FormatHTMLInput(in time.Time) string {
	if in.IsZero() {
		return ""
	}
	return in.Format(Format.HTMLInput)
}

// ParseHTMLInput parses a string formatted for HTML input controls.
func ParseHTMLInput(in string) (time.Time, error) {
	if in == "" {
		return time.Time{}, nil
	}
	return time.Parse(Format.HTMLInput, in)
}

// FormatDMY4 formats a time as dd/mm/yyyy using the configured DMY4 layout.
func FormatDMY4(in time.Time) string {
	if in.IsZero() {
		return ""
	}
	return in.Format(Format.DMY4)
}

// ParseDMY4 parses a dd/mm/yyyy string using the configured DMY4 layout.
func ParseDMY4(in string) (time.Time, error) {
	if in == "" {
		return time.Time{}, nil
	}
	return time.Parse(Format.DMY4, in)
}

// FormatDetail returns a detailed date/time string including time components.
func FormatDetail(in time.Time) string {
	if in.IsZero() {
		return ""
	}
	return in.Format(Format.Detail)
}

// ParseDetail parses a detailed date/time string using the Detail layout.
func ParseDetail(in string) (time.Time, error) {
	if in == "" {
		return time.Time{}, nil
	}
	return time.Parse(Format.Detail, in)
}

// FormatExternal returns a human/external formatted date/time string.
func FormatExternal(in time.Time) string {
	if in.IsZero() {
		return ""
	}
	return in.Format(Format.External)
}

// ParseExternal parses a human/external date/time string.
func ParseExternal(in string) (time.Time, error) {
	if in == "" {
		return time.Time{}, nil
	}
	return time.Parse(Format.External, in)
}

// FormatBackupDate formats a time for backup naming (date component).
func FormatBackupDate(in time.Time) string {
	if in.IsZero() {
		return ""
	}
	return in.Format(Format.BackupDate)
}

// ParseBackupDate parses a backup date string.
func ParseBackupDate(in string) (time.Time, error) {
	if in == "" {
		return time.Time{}, nil
	}
	return time.Parse(Format.BackupDate, in)
}

// FormatBackupFolder formats a time for backup folder naming.
func FormatBackupFolder(in time.Time) string {
	if in.IsZero() {
		return ""
	}
	return in.Format(Format.BackupFolder)
}

// ParseBackupFolder parses a backup folder date/time string.
func ParseBackupFolder(in string) (time.Time, error) {
	if in == "" {
		return time.Time{}, nil
	}
	return time.Parse(Format.BackupFolder, in)
}

// IsSameDay reports whether two times fall on the same calendar day.
func IsSameDay(t1, t2 time.Time) bool {
	return StartOfDay(t1).Equal(StartOfDay(t2))
}

// IsTomorrow reports whether the given time is tomorrow (relative to now).
func IsTomorrow(t time.Time) bool {
	tomorrow := time.Now().AddDate(0, 0, 1)
	return IsSameDay(t, tomorrow)
}

// IsYesterday reports whether the given time is yesterday (relative to now).
func IsYesterday(t time.Time) bool {
	yesterday := time.Now().AddDate(0, 0, -1)
	return IsSameDay(t, yesterday)
}

// IsToday reports whether the given time is today (relative to now).
func IsToday(t time.Time) bool {
	return IsSameDay(t, time.Now())
}

// Today returns the current date truncated to the start of day.
func Today() time.Time {
	return StartOfDay(time.Now())
}

// Tomorrow returns the time representing 24 hours from now.
func Tomorrow() time.Time {
	return time.Now().AddDate(0, 0, 1)
}

// StartOfDay returns the time truncated to midnight in the local timezone.
func StartOfDay(t time.Time) time.Time {
	w := t.Format(Format.DMY)
	r, err := time.Parse(Format.DMY, w)
	if err != nil {
		logHandler.WarningLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
		return t
	}
	if cfg.IsApplicationMode(commonConfig.MODE_DEVELOPMENT) {
		logHandler.TraceLogger.Printf("[%v] [DateStartOfDay] Date=[%v] Result=[%v]", strings.ToUpper(name), t, r)
	}
	return r
}

// EndOfDay returns the final second of the given date (23:59:59).
func EndOfDay(t time.Time) time.Time {
	w := t.Format(Format.DMY)
	r, err := time.Parse(Format.DMY, w)
	if err != nil {
		logHandler.WarningLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
		return t
	}
	r = r.AddDate(0, 0, 1)
	r = r.Add(-time.Second)
	if cfg.IsApplicationMode(commonConfig.MODE_DEVELOPMENT) {
		logHandler.TraceLogger.Printf("[%v] [DateEndOfDay] Date=[%v] Result=[%v]", strings.ToUpper(name), t, r)
	}
	return r
}

// IsBeforeOrEqualTo reports whether t1 is strictly before or equal to t2
// when both are compared at their respective start-of-day values.
func IsBeforeOrEqualTo(t1, t2 time.Time) bool {
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

// IsAfterOrEqualTo reports whether t1 is strictly after or equal to t2
// when compared at the start of day.
func IsAfterOrEqualTo(t1, t2 time.Time) bool {
	if t1.After(StartOfDay(t2)) || t1.Equal(StartOfDay(t2)) {
		return true
	}
	return false
}

func IsBetweenDates(testDate, startDate, endDate time.Time) bool {
	if IsAfterOrEqualTo(testDate, startDate) && IsBeforeOrEqualTo(testDate, endDate) {
		return true
	}
	return false
}

func IsWorkingDateBetween(testDate, startDate, endDate time.Time) bool {
	if !IsWorkingDay(testDate) {
		return false
	}
	if IsAfterOrEqualTo(testDate, startDate) && IsBeforeOrEqualTo(testDate, endDate) {
		return true
	}
	return false
}

func IsWeekend(t time.Time) bool {
	if IsSaturday(t) || IsSunday(t) {
		return true
	}
	return false
}

func IsWeekday(t time.Time) bool {
	if t.Weekday() >= time.Monday && t.Weekday() <= time.Friday {
		return true
	}
	return false
}

func IsMonday(t time.Time) bool {
	return t.Weekday() == time.Monday
}

func IsFriday(t time.Time) bool {
	return t.Weekday() == time.Friday
}

func IsSaturday(t time.Time) bool {
	return t.Weekday() == time.Saturday
}

func IsSunday(t time.Time) bool {
	return t.Weekday() == time.Sunday
}

func IsTuesday(t time.Time) bool {
	return t.Weekday() == time.Tuesday
}

func IsWednesday(t time.Time) bool {
	return t.Weekday() == time.Wednesday
}

func IsThursday(t time.Time) bool {
	return t.Weekday() == time.Thursday
}

func IsAfter(t1, t2 time.Time) bool {
	return t1.After(t2)
}

func IsBefore(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

func IsAfterDays(t1, t2 time.Time, days int) bool {
	t1sod := StartOfDay(t1)
	t2sod := StartOfDay(t2)
	return t1sod.After(t2sod)
}

func IsBeforeDays(t1, t2 time.Time, days int) bool {
	t1sod := StartOfDay(t1)
	t2sod := StartOfDay(t2)
	return t1sod.Before(t2sod)
}

func IsInXDays(t1, t2 time.Time, days int) bool {
	t1sod := StartOfDay(t1)
	t2sod := StartOfDay(t2)
	diff := t1sod.Sub(t2sod)
	return diff.Hours() <= float64(days*24)
}

func IsNotInXDays(t1, t2 time.Time, days int) bool {
	return !IsInXDays(t1, t2, days)
}

func DifferenceInDays(t1, t2 time.Time) int {
	t1sod := StartOfDay(t1)
	t2sod := StartOfDay(t2)
	diff := t1sod.Sub(t2sod)
	return int(diff.Hours() / 24)
}

// AddDays returns the time that is daysToAdd days after startDate.
func AddDays(startDate time.Time, daysToAdd int) time.Time {
	return startDate.AddDate(0, 0, daysToAdd)
}

// AddWorkingDays adds n working days (excluding weekends) to a given date.
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

func AddWeeks(startDate time.Time, weeksToAdd int) time.Time {
	return startDate.AddDate(0, 0, weeksToAdd*7)
}

func AddWorkingWeeks(startDate time.Time, weeksToAdd int) time.Time {
	daysToAdd := weeksToAdd * 5
	return AddWorkingDays(startDate, daysToAdd)
}

// AddMonths returns the time that is monthsToAdd months after startDate.
func AddMonths(startDate time.Time, monthsToAdd int) time.Time {
	return startDate.AddDate(0, monthsToAdd, 0)
}

func AddWorkingMonths(startDate time.Time, monthsToAdd int) time.Time {
	currentDate := startDate

	for i := 0; i < monthsToAdd; i++ {
		currentDate = currentDate.AddDate(0, 1, 0) // Add one month
		currentDate = AdjustToNextWorkingday(currentDate)
	}
	return currentDate
}

func AddYears(startDate time.Time, yearsToAdd int) time.Time {
	return startDate.AddDate(yearsToAdd, 0, 0)
}

func AddWorkingYears(startDate time.Time, yearsToAdd int) time.Time {
	currentDate := startDate

	for i := 0; i < yearsToAdd; i++ {
		currentDate = currentDate.AddDate(1, 0, 0) // Add one year
		currentDate = AdjustToNextWorkingday(currentDate)
	}
	return currentDate
}

func SubtractDays(startDate time.Time, daysToSubtract int) time.Time {
	return startDate.AddDate(0, 0, -daysToSubtract)
}

func SubtractWorkingDays(startDate time.Time, daysToSubtract int) time.Time {
	currentDate := startDate
	subtractedDays := 0

	for subtractedDays < daysToSubtract {
		currentDate = currentDate.AddDate(0, 0, -1) // Subtract one day

		// Skip weekends
		if IsWorkingDay(currentDate) {
			subtractedDays++
		}
	}
	return currentDate
}

func SubtractWeeks(startDate time.Time, weeksToSubtract int) time.Time {
	return startDate.AddDate(0, 0, -weeksToSubtract*7)
}

func SubtractWorkingWeeks(startDate time.Time, weeksToSubtract int) time.Time {
	daysToSubtract := weeksToSubtract * 5
	return SubtractWorkingDays(startDate, daysToSubtract)
}

func SubtractMonths(startDate time.Time, monthsToSubtract int) time.Time {
	return startDate.AddDate(0, -monthsToSubtract, 0)
}
func SubtractWorkingMonths(startDate time.Time, monthsToSubtract int) time.Time {
	currentDate := startDate

	for i := 0; i < monthsToSubtract; i++ {
		currentDate = currentDate.AddDate(0, -1, 0) // Subtract one month
		currentDate = AdjustToPreviousWorkingday(currentDate)
	}
	return currentDate
}

// SubtractYears returns the time that is yearsToSubtract years before startDate.
func SubtractYears(startDate time.Time, yearsToSubtract int) time.Time {
	return startDate.AddDate(-yearsToSubtract, 0, 0)
}

// SubtractWorkingYears subtracts full years and adjusts to the previous working day if needed.
func SubtractWorkingYears(startDate time.Time, yearsToSubtract int) time.Time {
	currentDate := startDate

	for i := 0; i < yearsToSubtract; i++ {
		currentDate = currentDate.AddDate(-1, 0, 0) // Subtract one year
		currentDate = AdjustToPreviousWorkingday(currentDate)
	}
	return currentDate
}

// IsWorkingDay reports whether the date falls on a weekday (Mon-Fri).
func IsWorkingDay(t time.Time) bool {
	if t.Weekday() == time.Saturday || t.Weekday() == time.Sunday {
		return false
	}
	return true
}

// NextWorkingDay returns the next working day after today.
func NextWorkingDay() time.Time {
	nextDay := Tomorrow()
	if IsWorkingDay(nextDay) {
		return nextDay
	}
	return AdjustToNextWorkingday(nextDay)
}

// Yesterday returns the time representing 24 hours before now.
func Yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

// PreviousWorkingDay returns the previous working day before today.
func PreviousWorkingDay() time.Time {
	previousDay := Yesterday()
	if IsWorkingDay(previousDay) {
		return previousDay
	}
	return AdjustToPreviousWorkingday(previousDay)
}

// AdjustToNextWorkingday moves the date forward to the next weekday if it's a weekend.
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

// AdjustToPreviousWorkingday moves the date backward to the previous weekday if it's a weekend.
func AdjustToPreviousWorkingday(date time.Time) time.Time {
	switch date.Weekday() {
	case time.Saturday:
		return date.AddDate(0, 0, -1) // Move to Friday
	case time.Sunday:
		return date.AddDate(0, 0, -2) // Move to Friday
	default:
		return date
	}
}

// AddMonthAndAdjust adds noMonths calendar months and adjusts to the next weekday if needed.
func AddMonthAndAdjust(startDate time.Time, noMonths int) time.Time {
	datePlusMonth := startDate.AddDate(0, noMonths, 0)
	return AdjustToNextWorkingday(datePlusMonth)
}

func OrdinaliseDay(x int) string {
	if x >= 10 && x < 19 {
		return fmt.Sprint(x, "th")
	}

	switch x % 10 { // the last digit
	case 1:
		return fmt.Sprint(x, "st")
	case 2:
		return fmt.Sprint(x, "nd")
	case 3:
		return fmt.Sprint(x, "rd")
	}

	return fmt.Sprint(x, "th")
}

func HumanDate(t time.Time) string {
	return OrdinaliseDay(t.Day()) + t.Format(" "+Format.MYH)
}
