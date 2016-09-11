package fa_AF

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type fa_AF struct {
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
	currencyPositivePrefix string
	currencyNegativePrefix string
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

// New returns a new instance of translator for the 'fa_AF' locale
func New() locales.Translator {
	return &fa_AF{
		locale:                 "fa_AF",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                "٫",
		group:                  "٬",
		percent:                "٪",
		perMille:               "؉",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP ", "AED ", "AFA ", "AFN ", "ALK ", "ALL ", "AMD ", "ANG ", "AOA ", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS ", "ATS ", "AUD ", "AWG ", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL ", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "BYR ", "BZD ", "CAD ", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CNY ", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "DOP ", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "EUR ", "FIM ", "FJD ", "FKP ", "FRF ", "GBP ", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HKD ", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "ILS ", "INR ", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JPY ", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "KRW ", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "MDL ", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "MXN ", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZD ", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR ", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON ", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "THB ", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TRY ", "TTD ", "TWD ", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "USD ", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "UZS ", "VEB ", "VEF ", "VND ", "VNN ", "VUV ", "WST ", "XAF ", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD ", "XDR ", "XEU ", "XFO ", "XFU ", "XOF ", "XPD ", "XPF ", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		currencyPositivePrefix: "‎",
		currencyNegativePrefix: "‎",
		monthsAbbreviated:      []string{"", "جنو", "فبروری", "مارچ", "اپریل", "می", "جون", "جول", "اگست", "سپتمبر", "اکتوبر", "نومبر", "دسم"},
		monthsNarrow:           []string{"", "ج", "ف", "م", "ا", "م", "ج", "ج", "ا", "س", "ا", "ن", "د"},
		monthsWide:             []string{"", "جنوری", "فبروری", "مارچ", "اپریل", "می", "جون", "جولای", "اگست", "سپتمبر", "اکتوبر", "نومبر", "دسمبر"},
		daysAbbreviated:        []string{"یکشنبه", "دوشنبه", "سه\u200cشنبه", "چهارشنبه", "پنجشنبه", "جمعه", "شنبه"},
		daysNarrow:             []string{"ی", "د", "س", "چ", "پ", "ج", "ش"},
		daysShort:              []string{"۱ش", "۲ش", "۳ش", "۴ش", "۵ش", "ج", "ش"},
		daysWide:               []string{"یکشنبه", "دوشنبه", "سه\u200cشنبه", "چهارشنبه", "پنجشنبه", "جمعه", "شنبه"},
		periodsAbbreviated:     []string{"ق.ظ.", "ب.ظ."},
		periodsNarrow:          []string{"ق", "ب"},
		periodsWide:            []string{"قبل\u200cازظهر", "بعدازظهر"},
		erasAbbreviated:        []string{"ق.م.", "م."},
		erasNarrow:             []string{"ق", "م"},
		erasWide:               []string{"قبل از میلاد", "میلادی"},
		timezones:              map[string]string{"MDT": "وقت تابستانی ماکائو", "WESZ": "وقت تابستانی غرب اروپا", "HADT": "وقت تابستانی هاوایی‐الوشن", "ChST": "وقت عادی چامورو", "JST": "وقت عادی ژاپن", "AEDT": "وقت تابستانی شرق استرالیا", "ARST": "وقت تابستانی آرژانتین", "BT": "وقت بوتان", "COST": "وقت تابستانی کلمبیا", "ACWDT": "وقت تابستانی مرکز-غرب استرالیا", "LHDT": "وقت تابستانی لردهو", "NZST": "وقت عادی زلاند نو", "AEST": "وقت عادی شرق استرالیا", "VET": "وقت ونزوئلا", "SGT": "وقت سنگاپور", "UYT": "وقت عادی اروگوئه", "GFT": "وقت گویان فرانسه", "MST": "وقت عادی ماکائو", "WEZ": "وقت عادی غرب اروپا", "MESZ": "وقت تابستانی مرکز اروپا", "TMST": "وقت تابستانی ترکمنستان", "HAST": "وقت عادی هاوایی‐الوشن", "PST": "وقت عادی غرب امریکا", "AWST": "وقت عادی غرب استرالیا", "EDT": "وقت تابستانی شرق امریکا", "AKST": "وقت عادی آلاسکا", "ECT": "وقت اکوادور", "ACWST": "وقت عادی مرکز-غرب استرالیا", "JDT": "وقت تابستانی ژاپن", "MYT": "وقت مالزی", "WAST": "وقت تابستانی غرب افریقا", "IST": "وقت هند", "WIT": "وقت شرق اندونزی", "COT": "وقت عادی کلمبیا", "EAT": "وقت شرق افریقا", "SAST": "وقت عادی جنوب افریقا", "WAT": "وقت عادی غرب افریقا", "WARST": "وقت تابستانی غرب آرژانتین", "WITA": "وقت مرکز اندونزی", "CHADT": "وقت تابستانی چت\u200cهام", "CLST": "وقت تابستانی شیلی", "WART": "وقت عادی غرب آرژانتین", "TMT": "وقت عادی ترکمنستان", "NZDT": "وقت تابستانی زلاند نو", "OESZ": "وقت تابستانی شرق اروپا", "ACST": "وقت عادی مرکز استرالیا", "HKST": "وقت تابستانی هنگ\u200cکنگ", "CHAST": "وقت عادی چت\u200cهام", "PDT": "وقت تابستانی غرب امریکا", "SRT": "وقت سورینام", "CST": "وقت عادی مرکز امریکا", "AKDT": "وقت تابستانی آلاسکا", "CLT": "وقت عادی شیلی", "AST": "وقت عادی آتلانتیک", "HNT": "وقت عادی نیوفاندلند", "BOT": "وقت بولیوی", "HKT": "وقت عادی هنگ\u200cکنگ", "WIB": "وقت غرب اندونزی", "OEZ": "وقت عادی شرق اروپا", "GYT": "وقت گویان", "ADT": "وقت تابستانی آتلانتیک", "∅∅∅": "وقت تابستانی آزور", "LHST": "وقت عادی لردهو", "EST": "وقت عادی شرق امریکا", "CAT": "وقت مرکز افریقا", "CDT": "وقت تابستانی مرکز امریکا", "ACDT": "وقت تابستانی مرکز استرالیا", "UYST": "وقت تابستانی اروگوئه", "ART": "وقت عادی آرژانتین", "GMT": "وقت گرینویچ", "MEZ": "وقت عادی مرکز اروپا", "AWDT": "وقت تابستانی غرب استرالیا", "HAT": "وقت تابستانی نیوفاندلند"},
	}
}

// Locale returns the current translators string locale
func (fa *fa_AF) Locale() string {
	return fa.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'fa_AF'
func (fa *fa_AF) PluralsCardinal() []locales.PluralRule {
	return fa.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'fa_AF'
func (fa *fa_AF) PluralsOrdinal() []locales.PluralRule {
	return fa.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'fa_AF'
func (fa *fa_AF) PluralsRange() []locales.PluralRule {
	return fa.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'fa_AF'
func (fa *fa_AF) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'fa_AF'
func (fa *fa_AF) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'fa_AF'
func (fa *fa_AF) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (fa *fa_AF) MonthAbbreviated(month time.Month) string {
	return fa.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (fa *fa_AF) MonthsAbbreviated() []string {
	return fa.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (fa *fa_AF) MonthNarrow(month time.Month) string {
	return fa.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (fa *fa_AF) MonthsNarrow() []string {
	return fa.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (fa *fa_AF) MonthWide(month time.Month) string {
	return fa.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (fa *fa_AF) MonthsWide() []string {
	return fa.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (fa *fa_AF) WeekdayAbbreviated(weekday time.Weekday) string {
	return fa.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (fa *fa_AF) WeekdaysAbbreviated() []string {
	return fa.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (fa *fa_AF) WeekdayNarrow(weekday time.Weekday) string {
	return fa.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (fa *fa_AF) WeekdaysNarrow() []string {
	return fa.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (fa *fa_AF) WeekdayShort(weekday time.Weekday) string {
	return fa.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (fa *fa_AF) WeekdaysShort() []string {
	return fa.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (fa *fa_AF) WeekdayWide(weekday time.Weekday) string {
	return fa.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (fa *fa_AF) WeekdaysWide() []string {
	return fa.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'fa_AF' and handles both Whole and Real numbers based on 'v'
func (fa *fa_AF) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(fa.decimal) + len(fa.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(fa.decimal) - 1; j >= 0; j-- {
				b = append(b, fa.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fa.group) - 1; j >= 0; j-- {
					b = append(b, fa.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, fa.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'fa_AF' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (fa *fa_AF) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(fa.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(fa.decimal) - 1; j >= 0; j-- {
				b = append(b, fa.decimal[j])
			}
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, fa.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, fa.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'fa_AF'
func (fa *fa_AF) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fa.currencies[currency]
	l := len(s) + len(fa.decimal) + len(fa.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(fa.decimal) - 1; j >= 0; j-- {
				b = append(b, fa.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fa.group) - 1; j >= 0; j-- {
					b = append(b, fa.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(fa.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, fa.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, fa.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fa.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'fa_AF'
// in accounting notation.
func (fa *fa_AF) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fa.currencies[currency]
	l := len(s) + len(fa.decimal) + len(fa.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(fa.decimal) - 1; j >= 0; j-- {
				b = append(b, fa.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fa.group) - 1; j >= 0; j-- {
					b = append(b, fa.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(fa.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, fa.currencyNegativePrefix[j])
		}

		b = append(b, fa.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(fa.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, fa.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fa.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'fa_AF'
func (fa *fa_AF) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'fa_AF'
func (fa *fa_AF) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fa.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'fa_AF'
func (fa *fa_AF) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fa.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'fa_AF'
func (fa *fa_AF) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, fa.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fa.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'fa_AF'
func (fa *fa_AF) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'fa_AF'
func (fa *fa_AF) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'fa_AF'
func (fa *fa_AF) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	b = append(b, []byte{0x29}...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'fa_AF'
func (fa *fa_AF) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := fa.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return string(b)
}
