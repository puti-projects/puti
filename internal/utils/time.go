package utils

import (
	"database/sql"
	"time"

	"github.com/puti-projects/puti/internal/pkg/config"

	"gorm.io/gorm"
)

// GetFormatTime get format time include nil value
func GetFormatTime(t *time.Time, layout string) string {
	if t == nil {
		return ""
	}

	formatedTime := t.In(config.TimeLoc()).Format(layout)
	return formatedTime
}

// GetFormatDeletedAtTime get format time from grom(v2) DeletedAt
func GetFormatDeletedAtTime(d *gorm.DeletedAt, layout string) string {
	t, _ := d.Value()
	if tt, ok := t.(time.Time); ok {
		return GetFormatTime(&tt, layout)
	}

	return ""
}

// GetFormatNullTime get format time which could be NULL
// it returns empty string if the time is NULL
func GetFormatNullTime(t *sql.NullTime, layout string) string {
	if !t.Valid {
		return ""
	}

	formatedTime := t.Time.In(config.TimeLoc()).Format(layout)
	return formatedTime
}

// StringToTime changesfer a time string to time.Time
func StringToTime(layout string, timeString string) time.Time {
	t, _ := time.ParseInLocation(layout, timeString, config.TimeLoc())

	return t
}

// StringToNullTime changesfer a time string to mysql.NullTime
func StringToNullTime(layout string, timeString string) sql.NullTime {
	var nullTime sql.NullTime
	if timeString == "" {
		nullTime.Valid = false
	} else {
		nullTime.Valid = true
	}

	nullTime.Time = StringToTime(layout, timeString)
	return nullTime
}

// SubNullTimeUnitlNowAsDay calculate the diff day until now
func SubNullTimeUnitlNowAsDay(t time.Time) float64 {
	now := time.Now()

	sub := now.Sub(t)
	subDay := sub.Hours() / 24

	return subDay
}
