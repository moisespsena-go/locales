package pl_PL

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type pl_PL struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositiveSuffix string
	currencyNegativePrefix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'pl_PL' locale
func New() locales.Translator {
	return &pl_PL{
		locale:                 "pl_PL",
		pluralsCardinal:        []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{5, 6, 2, 4},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP ", "AED ", "AFA ", "AFN ", "ALK ", "ALL ", "AMD ", "ANG ", "AOA ", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS ", "ATS ", "AUD ", "AWG ", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL ", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "BYR ", "BZD ", "CAD ", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CNY ", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "DOP ", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "EUR ", "FIM ", "FJD ", "FKP ", "FRF ", "GBP ", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HKD ", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "ILS ", "INR ", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JPY ", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "KRW ", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "MDL ", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "MXN ", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZD ", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR ", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON ", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "THB ", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TRY ", "TTD ", "TWD ", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "USD ", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "UZS ", "VEB ", "VEF ", "VND ", "VNN ", "VUV ", "WST ", "XAF ", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD ", "XDR ", "XEU ", "XFO ", "XFU ", "XOF ", "XPD ", "XPF ", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "sty", "lut", "mar", "kwi", "maj", "cze", "lip", "sie", "wrz", "paź", "lis", "gru"},
		monthsNarrow:           []string{"", "s", "l", "m", "k", "m", "c", "l", "s", "w", "p", "l", "g"},
		monthsWide:             []string{"", "stycznia", "lutego", "marca", "kwietnia", "maja", "czerwca", "lipca", "sierpnia", "września", "października", "listopada", "grudnia"},
		daysAbbreviated:        []string{"niedz.", "pon.", "wt.", "śr.", "czw.", "pt.", "sob."},
		daysNarrow:             []string{"N", "P", "W", "Ś", "C", "P", "S"},
		daysShort:              []string{"nie", "pon", "wto", "śro", "czw", "pią", "sob"},
		daysWide:               []string{"niedziela", "poniedziałek", "wtorek", "środa", "czwartek", "piątek", "sobota"},
		periodsAbbreviated:     []string{"AM", "PM"},
		periodsNarrow:          []string{"a", "p"},
		periodsWide:            []string{"AM", "PM"},
		erasAbbreviated:        []string{"p.n.e.", "n.e."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"p.n.e.", "n.e."},
		timezones:              map[string]string{"CLST": "Chile (czas letni)", "VET": "Wenezuela", "HAT": "Nowa Fundlandia (czas letni)", "GYT": "Gujana", "GFT": "Gujana Francuska", "LHST": "Lord Howe (czas standardowy)", "ACWDT": "Czas środkowo-zachodnioaustralijski letni", "CHADT": "Chatham (czas letni)", "JST": "Japonia (czas standardowy)", "AEDT": "Czas wschodnioaustraliski letni", "PST": "Czas pacyficzny standardowy", "∅∅∅": "Azory (czas letni)", "WIT": "Indonezja Wschodnia", "EST": "Czas wschodnioamerykański standardowy", "ChST": "Czamorro", "ACST": "Czas środkowoaustralijski standardowy", "WARST": "Argentyna Zachodnia (czas letni)", "MYT": "Malezja", "HNT": "Nowa Fundlandia (czas standardowy)", "OESZ": "Czas wschodnioeuropejski letni", "ACDT": "Czas środkowoaustralijski letni", "TMT": "Turkmenistan (czas standardowy)", "WESZ": "Czas zachodnioeuropejski letni", "ACWST": "Czas środkowo-zachodnioaustralijski standardowy", "MESZ": "Czas środkowoeuropejski letni", "SGT": "Singapur", "AWST": "Czas zachodnioaustralijski standardowy", "AWDT": "Czas zachodnioaustralijski letni", "MEZ": "Czas środkowoeuropejski standardowy", "EAT": "Czas wschodnioafrykański", "WAST": "Czas zachodnioafrykański letni", "MDT": "Czas górski letni", "HKST": "Hongkong (czas letni)", "ARST": "Argentyna (czas letni)", "CDT": "Czas środkowoamerykański letni", "AKDT": "Alaska (czas letni)", "LHDT": "Lord Howe (czas letni)", "WART": "Argentyna Zachodnia (czas standardowy)", "ECT": "Ekwador", "UYT": "Urugwaj (czas standardowy)", "WITA": "Indonezja Środkowa", "AKST": "Alaska (czas standardowy)", "PDT": "Czas pacyficzny letni", "ART": "Argentyna (czas standardowy)", "WIB": "Indonezja Zachodnia", "OEZ": "Czas wschodnioeuropejski standardowy", "BT": "Bhutan", "GMT": "Czas uniwersalny", "NZST": "Nowa Zelandia (czas standardowy)", "CHAST": "Chatham (czas standardowy)", "JDT": "Japonia (czas letni)", "AEST": "Czas standardowy wschodnioaustraliski", "MST": "Czas górski standardowy", "COT": "Kolumbia (czas standardowy)", "COST": "Kolumbia (czas letni)", "ADT": "Czas atlantycki letni", "HADT": "Hawaje-Aleuty (czas letni)", "NZDT": "Nowa Zelandia (czas letni)", "WEZ": "Czas zachodnioeuropejski standardowy", "BOT": "Boliwia", "SRT": "Surinam", "HKT": "Hongkong (czas standardowy)", "TMST": "Turkmenistan (czas letni)", "CST": "Czas środkowoamerykański standardowy", "HAST": "Hawaje-Aleuty (czas standardowy)", "IST": "Czas indyjski", "EDT": "Czas wschodnioamerykański letni", "SAST": "Czas południowoafrykański", "AST": "Czas standardowy atlantycki", "CAT": "Czas środkowoafrykański", "UYST": "Urugwaj (czas letni)", "WAT": "Czas zachodnioafrykański standardowy", "CLT": "Chile (czas standardowy)"},
	}
}

// Locale returns the current translators string locale
func (pl *pl_PL) Locale() string {
	return pl.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'pl_PL'
func (pl *pl_PL) PluralsCardinal() []locales.PluralRule {
	return pl.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'pl_PL'
func (pl *pl_PL) PluralsOrdinal() []locales.PluralRule {
	return pl.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'pl_PL'
func (pl *pl_PL) PluralsRange() []locales.PluralRule {
	return pl.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'pl_PL'
func (pl *pl_PL) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	iMod10 := i % 10
	iMod100 := i % 100

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	} else if v == 0 && iMod10 >= 2 && iMod10 <= 4 && iMod100 < 12 && iMod100 > 14 {
		return locales.PluralRuleFew
	} else if (v == 0 && i != 1 && iMod10 >= 0 && iMod10 <= 1) || (v == 0 && iMod10 >= 5 && iMod10 <= 9) || (v == 0 && iMod100 >= 12 && iMod100 <= 14) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'pl_PL'
func (pl *pl_PL) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'pl_PL'
func (pl *pl_PL) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := pl.CardinalPluralRule(num1, v1)
	end := pl.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (pl *pl_PL) MonthAbbreviated(month time.Month) string {
	return pl.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (pl *pl_PL) MonthsAbbreviated() []string {
	return pl.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (pl *pl_PL) MonthNarrow(month time.Month) string {
	return pl.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (pl *pl_PL) MonthsNarrow() []string {
	return pl.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (pl *pl_PL) MonthWide(month time.Month) string {
	return pl.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (pl *pl_PL) MonthsWide() []string {
	return pl.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (pl *pl_PL) WeekdayAbbreviated(weekday time.Weekday) string {
	return pl.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (pl *pl_PL) WeekdaysAbbreviated() []string {
	return pl.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (pl *pl_PL) WeekdayNarrow(weekday time.Weekday) string {
	return pl.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (pl *pl_PL) WeekdaysNarrow() []string {
	return pl.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (pl *pl_PL) WeekdayShort(weekday time.Weekday) string {
	return pl.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (pl *pl_PL) WeekdaysShort() []string {
	return pl.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (pl *pl_PL) WeekdayWide(weekday time.Weekday) string {
	return pl.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (pl *pl_PL) WeekdaysWide() []string {
	return pl.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'pl_PL' and handles both Whole and Real numbers based on 'v'
func (pl *pl_PL) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(pl.decimal) + len(pl.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(pl.group) - 1; j >= 0; j-- {
					b = append(b, pl.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'pl_PL' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (pl *pl_PL) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(pl.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pl.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, pl.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'pl_PL'
func (pl *pl_PL) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pl.currencies[currency]
	l := len(s) + len(pl.decimal) + len(pl.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(pl.group) - 1; j >= 0; j-- {
					b = append(b, pl.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, pl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, pl.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'pl_PL'
// in accounting notation.
func (pl *pl_PL) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pl.currencies[currency]
	l := len(s) + len(pl.decimal) + len(pl.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(pl.group) - 1; j >= 0; j-- {
					b = append(b, pl.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, pl.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, pl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, pl.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, pl.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'pl_PL'
func (pl *pl_PL) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'pl_PL'
func (pl *pl_PL) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'pl_PL'
func (pl *pl_PL) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'pl_PL'
func (pl *pl_PL) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, pl.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'pl_PL'
func (pl *pl_PL) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'pl_PL'
func (pl *pl_PL) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'pl_PL'
func (pl *pl_PL) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'pl_PL'
func (pl *pl_PL) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := pl.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
