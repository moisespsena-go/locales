package locales

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type (
	TimeTranslator interface {
		TimeSeparator() string

		// returns the locales wide months
		MonthsWide() []string
		MonthWide(time.Month) string

		// returns the locales abbreviated months
		MonthsAbbreviated() []string
		MonthAbbreviated(time.Month) string

		// returns the locales narrow months
		MonthsNarrow() []string
		MonthNarrow(time.Month) string

		// returns the locales wide weekdays
		WeekdaysWide() []string
		WeekdayWide(d time.Weekday) string

		// returns the locales abbreviated weekdays
		WeekdaysAbbreviated() []string
		WeekdayAbbreviated(d time.Weekday) string

		// WeekdaysNarrowreturns the locales narrow weekdays
		WeekdaysNarrow() []string
		WeekdayNarrow(d time.Weekday) string

		// returns the locales short weekdays
		WeekdaysShort() []string
		WeekdayShort(d time.Weekday) string

		PeriodsAbbreviated() []string
		PeriodsShort() []string
		PeriodsWide() []string
		PeriodsNarrow() []string

		ErasWide() []string
		ErasAbbreviated() []string
		ErasNarrow() []string

		Timezones() map[string]string

		// returns the short date representation of 't' for locale
		FmtDateShort(t time.Time) string

		// returns the medium date representation of 't' for locale
		FmtDateMedium(t time.Time) string

		//  returns the long date representation of 't' for locale
		FmtDateLong(t time.Time) string

		// returns the full date representation of 't' for locale
		FmtDateFull(t time.Time) string

		// returns the short time representation of 't' for locale
		FmtTimeShort(t time.Time) string

		// returns the medium time representation of 't' for locale
		FmtTimeMedium(t time.Time) string

		// returns the long time representation of 't' for locale
		FmtTimeLong(t time.Time) string

		// returns the full time representation of 't' for locale
		FmtTimeFull(t time.Time) string

		// DateFullLayout returns the full date layout for '{{ .Locale }}'
		DateFullLayout() string

		// DateLongLayout returns the long date layout for '{{ .Locale }}'
		DateLongLayout() string

		// DateMediumLayout returns the medium date layout for '{{ .Locale }}'
		DateMediumLayout() string

		// DateShortLayout returns the short date layout for '{{ .Locale }}'
		DateShortLayout() string

		// TimeFullLayout returns the full time layout for '{{ .Locale }}'
		TimeFullLayout() string

		// TimeLongLayout returns the full long layout for '{{ .Locale }}'
		TimeLongLayout() string

		// TimeMediumLayout returns the medium time layout for '{{ .Locale }}'
		TimeMediumLayout() string

		// TimeShortLayout returns the short time layout for '{{ .Locale }}'
		TimeShortLayout() string
	}

	TimeSpec struct {
		Separator string

		// returns the locales wide months
		MonthsWideValue []string

		// returns the locales abbreviated months
		MonthsAbbreviatedValue []string

		// returns the locales narrow months
		MonthsNarrowValue []string

		// returns the locales wide weekdays
		WeekdaysWideValue []string

		// returns the locales abbreviated weekdays
		WeekdaysAbbreviatedValue []string

		// WeekdaysNarrowreturns the locales narrow weekdays
		WeekdaysNarrowValue []string

		// returns the locales short weekdays
		WeekdaysShortValue []string

		PeriodsAbbreviatedValue []string

		PeriodsShortValue []string

		PeriodsNarrowValue []string

		PeriodsWideValue []string

		ErasWideValue []string

		ErasAbbreviatedValue []string

		ErasNarrowValue []string

		TimezonesValue map[string]string
	}
)

func (t *TimeSpec) TimeSeparator() string {
	return t.Separator
}

func (t *TimeSpec) MonthsWide() []string {
	return t.MonthsWideValue
}

func (t *TimeSpec) MonthWide(m time.Month) string {
	return t.MonthsWideValue[m]
}

func (t *TimeSpec) MonthsAbbreviated() []string {
	return t.MonthsAbbreviatedValue
}

func (t *TimeSpec) MonthAbbreviated(m time.Month) string {
	return t.MonthsAbbreviatedValue[m]
}

func (t *TimeSpec) MonthsNarrow() []string {
	return t.MonthsNarrowValue
}

func (t *TimeSpec) MonthNarrow(m time.Month) string {
	return t.MonthsNarrowValue[m]
}

func (t *TimeSpec) WeekdaysWide() []string {
	return t.WeekdaysWideValue
}

func (t *TimeSpec) WeekdayWide(d time.Weekday) string {
	return t.WeekdaysWideValue[d]
}

func (t *TimeSpec) WeekdaysAbbreviated() []string {
	return t.WeekdaysAbbreviatedValue
}

func (t *TimeSpec) WeekdayAbbreviated(d time.Weekday) string {
	return t.WeekdaysAbbreviatedValue[d]
}

func (t *TimeSpec) WeekdaysNarrow() []string {
	return t.WeekdaysNarrowValue
}

func (t *TimeSpec) WeekdayNarrow(d time.Weekday) string {
	return t.WeekdaysNarrowValue[d]
}

func (t *TimeSpec) WeekdaysShort() []string {
	return t.WeekdaysShortValue
}

func (t *TimeSpec) WeekdayShort(d time.Weekday) string {
	return t.WeekdaysShortValue[d]
}

func (t *TimeSpec) PeriodsAbbreviated() []string {
	return t.PeriodsAbbreviatedValue
}

func (t *TimeSpec) PeriodsNarrow() []string {
	return t.PeriodsNarrowValue
}

func (t *TimeSpec) PeriodsShort() []string {
	return t.PeriodsShortValue
}

func (t *TimeSpec) PeriodsWide() []string {
	return t.PeriodsWideValue
}

func (t *TimeSpec) ErasWide() []string {
	return t.ErasWideValue
}

func (t *TimeSpec) ErasAbbreviated() []string {
	return t.ErasAbbreviatedValue
}

func (t *TimeSpec) ErasNarrow() []string {
	return t.ErasNarrowValue
}

func (t *TimeSpec) Timezones() map[string]string {
	return t.TimezonesValue
}

type TimestampFormat uint8

const (
	TFNone TimestampFormat = 0
	TFDate TimestampFormat = 1 << iota
	TFTime
	TFFull
	TFLong
	TFMedium
	TFShort
	TFTimeFirst

	DateTime = TFDate | TFTime
)

func (f TimestampFormat) Set(flag TimestampFormat) TimestampFormat    { return f | flag }
func (f TimestampFormat) Clear(flag TimestampFormat) TimestampFormat  { return f &^ flag }
func (f TimestampFormat) Toggle(flag TimestampFormat) TimestampFormat { return f ^ flag }
func (f TimestampFormat) Has(flag TimestampFormat) bool               { return f&flag != 0 }

func (f *TimestampFormat) ParseAppend(p string) {
	switch p {
	case "time_first":
		*f |= TFTimeFirst
	case "time", "t":
		*f |= TFTime
	case "date", "d":
		*f |= TFDate
	case "full", "f":
		*f = f.Clear(TFLong|TFMedium|TFShort) | TFFull
	case "long", "l":
		*f = f.Clear(TFFull|TFMedium|TFShort) | TFLong
	case "medium", "m":
		*f = f.Clear(TFFull|TFLong|TFShort) | TFMedium
	case "short", "s":
		*f = f.Clear(TFFull|TFLong|TFMedium) | TFShort
	case "-date":
		*f = f.Clear(TFDate)
	case "-time":
		*f = f.Clear(TFTime)
	}
}

func (f *TimestampFormat) ParseAppendArray(p ...string) {
	for _, s := range p {
		f.ParseAppend(s)
	}
}

func (f *TimestampFormat) Parse(s string) {
	*f = TFNone
	f.ParseAppendArray(strings.Split(s, ",")...)
}

func (f TimestampFormat) Layout() TimestampFormat {
	if f.Has(TFFull) {
		return TFFull
	} else if f.Has(TFLong) {
		return TFLong
	} else if f.Has(TFMedium) {
		return TFMedium
	} else if f.Has(TFShort) {
		return TFShort
	}
	return 0
}

func (f *TimestampFormat) String() string {
	var s []string
	if f.Has(TFDate) {
		s = append(s, "date")
	}
	if f.Has(TFTime) {
		s = append(s, "date")
	}
	if f.Has(TFFull) {
		s = append(s, "full")
	} else if f.Has(TFLong) {
		s = append(s, "long")
	} else if f.Has(TFMedium) {
		s = append(s, "medium")
	} else if f.Has(TFShort) {
		s = append(s, "short")
	}
	return strings.Join(s, ",")
}

func (f *TimestampFormat) EraScore() int8 {
	if f.Has(TFFull) {
		return 0
	} else if f.Has(TFLong) {
		return 1
	} else if f.Has(TFMedium) || f.Has(TFShort) {
		return 2
	}
	return -1
}

func (f TimestampFormat) Layouts(t TimeTranslator) (s []string) {
	if f.Has(TFDate) {
		if f.Has(TFFull) {
			s = append(s, t.DateFullLayout())
		} else if f.Has(TFLong) {
			s = append(s, t.DateLongLayout())
		} else if f.Has(TFMedium) {
			s = append(s, t.DateMediumLayout())
		} else if f.Has(TFShort) {
			s = append(s, t.DateShortLayout())
		}
	}
	if f.Has(TFTime) {
		if f.Has(TFFull) {
			s = append(s, t.TimeFullLayout())
		} else if f.Has(TFLong) {
			s = append(s, t.TimeLongLayout())
		} else if f.Has(TFMedium) {
			s = append(s, t.TimeMediumLayout())
		} else if f.Has(TFShort) {
			s = append(s, t.TimeShortLayout())
		}
	}
	if f.Has(TFTimeFirst) && len(s) == 2 {
		s[0], s[1] = s[1], s[0]
	}
	return
}

func (f TimestampFormat) Format(Tr TimeTranslator, t time.Time) string {
	return f.FormatSep(Tr, t, ", ")
}

func (f TimestampFormat) FormatSep(Tr TimeTranslator, t time.Time, dateTimeSep string) string {
	var (
		eraScore = f.EraScore()
		s        = f.Layouts(Tr)
	)

	for i, l := range s {
		s[i] = FormatTimeEra(Tr, l, t, eraScore)
	}

	return strings.Join(s, dateTimeSep)
}

func FormatTime(t TimeTranslator, format string, T time.Time) string {
	return FormatTimeEra(t, format, T, -1)
}

func FormatTimeEra(t TimeTranslator, format string, T time.Time, eraScore int8) string {

	// rules:
	// y = four digit year
	// yy = two digit year

	var (
		inConstantText bool
		start          int
		b              bytes.Buffer
	)

	if format == "" {
		format = t.DateFullLayout()
	}

	for i := 0; i < len(format); i++ {

		switch format[i] {

		// time separator
		case ':':

			if inConstantText {
				inConstantText = false
				b.WriteString(format[start:i])
			}

			b.WriteString(t.TimeSeparator())
		case '\'':

			i++
			startI := i

			// peek to see if ''
			if len(format) != i && format[i] == '\'' {

				if inConstantText {
					inConstantText = false
					b.WriteString(format[start : i-1])
				} else {
					inConstantText = true
					start = i
				}

				continue
			}

			// not '' so whatever comes between '' is constant

			if len(format) != i {

				// advance i to the next single quote + 1
				for ; i < len(format); i++ {
					if format[i] == '\'' {

						if inConstantText {
							inConstantText = false
							b.WriteString(format[start : startI-1])
							b.WriteString(format[startI:i])
						} else {
							b.WriteString(format[startI:i])
						}

						break
					}
				}
			}

		// 24 hour
		case 'H':

			if inConstantText {
				inConstantText = false
				b.WriteString(format[start:i])
			}

			// peek
			// two digit hour required?
			if len(format) != i+1 && format[i+1] == 'H' {
				i++

				if T.Hour() < 10 {
					b.WriteByte('0')
				}
			}

			b.Write(strconv.AppendInt(nil, int64(T.Hour()), 10))

		// hour
		case 'h':

			if inConstantText {
				inConstantText = false
				b.WriteString(format[start:i])
			}

			h := T.Hour()

			if h > 12 {
				h -= 12
			}

			// peek
			// two digit hour required?
			if len(format) != i+1 && format[i+1] == 'h' {
				i++
				if h < 10 {
					b.WriteByte('0')
				}

			}

			b.Write(strconv.AppendInt(nil, int64(h), 10))

		// minute
		case 'm':

			if inConstantText {
				inConstantText = false
				b.WriteString(format[start:i])
			}

			// peek
			// two digit minute required?
			if len(format) != i+1 && format[i+1] == 'm' {
				i++
				if T.Minute() < 10 {
					b.WriteByte('0')
				}
			}

			b.Write(strconv.AppendInt(nil, int64(T.Minute()), 10))

		// second
		case 's':

			if inConstantText {
				inConstantText = false
				b.WriteString(format[start:i])
			}

			// peek
			// two digit minute required?
			if len(format) != i+1 && format[i+1] == 's' {
				i++
				if T.Second() < 10 {
					b.WriteByte('0')
				}
			}

			b.Write(strconv.AppendInt(nil, int64(T.Second()), 10))

		// day period
		case 'a':

			if inConstantText {
				inConstantText = false
				b.WriteString(format[start:i])
			}

			// only used with 'h', patterns should not contains 'a' without 'h' so not checking

			// choosing to use abbreviated, didn't see any rules about which should be used with which
			// date format....

			if T.Hour() < 12 {
				b.WriteString(t.PeriodsAbbreviated()[0])
			} else {
				b.WriteString(t.PeriodsAbbreviated()[1])
			}

		// timezone
		case 'z', 'v':

			if inConstantText {
				inConstantText = false
				b.WriteString(format[start:i])
			}

			// consume multiple, only handling Abbrev tz from time.Time for the moment...

			var count int

			if format[i] == 'z' {
				for j := i; j < len(format); j++ {
					if format[j] == 'z' {
						count++
					} else {
						break
					}
				}
			}

			if format[i] == 'v' {
				for j := i; j < len(format); j++ {
					if format[j] == 'v' {
						count++
					} else {
						break
					}
				}
			}

			i += count - 1

			tz, _ := T.Zone()

			// using the timezone on the Go time object, eg. EST, EDT, MST.....
			if count < 4 {
				b.WriteString(tz)
			} else {
				if btz := t.Timezones()[tz]; btz != "" {
					b.WriteString(btz)
				} else {
					b.WriteString(tz)
				}
			}

		// day
		case 'd':

			if inConstantText {
				inConstantText = false
				b.WriteString(format[start:i])
			}

			// peek
			// two digit day required?
			if len(format) != i+1 && format[i+1] == 'd' {
				i++
				if T.Day() < 10 {
					b.WriteByte('0')
				}
			}

			b.Write(strconv.AppendInt(nil, int64(T.Day()), 10))

		// month
		case 'M':

			if inConstantText {
				inConstantText = false
				b.WriteString(format[start:i])
			}

			var count int

			// count # of M's
			for j := i; j < len(format); j++ {
				if format[j] == 'M' {
					count++
				} else {
					break
				}
			}

			switch count {

			// Numeric form, at least 1 digit
			case 1:
				b.Write(strconv.AppendInt(nil, int64(T.Month()), 10))

			// Number form, at least 2 digits (padding with 0)
			case 2:
				if T.Month() < 10 {
					b.WriteByte('0')
				}

				b.Write(strconv.AppendInt(nil, int64(T.Month()), 10))

			// Abbreviated form
			case 3:
				b.WriteString(t.MonthAbbreviated(T.Month()))

			// Full/Wide form
			case 4:
				b.WriteString(t.MonthWide(T.Month()))

			// Narrow form - only used in where context makes it clear, such as headers in a calendar.
			// Should be one character wherever possible.
			case 5:
				b.WriteString(t.MonthNarrow(T.Month()))
			}

			// skip over M's
			i += count - 1

		// year
		case 'y':

			if inConstantText {
				inConstantText = false
				b.WriteString(format[start:i])
			}

			// peek
			// two digit year
			if len(format) != i+1 && format[i+1] == 'y' {
				i++

				if T.Year() > 9 {
					b.WriteString(strconv.Itoa(T.Year())[2:])
				} else {
					b.WriteString(strconv.Itoa(T.Year())[1:])
				}
			} else {
				// four digit year
				if T.Year() > 0 {
					b.Write(strconv.AppendInt(nil, int64(T.Year()), 10))
				} else {
					b.Write(strconv.AppendInt(nil, int64(-T.Year()), 10))
				}
			}

		// weekday
		// I know I only see 'EEEE' in the xml, but just in case handled all posibilities
		case 'E':

			if inConstantText {
				inConstantText = false
				b.WriteString(format[start:i])
			}

			var count int

			// count # of E's
			for j := i; j < len(format); j++ {
				if format[j] == 'E' {
					count++
				} else {
					break
				}
			}

			switch count {

			// Narrow
			case 1:

				b.WriteString(t.WeekdayNarrow(T.Weekday()))

			// Short
			case 2:

				b.WriteString(t.WeekdayShort(T.Weekday()))

			// Abbreviated
			case 3:

				b.WriteString(t.WeekdayAbbreviated(T.Weekday()))

			// Full/Wide
			case 4:

				b.WriteString(t.WeekdayWide(T.Weekday()))
			}

			// skip over E's
			i += count - 1

		// era eg. AD, BC
		case 'G':

			if inConstantText {
				inConstantText = false
				b.WriteString(format[start:i])
			}

			switch eraScore {
			case 0:
				if T.Year() < 0 {
					b.WriteString(t.ErasWide()[0])
				} else {
					b.WriteString(t.ErasWide()[1])
				}

			case 1, 2:
				if T.Year() < 0 {
					b.WriteString(t.ErasAbbreviated()[0])
				} else {
					b.WriteString(t.ErasAbbreviated()[1])
				}
			}

		default:
			// append all non matched text as they are constants
			if !inConstantText {
				inConstantText = true
				start = i
			}
		}
	}

	// if we were inConstantText when the string ended, add what's left.
	if inConstantText {
		// inContantText = false
		b.WriteString(format[start:])
	}

	return b.String()
}

type TimeValue struct {
	Translator TimeTranslator
	Formatter  TimestampFormat
	Value      time.Time
}

func NewTimeValue(translator TimeTranslator, value time.Time) *TimeValue {
	return &TimeValue{Translator: translator, Value: value}
}

func NewDateValue(translator TimeTranslator, value time.Time) *TimeValue {
	return &TimeValue{Translator: translator, Value: value, Formatter: TFShort | TFDate}
}

func NewDateTimeValue(translator TimeTranslator, value time.Time) *TimeValue {
	return &TimeValue{Translator: translator, Value: value, Formatter: DateTime}
}

func (t TimeValue) ToString() string {
	return fmt.Sprint(t)
}

func (t TimeValue) Format(f fmt.State, verb rune) {
	if t.Formatter == 0 {
		t.Formatter = TFShort | TFTime | TFDate
	}
	switch verb {
	case 's', 'v':
		io.WriteString(f, t.Formatter.Format(t.Translator, t.Value))
	default:
		fmt.Fprintf(f, "%v", t.Value)
	}
}

type TimeRangeValue struct {
	Translator Translator
	Formatter  TimestampFormat
	From, To   time.Time
}

func NewTimeRangeValue(translator Translator, from time.Time, to time.Time) *TimeRangeValue {
	return &TimeRangeValue{Translator: translator, From: from, To: to}
}

func NewdDateRangeValue(translator Translator, from time.Time, to time.Time) *TimeRangeValue {
	return &TimeRangeValue{Translator: translator, From: from, To: to, Formatter: TFShort | TFDate}
}

func (t TimeRangeValue) ToString() string {
	return fmt.Sprint(t)
}

func (t TimeRangeValue) Format(f fmt.State, verb rune) {
	if t.Formatter == 0 {
		t.Formatter = TFShort | TFTime | TFDate
	}
	switch verb {
	case 's', 'v':
		rg := t.Translator.GetMiscPatterns().Range
		fmt.Fprintf(f, rg, t.Formatter.Format(t.Translator, t.From), t.Formatter.Format(t.Translator, t.To))
	default:
		fmt.Fprintf(f, "TimeRangeValue{%v,%v,%v,%v}", t.Translator, t.Formatter, t.From, t.To)
	}
}
