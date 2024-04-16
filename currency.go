package locales

import (
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"

	"github.com/moisespsena-go/locales/currency"
	"github.com/shopspring/decimal"
)

func ws(out io.Writer, v ...string) {
	for _, s := range v {
		io.WriteString(out, s)
	}
}

type (
	CurrencyTranslator interface {
		Currencies() []Currency

		CurrencyFormatters() *CurrencyFormatters
	}

	CurrencyWithNumberTranslator interface {
		CurrencyTranslator
		NumberTranslator
	}

	CurrencySpec struct {
		CurrenciesValue []Currency
		Formatters      *CurrencyFormatters
	}

	CurrencyAccountingFormatterByExpPlural struct {
		Rules map[PluralRule]*NumberFormatProperties
	}

	CurrencyAccountingFormatterByExp struct {
		CurrencyFmt   map[uint8]*CurrencyAccountingFormatterByExpPlural
		AccountingFmt map[uint8]*CurrencyAccountingFormatterByExpPlural
	}

	CurrencyFormatters struct {
		CurrencyFmt,
		AccountingFmt *NumberFormatProperties

		Short CurrencyAccountingFormatterByExp
	}

	Direction           uint8
	CurrencyLevel       uint8
	CurrencySymbolLevel uint8

	CurrencyFormat struct {
		Raw                   string
		ThousandsLen          int
		DigitsLen             int
		SecondaryThousandsLen int
		MinDecimalLen         int
		SymbolDirection       Direction
		SymbolPrefix          string
		SymbolSuffix          string
		Prefix                string
		Suffix                string
	}
)

const (
	Left Direction = iota
	Right
)

const (
	CurrencyLong CurrencyLevel = iota
	CurrencyShort
)

const (
	CurrencySymbolStd CurrencySymbolLevel = 1 + iota
	CurrencySymbolISO
	CurrencySymbolName
	CurrencySymbolNarrow
)

var (
	digitsLenRegex         = regexp.MustCompile("[0-9]+")
	thousandsLenRegex      = regexp.MustCompile(",([0-9#]+)\\.")
	secondaryGroupLenRegex = regexp.MustCompile(",([0-9#]+),")
	requiredDecimalRegex   = regexp.MustCompile("\\.([0-9]+)")
)

func (c *CurrencySpec) Currencies() []Currency {
	return c.CurrenciesValue
}

func (c *CurrencySpec) CurrencyFormatters() *CurrencyFormatters {
	return c.Formatters
}

func (t *CurrencyAccountingFormatterByExpPlural) Dump(pkg string, out io.Writer) {
	ws(out, pkg+"CurrencyAccountingFormatterByExpPlural{Rules:map["+pkg+"PluralRule]*"+pkg+"NumberFormatProperties{\n")

	var (
		i    int
		keys = make([]PluralRule, len(t.Rules))
	)

	for k := range t.Rules {
		keys[i] = k
		i++
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, k := range keys {
		ws(out, pkg+"PluralRule"+k.String()+": ")
		t.Rules[k].Dump(pkg, out)
		ws(out, ",\n")
	}

	ws(out, "}}")
}

func (f *CurrencyAccountingFormatterByExp) Dump(pkg string, out io.Writer) {
	dump := func(m map[uint8]*CurrencyAccountingFormatterByExpPlural) {
		ws(out, "map[uint8]*"+pkg+"CurrencyAccountingFormatterByExpPlural {\n")

		var (
			i    int
			keys = make([]uint8, len(m))
		)

		for k := range m {
			keys[i] = k
			i++
		}

		sort.Slice(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})

		for _, k := range keys {
			ws(out, strconv.FormatUint(uint64(k), 10)+": &")
			m[k].Dump(pkg, out)
			ws(out, ",\n")
		}

		ws(out, "}")
	}

	ws(out, pkg+"CurrencyAccountingFormatterByExp{\n")

	if f.CurrencyFmt != nil {
		ws(out, "CurrencyFmt: ")
		dump(f.CurrencyFmt)
		ws(out, ",\n")
	}

	if f.AccountingFmt != nil {
		ws(out, "AccountingFmt: ")
		dump(f.AccountingFmt)
		ws(out, ",\n")
	}

	ws(out, "}")
}

func (f *CurrencyFormatters) Dump(pkg string, out io.Writer) {
	ws(out, pkg+"CurrencyFormatters{\n")

	if f.CurrencyFmt != nil {
		ws(out, "CurrencyFmt: ")
		f.CurrencyFmt.Dump(pkg, out)
		ws(out, ",\n")
	}
	if f.AccountingFmt != nil {
		ws(out, "AccountingFmt: ")
		f.AccountingFmt.Dump(pkg, out)
		ws(out, ",\n")
	}
	if len(f.Short.CurrencyFmt) > 0 || len(f.Short.AccountingFmt) > 0 {
		ws(out, "Short: ")
		f.Short.Dump(pkg, out)
		ws(out, ",\n")
	}
	ws(out, "}")
}

func (f *CurrencyFormatters) IsZero() bool {
	return f.CurrencyFmt == nil && f.AccountingFmt == nil && len(f.Short.CurrencyFmt) == 0 && len(f.Short.AccountingFmt) == 0
}

func (f *CurrencyFormatters) GetCurrency(level CurrencyLevel, num decimal.Decimal) (cf *NumberFormatProperties, n decimal.Decimal) {
	return f.Get(f.CurrencyFmt, f.Short.CurrencyFmt, level, num)
}

func (f *CurrencyFormatters) Get(defaul *NumberFormatProperties, short map[uint8]*CurrencyAccountingFormatterByExpPlural, level CurrencyLevel, num decimal.Decimal) (cf *NumberFormatProperties, n decimal.Decimal) {
	cf = defaul
	n = num
	if level == CurrencyShort && short != nil {
		var (
			intS = strconv.FormatInt(n.IntPart(), 10)
			neg  = intS[0] == '-'
			exp  = uint8(len(intS)) - 1
		)

		if neg {
			exp--
		}

		for exp > 0 {
			if v, ok := short[exp]; ok {
				if n.IntPart() == 1 {
					cf = v.Rules[PluralRuleOne]
				} else {
					cf = v.Rules[PluralRuleOther]
				}
				fmt.Println(n)
				return
			}
			n = n.Div(decimal.NewFromInt(10))
			exp--
		}
		n = num
	}
	return
}

func (f *CurrencyFormatters) Currency(t CurrencyWithNumberTranslator, level CurrencyLevel, cur *Currency, num decimal.Decimal, v int) string {
	cf, num := f.GetCurrency(level, num)
	return f.Format(cf, t, cur, num, v)
}

func (f *CurrencyFormatters) GetAccounting(level CurrencyLevel, num decimal.Decimal) (cf *NumberFormatProperties, n decimal.Decimal) {
	cf, num = f.Get(f.AccountingFmt, f.Short.AccountingFmt, level, num)
	if cf == nil {
		cf, num = f.GetCurrency(level, num)
	}
	return cf, num
}

func (f *CurrencyFormatters) Accounting(t CurrencyWithNumberTranslator, level CurrencyLevel, cur *Currency, num decimal.Decimal, v int) string {
	cf, num := f.GetAccounting(level, num)
	return f.Format(cf, t, cur, num, v)
}

func (f *CurrencyFormatters) Format(cf *NumberFormatProperties, t CurrencyWithNumberTranslator, cur *Currency, num decimal.Decimal, precision int) string {
	return cf.Format(num, NumberFormatOptions{
		Precision: precision,
		Decimal:   t.Decimal(),
		Group:     t.Group(),
		Minus:     t.Minus(),
		Symbol:    cur.Symbols.Default,
	})
}

func (f *CurrencyFormat) Dump(pkg string, out io.Writer) {
	ws(out, pkg+"CurrencyFormat{\n")
	ws(out, "Raw: "+strconv.Quote(f.Raw)+",\n")

	if f.DigitsLen > 0 {
		fmt.Fprintf(out, "DigitsLen: %d,\n", f.DigitsLen)
	}
	if f.ThousandsLen > 0 {
		fmt.Fprintf(out, "ThousandsLen: %d,\n", f.ThousandsLen)
	}
	if f.SecondaryThousandsLen > 0 {
		fmt.Fprintf(out, "SecondaryThousandsLen: %d,\n", f.SecondaryThousandsLen)
	}
	if f.MinDecimalLen > 0 {
		fmt.Fprintf(out, "MinDecimalLen: %d,\n", f.MinDecimalLen)
	}
	if f.SymbolDirection == Right {
		ws(out, "SymbolDirection: "+pkg+".Right,\n")
	}
	if f.SymbolPrefix != "" {
		ws(out, "SymbolPrefix: "+strconv.Quote(f.SymbolPrefix)+",\n")
	}
	if f.SymbolSuffix != "" {
		ws(out, "SymbolSuffix: "+strconv.Quote(f.SymbolSuffix)+",\n")
	}
	if f.Prefix != "" {
		ws(out, "Prefix: "+strconv.Quote(f.Prefix)+",\n")
	}
	if f.Suffix != "" {
		ws(out, "Suffix: "+strconv.Quote(f.Suffix)+",\n")
	}

	ws(out, "}")
}

func (f *CurrencyFormat) Parse(numberFormat string) {
	*f = CurrencyFormat{Raw: numberFormat}

	match := thousandsLenRegex.FindString(numberFormat)
	if len(match) > 0 {
		f.ThousandsLen = len(match) - 2
	}

	match = requiredDecimalRegex.FindString(numberFormat)
	if len(match) > 0 {
		f.MinDecimalLen = len(match) - 1
	}

	match = secondaryGroupLenRegex.FindString(numberFormat)
	if len(match) > 0 {
		f.SecondaryThousandsLen = len(match) - 2
	}

	if f.ThousandsLen == 0 && f.MinDecimalLen == 0 && f.SecondaryThousandsLen == 0 {
		match = digitsLenRegex.FindString(numberFormat)
		if len(match) > 0 {
			f.DigitsLen = len(match)
		}
	}

	idx := 0

	for idx = 0; idx < len(numberFormat); idx++ {
		if numberFormat[idx] == '#' || numberFormat[idx] == '0' {
			f.Prefix = numberFormat[:idx]
			break
		}
	}

	for idx = len(numberFormat) - 1; idx >= 0; idx-- {
		if numberFormat[idx] == '#' || numberFormat[idx] == '0' {
			idx++
			f.Suffix = numberFormat[idx:]
			break
		}
	}

	var runes = []rune(f.Prefix)

	for idx = 0; idx < len(runes); idx++ {
		if runes[idx] == '¤' {
			f.SymbolPrefix = string(runes[:idx])
			f.Prefix = string(runes[idx+1:])
			break
		}
	}

	runes = []rune(f.Suffix)
	for idx = 0; idx < len(runes); idx++ {
		if runes[idx] == '¤' {
			f.Suffix = string(runes[:idx])
			f.SymbolDirection = Right
			f.SymbolSuffix = string(runes[idx+1:])
			break
		}
	}

	return
}

func (f *CurrencyFormat) Exec(symbol, value string) string {
	if f.SymbolDirection == Left {
		return f.SymbolPrefix + symbol + f.SymbolSuffix + f.Prefix + value + f.Suffix
	}
	return f.Prefix + value + f.Suffix + f.SymbolPrefix + symbol + f.SymbolSuffix
}

// Format returns the currency representation of 'num' with digits/precision of 'v' for '{{ .Locale }}'
func (f *CurrencyFormat) Format(t NumberTranslator, cur *Currency, num string, v uint64) string {
	val := FormatNumberString(num, int(v), t.Minus(), t.Group(), t.Decimal())
	return f.Exec(cur.Symbols.Default, val)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'pt'
func FmtCurrency(t CurrencyWithNumberTranslator, level CurrencyLevel, currency currency.Type, num decimal.Decimal, precision int) string {
	return t.CurrencyFormatters().Currency(t, level, &t.Currencies()[currency.Index()], num, precision)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'pt'
// in accounting notation.
func FmtAccounting(t CurrencyWithNumberTranslator, level CurrencyLevel, currency currency.Type, num decimal.Decimal, precision int) string {
	return t.CurrencyFormatters().Accounting(t, level, &t.Currencies()[currency.Index()], num, precision)
}
