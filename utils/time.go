package utils

import (
	"time"

	"puti/config"
)

// GetFormatTime get format time include nil value
func GetFormatTime(t *time.Time, layout string) string {
	if t != nil {
		formatTime := t.In(config.TimeLoc()).Format(layout)
		return formatTime
	}

	return ""
}

// StringToTime changesfer a time string to time.Time
func StringToTime(layout string, timeString string) time.Time {
	t, _ := time.ParseInLocation(layout, timeString, config.TimeLoc())

	return t
}
