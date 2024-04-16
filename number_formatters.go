package locales

import (
	"math/big"
	"reflect"

	"github.com/shopspring/decimal"
)

type NumberFormatter func(value interface{}, precision int, minus, thousand string, decimalStr string) string

var NumberForamtters = map[reflect.Type]NumberFormatter{
	reflect.TypeOf(big.Rat{}): func(value interface{}, precision int, minus, thousand string, decimalStr string) string {
		var rat, _ = value.(*big.Rat)
		if rat == nil {
			var v2 any = value
			rat = v2.(*big.Rat)
		}

		if precision > 0 {
			return value.(*big.Rat).FloatString(precision)
		} else {
			return value.(*big.Rat).String()
		}
	},

	reflect.TypeOf(decimal.Decimal{}): func(value interface{}, precision int, minus, thousand string, decimalStr string) string {
		if precision > 0 {
			return value.(*decimal.Decimal).StringFixed(int32(precision))
		} else {
			return value.(*decimal.Decimal).String()
		}
	},
}

func ptrOfType[T any](v T) *T {
	return &v
}
