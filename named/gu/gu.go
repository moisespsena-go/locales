package gu

import (
	"math"
	"time"

	"github.com/moisespsena-go/locales"
)

type gu struct {
	*locales.NumberSpec
	*locales.CurrencySpec
	*locales.TimeSpec
	*locales.DurationSpec

	locale           string
	pluralsCardinal  []locales.PluralRule
	pluralsOrdinal   []locales.PluralRule
	pluralsRange     []locales.PluralRule
	percent          string
	perMille         string
	inifinity        string
	dateFullLayout   string
	dateLongLayout   string
	dateMediumLayout string
	dateShortLayout  string
	timeFullLayout   string
	timeLongLayout   string
	timeMediumLayout string
	timeShortLayout  string
	listPatterns     *locales.ListPatterns
	miscPatterns     locales.MiscPatterns
}

// New returns a new instance of translator for the 'gu' locale
func New() locales.Translator {
	return &gu{
		NumberSpec: &locales.NumberSpec{
			MinusValue:   "-",
			GroupValue:   ",",
			DecimalValue: ".",
		},
		CurrencySpec: &locales.CurrencySpec{
			CurrenciesValue: []locales.Currency{{Names: map[locales.PluralRule]string{0: "ADP"}, Symbols: locales.CurrencySymbols{Default: "ADP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AED"}, Symbols: locales.CurrencySymbols{Default: "AED", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AFA"}, Symbols: locales.CurrencySymbols{Default: "AFA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AFN"}, Symbols: locales.CurrencySymbols{Default: "AFN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ALK"}, Symbols: locales.CurrencySymbols{Default: "ALK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ALL"}, Symbols: locales.CurrencySymbols{Default: "ALL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AMD"}, Symbols: locales.CurrencySymbols{Default: "AMD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ANG"}, Symbols: locales.CurrencySymbols{Default: "ANG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AOA"}, Symbols: locales.CurrencySymbols{Default: "AOA", Narrow: "Kz"}}, {Names: map[locales.PluralRule]string{0: "AOK"}, Symbols: locales.CurrencySymbols{Default: "AOK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AON"}, Symbols: locales.CurrencySymbols{Default: "AON", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AOR"}, Symbols: locales.CurrencySymbols{Default: "AOR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARA"}, Symbols: locales.CurrencySymbols{Default: "ARA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARL"}, Symbols: locales.CurrencySymbols{Default: "ARL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARM"}, Symbols: locales.CurrencySymbols{Default: "ARM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARP"}, Symbols: locales.CurrencySymbols{Default: "ARP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARS"}, Symbols: locales.CurrencySymbols{Default: "ARS", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "ATS"}, Symbols: locales.CurrencySymbols{Default: "ATS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AUD"}, Symbols: locales.CurrencySymbols{Default: "A$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "AWG"}, Symbols: locales.CurrencySymbols{Default: "AWG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AZM"}, Symbols: locales.CurrencySymbols{Default: "AZM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AZN"}, Symbols: locales.CurrencySymbols{Default: "AZN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BAD"}, Symbols: locales.CurrencySymbols{Default: "BAD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BAM"}, Symbols: locales.CurrencySymbols{Default: "BAM", Narrow: "KM"}}, {Names: map[locales.PluralRule]string{0: "BAN"}, Symbols: locales.CurrencySymbols{Default: "BAN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BBD"}, Symbols: locales.CurrencySymbols{Default: "BBD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "BDT"}, Symbols: locales.CurrencySymbols{Default: "BDT", Narrow: "৳"}}, {Names: map[locales.PluralRule]string{0: "BEC"}, Symbols: locales.CurrencySymbols{Default: "BEC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BEF"}, Symbols: locales.CurrencySymbols{Default: "BEF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BEL"}, Symbols: locales.CurrencySymbols{Default: "BEL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGL"}, Symbols: locales.CurrencySymbols{Default: "BGL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGM"}, Symbols: locales.CurrencySymbols{Default: "BGM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGN"}, Symbols: locales.CurrencySymbols{Default: "BGN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGO"}, Symbols: locales.CurrencySymbols{Default: "BGO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BHD"}, Symbols: locales.CurrencySymbols{Default: "BHD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BIF"}, Symbols: locales.CurrencySymbols{Default: "BIF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BMD"}, Symbols: locales.CurrencySymbols{Default: "BMD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "BND"}, Symbols: locales.CurrencySymbols{Default: "BND", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "BOB"}, Symbols: locales.CurrencySymbols{Default: "BOB", Narrow: "Bs"}}, {Names: map[locales.PluralRule]string{0: "BOL"}, Symbols: locales.CurrencySymbols{Default: "BOL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BOP"}, Symbols: locales.CurrencySymbols{Default: "BOP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BOV"}, Symbols: locales.CurrencySymbols{Default: "BOV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRB"}, Symbols: locales.CurrencySymbols{Default: "BRB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRC"}, Symbols: locales.CurrencySymbols{Default: "BRC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRE"}, Symbols: locales.CurrencySymbols{Default: "BRE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRL"}, Symbols: locales.CurrencySymbols{Default: "R$", Narrow: "R$"}}, {Names: map[locales.PluralRule]string{0: "BRN"}, Symbols: locales.CurrencySymbols{Default: "BRN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRR"}, Symbols: locales.CurrencySymbols{Default: "BRR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRZ"}, Symbols: locales.CurrencySymbols{Default: "BRZ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BSD"}, Symbols: locales.CurrencySymbols{Default: "BSD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "BTN"}, Symbols: locales.CurrencySymbols{Default: "BTN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BUK"}, Symbols: locales.CurrencySymbols{Default: "BUK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BWP"}, Symbols: locales.CurrencySymbols{Default: "BWP", Narrow: "P"}}, {Names: map[locales.PluralRule]string{0: "BYB"}, Symbols: locales.CurrencySymbols{Default: "BYB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BYN"}, Symbols: locales.CurrencySymbols{Default: "BYN", Narrow: "р."}}, {Names: map[locales.PluralRule]string{0: "BYR"}, Symbols: locales.CurrencySymbols{Default: "BYR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BZD"}, Symbols: locales.CurrencySymbols{Default: "BZD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CAD"}, Symbols: locales.CurrencySymbols{Default: "CA$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CDF"}, Symbols: locales.CurrencySymbols{Default: "CDF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CHE"}, Symbols: locales.CurrencySymbols{Default: "CHE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CHF"}, Symbols: locales.CurrencySymbols{Default: "CHF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CHW"}, Symbols: locales.CurrencySymbols{Default: "CHW", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CLE"}, Symbols: locales.CurrencySymbols{Default: "CLE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CLF"}, Symbols: locales.CurrencySymbols{Default: "CLF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CLP"}, Symbols: locales.CurrencySymbols{Default: "CLP", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CNH"}, Symbols: locales.CurrencySymbols{Default: "CNH", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CNX"}, Symbols: locales.CurrencySymbols{Default: "CNX", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CNY"}, Symbols: locales.CurrencySymbols{Default: "CN¥", Narrow: "¥"}}, {Names: map[locales.PluralRule]string{0: "COP"}, Symbols: locales.CurrencySymbols{Default: "COP", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "COU"}, Symbols: locales.CurrencySymbols{Default: "COU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CRC"}, Symbols: locales.CurrencySymbols{Default: "CRC", Narrow: "₡"}}, {Names: map[locales.PluralRule]string{0: "CSD"}, Symbols: locales.CurrencySymbols{Default: "CSD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CSK"}, Symbols: locales.CurrencySymbols{Default: "CSK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CUC"}, Symbols: locales.CurrencySymbols{Default: "CUC", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CUP"}, Symbols: locales.CurrencySymbols{Default: "CUP", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CVE"}, Symbols: locales.CurrencySymbols{Default: "CVE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CYP"}, Symbols: locales.CurrencySymbols{Default: "CYP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CZK"}, Symbols: locales.CurrencySymbols{Default: "CZK", Narrow: "Kč"}}, {Names: map[locales.PluralRule]string{0: "DDM"}, Symbols: locales.CurrencySymbols{Default: "DDM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DEM"}, Symbols: locales.CurrencySymbols{Default: "DEM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DJF"}, Symbols: locales.CurrencySymbols{Default: "DJF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DKK"}, Symbols: locales.CurrencySymbols{Default: "DKK", Narrow: "kr"}}, {Names: map[locales.PluralRule]string{0: "DOP"}, Symbols: locales.CurrencySymbols{Default: "DOP", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "DZD"}, Symbols: locales.CurrencySymbols{Default: "DZD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ECS"}, Symbols: locales.CurrencySymbols{Default: "ECS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ECV"}, Symbols: locales.CurrencySymbols{Default: "ECV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "EEK"}, Symbols: locales.CurrencySymbols{Default: "EEK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "EGP"}, Symbols: locales.CurrencySymbols{Default: "EGP", Narrow: "E£"}}, {Names: map[locales.PluralRule]string{0: "ERN"}, Symbols: locales.CurrencySymbols{Default: "ERN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ESA"}, Symbols: locales.CurrencySymbols{Default: "ESA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ESB"}, Symbols: locales.CurrencySymbols{Default: "ESB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ESP"}, Symbols: locales.CurrencySymbols{Default: "ESP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ETB"}, Symbols: locales.CurrencySymbols{Default: "ETB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "EUR"}, Symbols: locales.CurrencySymbols{Default: "€", Narrow: "€"}}, {Names: map[locales.PluralRule]string{0: "FIM"}, Symbols: locales.CurrencySymbols{Default: "FIM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "FJD"}, Symbols: locales.CurrencySymbols{Default: "FJD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "FKP"}, Symbols: locales.CurrencySymbols{Default: "FKP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "FRF"}, Symbols: locales.CurrencySymbols{Default: "FRF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GBP"}, Symbols: locales.CurrencySymbols{Default: "£", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "GEK"}, Symbols: locales.CurrencySymbols{Default: "GEK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GEL"}, Symbols: locales.CurrencySymbols{Default: "GEL", Narrow: "₾"}}, {Names: map[locales.PluralRule]string{0: "GHC"}, Symbols: locales.CurrencySymbols{Default: "GHC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GHS"}, Symbols: locales.CurrencySymbols{Default: "GHS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GIP"}, Symbols: locales.CurrencySymbols{Default: "GIP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "GMD"}, Symbols: locales.CurrencySymbols{Default: "GMD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GNF"}, Symbols: locales.CurrencySymbols{Default: "GNF", Narrow: "FG"}}, {Names: map[locales.PluralRule]string{0: "GNS"}, Symbols: locales.CurrencySymbols{Default: "GNS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GQE"}, Symbols: locales.CurrencySymbols{Default: "GQE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GRD"}, Symbols: locales.CurrencySymbols{Default: "GRD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GTQ"}, Symbols: locales.CurrencySymbols{Default: "GTQ", Narrow: "Q"}}, {Names: map[locales.PluralRule]string{0: "GWE"}, Symbols: locales.CurrencySymbols{Default: "GWE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GWP"}, Symbols: locales.CurrencySymbols{Default: "GWP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GYD"}, Symbols: locales.CurrencySymbols{Default: "GYD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "HKD"}, Symbols: locales.CurrencySymbols{Default: "HK$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "HNL"}, Symbols: locales.CurrencySymbols{Default: "HNL", Narrow: "L"}}, {Names: map[locales.PluralRule]string{0: "HRD"}, Symbols: locales.CurrencySymbols{Default: "HRD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "HRK"}, Symbols: locales.CurrencySymbols{Default: "HRK", Narrow: "kn"}}, {Names: map[locales.PluralRule]string{0: "HTG"}, Symbols: locales.CurrencySymbols{Default: "HTG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "HUF"}, Symbols: locales.CurrencySymbols{Default: "HUF", Narrow: "Ft"}}, {Names: map[locales.PluralRule]string{0: "IDR"}, Symbols: locales.CurrencySymbols{Default: "IDR", Narrow: "Rp"}}, {Names: map[locales.PluralRule]string{0: "IEP"}, Symbols: locales.CurrencySymbols{Default: "IEP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ILP"}, Symbols: locales.CurrencySymbols{Default: "ILP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ILR"}, Symbols: locales.CurrencySymbols{Default: "ILR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ILS"}, Symbols: locales.CurrencySymbols{Default: "₪", Narrow: "₪"}}, {Names: map[locales.PluralRule]string{0: "INR"}, Symbols: locales.CurrencySymbols{Default: "₹", Narrow: "₹"}}, {Names: map[locales.PluralRule]string{0: "IQD"}, Symbols: locales.CurrencySymbols{Default: "IQD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "IRR"}, Symbols: locales.CurrencySymbols{Default: "IRR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ISJ"}, Symbols: locales.CurrencySymbols{Default: "ISJ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ISK"}, Symbols: locales.CurrencySymbols{Default: "ISK", Narrow: "kr"}}, {Names: map[locales.PluralRule]string{0: "ITL"}, Symbols: locales.CurrencySymbols{Default: "ITL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "JMD"}, Symbols: locales.CurrencySymbols{Default: "JMD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "JOD"}, Symbols: locales.CurrencySymbols{Default: "JOD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "JPY"}, Symbols: locales.CurrencySymbols{Default: "JP¥", Narrow: "¥"}}, {Names: map[locales.PluralRule]string{0: "KES"}, Symbols: locales.CurrencySymbols{Default: "KES", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KGS"}, Symbols: locales.CurrencySymbols{Default: "KGS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KHR"}, Symbols: locales.CurrencySymbols{Default: "KHR", Narrow: "៛"}}, {Names: map[locales.PluralRule]string{0: "KMF"}, Symbols: locales.CurrencySymbols{Default: "KMF", Narrow: "CF"}}, {Names: map[locales.PluralRule]string{0: "KPW"}, Symbols: locales.CurrencySymbols{Default: "KPW", Narrow: "₩"}}, {Names: map[locales.PluralRule]string{0: "KRH"}, Symbols: locales.CurrencySymbols{Default: "KRH", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KRO"}, Symbols: locales.CurrencySymbols{Default: "KRO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KRW"}, Symbols: locales.CurrencySymbols{Default: "₩", Narrow: "₩"}}, {Names: map[locales.PluralRule]string{0: "KWD"}, Symbols: locales.CurrencySymbols{Default: "KWD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KYD"}, Symbols: locales.CurrencySymbols{Default: "KYD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "KZT"}, Symbols: locales.CurrencySymbols{Default: "KZT", Narrow: "₸"}}, {Names: map[locales.PluralRule]string{0: "LAK"}, Symbols: locales.CurrencySymbols{Default: "LAK", Narrow: "₭"}}, {Names: map[locales.PluralRule]string{0: "LBP"}, Symbols: locales.CurrencySymbols{Default: "LBP", Narrow: "L£"}}, {Names: map[locales.PluralRule]string{0: "LKR"}, Symbols: locales.CurrencySymbols{Default: "LKR", Narrow: "Rs"}}, {Names: map[locales.PluralRule]string{0: "LRD"}, Symbols: locales.CurrencySymbols{Default: "LRD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "LSL"}, Symbols: locales.CurrencySymbols{Default: "LSL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LTL"}, Symbols: locales.CurrencySymbols{Default: "LTL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LTT"}, Symbols: locales.CurrencySymbols{Default: "LTT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LUC"}, Symbols: locales.CurrencySymbols{Default: "LUC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LUF"}, Symbols: locales.CurrencySymbols{Default: "LUF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LUL"}, Symbols: locales.CurrencySymbols{Default: "LUL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LVL"}, Symbols: locales.CurrencySymbols{Default: "LVL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LVR"}, Symbols: locales.CurrencySymbols{Default: "LVR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LYD"}, Symbols: locales.CurrencySymbols{Default: "LYD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MAD"}, Symbols: locales.CurrencySymbols{Default: "MAD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MAF"}, Symbols: locales.CurrencySymbols{Default: "MAF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MCF"}, Symbols: locales.CurrencySymbols{Default: "MCF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MDC"}, Symbols: locales.CurrencySymbols{Default: "MDC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MDL"}, Symbols: locales.CurrencySymbols{Default: "MDL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MGA"}, Symbols: locales.CurrencySymbols{Default: "MGA", Narrow: "Ar"}}, {Names: map[locales.PluralRule]string{0: "MGF"}, Symbols: locales.CurrencySymbols{Default: "MGF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MKD"}, Symbols: locales.CurrencySymbols{Default: "MKD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MKN"}, Symbols: locales.CurrencySymbols{Default: "MKN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MLF"}, Symbols: locales.CurrencySymbols{Default: "MLF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MMK"}, Symbols: locales.CurrencySymbols{Default: "MMK", Narrow: "K"}}, {Names: map[locales.PluralRule]string{0: "MNT"}, Symbols: locales.CurrencySymbols{Default: "MNT", Narrow: "₮"}}, {Names: map[locales.PluralRule]string{0: "MOP"}, Symbols: locales.CurrencySymbols{Default: "MOP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MRO"}, Symbols: locales.CurrencySymbols{Default: "MRO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MRU"}, Symbols: locales.CurrencySymbols{Default: "MRU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MTL"}, Symbols: locales.CurrencySymbols{Default: "MTL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MTP"}, Symbols: locales.CurrencySymbols{Default: "MTP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MUR"}, Symbols: locales.CurrencySymbols{Default: "MUR", Narrow: "રૂ."}}, {Names: map[locales.PluralRule]string{0: "MVP"}, Symbols: locales.CurrencySymbols{Default: "MVP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MVR"}, Symbols: locales.CurrencySymbols{Default: "MVR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MWK"}, Symbols: locales.CurrencySymbols{Default: "MWK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MXN"}, Symbols: locales.CurrencySymbols{Default: "MX$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "MXP"}, Symbols: locales.CurrencySymbols{Default: "MXP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MXV"}, Symbols: locales.CurrencySymbols{Default: "MXV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MYR"}, Symbols: locales.CurrencySymbols{Default: "MYR", Narrow: "RM"}}, {Names: map[locales.PluralRule]string{0: "MZE"}, Symbols: locales.CurrencySymbols{Default: "MZE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MZM"}, Symbols: locales.CurrencySymbols{Default: "MZM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MZN"}, Symbols: locales.CurrencySymbols{Default: "MZN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NAD"}, Symbols: locales.CurrencySymbols{Default: "NAD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "NGN"}, Symbols: locales.CurrencySymbols{Default: "NGN", Narrow: "₦"}}, {Names: map[locales.PluralRule]string{0: "NIC"}, Symbols: locales.CurrencySymbols{Default: "NIC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NIO"}, Symbols: locales.CurrencySymbols{Default: "NIO", Narrow: "C$"}}, {Names: map[locales.PluralRule]string{0: "NLG"}, Symbols: locales.CurrencySymbols{Default: "NLG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NOK"}, Symbols: locales.CurrencySymbols{Default: "NOK", Narrow: "kr"}}, {Names: map[locales.PluralRule]string{0: "NPR"}, Symbols: locales.CurrencySymbols{Default: "NPR", Narrow: "Rs"}}, {Names: map[locales.PluralRule]string{0: "NZD"}, Symbols: locales.CurrencySymbols{Default: "NZ$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "OMR"}, Symbols: locales.CurrencySymbols{Default: "OMR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PAB"}, Symbols: locales.CurrencySymbols{Default: "PAB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PEI"}, Symbols: locales.CurrencySymbols{Default: "PEI", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PEN"}, Symbols: locales.CurrencySymbols{Default: "PEN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PES"}, Symbols: locales.CurrencySymbols{Default: "PES", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PGK"}, Symbols: locales.CurrencySymbols{Default: "PGK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PHP"}, Symbols: locales.CurrencySymbols{Default: "PHP", Narrow: "₱"}}, {Names: map[locales.PluralRule]string{0: "PKR"}, Symbols: locales.CurrencySymbols{Default: "PKR", Narrow: "Rs"}}, {Names: map[locales.PluralRule]string{0: "PLN"}, Symbols: locales.CurrencySymbols{Default: "PLN", Narrow: "zł"}}, {Names: map[locales.PluralRule]string{0: "PLZ"}, Symbols: locales.CurrencySymbols{Default: "PLZ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PTE"}, Symbols: locales.CurrencySymbols{Default: "PTE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PYG"}, Symbols: locales.CurrencySymbols{Default: "PYG", Narrow: "₲"}}, {Names: map[locales.PluralRule]string{0: "QAR"}, Symbols: locales.CurrencySymbols{Default: "QAR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RHD"}, Symbols: locales.CurrencySymbols{Default: "RHD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ROL"}, Symbols: locales.CurrencySymbols{Default: "ROL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RON"}, Symbols: locales.CurrencySymbols{Default: "RON", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RSD"}, Symbols: locales.CurrencySymbols{Default: "RSD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RUB"}, Symbols: locales.CurrencySymbols{Default: "RUB", Narrow: "₽"}}, {Names: map[locales.PluralRule]string{0: "RUR"}, Symbols: locales.CurrencySymbols{Default: "RUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RWF"}, Symbols: locales.CurrencySymbols{Default: "RWF", Narrow: "RF"}}, {Names: map[locales.PluralRule]string{0: "SAR"}, Symbols: locales.CurrencySymbols{Default: "SAR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SBD"}, Symbols: locales.CurrencySymbols{Default: "SBD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "SCR"}, Symbols: locales.CurrencySymbols{Default: "SCR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SDD"}, Symbols: locales.CurrencySymbols{Default: "SDD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SDG"}, Symbols: locales.CurrencySymbols{Default: "SDG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SDP"}, Symbols: locales.CurrencySymbols{Default: "SDP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SEK"}, Symbols: locales.CurrencySymbols{Default: "SEK", Narrow: "kr"}}, {Names: map[locales.PluralRule]string{0: "SGD"}, Symbols: locales.CurrencySymbols{Default: "SGD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "SHP"}, Symbols: locales.CurrencySymbols{Default: "SHP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "SIT"}, Symbols: locales.CurrencySymbols{Default: "SIT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SKK"}, Symbols: locales.CurrencySymbols{Default: "SKK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SLE"}, Symbols: locales.CurrencySymbols{Default: "SLE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SLL"}, Symbols: locales.CurrencySymbols{Default: "SLL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SOS"}, Symbols: locales.CurrencySymbols{Default: "SOS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SRD"}, Symbols: locales.CurrencySymbols{Default: "SRD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "SRG"}, Symbols: locales.CurrencySymbols{Default: "SRG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SSP"}, Symbols: locales.CurrencySymbols{Default: "SSP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "STD"}, Symbols: locales.CurrencySymbols{Default: "STD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "STN"}, Symbols: locales.CurrencySymbols{Default: "STN", Narrow: "Db"}}, {Names: map[locales.PluralRule]string{0: "SUR"}, Symbols: locales.CurrencySymbols{Default: "SUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SVC"}, Symbols: locales.CurrencySymbols{Default: "SVC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SYP"}, Symbols: locales.CurrencySymbols{Default: "SYP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "SZL"}, Symbols: locales.CurrencySymbols{Default: "SZL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "THB"}, Symbols: locales.CurrencySymbols{Default: "฿", Narrow: "฿"}}, {Names: map[locales.PluralRule]string{0: "TJR"}, Symbols: locales.CurrencySymbols{Default: "TJR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TJS"}, Symbols: locales.CurrencySymbols{Default: "TJS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TMM"}, Symbols: locales.CurrencySymbols{Default: "TMM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TMT"}, Symbols: locales.CurrencySymbols{Default: "TMT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TND"}, Symbols: locales.CurrencySymbols{Default: "TND", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TOP"}, Symbols: locales.CurrencySymbols{Default: "TOP", Narrow: "T$"}}, {Names: map[locales.PluralRule]string{0: "TPE"}, Symbols: locales.CurrencySymbols{Default: "TPE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TRL"}, Symbols: locales.CurrencySymbols{Default: "TRL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TRY"}, Symbols: locales.CurrencySymbols{Default: "TRY", Narrow: "₺"}}, {Names: map[locales.PluralRule]string{0: "TTD"}, Symbols: locales.CurrencySymbols{Default: "TTD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "TWD"}, Symbols: locales.CurrencySymbols{Default: "NT$", Narrow: "NT$"}}, {Names: map[locales.PluralRule]string{0: "TZS"}, Symbols: locales.CurrencySymbols{Default: "TZS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UAH"}, Symbols: locales.CurrencySymbols{Default: "UAH", Narrow: "₴"}}, {Names: map[locales.PluralRule]string{0: "UAK"}, Symbols: locales.CurrencySymbols{Default: "UAK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UGS"}, Symbols: locales.CurrencySymbols{Default: "UGS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UGX"}, Symbols: locales.CurrencySymbols{Default: "UGX", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "USD"}, Symbols: locales.CurrencySymbols{Default: "US$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "USN"}, Symbols: locales.CurrencySymbols{Default: "USN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "USS"}, Symbols: locales.CurrencySymbols{Default: "USS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UYI"}, Symbols: locales.CurrencySymbols{Default: "UYI", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UYP"}, Symbols: locales.CurrencySymbols{Default: "UYP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UYU"}, Symbols: locales.CurrencySymbols{Default: "UYU", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "UYW"}, Symbols: locales.CurrencySymbols{Default: "UYW", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UZS"}, Symbols: locales.CurrencySymbols{Default: "UZS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VEB"}, Symbols: locales.CurrencySymbols{Default: "VEB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VED"}, Symbols: locales.CurrencySymbols{Default: "VED", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VEF"}, Symbols: locales.CurrencySymbols{Default: "VEF", Narrow: "Bs"}}, {Names: map[locales.PluralRule]string{0: "VES"}, Symbols: locales.CurrencySymbols{Default: "VES", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VND"}, Symbols: locales.CurrencySymbols{Default: "₫", Narrow: "₫"}}, {Names: map[locales.PluralRule]string{0: "VNN"}, Symbols: locales.CurrencySymbols{Default: "VNN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VUV"}, Symbols: locales.CurrencySymbols{Default: "VUV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "WST"}, Symbols: locales.CurrencySymbols{Default: "WST", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XAF"}, Symbols: locales.CurrencySymbols{Default: "FCFA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XAG"}, Symbols: locales.CurrencySymbols{Default: "XAG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XAU"}, Symbols: locales.CurrencySymbols{Default: "XAU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBA"}, Symbols: locales.CurrencySymbols{Default: "XBA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBB"}, Symbols: locales.CurrencySymbols{Default: "XBB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBC"}, Symbols: locales.CurrencySymbols{Default: "XBC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBD"}, Symbols: locales.CurrencySymbols{Default: "XBD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XCD"}, Symbols: locales.CurrencySymbols{Default: "EC$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "XDR"}, Symbols: locales.CurrencySymbols{Default: "XDR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XEU"}, Symbols: locales.CurrencySymbols{Default: "XEU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XFO"}, Symbols: locales.CurrencySymbols{Default: "XFO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XFU"}, Symbols: locales.CurrencySymbols{Default: "XFU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XOF"}, Symbols: locales.CurrencySymbols{Default: "F\u202fCFA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XPD"}, Symbols: locales.CurrencySymbols{Default: "XPD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XPF"}, Symbols: locales.CurrencySymbols{Default: "CFPF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XPT"}, Symbols: locales.CurrencySymbols{Default: "XPT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XRE"}, Symbols: locales.CurrencySymbols{Default: "XRE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XSU"}, Symbols: locales.CurrencySymbols{Default: "XSU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XTS"}, Symbols: locales.CurrencySymbols{Default: "XTS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XUA"}, Symbols: locales.CurrencySymbols{Default: "XUA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XXX"}, Symbols: locales.CurrencySymbols{Default: "XXX", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YDD"}, Symbols: locales.CurrencySymbols{Default: "YDD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YER"}, Symbols: locales.CurrencySymbols{Default: "YER", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUD"}, Symbols: locales.CurrencySymbols{Default: "YUD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUM"}, Symbols: locales.CurrencySymbols{Default: "YUM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUN"}, Symbols: locales.CurrencySymbols{Default: "YUN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUR"}, Symbols: locales.CurrencySymbols{Default: "YUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZAL"}, Symbols: locales.CurrencySymbols{Default: "ZAL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZAR"}, Symbols: locales.CurrencySymbols{Default: "ZAR", Narrow: "R"}}, {Names: map[locales.PluralRule]string{0: "ZMK"}, Symbols: locales.CurrencySymbols{Default: "ZMK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZMW"}, Symbols: locales.CurrencySymbols{Default: "ZMW", Narrow: "ZK"}}, {Names: map[locales.PluralRule]string{0: "ZRN"}, Symbols: locales.CurrencySymbols{Default: "ZRN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZRZ"}, Symbols: locales.CurrencySymbols{Default: "ZRZ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZWD"}, Symbols: locales.CurrencySymbols{Default: "ZWD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZWL"}, Symbols: locales.CurrencySymbols{Default: "ZWL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZWR"}, Symbols: locales.CurrencySymbols{Default: "ZWR", Narrow: ""}}},
			Formatters: &locales.CurrencyFormatters{
				CurrencyFmt: locales.MustParseNumberFormatPatterns("¤#,##0.00"),
			},
		},
		TimeSpec: &locales.TimeSpec{
			Separator:                ":",
			MonthsAbbreviatedValue:   []string{"", "જાન્યુ", "ફેબ્રુ", "માર્ચ", "એપ્રિલ", "મે", "જૂન", "જુલાઈ", "ઑગસ્ટ", "સપ્ટે", "ઑક્ટો", "નવે", "ડિસે"},
			MonthsNarrowValue:        []string{"", "જા", "ફે", "મા", "એ", "મે", "જૂ", "જુ", "ઑ", "સ", "ઑ", "ન", "ડિ"},
			MonthsWideValue:          []string{"", "જાન્યુઆરી", "ફેબ્રુઆરી", "માર્ચ", "એપ્રિલ", "મે", "જૂન", "જુલાઈ", "ઑગસ્ટ", "સપ્ટેમ્બર", "ઑક્ટોબર", "નવેમ્બર", "ડિસેમ્બર"},
			WeekdaysAbbreviatedValue: []string{"રવિ", "સોમ", "મંગળ", "બુધ", "ગુરુ", "શુક્ર", "શનિ"},
			WeekdaysNarrowValue:      []string{"ર", "સો", "મં", "બુ", "ગુ", "શુ", "શ"},
			WeekdaysShortValue:       []string{"ર", "સો", "મં", "બુ", "ગુ", "શુ", "શ"},
			WeekdaysWideValue:        []string{"રવિવાર", "સોમવાર", "મંગળવાર", "બુધવાર", "ગુરુવાર", "શુક્રવાર", "શનિવાર"},
			PeriodsAbbreviatedValue:  []string{"AM", "PM"},
			PeriodsNarrowValue:       []string{"AM", "PM"},
			PeriodsWideValue:         []string{"AM", "PM"},
			ErasAbbreviatedValue:     []string{"", ""},
			ErasNarrowValue:          []string{"ઇ સ પુ", "ઇસ"},
			ErasWideValue:            []string{"ઈસવીસન પૂર્વે", "ઇસવીસન"},
			TimezonesValue:           map[string]string{"": "પેરાગ્વે માનક સમય", "ACDT": "ઓસ્ટ્રેલિયન કેન્દ્રીય દિવસ સમય", "ACST": "ઑસ્ટ્રેલિયન કેન્દ્રીય માનક સમય", "ACT": "એકર પ્રમાણભૂત સમય", "ACWDT": "ઓસ્ટ્રેલિયન કેન્દ્રીય પશ્ચિમી દિવસ સમય", "ACWST": "ઑસ્ટ્રેલિયન કેન્દ્રીય પશ્ચિમી માનક સમય", "ADT": "એટલાન્ટિક દિવસ સમય", "AEDT": "ઓસ્ટ્રેલિયન પૂર્વીય દિવસ સમય", "AEST": "ઑસ્ટ્રેલિયન પૂર્વીય માનક સમય", "AFT": "અફઘાનિસ્તાન સમય", "AKDT": "અલાસ્કા દિવસ સમય", "AKST": "અલાસ્કા માનક સમય", "AMST": "એમેઝોન ગ્રીષ્મ સમય", "AMT": "એમેઝોન માનક સમય", "ARST": "આર્જેન્ટીના ગ્રીષ્મ સમય", "ART": "આર્જેન્ટીના માનક સમય", "AST": "એટલાન્ટિક માનક સમય", "AWDT": "ઓસ્ટ્રેલિયન પશ્ચિમી દિવસ સમય", "AWST": "ઑસ્ટ્રેલિયન પશ્ચિમી માનક સમય", "BNT": "બ્રુનેઇ દરુસલામ સમય", "BOT": "બોલિવિયા સમય", "BRST": "બ્રાઝિલિયા ગ્રીષ્મ સમય", "BRT": "બ્રાઝિલિયા માનક સમય", "BST": "બાંગ્લાદેશ ગ્રીષ્મ સમય", "BTT": "ભૂટાન સમય", "CAT": "મધ્ય આફ્રિકા સમય", "CCT": "કોકોઝ આઇલેન્ડ્સ સમય", "CDT": "ઉત્તર અમેરિકન કેન્દ્રીય દિવસ સમય", "CDT (Kuba)": "ક્યુબા દિવસ સમય", "CHADT": "ચેતહામ દિવસ સમય", "CHAST": "ચેતહામ માનક સમય", "CLST": "ચિલી ગ્રીષ્મ સમય", "CLT": "ચિલી માનક સમય", "COST": "કોલમ્બિયા ગ્રીષ્મ સમય", "COT": "કોલમ્બિયા માનક સમય", "CST": "ઉત્તર અમેરિકન કેન્દ્રીય માનક સમય", "CST (Kuba)": "ક્યુબા માનક સમય", "CXT": "ક્રિસમસ આઇલેન્ડ સમય", "ChST": "કેમોરો માનક સમય", "EASST": "ઇસ્ટર આઇલેન્ડ ગ્રીષ્મ સમય", "EAST": "ઇસ્ટર આઇલેન્ડ માનક સમય", "EAT": "પૂર્વ આફ્રિકા સમય", "ECT": "એક્વાડોર સમય", "EDT": "ઉત્તર અમેરિકન પૂર્વી દિવસ સમય", "EGDT": "પૂર્વ ગ્રીનલેન્ડ ગ્રીષ્મ સમય", "EGST": "પૂર્વ ગ્રીનલેન્ડ માનક સમય", "EST": "ઉત્તર અમેરિકન પૂર્વી માનક સમય", "FKST": "ફૉકલેન્ડ આઇલેન્ડ્સ ગ્રીષ્મ સમય", "GALT": "ગાલાપાગોસ સમય", "GFT": "ફ્રેંચ ગયાના સમય", "GMT": "ગ્રીનવિચ મધ્યમ સમય", "GST": "ગલ્ફ માનક સમય", "GYT": "ગયાના સમય", "HADT": "હવાઇ-એલ્યુશિઅન દિવસ સમય", "HAST": "હવાઇ-એલ્યુશિઅન માનક સમય", "HENOMX": "ઉત્તરપશ્ચિમ મેક્સિકો દિવસ સમય", "HEOG": "પશ્ચિમ ગ્રીનલેન્ડ ગ્રીષ્મ સમય", "HEPMX": "મેક્સિકન પેસિફિક દિવસ સમય", "HKST": "હોંગ કોંગ ગ્રીષ્મ સમય", "HKT": "હોંગ કોંગ માનક સમય", "HNNOMX": "ઉત્તરપશ્ચિમ મેક્સિકો માનક સમય", "HNOG": "પશ્ચિમ ગ્રીનલેન્ડ માનક સમય", "HNPMX": "મેક્સિકન પેસિફિક માનક સમય", "ICT": "ઇન્ડોચાઇના સમય", "IRDT": "ઈરાન દિવસ સમય", "IRST": "ઈરાન માનક સમય", "IST": "ભારતીય માનક સમય", "JDT": "જાપાન દિવસ સમય", "JST": "જાપાન માનક સમય", "LHDT": "લોર્ડ હોવ દિવસ સમય", "LHST": "લોર્ડ હોવ માનક સમય", "MDT": "ઉત્તર અમેરિકન માઉન્ટન દિવસ સમય", "MESZ": "મધ્ય યુરોપિયન ગ્રીષ્મ સમય", "MEZ": "મધ્ય યુરોપિયન માનક સમય", "MST": "ઉત્તર અમેરિકન માઉન્ટન માનક સમય", "MVT": "માલદીવ સમય", "MYT": "મલેશિયા સમય", "NDT": "ન્યૂફાઉન્ડલેન્ડ દિવસ સમય", "NPT": "નેપાળ સમય", "NST": "ન્યૂફાઉન્ડલેન્ડ માનક સમય", "NZDT": "ન્યુઝીલેન્ડ દિવસ સમય", "NZST": "ન્યુઝીલેન્ડ માનક સમય", "OESZ": "પૂર્વી યુરોપિયન ગ્રીષ્મ સમય", "OEZ": "પૂર્વી યુરોપિયન માનક સમય", "PDT": "ઉત્તર અમેરિકન પેસિફિક દિવસ સમય", "PKT": "પાકિસ્તાન ગ્રીષ્મ સમય", "PMDT": "સેંટ પીએરી અને મિક્યુલોન દિવસ સમય", "PMST": "સેંટ પીએરી અને મિક્યુલોન માનક સમય", "PST": "ઉત્તર અમેરિકન પેસિફિક માનક સમય", "PYST": "પેરાગ્વે ગ્રીષ્મ સમય", "SAST": "દક્ષિણ આફ્રિકા માનક સમય", "SGT": "સિંગાપુર માનક સમય", "SRT": "સુરીનામ સમય", "TLT": "પૂર્વ તિમોર સમય", "TMST": "તુર્કમેનિસ્તાન ગ્રીષ્મ સમય", "TMT": "તુર્કમેનિસ્તાન માનક સમય", "UYST": "ઉરુગ્વે ગ્રીષ્મ સમય", "UYT": "ઉરુગ્વે માનક સમય", "VET": "વેનેઝુએલા સમય", "WARST": "પશ્ચિમી આર્જેન્ટીના ગ્રીષ્મ સમય", "WART": "પશ્ચિમી આર્જેન્ટીના માનક સમય", "WAST": "પશ્ચિમ આફ્રિકા ગ્રીષ્મ સમય", "WAT": "પશ્ચિમ આફ્રિકા માનક સમય", "WESZ": "પશ્ચિમી યુરોપિયન ગ્રીષ્મ સમય", "WEZ": "પશ્ચિમી યુરોપિયન માનક સમય", "WIB": "પશ્ચિમી ઇન્ડોનેશિયા સમય", "WIT": "પૂર્વીય ઇન્ડોનેશિયા સમય", "WITA": "મધ્ય ઇન્ડોનેશિયા સમય", "∅∅∅": "પેરુ ગ્રીષ્મ સમય"},
		},

		locale:          "gu",
		pluralsCardinal: []locales.PluralRule{2, 6},
		pluralsOrdinal:  []locales.PluralRule{2, 3, 4, 5, 6},
		pluralsRange:    []locales.PluralRule{2, 6},
		percent:         "%",
		perMille:        "‰",
		inifinity:       "∞",

		dateFullLayout:   "EEEE, d MMMM, y",
		dateLongLayout:   "d MMMM, y",
		dateMediumLayout: "d MMM, y",
		dateShortLayout:  "d/M/yy",
		timeFullLayout:   "hh:mm:ss a zzzz",
		timeLongLayout:   "hh:mm:ss a z",
		timeMediumLayout: "hh:mm:ss a",
		timeShortLayout:  "hh:mm a",
		listPatterns:     locales.NewListPatternsFromSlice([][]string{{"{0} અને {1}", "{0}, {1}", "{0}, {1}", "{0} અને {1}"}, {"{0} અથવા {1}", "{0}, {1}", "{0}, {1}", "{0}, અથવા {1}"}}),
		DurationSpec: &locales.DurationSpec{" ",
			"",
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"સદી", "%v સદી", "%v સદી", ""},
				Long:   &locales.CounterFormat{"સદીઓ", "%v સદી", "%v સદીઓ", ""}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"દાયકા", "%v દા", "%v દા", ""},
				Long:   &locales.CounterFormat{"દાયકા", "%v દાયકો", "%v દાયકા", ""}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"વર્ષ", "%v વ", "%v વ", ""},
				Long:   &locales.CounterFormat{"વર્ષ", "%v વર્ષ", "%v વર્ષ", "%v પ્રતિ વર્ષ"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"મહિના", "%v મ", "%v મ", ""},
				Long:   &locales.CounterFormat{"મહિના", "%v મહિનો", "%v મહિના", "%v પ્રતિ મહિનો"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"અઠવાડિયા", "%v અઠ.", "%v અઠ.", ""},
				Long:   &locales.CounterFormat{"અઠવાડિયા", "%v અઠવાડિયું", "%v અઠવાડિયા", "%v પ્રતિ અઠવાડિયું"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"દિવસ", "%v દિ", "%v દિ", "%v/ દિ"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"કલાક", "%v ક", "%v ક", "%v/ક"},
				Long:   &locales.CounterFormat{"કલાક", "%v કલાક", "%v કલાક", "%v પ્રતિ કલાક"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"મિનિટ", "%v મિ", "%v મિ", "%v/મિ"},
				Long:   &locales.CounterFormat{"મિનિટ", "%v મિનિટ", "%v મિનિટ", "%v પ્રતિ મિનિટ"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"સેકંડ", "%v સે", "%v સે", "%v/સે"},
				Long:   &locales.CounterFormat{"સેકંડ", "%v સેકંડ", "%v સેકંડ", "%v પ્રતિ સેકંડ"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"મિલિસેકન્ડ", "%v મિસે", "%v મિસે", ""},
				Long:   &locales.CounterFormat{"મિલિસેકન્ડ", "%v મિલિસેકન્ડ", "%v મિલિસેકન્ડ", ""}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"μs", "%v μs", "%v μs", ""},
				Long:   &locales.CounterFormat{"માઇક્રોસેકંડ", "%v માઇક્રોસેકંડ", "%v માઇક્રોસેકંડ", ""}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"નેનોસેકંડ", "%v ns", "%v ns", ""},
				Long:   &locales.CounterFormat{"નેનોસેકંડ", "%v નેનોસેકંડ", "%v નેનોસેકંડ", ""}},
		},
		miscPatterns: locales.MiscPatterns{
			"~%[1]v",
			"%[1]v+",
			"≤%[1]v",
			"%[1]v-%[2]v",
		},
	}
}

// Locale returns the current translators string locale
func (gu *gu) Locale() string {
	return gu.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'gu'
func (gu *gu) PluralsCardinal() []locales.PluralRule {
	return gu.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'gu'
func (gu *gu) PluralsOrdinal() []locales.PluralRule {
	return gu.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'gu'
func (gu *gu) PluralsRange() []locales.PluralRule {
	return gu.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'gu'
func (gu *gu) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'gu'
func (gu *gu) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 || n == 3 {
		return locales.PluralRuleTwo
	} else if n == 4 {
		return locales.PluralRuleFew
	} else if n == 6 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'gu'
func (gu *gu) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := gu.CardinalPluralRule(num1, v1)
	end := gu.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther

}

// FmtDateShort returns the short date representation of 't' for 'gu'
func (gu *gu) FmtDateShort(t time.Time) string {
	return locales.FormatTimeEra(gu, gu.dateShortLayout, t, 2)
}

// FmtDateMedium returns the medium date representation of 't' for 'gu'
func (gu *gu) FmtDateMedium(t time.Time) string {
	return locales.FormatTimeEra(gu, gu.dateMediumLayout, t, 2)
}

// FmtDateLong returns the long date representation of 't' for 'gu'
func (gu *gu) FmtDateLong(t time.Time) string {
	return locales.FormatTimeEra(gu, gu.dateLongLayout, t, 1)
}

// FmtDateFull returns the full date representation of 't' for 'gu'
func (gu *gu) FmtDateFull(t time.Time) string {
	return locales.FormatTimeEra(gu, gu.dateFullLayout, t, 0)
}

// FmtTimeShort returns the short time representation of 't' for 'gu'
func (gu *gu) FmtTimeShort(t time.Time) string {
	return locales.FormatTimeEra(gu, gu.timeShortLayout, t, 2)
}

// FmtTimeMedium returns the medium time representation of 't' for 'gu'
func (gu *gu) FmtTimeMedium(t time.Time) string {
	return locales.FormatTimeEra(gu, gu.timeMediumLayout, t, 2)
}

// FmtTimeLong returns the long time representation of 't' for 'gu'
func (gu *gu) FmtTimeLong(t time.Time) string {
	return locales.FormatTimeEra(gu, gu.timeLongLayout, t, 1)
}

// FmtTimeFull returns the full time representation of 't' for 'gu'
func (gu *gu) FmtTimeFull(t time.Time) string {
	return locales.FormatTimeEra(gu, gu.timeFullLayout, t, 0)
}

// DateFullLayout returns the full date layout for 'gu'
func (gu *gu) DateFullLayout() string {
	return gu.dateFullLayout
}

// DateLongLayout returns the long date layout for 'gu'
func (gu *gu) DateLongLayout() string {
	return gu.dateLongLayout
}

// DateMediumLayout returns the medium date layout for 'gu'
func (gu *gu) DateMediumLayout() string {
	return gu.dateMediumLayout
}

// DateShortLayout returns the short date layout for 'gu'
func (gu *gu) DateShortLayout() string {
	return gu.dateShortLayout
}

// TimeFullLayout returns the full time layout for 'gu'
func (gu *gu) TimeFullLayout() string {
	return gu.timeFullLayout
}

// TimeLongLayout returns the full long layout for 'gu'
func (gu *gu) TimeLongLayout() string {
	return gu.timeLongLayout
}

// TimeMediumLayout returns the medium time layout for 'gu'
func (gu *gu) TimeMediumLayout() string {
	return gu.timeMediumLayout
}

// TimeShortLayout returns the short time layout for 'gu'
func (gu *gu) TimeShortLayout() string {
	return gu.timeShortLayout
}

func (gu *gu) ListPatterns() *locales.ListPatterns {
	return gu.listPatterns
}

func (gu *gu) GetDurationSpec() *locales.DurationSpec {
	return gu.DurationSpec
}

func (gu *gu) GetMiscPatterns() *locales.MiscPatterns {
	return &gu.miscPatterns
}
