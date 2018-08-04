package shorts

import "time"

// FormatTimeDay Format time.Time to dd-MM-yyyy
func FormatTimeDay(time time.Time) string {
	return time.Format("02.01.2006")
}

// FormatTimeDayHour Format time.Time to dd-MM-yyyy hh:mm
func FormatTimeDayHour(time time.Time) string {
	return time.Format("02.01.2006 15:04")
}

// FormatTimeDaySecond Format time.Time to dd-MM-yyyy hh:mm:ss
func FormatTimeDaySecond(time time.Time) string {
	return time.Format("02.01.2006 15:04:05")
}

// FormatTimeYear Format time.Time to yyyy
func FormatTimeYear(time time.Time) string {
	return time.Format("2006")
}

// FormatTimeDayName Get the name of the day
func FormatTimeDayName(time time.Time) string {
	return time.Format("Mon")
}

// FormatTimeMonthName Get the name of the day
func FormatTimeMonthName(time time.Time) string {
	return time.Format("Jun")
}
