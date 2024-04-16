package locales

import (
	"errors"
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/shopspring/decimal"
)

type NumberFlag uint8

const (
	PercentNumber NumberFlag = 1 + iota
	MilleNumber
	CurrencyNumber
)

type NumberPatternProperties struct {
	// Pattern Raw pattern value
	Pattern string

	Prefix string
	Padding,
	ScientificNotation string

	// MinimumIntegerDigits non-negative integer Number value indicating the minimum integer
	//        digits to be used. Numbers will be padded with leading zeroes if necessary.
	MinimumIntegerDigits,

	// MinimumFractionDigits and MaximumFractionDigits are non-negative integer Number values indicating the minimum and
	//	//        maximum fraction digits to be used. Numbers will be rounded or padded with trailing
	//	//        zeroes if necessary.
	MinimumFractionDigits,
	MaximumFractionDigits,

	// MinimumSignificantDigits and MaximumSignificantDigits are positive integer Number values indicating the minimum and
	//        maximum fraction digits to be shown. Either none or both of these properties are
	//        present; if they are, they override minimum and maximum integer and fraction digits
	//        – the formatter uses however many integer and fraction digits are required to display
	//        the specified number of significant digits.
	MinimumSignificantDigits,
	MaximumSignificantDigits,
	PrimaryGroupingSize,
	SecondaryGroupingSize int
	Suffix string

	// RoundIncrement Decimal round increment or null
	RoundIncrement float64

	Type NumberFlag

	CurrencySymbol CurrencySymbolLevel

	DivExp uint8
}

func (f NumberPatternProperties) Dump(pkg string, out io.Writer) {
	var s = []string{f.Pattern}
	if f.DivExp > 0 {
		s = append(s, strconv.FormatUint(uint64(f.DivExp), 10))
	}
	ws(out, pkg+"MustParseNumberPattern(")
	ws(out, strconv.Quote(strings.Join(s, "÷")))
	ws(out, ")")
}

func (f NumberPatternProperties) Neg() NumberPatternProperties {
	if len(f.Prefix) == 0 || f.Prefix[len(f.Prefix)-1] != '-' {
		f.Prefix += "-"
	} else {
		f.Prefix = f.Prefix[:len(f.Prefix)-1]
	}

	rp := []rune(f.Pattern)

	for i, r := range rp {
		switch r {
		case '0', '#':
			if i > 0 && rp[i-1] == '-' {
				f.Pattern = string(append(rp[:i-1], rp[i:]...))
				return f
			}
			f.Pattern = string(append(rp[:i], append([]rune{'-'}, rp[i:]...)...))
			return f
		}
	}
	return f
}

/**
 * format( pattern )
 *
 * @pattern [String] raw pattern for numbers.
 *
 * Return the formatted number.
 * ref: http://www.unicode.org/reports/tr35/tr35-numbers.html
 */
func (p *NumberPatternProperties) Parse(pattern string) (err error) {
	var (
		m = numberPatternRe.FindStringSubmatch(pattern)

		fractionPattern,
		integerFractionOrSignificantPattern,
		integerPattern,
		padding,
		prefix,
		significantPattern,
		scientificNotation,
		suffix string

		roundIncrement float64

		minimumIntegerDigits,
		primaryGroupingSize,
		secondaryGroupingSize,
		maximumFractionDigits,
		maximumSignificantDigits,
		minimumFractionDigits,
		minimumSignificantDigits int
	)
	if len(m) != 13 {
		return errors.New("Invalid pattern: " + pattern)
	}

	prefix = m[1]
	padding = m[4]
	integerFractionOrSignificantPattern = m[5]
	significantPattern = m[9]
	scientificNotation = m[10]
	suffix = m[11]
	rsuffix := []rune(suffix)
	for i, s := range rsuffix {
		if s == '÷' {
			var (
				i2 = i + 1
				r2 rune
			)
			for _, r2 = range rsuffix[i+1:] {
				if !unicode.IsDigit(r2) {
					break
				}
				i2++
			}
			var u uint64
			if u, err = strconv.ParseUint(string(rsuffix[i+1:i2]), 10, 8); err != nil {
				return
			}
			suffix = string(rsuffix[:i])
			p.DivExp = uint8(u)
		}
	}

	if significantPattern != "" {
		m := regexp.MustCompile(`(@+)(#*)`).FindStringSubmatch(significantPattern)
		minimumSignificantDigits = len(m[1])
		maximumSignificantDigits = minimumSignificantDigits + len(m[2])
	} else {
		fractionPattern = m[8]
		integerPattern = m[7]

		if fractionPattern != "" {
			s := regexp.MustCompile(`[0-9]+`).FindString(fractionPattern)
			if s != "" {
				f, err := strconv.ParseFloat("0."+s, 10)
				if err != nil {
					panic(err)
				}
				roundIncrement = +f
				minimumFractionDigits = len(s)
			}

			// Maximum fraction digits
			// 1: ignore decimal character
			maximumFractionDigits = len(fractionPattern) - 1 /* 1 */
		}

		minimumIntegerDigits = len(regexp.MustCompile(`0+$`).FindString(integerPattern))
	}

	if padding != "" {
		return errors.New("padding (not implemented)")
	}

	// Grouping
	if aux1 := strings.LastIndexByte(integerFractionOrSignificantPattern, ','); aux1 != -1 {

		// Primary grouping size is the interval between the last group separator and the end of
		// the integer (or the end of the significant pattern).
		aux2 := strings.Split(integerFractionOrSignificantPattern, ".")[0]
		primaryGroupingSize = len(aux2) - aux1 - 1

		// Secondary grouping size is the interval between the last two group separators.
		if aux3 := strings.LastIndexByte(integerFractionOrSignificantPattern[aux1-1:], ','); aux3 != -1 {
			secondaryGroupingSize = aux1 - 1 - aux3
		}
	}

	if minimumSignificantDigits > 0 && maximumSignificantDigits > 0 {
		if err = validateParameterRange("minimumSignificantDigits", minimumSignificantDigits, 1, 21); err != nil {
			return
		}
		if err = validateParameterRange("maximumSignificantDigits", maximumSignificantDigits, minimumSignificantDigits, 21); err != nil {
			return
		}
	} else if minimumSignificantDigits > 0 || maximumSignificantDigits > 0 {
		return errors.New("Neither or both the minimum and maximum significant digits must be present")
	} else {
		if err = validateParameterRange("minimumIntegerDigits", minimumIntegerDigits, 1, 21); err != nil {
			return
		}
		if err = validateParameterRange("minimumFractionDigits", minimumFractionDigits, 0, 21); err != nil {
			return
		}
		if err = validateParameterRange("maximumFractionDigits", maximumFractionDigits, minimumFractionDigits, 20); err != nil {
			return
		}
	}

	if strings.IndexByte(pattern, '%') != -1 {
		p.Type = PercentNumber
	} else if strings.IndexRune(pattern, '\u2030') != -1 {
		p.Type = MilleNumber
	} else if i := strings.IndexRune(pattern, '¤'); i != -1 {
		p.Type = CurrencyNumber
		var symbols []rune

		for _, r := range []rune(pattern[i:]) {
			if r == '¤' {
				symbols = append(symbols, r)
				p.CurrencySymbol++
			} else {
				break
			}
		}

		prefix = strings.ReplaceAll(prefix, string(symbols), "¤")
		suffix = strings.ReplaceAll(suffix, string(symbols), "¤")
	}

	p.Pattern = pattern
	p.Prefix = prefix
	p.Padding = padding
	p.MinimumIntegerDigits = minimumIntegerDigits
	p.MinimumFractionDigits = minimumFractionDigits
	p.MaximumFractionDigits = maximumFractionDigits
	p.MinimumSignificantDigits = minimumSignificantDigits
	p.MaximumSignificantDigits = maximumSignificantDigits
	p.RoundIncrement = roundIncrement
	p.PrimaryGroupingSize = primaryGroupingSize
	p.SecondaryGroupingSize = secondaryGroupingSize
	p.ScientificNotation = scientificNotation
	p.Suffix = suffix
	return
}

type NumberFormatProperties struct {
	Pos,
	Neg NumberPatternProperties
}

func (f NumberFormatProperties) String() string {
	if f.Pos.Neg().Pattern == f.Neg.Pattern {
		return f.Pos.Pattern
	}
	return f.Pos.Pattern + ";" + f.Neg.Pattern
}

func (f NumberFormatProperties) Dump(pkg string, out io.Writer) {
	ws(out, pkg+"MustParseNumberFormatPatterns("+strconv.Quote(f.String())+")")
}

func (p *NumberFormatProperties) Parse(value string) (err error) {
	var (
		parts    = strings.Split(value, ";")
		pos, neg NumberPatternProperties
	)

	if err = pos.Parse(parts[0]); err != nil {
		return
	}

	if len(parts) == 1 || parts[1] == "" {
		neg = pos.Neg()
	} else if err = neg.Parse(parts[1]); err != nil {
		return
	}

	p.Pos = pos
	p.Neg = neg
	return
}

func MustParseNumberFormatPatterns(pattern string) *NumberFormatProperties {
	var p NumberFormatProperties
	if err := p.Parse(pattern); err != nil {
		panic(err)
	}
	return &p
}

func (f *NumberFormatProperties) Format(value decimal.Decimal, opts NumberFormatOptions) string {
	if value.IsZero() {
		return opts.Zero
	}
	if value.IsNegative() {
		return f.Neg.Format(value, opts)
	}
	return f.Pos.Format(value, opts)
}

func (f *NumberFormatProperties) FormatS(s string, opts NumberFormatOptions) (ret string) {
	if s[0] == '-' {
		return f.Neg.FormatS(s, opts)
	}
	return f.Pos.FormatS(s, opts)
}

func (f *NumberFormatProperties) FormatAny(value any, opts NumberFormatOptions) (s string, err error) {
	var d decimal.Decimal
try:
	switch t := value.(type) {
	case string:
		s = f.FormatS(t, opts)
		return
	case uint8:
		value = uint64(t)
		goto try
	case uint16:
		value = uint64(t)
		goto try
	case uint32:
		value = uint64(t)
		goto try
	case uint64:
		if d, err = decimal.NewFromString(strconv.FormatUint(t, 10)); err != nil {
			return
		}
	case float32:
		value = float64(t)
		goto try
	case float64:
		d = decimal.NewFromFloat(t)
	case decimal.Decimal:
		d = t
	}
	s = f.Format(d, opts)
	return
}

func (f *NumberPatternProperties) Template(opts NumberFormatOptions) (prefix, suffix string, precision int) {
	prefix, suffix, precision = f.Prefix, f.Suffix, opts.Precision
	if precision == 0 && !opts.DefaultPrecisionDisabled {
		precision = f.MinimumFractionDigits
	}
	switch f.Type {
	case CurrencyNumber:
		prefix = strings.Replace(prefix, "¤", opts.Symbol, 1)
		suffix = strings.Replace(suffix, "¤", opts.Symbol, 1)
	}
	return
}

func (f *NumberPatternProperties) FormatS(s string, opts NumberFormatOptions) (ret string) {
	var (
		neg            = s[0] == '-'
		prefix, suffix = f.Prefix, f.Suffix
	)

	// Remove the possible number minus sign
	if neg {
		s = s[1:]
	}

	parts := strings.Split(s, ".")

	if neg {
		if pos := strings.IndexByte(prefix, '-'); pos >= 0 {
			prefix = prefix[0:pos] + opts.Minus + prefix[pos+1:]
		} else if pos = strings.IndexByte(suffix, '-'); pos >= 0 {
			suffix = suffix[0:pos] + opts.Minus + suffix[pos+1:]
		}
	}

	if len(parts) == 1 {
		parts = append(parts, "")
	}

	// Minimum fraction digits (post string phase)
	if f.MinimumFractionDigits > 0 {
		parts[1] = stringPad(parts[1], f.MinimumFractionDigits, true)
	}

	// Minimum integer digits
	if f.MinimumIntegerDigits > 0 {
		parts[0] = stringPad(parts[0], f.MinimumIntegerDigits, false)
	}

	if len(parts[0]) > f.PrimaryGroupingSize {
		parts[0] = numberFormatGroupingSeparator(parts[0], opts.Group, f.PrimaryGroupingSize, f.SecondaryGroupingSize)
	}

	switch f.Type {
	case CurrencyNumber:
		prefix = strings.Replace(prefix, "¤", opts.Symbol, 1)
		suffix = strings.Replace(suffix, "¤", opts.Symbol, 1)
	}

	if parts[1] == "" {
		return strings.TrimSpace(prefix + parts[0] + suffix)
	}
	return strings.TrimSpace(prefix + strings.Join(parts, opts.Decimal) + suffix)
}

func (f *NumberPatternProperties) Format(value decimal.Decimal, opts NumberFormatOptions) (ret string) {
	if f.DivExp > 0 {
		value = value.Div(decimal.NewFromInt(int64(math.Pow10(int(f.DivExp)))))
	}

	if f.MinimumSignificantDigits > 0 && f.MaximumSignificantDigits > 0 {
		panic("NumberPatternProperties.Format: significantDigits (not implemented)")
	}

	if opts.Precision == 0 {
		opts.Precision = f.MaximumFractionDigits
	}

	if opts.Precision > 0 {
		if f.RoundIncrement > 0 {
			value = value.Round(int32(f.RoundIncrement))
		} else {
			value = value.Round(int32(opts.Precision))
		}
	}

	return f.FormatS(value.String(), opts)
}

type NumberFormatOptions struct {
	Minus                    string
	Group                    string
	Decimal                  string
	Symbol                   string
	Infinity                 string
	Percent                  string
	Zero                     string
	Precision                int
	DefaultPrecisionDisabled bool
}

var ComputedNumberFormatOptions = NumberFormatOptions{
	Decimal: ".",
	Group:   "",
	Minus:   "-",
}

/**
 * EBNF representation:
 *
 * number_pattern_re =        prefix?
 *                            padding?
 *                            (integer_fraction_pattern | significant_pattern)
 *                            scientific_notation?
 *                            suffix?
 *
 * prefix =                   non_number_stuff
 *
 * padding =                  "*" regexp(.)
 *
 * integer_fraction_pattern = integer_pattern
 *                            fraction_pattern?
 *
 * integer_pattern =          regexp([#,]*[0,]*0+)
 *
 * fraction_pattern =         "." regexp(0*[0-9]*#*)
 *
 * significant_pattern =      regexp([#,]*@+#*)
 *
 * scientific_notation =      regexp(E\+?0+)
 *
 * suffix =                   non_number_stuff
 *
 * non_number_stuff =         regexp(('[^']+'|''|[^*#@0,.E])*)
 *
 *
 * Regexp groups:
 *
 *  0: number_pattern_re
 *  1: prefix
 *  2: -
 *  3: -
 *  4: padding
 *  5: (integer_fraction_pattern | significant_pattern)
 *  6: integer_fraction_pattern
 *  7: integer_pattern
 *  8: fraction_pattern
 *  9: significant_pattern
 * 10: scientific_notation
 * 11: suffix
 * 12: -
 */
var numberPatternRe = regexp.MustCompile(`^(('([^']|'')*'|[^*#@0,.E])*)(\*.)?((([#,]*[0,]*0+)(\.0*[0-9]*#*)?)|([#,]*@+#*)) ?(E[\+\-]?0*)?(('[^']+'|''|[^*#@0,.E])*)$`)

/**
 * range( value, name, minimum, maximum )
 *
 * @value [Number].
 *
 * @name [String] name of variable.
 *
 * @minimum [Number]. The lowest valid value, inclusive.
 *
 * @maximum [Number]. The greatest valid value, inclusive.
 */
func validateParameterRange[T int | float32 | float64](name string, value, minimum, maximum T) (err error) {
	if value < 0 || value >= minimum && value <= maximum {
		return nil
	}
	return fmt.Errorf("E_PAR_OUT_OF_RANGE: Parameter %q has value %v out of range [%v, %v].", name, value, minimum, maximum)
}

/**
 * toPrecision( number, precision, round )
 *
 * @number (Number)
 *
 * @precision (Number) significant figures precision (not decimal precision).
 *
 * @round (Function)
 *
 * Return number.toPrecision( precision ) using the given round function.
 */
func numberToPrecision(number float64, precision int64, round func(f float64, exp *float64) float64) float64 {
	var roundOrder float64

	if number == 0 {
		return number
	}

	roundOrder = math.Ceil(math.Log(math.Abs(number)) / math.Log(10))
	roundOrder -= float64(precision)

	return round(number, &roundOrder)
}

func stringPad(str string, count int, right bool) string {
	for len(str) < count {
		if right {
			str += "0"
		} else {
			str = "0" + str
		}
	}
	return str
}

func numberFormatGroupingSeparator(intPart string, sep string, primaryGroupingSize, secondaryGroupingSize int) (ret string) {
	var (
		currentGroupingSize = primaryGroupingSize
		switchToSecondary   = secondaryGroupingSize > 0
		buf                 []byte
		in                  = []byte(intPart)
		i                   int
	)

	for i, j := 0, len(in)-1; i <= j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}

	for len(in) > 0 {
		if i == currentGroupingSize {
			i = 0
			if switchToSecondary {
				currentGroupingSize = secondaryGroupingSize
			}
			switchToSecondary = false
			buf = append(buf, []byte(sep)...)
		}

		buf = append(buf, in[0])
		in = in[1:]
		i++
	}

	for i, j := 0, len(buf)-1; i <= j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}

	return string(buf)
}
