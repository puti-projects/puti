package config

import "time"

var defaultTimeZone = "Asia/Shanghai"

// TimeLoc sets the time zone location
func TimeLoc() *time.Location {
	// TODO load settings and set defult
	loc, err := time.LoadLocation(defaultTimeZone)
	if err != nil {
		return nil
	}

	return loc
}
