package wdate

import (
	"testing"
	"time"
)

var timeZero = time.Time{}

func TestParseDate(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wantTime time.Time
		wantErr  bool
	}{
		{"start of time", `{1900-01-01}`, simpleDate(1900, 1, 1), false},
		{"end of time", `{3000-12-31}`, simpleDate(3000, 12, 31), false},
		{"too early", `{1899-12-31}`, timeZero, true},
		{"too late", `{3001-01-01}`, timeZero, true},
		{"extra char start", ` {3001-01-01}`, timeZero, true},
		{"extra char end", `{3001-01-01} `, timeZero, true},
		{"missing braces", `3001-01-01`, timeZero, true},
		{"not a date", `{23:59}`, timeZero, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have, err := ParseDate(tt.s)
			if err != nil && !tt.wantErr {
				t.Fatalf("unexpected error: %s", err.Error())
			}
			if err == nil && tt.wantErr {
				t.Fatalf("did not error when wanted")
			}
			if have != tt.wantTime {
				t.Fatalf("have %s, want %s", have.String(), tt.wantTime.String())
			}
		})
	}
}

func TestParseTime(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wantTime time.Time
		wantErr  bool
	}{
		{"start of day", `{00:00}`, simpleTime(0, 0), false},
		{"end of day", `{23:59}`, simpleTime(23, 59), false},
		{"bad hours value", `{24:00}`, timeZero, true},
		{"bad minutes value", `{22:60}`, timeZero, true},
		{"extra char start", ` {23:59}`, timeZero, true},
		{"extra char end", `{23:59} `, timeZero, true},
		{"missing braces", `23:59`, timeZero, true},
		{"not a date", `{1900-01-01}`, timeZero, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have, err := ParseTime(tt.s)
			if err != nil && !tt.wantErr {
				t.Fatalf("unexpected error: %s", err.Error())
			}
			if err == nil && tt.wantErr {
				t.Fatalf("did not error when wanted")
			}
			if have != tt.wantTime {
				t.Fatalf("have %s, want %s", have.String(), tt.wantTime.String())
			}
		})
	}
}

func simpleDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func simpleTime(hours, minutes int) time.Time {
	return time.Date(0, 1, 1, hours, minutes, 0, 0, time.UTC)
}
