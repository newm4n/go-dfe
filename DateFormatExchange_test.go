package DateFormatExchange

import "testing"

func NewConvertTest(from, to string) *ConvertTest {
	return &ConvertTest{
		From: from,
		To:   to,
	}
}

type ConvertTest struct {
	From string
	To   string
}

func TestGoToJava(t *testing.T) {
	tests := make([]*ConvertTest, 0)

	tests = append(tests, NewConvertTest("02 January 2006 15:04:05", "dd MMMM yyyy HH:mm:ss"))
	tests = append(tests, NewConvertTest("02 Jan 2006 15:04:05", "dd MMM yyyy HH:mm:ss"))
	tests = append(tests, NewConvertTest("Mon, 02 Jan 2006 15:04:05", "EEE, dd MMM yyyy HH:mm:ss"))
	tests = append(tests, NewConvertTest("Mon, 02 Jan 2006 15:04:05 PM", "EEE, dd MMM yyyy HH:mm:ss aa"))
	tests = append(tests, NewConvertTest("Mon, 02 Jan 2006 03:04:05PM", "EEE, dd MMM yyyy KK:mm:ssaa"))
	tests = append(tests, NewConvertTest("Monday, 02 Jan 2006 15:04:05", "EEEE, dd MMM yyyy HH:mm:ss"))

	tests = append(tests, NewConvertTest("Mon Jan _2 15:04:05 2006", "EEE MMM _d HH:mm:ss yyyy"))
	tests = append(tests, NewConvertTest("Mon Jan _2 15:04:05 MST 2006", "EEE MMM _d HH:mm:ss Z yyyy"))
	tests = append(tests, NewConvertTest("Mon Jan 02 15:04:05 -0700 2006", "EEE MMM dd HH:mm:ss XX yyyy"))
	tests = append(tests, NewConvertTest("02 Jan 06 15:04 MST", "dd MMM yy HH:mm Z"))
	tests = append(tests, NewConvertTest("02 Jan 06 15:04 -0700", "dd MMM yy HH:mm XX"))
	tests = append(tests, NewConvertTest("Monday, 02-Jan-06 15:04:05 MST", "EEEE, dd-MMM-yy HH:mm:ss Z"))
	tests = append(tests, NewConvertTest("Mon, 02 Jan 2006 15:04:05 MST", "EEE, dd MMM yyyy HH:mm:ss Z"))
	tests = append(tests, NewConvertTest("Mon, 02 Jan 2006 15:04:05 -0700", "EEE, dd MMM yyyy HH:mm:ss XX"))
	tests = append(tests, NewConvertTest("2006-01-02T15:04:05Z07:00", "yyyy-MM-ddTHH:mm:ss'Z'XXX"))
	tests = append(tests, NewConvertTest("2006-01-02T15:04:05.999999999Z07:00", "yyyy-MM-ddTHH:mm:ss.SSS'Z'XXX"))
	tests = append(tests, NewConvertTest("3:04PM", "K:mmaa"))
	// Handy time stamps.
	tests = append(tests, NewConvertTest("Jan _2 15:04:05", "MMM _d HH:mm:ss"))
	tests = append(tests, NewConvertTest("Jan _2 15:04:05.000", "MMM _d HH:mm:ss.000"))
	tests = append(tests, NewConvertTest("Jan _2 15:04:05.000000", "MMM _d HH:mm:ss.000000"))
	tests = append(tests, NewConvertTest("Jan _2 15:04:05.000000000", "MMM _d HH:mm:ss.000000000"))

	pt := NewPatternTranslation()

	for _, ct := range tests {
		result := pt.GoToJavaFormat(ct.From)
		if result != ct.To {
			t.Logf("GoToPattern Expect \"%s\" == \"%s\", But convert result is \"%s\" ... Fail", ct.From, ct.To, result)
			t.Fail()
		} else {
			t.Logf("Conversion from \"%s\" to \"%s\" ... OK", ct.From, ct.To)
		}
	}
}

func TestJavaToGo(t *testing.T) {
	tests := make([]*ConvertTest, 0)

	tests = append(tests, NewConvertTest("dd MMMM yyyy HH:mm:ss", "02 January 2006 15:04:05"))
	tests = append(tests, NewConvertTest("dd MMM yyyy HH:mm:ss", "02 Jan 2006 15:04:05"))
	tests = append(tests, NewConvertTest("EEEE, dd MMM yyyy HH:mm:ss", "Monday, 02 Jan 2006 15:04:05"))
	tests = append(tests, NewConvertTest("EEE, dd MMM yyyy HH:mm:ss aa", "Mon, 02 Jan 2006 15:04:05 PM"))
	tests = append(tests, NewConvertTest("EEE, dd MMM yyyy HH:mm:ssaa", "Mon, 02 Jan 2006 15:04:05PM"))
	tests = append(tests, NewConvertTest("yyyy-MM-ddTHH:mm:ss.SSSXXX", "2006-01-02T15:04:05.999999999-07:00"))

	pt := NewPatternTranslation()

	for _, ct := range tests {
		result := pt.JavaToGoFormat(ct.From)
		if result != ct.To {
			t.Logf("PatternToGo Expect \"%s\" == \"%s\", But convert result is \"%s\" ... Fail", ct.From, ct.To, result)
			t.Fail()
		} else {
			t.Logf("Conversion from \"%s\" to \"%s\" ... OK", ct.From, ct.To)
		}
	}
}
