package utils

import "time"

// CustomDate to handle the "2006-01-02" format for date-only fields
type CustomDate struct {
	time.Time
}

const customDateFormat = "2006-01-02"