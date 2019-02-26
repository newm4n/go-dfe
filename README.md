# DateFormatExchange (DFE)

## Get it

```text
$ go get github.com/newm4n/go-dfe
```

## About

Golang formats `time.Time` into `string` or parses a `string` back to `time.Time` using a speciffic date format.
While the format it self resembles the actual date so you get the feeling on how it would look like, making these format
can be some what *un-intuitive*.

Imagine if we have a date string of `"February 21, 2018. At 21:32:12"`,  And the format to parse
this string is `"January 02, 2006. At 15:04:05"`. It appears that `January`, `02`, `2006`,
`15`, `04`, `05` is actually a sub pattern to recognize for their respected date element.

### Date format in Golang

According to **[Golang time.Format() Code](https://golang.org/src/time/format.go)** :

```go
const (
	_                        = iota
	stdLongMonth             = iota + stdNeedDate  // "January"
	stdMonth                                       // "Jan"
	stdNumMonth                                    // "1"
	stdZeroMonth                                   // "01"
	stdLongWeekDay                                 // "Monday"
	stdWeekDay                                     // "Mon"
	stdDay                                         // "2"
	stdUnderDay                                    // "_2"
	stdZeroDay                                     // "02"
	stdHour                  = iota + stdNeedClock // "15"
	stdHour12                                      // "3"
	stdZeroHour12                                  // "03"
	stdMinute                                      // "4"
	stdZeroMinute                                  // "04"
	stdSecond                                      // "5"
	stdZeroSecond                                  // "05"
	stdLongYear              = iota + stdNeedDate  // "2006"
	stdYear                                        // "06"
	stdPM                    = iota + stdNeedClock // "PM"
	stdpm                                          // "pm"
	stdTZ                    = iota                // "MST"
	stdISO8601TZ                                   // "Z0700"  // prints Z for UTC
	stdISO8601SecondsTZ                            // "Z070000"
	stdISO8601ShortTZ                              // "Z07"
	stdISO8601ColonTZ                              // "Z07:00" // prints Z for UTC
	stdISO8601ColonSecondsTZ                       // "Z07:00:00"
	stdNumTZ                                       // "-0700"  // always numeric
	stdNumSecondsTz                                // "-070000"
	stdNumShortTZ                                  // "-07"    // always numeric
	stdNumColonTZ                                  // "-07:00" // always numeric
	stdNumColonSecondsTZ                           // "-07:00:00"
	stdFracSecond0                                 // ".0", ".00", ... , trailing zeros included
	stdFracSecond9                                 // ".9", ".99", ..., trailing zeros omitted

	stdNeedDate  = 1 << 8             // need month, day, year
	stdNeedClock = 2 << 8             // need hour, minute, second
	stdArgShift  = 16                 // extra argument in high bits, above low stdArgShift
	stdMask      = 1<<stdArgShift - 1 // mask out argument
)
```

By looking to the above `const` snippet, you will understand how to create a date format in Go.

### Date format in Java

Date format in Java have a more *intuitive* way to define date format.

According to **[SimpleDateFormat Java Doc](https://docs.oracle.com/javase/7/docs/api/java/text/SimpleDateFormat.html)** :

| Letter | Date or Time Component | Presentation | Examples  |
|-------| --------| -------| -----|
| G | Era designator | Text | AD  |
| y | Year | Year | 1996; 96  |
| Y | Week year | Year | 2009; 09  |
| M | Month in year | Month | July; Jul; 07  |
| w | Week in year | Number | 27  |
| W | Week in month | Number | 2  |
| D | Day in year | Number | 189  |
| d | Day in month | Number | 10  |
| F | Day of week in month | Number | 2  |
| E | Day name in week | Text | Tuesday; Tue  |
| u | Day number of week (1 = Monday, ..., 7 = Sunday) | Number | 1  |
| a | Am/pm marker | Text | PM  |
| H | Hour in day (0-23) | Number | 0  |
| k | Hour in day (1-24) | Number | 24  |
| K | Hour in am/pm (0-11) | Number | 0  |
| h | Hour in am/pm (1-12) | Number | 12  |
| m | Minute in hour | Number | 30  |
| s | Second in minute | Number | 55  |
| S | Millisecond | Number | 978  |
| z | Time zone | General time zone | Pacific Standard Time; PST; GMT-08:00  |
| Z | Time zone | RFC 822 time zone | -0800  |
| X | Time zone | ISO 8601 time zone | -08; -0800; -08:00 |

### How do we map date format in Golang and Java

Thus the following mapping can be created between golang and java date format.

| Const | Go Format | Java Format | 
| --------- | ----------- | ---|
| `stdLongMonth ` | `January` | `MMMM` | 
| `stdMonth` | `Jan` | `MMM` | 
| `stdNumMonth` | `1` | `M` | 
| `stdZeroMonth` | `01` | `MM` | 
| `stdLongWeekDay` | `Monday` | `EEEE` | 
| `stdWeekDay` | `Mon` | `EEE` | 
| `stdDay` | `2` | `d` | 
| `stdUnderDay` | `_2` | n/a fallback to `_d` | 
| `stdZeroDay` | `02` | `dd` | 
| `stdHour` | `15` | `HH` | 
| `stdHour12` | `3` | `K` | 
| `stdZeroHour12` | `03` | `KK` | 
| `stdMinute` | `4` | `m` | 
| `stdZeroMinute` | `04` | `mm` | 
| `stdSecond` | `5` | `s` | 
| `stdZeroSecond` | `05` | `ss` | 
| `stdLongYear` | `2006` | `yyyy` | 
| `stdYear` | `06` | `yy` | 
| `stdPM` | `PM` | `aa` | 
| `stdpm` | `pm` | n/a fallback to `aa` | 
| `stdTZ` | `MST` | `Z` | 
| `stdISO8601TZ` | `Z0700` | `'Z'XX` | 
| `stdISO8601SecondsTZ` | `Z070000` | n/a fallback to `'Z'XX` | 
| `stdISO8601ShortTZ` | `Z07` | `'Z'X` | 
| `stdISO8601ColonTZ` | `Z07:00` | `'Z'XXX` | 
| `stdISO8601ColonSecondsTZ` | `Z07:00:00` | n/a fallback to `'Z'XXX` | 
| `stdNumTZ` | `-0700` | `XX` | 
| `stdNumSecondsTz` | `-070000` | n/a fallback to `'Z'XX` | 
| `stdNumShortTZ` | `-07` | `X` | 
| `stdNumColonTZ` | `-07:00` | `XXX` | 
| `stdNumColonSecondsTZ` | `-07:00:00` | n/a fallback to `XXX` |

## Usage

```go
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
```