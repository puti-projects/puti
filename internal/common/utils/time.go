package utils

import (
	"time"

	"github.com/puti-projects/puti/internal/common/config"

	"github.com/go-sql-driver/mysql"
)

// GetFormatTime get format time include nil value
func GetFormatTime(t *time.Time, layout string) string {
	if t != nil {
		formatTime := t.In(config.TimeLoc()).Format(layout)
		return formatTime
	}

	return ""
}

// GetFormatNullTime get format time which could be null
// it returns empty string if the time is NULL
func GetFormatNullTime(t *mysql.NullTime, layout string) string {
	if t.Valid {
		formatTime := t.Time.In(config.TimeLoc()).Format(layout)
		return formatTime
	}

	return ""
}

// StringToTime changesfer a time string to time.Time
func StringToTime(layout string, timeString string) time.Time {
	t, _ := time.ParseInLocation(layout, timeString, config.TimeLoc())

	return t
}

// StringToNullTime changesfer a time string to mysql.NullTime
func StringToNullTime(layout string, timeString string) mysql.NullTime {
	var nullTime mysql.NullTime
	if timeString == "" {
		nullTime.Valid = false
	} else {
		nullTime.Valid = true
	}

	t, _ := time.ParseInLocation(layout, timeString, config.TimeLoc())
	nullTime.Time = t

	return nullTime
}
