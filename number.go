package locales

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type (
	NumberTranslator interface {
		Group() string

		Decimal() string

		Minus() string
	}

	NumberSpec struct {
		GroupValue,
		DecimalValue,
		MinusValue string
	}
)

func (n *NumberSpec) Group() string {
	return n.GroupValue
}

func (n *NumberSpec) Decimal() string {
	return n.DecimalValue
}

func (n *NumberSpec) Minus() string {
	return n.MinusValue
}

func FormatNumberT(t Translator, value interface{}, precision int) string {
	return FormatNumber(value, precision, t.Minus(), t.Group(), t.Decimal())
}

// FormatNumber is a base function of the library which formats a number with custom precision and separators.
// FormatNumber supports various types of value by runtime reflection.
// If you don't need runtime type evaluation, please refer to FormatNumberInt or FormatNumberFloat64.
// (supported value types : int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, *big.Rat)
// (also supported value types : decimal.Decimal, *decimal.Decimal *apd.Decimal)
func FormatNumber(value interface{}, precision int, minus, thousand string, decimalStr string) string {
	v := reflect.ValueOf(value)
	var x string
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		x = fmt.Sprintf("%d", v.Int())
		if precision > 0 {
			x += "." + strings.Repeat("0", precision)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		x = fmt.Sprintf("%d", v.Uint())
		if precision > 0 {
			x += "." + strings.Repeat("0", precision)
		}
	case reflect.Float32, reflect.Float64:
		if precision > 0 {
			x = fmt.Sprintf(fmt.Sprintf("%%.%df", precision), v.Float())
		} else {
			x = fmt.Sprint(v.Float())
		}
	case reflect.Struct:
		if fmtr, ok := NumberForamtters[v.Type()]; ok {
			return fmtr(ptrOfType(value), precision, minus, thousand, decimalStr)
		} else {
			panic("Unsupported type - " + v.Type().String() + " - kind " + v.Kind().String())
		}
	case reflect.Ptr:
		if fmtr, ok := NumberForamtters[v.Type().Elem()]; ok {
			return fmtr(value, precision, minus, thousand, decimalStr)
		} else {
			panic("Unsupported type - " + v.Type().String() + " - kind " + v.Kind().String())
		}
	default:
		panic("Unsupported type - " + v.Type().String() + " - kind " + v.Kind().String())
	}

	return FormatNumberString(x, precision, minus, thousand, decimalStr)
}

func FormatNumberString(x string, precision int, minus, thousand string, decimalStr string) string {
	var (
		buffer    []byte
		strBuffer bytes.Buffer
	)

	lastIndex := strings.Index(x, ".") - 1

	if lastIndex < 0 {
		lastIndex = len(x) - 1
	}

	j := 0
	for i := lastIndex; i >= 0; i-- {
		j++

		if x[i] == '-' {
			minus := []byte(minus)
			for i := 0; i < len(minus)/2; i++ {
				j := len(minus) - i - 1
				minus[i], minus[j] = minus[j], minus[i]
			}
			buffer = append(buffer, minus...)
		} else {
			buffer = append(buffer, x[i])
		}

		if j == 3 && i > 0 && !(i == 1 && x[0] == '-') {
			buffer = append(buffer, ',')
			j = 0
		}
	}

	for i := len(buffer) - 1; i >= 0; i-- {
		strBuffer.WriteByte(buffer[i])
	}

	result := strBuffer.String()

	if thousand != "," {
		result = strings.Replace(result, ",", thousand, -1)
	}

	extra := x[lastIndex+1:]
	if decimalStr != "." {
		if precision > 0 {
			if len(extra) > precision {
				extra = extra[:precision+1]
			} else {
				extra += strings.Repeat("0", precision+1-len(extra))
			}
		}
		extra = strings.Replace(extra, ".", decimalStr, 1)
	}

	return result + extra
}

// UnformatNumber takes a string of the number to strip currency info on
// and precision for decimals.
// It pulls the currency descripter from the LocaleInfo map and uses it to return an unformatted value
// based on thous sep and decimal sep
func UnformatNumber(t Translator, n string, precision int) string {
	r := regexp.MustCompile(`[^0-9-., ]`) // Remove anything thats not a space, comma, or decimal
	num := r.ReplaceAllString(n, "${1}")

	r = regexp.MustCompile(fmt.Sprintf("\\%v", t.Group())) // Strip out thousands seperator, whatever it is
	num = r.ReplaceAllString(num, "${1}")

	// Replace decimal seperator with a decimal at specified precision
	if t.Decimal() != "." {
		r = regexp.MustCompile(`,`)
		num = r.ReplaceAllString(num, ".")
	}

	num = setPrecision(num, uint64(precision))
	return num
}

func setPrecision(num string, precision uint64) string {
	p := fmt.Sprintf("%%.%vf", precision)
	num = strings.Trim(num, " ")
	v, _ := strconv.ParseFloat(num, 64)
	return fmt.Sprintf(p, v)
}
