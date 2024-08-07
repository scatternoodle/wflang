// Package wdate parses dates and times with WFLangs's required date and time string
// literal formats.
package wdate

import (
	"errors"
	"regexp"
	"time"
)

const (
	datePattern string = `\A{\d{4}-\d{2}-\d{2}}\z`
	timePattern string = `\A{\d{2}:\d{2}}\z`
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
	if !IsDateLiteral(s) {
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
	if !IsTimeLiteral(s) {
		return time.Time{}, ErrTimeFormat
	}
	return time.Parse(timeLayout, s)
}

// IsDateLiteral returns true if the given string matches the WFLang regex pattern
// for date literals.
func IsDateLiteral(s string) bool {
	ok, _ := regexp.MatchString(datePattern, s)
	return ok
}

// IsDateLiteral returns true if the given string matches the WFLang regex pattern
// for time literals.
func IsTimeLiteral(s string) bool {
	ok, _ := regexp.MatchString(timePattern, s)
	return ok
}
