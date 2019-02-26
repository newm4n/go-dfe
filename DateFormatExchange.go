package DateFormatExchange

import (
	"sort"
	"strings"
	"sync"
)

// newFormatMapping create new instance of FormatMapping
func newFormatMapping(gof, jaf string, fallback bool) *FormatMapping {
	return &FormatMapping{
		GoForm:   gof,
		JavaForm: jaf,
		Fallback: fallback,
	}
}

// FormatMapping stores one on one translation between date format element in Golang and Java
type FormatMapping struct {
	GoForm   string
	JavaForm string
	Fallback bool
}

// NewPatternTranslation create new PatternTranslation instance, initialize the translation table.
func NewPatternTranslation() *PatternTranslation {
	ret := &PatternTranslation{
		Mappings: []*FormatMapping{
			newFormatMapping("January", "MMMM", false),
			newFormatMapping("Jan", "MMM", false),
			newFormatMapping("1", "M", false),
			newFormatMapping("01", "MM", false),
			newFormatMapping("Monday", "EEEE", false),
			newFormatMapping("Mon", "EEE", false),
			newFormatMapping("2", "d", false),
			newFormatMapping("_2", "_d", true),
			newFormatMapping("02", "dd", false),
			newFormatMapping("15", "HH", false),
			newFormatMapping("3", "K", false),
			newFormatMapping("03", "KK", false),
			newFormatMapping("4", "m", false),
			newFormatMapping("04", "mm", false),
			newFormatMapping("5", "s", false),
			newFormatMapping("05", "ss", false),
			newFormatMapping("2006", "yyyy", false),
			newFormatMapping("06", "yy", false),
			newFormatMapping("PM", "aa", false),
			newFormatMapping("pm", "aa", true),
			newFormatMapping("MST", "Z", false),
			newFormatMapping("Z0700", "'Z'XX", false),
			newFormatMapping("Z070000", "'Z'XX", true),
			newFormatMapping("Z07", "'Z'X", false),
			newFormatMapping("Z07:00", "'Z'XXX", false),
			newFormatMapping("Z07:00:00", "'Z'XXX", true),
			newFormatMapping("-0700", "XX", false),
			newFormatMapping("-070000", "'Z'XX", true),
			newFormatMapping("-07", "X", false),
			newFormatMapping("-07:00", "XXX", false),
			newFormatMapping("-07:00:00", "XXX", true),
		},
	}
	return ret
}

// PatternTranslation is a struct that do the translation.
type PatternTranslation struct {
	Mappings []*FormatMapping
	mutex    sync.Mutex
}

// GoToJavaFormat will translate Golang style date format into Java style.
func (pt *PatternTranslation) GoToJavaFormat(format string) string {
	pt.mutex.Lock()
	defer pt.mutex.Unlock()
	sort.Slice(pt.Mappings, func(i, j int) bool {
		return len(pt.Mappings[i].GoForm) > len(pt.Mappings[j].GoForm)
	})
	fstr := format
	for _, m := range pt.Mappings {
		fstr = strings.Replace(fstr, m.GoForm, m.JavaForm, -1)
	}
	return fstr
}

// JavaToGoFormat will translate Java style date format into Golang style.
func (pt *PatternTranslation) JavaToGoFormat(format string) string {
	pt.mutex.Lock()
	defer pt.mutex.Unlock()
	sort.Slice(pt.Mappings, func(i, j int) bool {
		return len(pt.Mappings[i].JavaForm) > len(pt.Mappings[j].JavaForm)
	})
	fstr := format
	for _, m := range pt.Mappings {
		if m.Fallback {
			continue
		}
		if m.JavaForm == "M" && (strings.Contains(fstr, "Mon") || strings.Contains(fstr, "AM") || strings.Contains(fstr, "PM")) {
			continue
		}
		if m.JavaForm == "d" && strings.Contains(fstr, "Monday") {
			continue
		}
		fstr = strings.Replace(fstr, m.JavaForm, m.GoForm, -1)
	}
	return fstr
}
