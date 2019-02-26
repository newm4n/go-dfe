package example

import (
	"github.com/newm4n/go-dfe"
	"time"
)

func FormatTime(t *time.Time) string {
	translation := DateFormatExchange.NewPatternTranslation()
	return t.Format(translation.JavaToGoFormat("dd MMMM yyyy HH:mm:ss"))
}

func ParseTime(date string) (*time.Time, error) {
	translation := DateFormatExchange.NewPatternTranslation()
	t, err := time.Parse(translation.JavaToGoFormat("dd MMMM yyyy HH:mm:ss"), date)
	if err != nil {
		return nil, err
	}
	return &t, err
}
