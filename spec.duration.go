package locales

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"time"
)

type Gender uint8

func (g Gender) String() string {
	switch g {
	case Masculine:
		return "Masculine"
	case Feminine:
		return "Feminine"
	default:
		return "Gender(" + strconv.Itoa(int(g)) + ")"
	}
}

const (
	Masculine Gender = iota
	Feminine
)

var DefaultDurationSpec = DurationSpec{
	LongSep:  " ",
	ShortSep: "",
	Year: DurationSpecPair{
		Masculine,
		&CounterFormat{
			Other: "%vY",
			Per:   "%v/Y",
		},
		&CounterFormat{
			"Years",
			"%d year",
			"%d years",
			"%v per year",
		},
	},
	Month: DurationSpecPair{
		Masculine,
		&CounterFormat{
			Other: "%vM",
			Per:   "%v/M",
		},
		&CounterFormat{
			"Months",
			"%d month",
			"%d months",
			"%v per month",
		},
	},
	Day: DurationSpecPair{
		Masculine,
		&CounterFormat{
			Other: "%vd",
			Per:   "%v/d",
		},
		&CounterFormat{
			"Days",
			"%d day",
			"%d days",
			"%v per day",
		},
	},
	Hour: DurationSpecPair{
		Masculine,
		&CounterFormat{
			Other: "%dh",
			Per:   "%v/h",
		},
		&CounterFormat{
			"Hours",
			"%d hour",
			"%d hours",
			"%v per hour",
		},
	},
	Minute: DurationSpecPair{
		Masculine,
		&CounterFormat{
			Other: "%dm",
			Per:   "%v/m",
		},
		&CounterFormat{
			"minutes",
			"%d minute",
			"%d minutes",
			"%v per minute",
		},
	},
	Second: DurationSpecPair{
		Masculine,
		&CounterFormat{
			Label: "sec",
			Other: "%vs",
			Per:   "%v/s",
		},
		&CounterFormat{
			"Seconds",
			"%v second",
			"%v seconds",
			"%v per second",
		},
	},
}

type CounterFormat struct {
	Label, One, Other, Per string
}

func (f *CounterFormat) Dump(pkg string, out io.Writer) {
	ws(out, pkg+"CounterFormat{")
	fmt.Fprintf(out, "%q, %q, %q, %q", f.Label, f.One, f.Other, f.Per)
	ws(out, "}")
}

func (f *CounterFormat) Get(qnt uint64, defaul ...*CounterFormat) (v string) {
	if qnt == 1 && f.One != "" {
		return f.One
	}

	if f.Other != "" {
		return f.Other
	}
	for _, d := range defaul {
		if v = d.Get(qnt); v != "" {
			return
		}
	}
	return
}

func (f *CounterFormat) Perf(qnt any) (v string) {
	return fmt.Sprintf(f.Per, qnt)
}

func (f *CounterFormat) GetF(qnt float64, defaul ...*CounterFormat) (v string) {
	if qnt == 1 && f.One != "" {
		return f.One
	}

	if f.Other != "" {
		return f.Other
	}
	for _, d := range defaul {
		if v = d.GetF(qnt); v != "" {
			return
		}
	}
	return
}

type DurationSpecPair struct {
	Gender      Gender
	Short, Long *CounterFormat
}

func (s *DurationSpecPair) Dump(pkg string, out io.Writer) {
	ws(out, pkg+"DurationSpecPair{")
	if s.Short != nil || s.Long != nil {
		fmt.Fprintf(out, "\nGender: "+pkg+"%s", s.Gender)
		if s.Short != nil && s.Short.One != "" {
			ws(out, ",\nShort: &")
			s.Short.Dump(pkg, out)
		}
		if s.Long != nil && s.Long.One != "" {
			ws(out, ",\nLong: &")
			s.Long.Dump(pkg, out)
		}
	}
	ws(out, "}")
}

func (p *DurationSpecPair) Of(mode DurationMode) *CounterFormat {
	switch mode {
	case DurationShort:
		if p.Short == nil {
			return p.Long
		}
		return p.Short
	default:
		if p.Long == nil {
			return p.Short
		}
		return p.Long
	}
}

type DurationSpec struct {
	LongSep, ShortSep string
	Century           DurationSpecPair
	Decade            DurationSpecPair
	Year              DurationSpecPair
	Month             DurationSpecPair
	Week              DurationSpecPair
	Day               DurationSpecPair
	Hour              DurationSpecPair
	Minute            DurationSpecPair
	Second            DurationSpecPair
	Millisecond       DurationSpecPair
	Microsecond       DurationSpecPair
	Nanosecond        DurationSpecPair
}

func (s *DurationSpec) Dump(pkg string, out io.Writer) {
	ws(out, pkg+"DurationSpec{")
	fmt.Fprintf(out, "%q,\n%q,\n", s.LongSep, s.ShortSep)
	s.Century.Dump(pkg, out)
	ws(out, ",\n")
	s.Decade.Dump(pkg, out)
	ws(out, ",\n")
	s.Year.Dump(pkg, out)
	ws(out, ",\n")
	s.Month.Dump(pkg, out)
	ws(out, ",\n")
	s.Week.Dump(pkg, out)
	ws(out, ",\n")
	s.Day.Dump(pkg, out)
	ws(out, ",\n")
	s.Hour.Dump(pkg, out)
	ws(out, ",\n")
	s.Minute.Dump(pkg, out)
	ws(out, ",\n")
	s.Second.Dump(pkg, out)
	ws(out, ",\n")
	s.Millisecond.Dump(pkg, out)
	ws(out, ",\n")
	s.Microsecond.Dump(pkg, out)
	ws(out, ",\n")
	s.Nanosecond.Dump(pkg, out)
	ws(out, ",\n}")
}

type DurationMode uint8

const (
	DurationShort DurationMode = iota
	DurationLong

	HoursInDay   = 24
	HoursInMonth = HoursInDay * 30
	HoursInYear  = HoursInMonth * 12
)

func (d *DurationSpec) Format(mode DurationMode, dur time.Duration) string {
	var (
		s []string
		f = func(p *DurationSpecPair, v any) {
			var f string
			switch t := v.(type) {
			case uint64:
				if t > 0 {
					f = p.Of(mode).Get(t)
				}
			default:
				if v := v.(float64); v != 0 {
					f = p.Of(mode).GetF(v)
				}
			}
			if f != "" {
				s = append(s, fmt.Sprintf(f, v))
			}
		}
	)

	if d == nil || d.Second.Long.Label == "" {
		d = &DefaultDurationSpec
	}

	if v := dur.Seconds(); v < 60.0 {
		f(&d.Second, v)
	} else if v = dur.Minutes(); v < 60.0 {
		remainingSeconds := math.Mod(dur.Seconds(), 60)

		f(&d.Minute, uint64(v))
		f(&d.Second, remainingSeconds)
	} else if v = dur.Hours(); v < HoursInDay {
		remainingMinutes := math.Mod(dur.Minutes(), 60)
		remainingSeconds := math.Mod(dur.Seconds(), 60)

		f(&d.Hour, uint64(v))
		f(&d.Minute, uint64(remainingMinutes))
		f(&d.Second, remainingSeconds)
	} else if v < HoursInMonth {
		remainingHours := math.Mod(dur.Hours(), 24)
		remainingMinutes := math.Mod(dur.Minutes(), 60)
		remainingSeconds := math.Mod(dur.Seconds(), 60)

		f(&d.Day, uint64(dur.Hours()/24))
		f(&d.Hour, uint64(remainingHours))
		f(&d.Minute, uint64(remainingMinutes))
		f(&d.Second, remainingSeconds)
	} else if v < HoursInYear {
		remainingDays := math.Mod(dur.Hours()/HoursInDay, 30)
		remainingHours := math.Mod(dur.Hours(), 24)
		remainingMinutes := math.Mod(dur.Minutes(), 60)
		remainingSeconds := math.Mod(dur.Seconds(), 60)

		f(&d.Month, uint64(dur.Hours()/HoursInMonth))
		f(&d.Day, uint64(remainingDays))
		f(&d.Hour, uint64(remainingHours))
		f(&d.Minute, uint64(remainingMinutes))
		f(&d.Second, remainingSeconds)
	} else {
		remainingMonths := math.Mod(dur.Hours()/HoursInMonth, 12)
		remainingDays := math.Mod(dur.Hours()/HoursInDay, 30)
		remainingHours := math.Mod(dur.Hours(), 24)
		remainingMinutes := math.Mod(dur.Minutes(), 60)
		remainingSeconds := math.Mod(dur.Seconds(), 60)

		f(&d.Year, uint64(dur.Hours()/HoursInYear))
		f(&d.Month, uint64(remainingMonths))
		f(&d.Day, uint64(remainingDays))
		f(&d.Hour, uint64(remainingHours))
		f(&d.Minute, uint64(remainingMinutes))
		f(&d.Second, remainingSeconds)
	}

	return strings.Join(s, d.LongSep)
}

type DurationTranslator interface {
	GetDurationSpec() *DurationSpec
}

type DurationSpecValue struct {
	D    time.Duration
	Spec *DurationSpec
}

func (d DurationSpecValue) Format(f fmt.State, verb rune) {
	switch verb {
	case 'd':
		fmt.Fprintf(f, "%d", d.D)
	case 's':
		io.WriteString(f, d.Spec.Format(DurationShort, d.D))
	default:
		io.WriteString(f, d.Spec.Format(DurationLong, d.D))
	}
}

type DurationValue struct {
	Translator Translator
	Value      time.Duration
}

func (d DurationValue) Format(f fmt.State, verb rune) {
	switch verb {
	case 'd':
		fmt.Fprintf(f, "%d", d.Value)
	case 's':
		io.WriteString(f, d.Translator.GetDurationSpec().Format(DurationShort, d.Value))
	default:
		io.WriteString(f, d.Translator.GetDurationSpec().Format(DurationLong, d.Value))
	}
}
