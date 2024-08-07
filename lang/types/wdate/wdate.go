// Package wdate parses dates and times with WFLangs's required date and time string
// literal formats.
package wdate

import (
	"errors"
	"regexp"
	"time"
)

const (
	dateRegexp string = `\A{\d{4}-\d{2}-\d{2}}\z`
	timeRegexp string = `\A{\d{2}:\d{2}}\z`
)

// For parsing into golang objects
const (
	dateLayout string = `{2006-01-02}`
	timeLayout string = `{15:04}`
)

var (
	ErrDateFormat  = errors.New("invalid date format, must match pattern {yyyy-MM-dd} (including braces)")
	ErrTimeFormat  = errors.New("invalid time format, must match pattern {hh:mm} (including braces)")
	ErrOutOfBounds = errors.New("invalid dateTime, must be between {1900-01-01 00:00} and {3000-12-31 23:59}")
)

// ParseDate returns the time represented by s.
func ParseDate(s string) (time.Time, error) {
	if ok, _ := regexp.MatchString(dateRegexp, s); !ok {
		return time.Time{}, ErrDateFormat
	}

	t, err := time.Parse(dateLayout, s)
	if err != nil {
		return time.Time{}, err
	}

	start := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(3000, 12, 31, 23, 59, 59, 0, time.UTC)
	if start.After(t) || t.After(end) {
		return time.Time{}, ErrOutOfBounds
	}

	return t, nil
}

// ParseDate returns the date represented by s.
func ParseTime(s string) (time.Time, error) {
	if ok, _ := regexp.MatchString(timeRegexp, s); !ok {
		return time.Time{}, ErrTimeFormat
	}
	return time.Parse(timeLayout, s)
}
