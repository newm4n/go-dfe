package example

import (
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {
	someTime := time.Date(2018, time.February, 12, 12, 13, 14, 0, time.Local)
	if FormatTime(&someTime) != "12 February 2018 12:13:14" {
		t.FailNow()
	}
}

func TestParseTime(t *testing.T) {
	someTime := time.Date(2018, time.February, 12, 12, 13, 14, 0, time.Local)
	time, err := ParseTime("12 February 2018 12:13:14")
	if err != nil {
		t.FailNow()
	} else {
		if time.Year() != someTime.Year() ||
			time.Month() != someTime.Month() ||
			time.Day() != someTime.Day() ||
			time.Hour() != someTime.Hour() ||
			time.Minute() != someTime.Minute() ||
			time.Second() != someTime.Second() {
			t.FailNow()
		}
	}
}
