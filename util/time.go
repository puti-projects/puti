package util

import (
	"puti/config"
	"time"
)

// GetFormatTime get format time include nil value
func GetFormatTime(t *time.Time, laypout string) string {
	if t != nil {
		formatTime := t.In(config.TimeLoc()).Format(laypout)
		return formatTime
	}

	return ""
}
