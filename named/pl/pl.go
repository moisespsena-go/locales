package pl

import (
	"math"
	"time"

	"github.com/moisespsena-go/locales"
)

type pl struct {
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

// New returns a new instance of translator for the 'pl' locale
func New() locales.Translator {
	return &pl{
		NumberSpec: &locales.NumberSpec{
			MinusValue:   "-",
			GroupValue:   " ",
			DecimalValue: ",",
		},
		CurrencySpec: &locales.CurrencySpec{
			CurrenciesValue: []locales.Currency{{Names: map[locales.PluralRule]string{0: "ADP"}, Symbols: locales.CurrencySymbols{Default: "ADP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AED"}, Symbols: locales.CurrencySymbols{Default: "AED", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AFA"}, Symbols: locales.CurrencySymbols{Default: "AFA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AFN"}, Symbols: locales.CurrencySymbols{Default: "AFN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ALK"}, Symbols: locales.CurrencySymbols{Default: "ALK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ALL"}, Symbols: locales.CurrencySymbols{Default: "ALL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AMD"}, Symbols: locales.CurrencySymbols{Default: "AMD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ANG"}, Symbols: locales.CurrencySymbols{Default: "ANG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AOA"}, Symbols: locales.CurrencySymbols{Default: "AOA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AOK"}, Symbols: locales.CurrencySymbols{Default: "AOK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AON"}, Symbols: locales.CurrencySymbols{Default: "AON", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AOR"}, Symbols: locales.CurrencySymbols{Default: "AOR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARA"}, Symbols: locales.CurrencySymbols{Default: "ARA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARL"}, Symbols: locales.CurrencySymbols{Default: "ARL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARM"}, Symbols: locales.CurrencySymbols{Default: "ARM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARP"}, Symbols: locales.CurrencySymbols{Default: "ARP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARS"}, Symbols: locales.CurrencySymbols{Default: "ARS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ATS"}, Symbols: locales.CurrencySymbols{Default: "ATS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AUD"}, Symbols: locales.CurrencySymbols{Default: "AUD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AWG"}, Symbols: locales.CurrencySymbols{Default: "AWG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AZM"}, Symbols: locales.CurrencySymbols{Default: "AZM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AZN"}, Symbols: locales.CurrencySymbols{Default: "AZN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BAD"}, Symbols: locales.CurrencySymbols{Default: "BAD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BAM"}, Symbols: locales.CurrencySymbols{Default: "BAM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BAN"}, Symbols: locales.CurrencySymbols{Default: "BAN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BBD"}, Symbols: locales.CurrencySymbols{Default: "BBD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BDT"}, Symbols: locales.CurrencySymbols{Default: "BDT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BEC"}, Symbols: locales.CurrencySymbols{Default: "BEC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BEF"}, Symbols: locales.CurrencySymbols{Default: "BEF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BEL"}, Symbols: locales.CurrencySymbols{Default: "BEL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGL"}, Symbols: locales.CurrencySymbols{Default: "BGL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGM"}, Symbols: locales.CurrencySymbols{Default: "BGM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGN"}, Symbols: locales.CurrencySymbols{Default: "BGN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGO"}, Symbols: locales.CurrencySymbols{Default: "BGO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BHD"}, Symbols: locales.CurrencySymbols{Default: "BHD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BIF"}, Symbols: locales.CurrencySymbols{Default: "BIF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BMD"}, Symbols: locales.CurrencySymbols{Default: "BMD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BND"}, Symbols: locales.CurrencySymbols{Default: "BND", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BOB"}, Symbols: locales.CurrencySymbols{Default: "BOB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BOL"}, Symbols: locales.CurrencySymbols{Default: "BOL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BOP"}, Symbols: locales.CurrencySymbols{Default: "BOP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BOV"}, Symbols: locales.CurrencySymbols{Default: "BOV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRB"}, Symbols: locales.CurrencySymbols{Default: "BRB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRC"}, Symbols: locales.CurrencySymbols{Default: "BRC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRE"}, Symbols: locales.CurrencySymbols{Default: "BRE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRL"}, Symbols: locales.CurrencySymbols{Default: "BRL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRN"}, Symbols: locales.CurrencySymbols{Default: "BRN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRR"}, Symbols: locales.CurrencySymbols{Default: "BRR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRZ"}, Symbols: locales.CurrencySymbols{Default: "BRZ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BSD"}, Symbols: locales.CurrencySymbols{Default: "BSD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BTN"}, Symbols: locales.CurrencySymbols{Default: "BTN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BUK"}, Symbols: locales.CurrencySymbols{Default: "BUK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BWP"}, Symbols: locales.CurrencySymbols{Default: "BWP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BYB"}, Symbols: locales.CurrencySymbols{Default: "BYB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BYN"}, Symbols: locales.CurrencySymbols{Default: "BYN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BYR"}, Symbols: locales.CurrencySymbols{Default: "BYR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BZD"}, Symbols: locales.CurrencySymbols{Default: "BZD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CAD"}, Symbols: locales.CurrencySymbols{Default: "CAD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CDF"}, Symbols: locales.CurrencySymbols{Default: "CDF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CHE"}, Symbols: locales.CurrencySymbols{Default: "CHE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CHF"}, Symbols: locales.CurrencySymbols{Default: "CHF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CHW"}, Symbols: locales.CurrencySymbols{Default: "CHW", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CLE"}, Symbols: locales.CurrencySymbols{Default: "CLE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CLF"}, Symbols: locales.CurrencySymbols{Default: "CLF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CLP"}, Symbols: locales.CurrencySymbols{Default: "CLP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CNH"}, Symbols: locales.CurrencySymbols{Default: "CNH", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CNX"}, Symbols: locales.CurrencySymbols{Default: "CNX", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CNY"}, Symbols: locales.CurrencySymbols{Default: "CNY", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "COP"}, Symbols: locales.CurrencySymbols{Default: "COP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "COU"}, Symbols: locales.CurrencySymbols{Default: "COU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CRC"}, Symbols: locales.CurrencySymbols{Default: "CRC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CSD"}, Symbols: locales.CurrencySymbols{Default: "CSD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CSK"}, Symbols: locales.CurrencySymbols{Default: "CSK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CUC"}, Symbols: locales.CurrencySymbols{Default: "CUC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CUP"}, Symbols: locales.CurrencySymbols{Default: "CUP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CVE"}, Symbols: locales.CurrencySymbols{Default: "CVE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CYP"}, Symbols: locales.CurrencySymbols{Default: "CYP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CZK"}, Symbols: locales.CurrencySymbols{Default: "CZK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DDM"}, Symbols: locales.CurrencySymbols{Default: "DDM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DEM"}, Symbols: locales.CurrencySymbols{Default: "DEM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DJF"}, Symbols: locales.CurrencySymbols{Default: "DJF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DKK"}, Symbols: locales.CurrencySymbols{Default: "DKK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DOP"}, Symbols: locales.CurrencySymbols{Default: "DOP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DZD"}, Symbols: locales.CurrencySymbols{Default: "DZD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ECS"}, Symbols: locales.CurrencySymbols{Default: "ECS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ECV"}, Symbols: locales.CurrencySymbols{Default: "ECV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "EEK"}, Symbols: locales.CurrencySymbols{Default: "EEK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "EGP"}, Symbols: locales.CurrencySymbols{Default: "EGP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ERN"}, Symbols: locales.CurrencySymbols{Default: "ERN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ESA"}, Symbols: locales.CurrencySymbols{Default: "ESA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ESB"}, Symbols: locales.CurrencySymbols{Default: "ESB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ESP"}, Symbols: locales.CurrencySymbols{Default: "ESP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ETB"}, Symbols: locales.CurrencySymbols{Default: "ETB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "EUR"}, Symbols: locales.CurrencySymbols{Default: "EUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "FIM"}, Symbols: locales.CurrencySymbols{Default: "FIM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "FJD"}, Symbols: locales.CurrencySymbols{Default: "FJD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "FKP"}, Symbols: locales.CurrencySymbols{Default: "FKP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "FRF"}, Symbols: locales.CurrencySymbols{Default: "FRF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GBP"}, Symbols: locales.CurrencySymbols{Default: "GBP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GEK"}, Symbols: locales.CurrencySymbols{Default: "GEK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GEL"}, Symbols: locales.CurrencySymbols{Default: "GEL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GHC"}, Symbols: locales.CurrencySymbols{Default: "GHC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GHS"}, Symbols: locales.CurrencySymbols{Default: "GHS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GIP"}, Symbols: locales.CurrencySymbols{Default: "GIP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GMD"}, Symbols: locales.CurrencySymbols{Default: "GMD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GNF"}, Symbols: locales.CurrencySymbols{Default: "GNF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GNS"}, Symbols: locales.CurrencySymbols{Default: "GNS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GQE"}, Symbols: locales.CurrencySymbols{Default: "GQE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GRD"}, Symbols: locales.CurrencySymbols{Default: "GRD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GTQ"}, Symbols: locales.CurrencySymbols{Default: "GTQ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GWE"}, Symbols: locales.CurrencySymbols{Default: "GWE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GWP"}, Symbols: locales.CurrencySymbols{Default: "GWP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GYD"}, Symbols: locales.CurrencySymbols{Default: "GYD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "HKD"}, Symbols: locales.CurrencySymbols{Default: "HKD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "HNL"}, Symbols: locales.CurrencySymbols{Default: "HNL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "HRD"}, Symbols: locales.CurrencySymbols{Default: "HRD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "HRK"}, Symbols: locales.CurrencySymbols{Default: "HRK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "HTG"}, Symbols: locales.CurrencySymbols{Default: "HTG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "HUF"}, Symbols: locales.CurrencySymbols{Default: "HUF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "IDR"}, Symbols: locales.CurrencySymbols{Default: "IDR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "IEP"}, Symbols: locales.CurrencySymbols{Default: "IEP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ILP"}, Symbols: locales.CurrencySymbols{Default: "ILP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ILR"}, Symbols: locales.CurrencySymbols{Default: "ILR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ILS"}, Symbols: locales.CurrencySymbols{Default: "ILS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "INR"}, Symbols: locales.CurrencySymbols{Default: "INR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "IQD"}, Symbols: locales.CurrencySymbols{Default: "IQD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "IRR"}, Symbols: locales.CurrencySymbols{Default: "IRR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ISJ"}, Symbols: locales.CurrencySymbols{Default: "ISJ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ISK"}, Symbols: locales.CurrencySymbols{Default: "ISK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ITL"}, Symbols: locales.CurrencySymbols{Default: "ITL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "JMD"}, Symbols: locales.CurrencySymbols{Default: "JMD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "JOD"}, Symbols: locales.CurrencySymbols{Default: "JOD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "JPY"}, Symbols: locales.CurrencySymbols{Default: "JPY", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KES"}, Symbols: locales.CurrencySymbols{Default: "KES", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KGS"}, Symbols: locales.CurrencySymbols{Default: "KGS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KHR"}, Symbols: locales.CurrencySymbols{Default: "KHR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KMF"}, Symbols: locales.CurrencySymbols{Default: "KMF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KPW"}, Symbols: locales.CurrencySymbols{Default: "KPW", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KRH"}, Symbols: locales.CurrencySymbols{Default: "KRH", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KRO"}, Symbols: locales.CurrencySymbols{Default: "KRO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KRW"}, Symbols: locales.CurrencySymbols{Default: "KRW", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KWD"}, Symbols: locales.CurrencySymbols{Default: "KWD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KYD"}, Symbols: locales.CurrencySymbols{Default: "KYD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KZT"}, Symbols: locales.CurrencySymbols{Default: "KZT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LAK"}, Symbols: locales.CurrencySymbols{Default: "LAK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LBP"}, Symbols: locales.CurrencySymbols{Default: "LBP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LKR"}, Symbols: locales.CurrencySymbols{Default: "LKR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LRD"}, Symbols: locales.CurrencySymbols{Default: "LRD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LSL"}, Symbols: locales.CurrencySymbols{Default: "LSL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LTL"}, Symbols: locales.CurrencySymbols{Default: "LTL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LTT"}, Symbols: locales.CurrencySymbols{Default: "LTT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LUC"}, Symbols: locales.CurrencySymbols{Default: "LUC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LUF"}, Symbols: locales.CurrencySymbols{Default: "LUF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LUL"}, Symbols: locales.CurrencySymbols{Default: "LUL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LVL"}, Symbols: locales.CurrencySymbols{Default: "LVL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LVR"}, Symbols: locales.CurrencySymbols{Default: "LVR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LYD"}, Symbols: locales.CurrencySymbols{Default: "LYD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MAD"}, Symbols: locales.CurrencySymbols{Default: "MAD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MAF"}, Symbols: locales.CurrencySymbols{Default: "MAF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MCF"}, Symbols: locales.CurrencySymbols{Default: "MCF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MDC"}, Symbols: locales.CurrencySymbols{Default: "MDC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MDL"}, Symbols: locales.CurrencySymbols{Default: "MDL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MGA"}, Symbols: locales.CurrencySymbols{Default: "MGA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MGF"}, Symbols: locales.CurrencySymbols{Default: "MGF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MKD"}, Symbols: locales.CurrencySymbols{Default: "MKD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MKN"}, Symbols: locales.CurrencySymbols{Default: "MKN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MLF"}, Symbols: locales.CurrencySymbols{Default: "MLF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MMK"}, Symbols: locales.CurrencySymbols{Default: "MMK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MNT"}, Symbols: locales.CurrencySymbols{Default: "MNT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MOP"}, Symbols: locales.CurrencySymbols{Default: "MOP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MRO"}, Symbols: locales.CurrencySymbols{Default: "MRO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MRU"}, Symbols: locales.CurrencySymbols{Default: "MRU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MTL"}, Symbols: locales.CurrencySymbols{Default: "MTL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MTP"}, Symbols: locales.CurrencySymbols{Default: "MTP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MUR"}, Symbols: locales.CurrencySymbols{Default: "MUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MVP"}, Symbols: locales.CurrencySymbols{Default: "MVP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MVR"}, Symbols: locales.CurrencySymbols{Default: "MVR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MWK"}, Symbols: locales.CurrencySymbols{Default: "MWK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MXN"}, Symbols: locales.CurrencySymbols{Default: "MXN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MXP"}, Symbols: locales.CurrencySymbols{Default: "MXP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MXV"}, Symbols: locales.CurrencySymbols{Default: "MXV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MYR"}, Symbols: locales.CurrencySymbols{Default: "MYR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MZE"}, Symbols: locales.CurrencySymbols{Default: "MZE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MZM"}, Symbols: locales.CurrencySymbols{Default: "MZM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MZN"}, Symbols: locales.CurrencySymbols{Default: "MZN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NAD"}, Symbols: locales.CurrencySymbols{Default: "NAD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NGN"}, Symbols: locales.CurrencySymbols{Default: "NGN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NIC"}, Symbols: locales.CurrencySymbols{Default: "NIC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NIO"}, Symbols: locales.CurrencySymbols{Default: "NIO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NLG"}, Symbols: locales.CurrencySymbols{Default: "NLG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NOK"}, Symbols: locales.CurrencySymbols{Default: "NOK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NPR"}, Symbols: locales.CurrencySymbols{Default: "NPR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NZD"}, Symbols: locales.CurrencySymbols{Default: "NZD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "OMR"}, Symbols: locales.CurrencySymbols{Default: "OMR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PAB"}, Symbols: locales.CurrencySymbols{Default: "PAB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PEI"}, Symbols: locales.CurrencySymbols{Default: "PEI", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PEN"}, Symbols: locales.CurrencySymbols{Default: "PEN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PES"}, Symbols: locales.CurrencySymbols{Default: "PES", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PGK"}, Symbols: locales.CurrencySymbols{Default: "PGK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PHP"}, Symbols: locales.CurrencySymbols{Default: "PHP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PKR"}, Symbols: locales.CurrencySymbols{Default: "PKR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PLN"}, Symbols: locales.CurrencySymbols{Default: "zł", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PLZ"}, Symbols: locales.CurrencySymbols{Default: "PLZ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PTE"}, Symbols: locales.CurrencySymbols{Default: "PTE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PYG"}, Symbols: locales.CurrencySymbols{Default: "PYG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "QAR"}, Symbols: locales.CurrencySymbols{Default: "QAR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RHD"}, Symbols: locales.CurrencySymbols{Default: "RHD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ROL"}, Symbols: locales.CurrencySymbols{Default: "ROL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RON"}, Symbols: locales.CurrencySymbols{Default: "RON", Narrow: "lej"}}, {Names: map[locales.PluralRule]string{0: "RSD"}, Symbols: locales.CurrencySymbols{Default: "RSD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RUB"}, Symbols: locales.CurrencySymbols{Default: "RUB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RUR"}, Symbols: locales.CurrencySymbols{Default: "RUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RWF"}, Symbols: locales.CurrencySymbols{Default: "RWF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SAR"}, Symbols: locales.CurrencySymbols{Default: "SAR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SBD"}, Symbols: locales.CurrencySymbols{Default: "SBD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SCR"}, Symbols: locales.CurrencySymbols{Default: "SCR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SDD"}, Symbols: locales.CurrencySymbols{Default: "SDD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SDG"}, Symbols: locales.CurrencySymbols{Default: "SDG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SDP"}, Symbols: locales.CurrencySymbols{Default: "SDP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SEK"}, Symbols: locales.CurrencySymbols{Default: "SEK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SGD"}, Symbols: locales.CurrencySymbols{Default: "SGD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SHP"}, Symbols: locales.CurrencySymbols{Default: "SHP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SIT"}, Symbols: locales.CurrencySymbols{Default: "SIT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SKK"}, Symbols: locales.CurrencySymbols{Default: "SKK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SLE"}, Symbols: locales.CurrencySymbols{Default: "SLE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SLL"}, Symbols: locales.CurrencySymbols{Default: "SLL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SOS"}, Symbols: locales.CurrencySymbols{Default: "SOS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SRD"}, Symbols: locales.CurrencySymbols{Default: "SRD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SRG"}, Symbols: locales.CurrencySymbols{Default: "SRG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SSP"}, Symbols: locales.CurrencySymbols{Default: "SSP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "STD"}, Symbols: locales.CurrencySymbols{Default: "STD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "STN"}, Symbols: locales.CurrencySymbols{Default: "STN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SUR"}, Symbols: locales.CurrencySymbols{Default: "SUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SVC"}, Symbols: locales.CurrencySymbols{Default: "SVC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SYP"}, Symbols: locales.CurrencySymbols{Default: "SYP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SZL"}, Symbols: locales.CurrencySymbols{Default: "SZL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "THB"}, Symbols: locales.CurrencySymbols{Default: "THB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TJR"}, Symbols: locales.CurrencySymbols{Default: "TJR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TJS"}, Symbols: locales.CurrencySymbols{Default: "TJS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TMM"}, Symbols: locales.CurrencySymbols{Default: "TMM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TMT"}, Symbols: locales.CurrencySymbols{Default: "TMT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TND"}, Symbols: locales.CurrencySymbols{Default: "TND", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TOP"}, Symbols: locales.CurrencySymbols{Default: "TOP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TPE"}, Symbols: locales.CurrencySymbols{Default: "TPE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TRL"}, Symbols: locales.CurrencySymbols{Default: "TRL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TRY"}, Symbols: locales.CurrencySymbols{Default: "TRY", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TTD"}, Symbols: locales.CurrencySymbols{Default: "TTD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TWD"}, Symbols: locales.CurrencySymbols{Default: "TWD", Narrow: "NT$"}}, {Names: map[locales.PluralRule]string{0: "TZS"}, Symbols: locales.CurrencySymbols{Default: "TZS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UAH"}, Symbols: locales.CurrencySymbols{Default: "UAH", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UAK"}, Symbols: locales.CurrencySymbols{Default: "UAK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UGS"}, Symbols: locales.CurrencySymbols{Default: "UGS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UGX"}, Symbols: locales.CurrencySymbols{Default: "UGX", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "USD"}, Symbols: locales.CurrencySymbols{Default: "USD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "USN"}, Symbols: locales.CurrencySymbols{Default: "USN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "USS"}, Symbols: locales.CurrencySymbols{Default: "USS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UYI"}, Symbols: locales.CurrencySymbols{Default: "UYI", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UYP"}, Symbols: locales.CurrencySymbols{Default: "UYP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UYU"}, Symbols: locales.CurrencySymbols{Default: "UYU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UYW"}, Symbols: locales.CurrencySymbols{Default: "UYW", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UZS"}, Symbols: locales.CurrencySymbols{Default: "UZS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VEB"}, Symbols: locales.CurrencySymbols{Default: "VEB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VED"}, Symbols: locales.CurrencySymbols{Default: "VED", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VEF"}, Symbols: locales.CurrencySymbols{Default: "VEF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VES"}, Symbols: locales.CurrencySymbols{Default: "VES", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VND"}, Symbols: locales.CurrencySymbols{Default: "VND", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VNN"}, Symbols: locales.CurrencySymbols{Default: "VNN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VUV"}, Symbols: locales.CurrencySymbols{Default: "VUV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "WST"}, Symbols: locales.CurrencySymbols{Default: "WST", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XAF"}, Symbols: locales.CurrencySymbols{Default: "XAF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XAG"}, Symbols: locales.CurrencySymbols{Default: "XAG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XAU"}, Symbols: locales.CurrencySymbols{Default: "XAU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBA"}, Symbols: locales.CurrencySymbols{Default: "XBA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBB"}, Symbols: locales.CurrencySymbols{Default: "XBB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBC"}, Symbols: locales.CurrencySymbols{Default: "XBC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBD"}, Symbols: locales.CurrencySymbols{Default: "XBD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XCD"}, Symbols: locales.CurrencySymbols{Default: "XCD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XDR"}, Symbols: locales.CurrencySymbols{Default: "XDR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XEU"}, Symbols: locales.CurrencySymbols{Default: "XEU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XFO"}, Symbols: locales.CurrencySymbols{Default: "XFO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XFU"}, Symbols: locales.CurrencySymbols{Default: "XFU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XOF"}, Symbols: locales.CurrencySymbols{Default: "XOF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XPD"}, Symbols: locales.CurrencySymbols{Default: "XPD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XPF"}, Symbols: locales.CurrencySymbols{Default: "XPF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XPT"}, Symbols: locales.CurrencySymbols{Default: "XPT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XRE"}, Symbols: locales.CurrencySymbols{Default: "XRE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XSU"}, Symbols: locales.CurrencySymbols{Default: "XSU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XTS"}, Symbols: locales.CurrencySymbols{Default: "XTS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XUA"}, Symbols: locales.CurrencySymbols{Default: "XUA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XXX"}, Symbols: locales.CurrencySymbols{Default: "XXX", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YDD"}, Symbols: locales.CurrencySymbols{Default: "YDD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YER"}, Symbols: locales.CurrencySymbols{Default: "YER", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUD"}, Symbols: locales.CurrencySymbols{Default: "YUD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUM"}, Symbols: locales.CurrencySymbols{Default: "YUM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUN"}, Symbols: locales.CurrencySymbols{Default: "YUN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUR"}, Symbols: locales.CurrencySymbols{Default: "YUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZAL"}, Symbols: locales.CurrencySymbols{Default: "ZAL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZAR"}, Symbols: locales.CurrencySymbols{Default: "ZAR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZMK"}, Symbols: locales.CurrencySymbols{Default: "ZMK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZMW"}, Symbols: locales.CurrencySymbols{Default: "ZMW", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZRN"}, Symbols: locales.CurrencySymbols{Default: "ZRN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZRZ"}, Symbols: locales.CurrencySymbols{Default: "ZRZ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZWD"}, Symbols: locales.CurrencySymbols{Default: "ZWD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZWL"}, Symbols: locales.CurrencySymbols{Default: "ZWL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZWR"}, Symbols: locales.CurrencySymbols{Default: "ZWR", Narrow: ""}}},
			Formatters: &locales.CurrencyFormatters{
				CurrencyFmt:   locales.MustParseNumberFormatPatterns("#,##0.00 ¤"),
				AccountingFmt: locales.MustParseNumberFormatPatterns("#,##0.00 ¤;(#,##0.00 ¤)"),
				Short: locales.CurrencyAccountingFormatterByExp{
					CurrencyFmt: map[uint8]*locales.CurrencyAccountingFormatterByExpPlural{
						3: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("0 tys'.' ¤"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("0 tys'.' ¤"),
						}},
						4: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("00 tys'.' ¤"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("00 tys'.' ¤"),
						}},
						5: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("000 tys'.' ¤"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("000 tys'.' ¤"),
						}},
						6: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("0 mln ¤"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("0 mln ¤"),
						}},
						7: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("00 mln ¤"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("00 mln ¤"),
						}},
						8: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("000 mln ¤"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("000 mln ¤"),
						}},
						9: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("0 mld ¤"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("0 mld ¤"),
						}},
						10: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("00 mld ¤"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("00 mld ¤"),
						}},
						11: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("000 mld ¤"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("000 mld ¤"),
						}},
						12: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("0 bln ¤"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("0 bln ¤"),
						}},
						13: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("00 bln ¤"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("00 bln ¤"),
						}},
						14: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("000 bln ¤"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("000 bln ¤"),
						}},
					},
					AccountingFmt: map[uint8]*locales.CurrencyAccountingFormatterByExpPlural{},
				},
			},
		},
		TimeSpec: &locales.TimeSpec{
			Separator:                ":",
			MonthsAbbreviatedValue:   []string{"", "sty", "lut", "mar", "kwi", "maj", "cze", "lip", "sie", "wrz", "paź", "lis", "gru"},
			MonthsNarrowValue:        []string{"", "s", "l", "m", "k", "m", "c", "l", "s", "w", "p", "l", "g"},
			MonthsWideValue:          []string{"", "stycznia", "lutego", "marca", "kwietnia", "maja", "czerwca", "lipca", "sierpnia", "września", "października", "listopada", "grudnia"},
			WeekdaysAbbreviatedValue: []string{"niedz.", "pon.", "wt.", "śr.", "czw.", "pt.", "sob."},
			WeekdaysNarrowValue:      []string{"n", "p", "w", "ś", "c", "p", "s"},
			WeekdaysShortValue:       []string{"nie", "pon", "wto", "śro", "czw", "pią", "sob"},
			WeekdaysWideValue:        []string{"niedziela", "poniedziałek", "wtorek", "środa", "czwartek", "piątek", "sobota"},
			PeriodsAbbreviatedValue:  []string{"AM", "PM"},
			PeriodsNarrowValue:       []string{"a", "p"},
			PeriodsWideValue:         []string{"AM", "PM"},
			ErasAbbreviatedValue:     []string{"p.n.e.", "n.e."},
			ErasNarrowValue:          []string{"", ""},
			ErasWideValue:            []string{"przed naszą erą", "naszej ery"},
			TimezonesValue:           map[string]string{"": "Falklandy (czas standardowy)", "ACDT": "czas środkowoaustralijski letni", "ACST": "czas środkowoaustralijski standardowy", "ACT": "ACT", "ACWDT": "czas środkowo-zachodnioaustralijski letni", "ACWST": "czas środkowo-zachodnioaustralijski standardowy", "ADT": "czas atlantycki letni", "AEDT": "czas wschodnioaustralijski letni", "AEST": "czas wschodnioaustralijski standardowy", "AFT": "czas Afganistan", "AKDT": "Alaska (czas letni)", "AKST": "Alaska (czas standardowy)", "AMST": "czas amazoński letni", "AMT": "czas amazoński standardowy", "ARST": "Argentyna (czas letni)", "ART": "Argentyna (czas standardowy)", "AST": "czas atlantycki standardowy", "AWDT": "czas zachodnioaustralijski letni", "AWST": "czas zachodnioaustralijski standardowy", "BNT": "czas Brunei", "BOT": "czas Boliwia", "BRST": "Brasília (czas letni)", "BRT": "Brasília (czas standardowy)", "BST": "Bangladesz (czas letni)", "BTT": "czas Bhutan", "CAT": "czas środkowoafrykański", "CCT": "czas Wyspy Kokosowe", "CDT": "czas środkowoamerykański letni", "CDT (Kuba)": "Kuba (czas letni)", "CHADT": "Chatham (czas letni)", "CHAST": "Chatham (czas standardowy)", "CLST": "Chile (czas letni)", "CLT": "Chile (czas standardowy)", "COST": "Kolumbia (czas letni)", "COT": "Kolumbia (czas standardowy)", "CST": "czas środkowoamerykański standardowy", "CST (Kuba)": "Kuba (czas standardowy)", "CXT": "czas Wyspa Bożego Narodzenia", "ChST": "czas Czamorro", "EASST": "Wyspa Wielkanocna (czas letni)", "EAST": "Wyspa Wielkanocna (czas standardowy)", "EAT": "czas wschodnioafrykański", "ECT": "czas Ekwador", "EDT": "czas wschodnioamerykański letni", "EGDT": "Grenlandia Wschodnia (czas letni)", "EGST": "Grenlandia Wschodnia (czas standardowy)", "EST": "czas wschodnioamerykański standardowy", "FKST": "Falklandy (czas letni)", "GALT": "czas Galapagos", "GFT": "czas Gujana Francuska", "GMT": "czas uniwersalny", "GST": "czas Zatoka Perska", "GYT": "czas Gujana", "HADT": "Hawaje-Aleuty (czas letni)", "HAST": "Hawaje-Aleuty (czas standardowy)", "HENOMX": "Meksyk Północno-Zachodni (czas letni)", "HEOG": "Grenlandia Zachodnia (czas letni)", "HEPMX": "Meksyk (czas pacyficzny letni)", "HKST": "Hongkong (czas letni)", "HKT": "Hongkong (czas standardowy)", "HNNOMX": "Meksyk Północno-Zachodni (czas standardowy)", "HNOG": "Grenlandia Zachodnia (czas standardowy)", "HNPMX": "Meksyk (czas pacyficzny standardowy)", "ICT": "czas indochiński", "IRDT": "Iran (czas letni)", "IRST": "Iran (czas standardowy)", "IST": "czas indyjski standardowy", "JDT": "Japonia (czas letni)", "JST": "Japonia (czas standardowy)", "LHDT": "Lord Howe (czas letni)", "LHST": "Lord Howe (czas standardowy)", "MDT": "MDT", "MESZ": "czas środkowoeuropejski letni", "MEZ": "czas środkowoeuropejski standardowy", "MST": "MST", "MVT": "czas Malediwy", "MYT": "czas Malezja", "NDT": "Nowa Fundlandia (czas letni)", "NPT": "czas Nepal", "NST": "Nowa Fundlandia (czas standardowy)", "NZDT": "Nowa Zelandia (czas letni)", "NZST": "Nowa Zelandia (czas standardowy)", "OESZ": "czas wschodnioeuropejski letni", "OEZ": "czas wschodnioeuropejski standardowy", "PDT": "czas pacyficzny letni", "PKT": "Pakistan (czas letni)", "PMDT": "Saint-Pierre i Miquelon (czas letni)", "PMST": "Saint-Pierre i Miquelon (czas standardowy)", "PST": "czas pacyficzny standardowy", "PYST": "Paragwaj (czas letni)", "SAST": "czas południowoafrykański", "SGT": "czas Singapur", "SRT": "czas Surinam", "TLT": "czas Timor Wschodni", "TMST": "Turkmenistan (czas letni)", "TMT": "Turkmenistan (czas standardowy)", "UYST": "Urugwaj (czas letni)", "UYT": "Urugwaj (czas standardowy)", "VET": "czas Wenezuela", "WARST": "Argentyna Zachodnia (czas letni)", "WART": "Argentyna Zachodnia (czas standardowy)", "WAST": "czas zachodnioafrykański letni", "WAT": "czas zachodnioafrykański standardowy", "WESZ": "czas zachodnioeuropejski letni", "WEZ": "czas zachodnioeuropejski standardowy", "WIB": "czas Indonezja Zachodnia", "WIT": "czas Indonezja Wschodnia", "WITA": "czas Indonezja Środkowa", "∅∅∅": "Peru (czas letni)"},
		},

		locale:          "pl",
		pluralsCardinal: []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:  []locales.PluralRule{6},
		pluralsRange:    []locales.PluralRule{2, 4, 5, 6},
		percent:         "%",
		perMille:        "‰",
		inifinity:       "∞",

		dateFullLayout:   "EEEE, d MMMM y",
		dateLongLayout:   "d MMMM y",
		dateMediumLayout: "d MMM y",
		dateShortLayout:  "d.MM.y",
		timeFullLayout:   "HH:mm:ss zzzz",
		timeLongLayout:   "HH:mm:ss z",
		timeMediumLayout: "HH:mm:ss",
		timeShortLayout:  "HH:mm",
		listPatterns:     locales.NewListPatternsFromSlice([][]string{{"{0} i {1}", "{0}, {1}", "{0}, {1}", "{0} i {1}"}, {"{0} lub {1}", "{0}, {1}", "{0}, {1}", "{0} lub {1}"}}),
		DurationSpec: &locales.DurationSpec{" ",
			"",
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"w.", "%v w.", "%v w.", ""},
				Long:   &locales.CounterFormat{"wieki", "%v wieku", "%v wieku", ""}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"dek", "%v dek", "%v dek", ""},
				Long:   &locales.CounterFormat{"dekady", "%v dekady", "%v dekady", ""}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"r.", "%v r.", "%v r.", ""},
				Long:   &locales.CounterFormat{"lata", "%v roku", "%v roku", "%v na rok"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"m-c", "%v m-c", "%v m-ca", "%v/m-c"},
				Long:   &locales.CounterFormat{"miesiące", "%v miesiącu", "%v miesiąca", "%v na miesiąc"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"tydz.", "%v tydz.", "%v tyg.", "%v/t."},
				Long:   &locales.CounterFormat{"tygodnie", "%v tygodniu", "%v tygodnia", "%v na tydzień"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"doba", "%v d.", "%v d.", "%v/d."}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"h", "%v h", "%v h", "%v/h"},
				Long:   &locales.CounterFormat{"godziny", "%v godzino", "%v godziny", "%v na godzinę"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Long:   &locales.CounterFormat{"minuty", "%v minuto", "%v minuty", "%v na minutę"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"s", "%v s", "%v s", ""},
				Long:   &locales.CounterFormat{"sekundy", "%v sekundo", "%v sekundy", "%v na sekundę"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Long:   &locales.CounterFormat{"milisekundy", "%v milisekundy", "%v milisekundy", ""}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Long:   &locales.CounterFormat{"mikrosekundy", "%v mikrosekundy", "%v mikrosekundy", ""}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Long:   &locales.CounterFormat{"nanosekundy", "%v nanosekundy", "%v nanosekundy", ""}},
		},
		miscPatterns: locales.MiscPatterns{
			"~%[1]v",
			"≥%[1]v",
			"≤%[1]v",
			"%[1]v–%[2]v",
		},
	}
}

// Locale returns the current translators string locale
func (pl *pl) Locale() string {
	return pl.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'pl'
func (pl *pl) PluralsCardinal() []locales.PluralRule {
	return pl.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'pl'
func (pl *pl) PluralsOrdinal() []locales.PluralRule {
	return pl.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'pl'
func (pl *pl) PluralsRange() []locales.PluralRule {
	return pl.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'pl'
func (pl *pl) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	iMod10 := i % 10
	iMod100 := i % 100

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	} else if v == 0 && iMod10 >= 2 && iMod10 <= 4 && (iMod100 < 12 || iMod100 > 14) {
		return locales.PluralRuleFew
	} else if (v == 0 && i != 1 && iMod10 >= 0 && iMod10 <= 1) || (v == 0 && iMod10 >= 5 && iMod10 <= 9) || (v == 0 && iMod100 >= 12 && iMod100 <= 14) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'pl'
func (pl *pl) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'pl'
func (pl *pl) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

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

// FmtDateShort returns the short date representation of 't' for 'pl'
func (pl *pl) FmtDateShort(t time.Time) string {
	return locales.FormatTimeEra(pl, pl.dateShortLayout, t, 2)
}

// FmtDateMedium returns the medium date representation of 't' for 'pl'
func (pl *pl) FmtDateMedium(t time.Time) string {
	return locales.FormatTimeEra(pl, pl.dateMediumLayout, t, 2)
}

// FmtDateLong returns the long date representation of 't' for 'pl'
func (pl *pl) FmtDateLong(t time.Time) string {
	return locales.FormatTimeEra(pl, pl.dateLongLayout, t, 1)
}

// FmtDateFull returns the full date representation of 't' for 'pl'
func (pl *pl) FmtDateFull(t time.Time) string {
	return locales.FormatTimeEra(pl, pl.dateFullLayout, t, 0)
}

// FmtTimeShort returns the short time representation of 't' for 'pl'
func (pl *pl) FmtTimeShort(t time.Time) string {
	return locales.FormatTimeEra(pl, pl.timeShortLayout, t, 2)
}

// FmtTimeMedium returns the medium time representation of 't' for 'pl'
func (pl *pl) FmtTimeMedium(t time.Time) string {
	return locales.FormatTimeEra(pl, pl.timeMediumLayout, t, 2)
}

// FmtTimeLong returns the long time representation of 't' for 'pl'
func (pl *pl) FmtTimeLong(t time.Time) string {
	return locales.FormatTimeEra(pl, pl.timeLongLayout, t, 1)
}

// FmtTimeFull returns the full time representation of 't' for 'pl'
func (pl *pl) FmtTimeFull(t time.Time) string {
	return locales.FormatTimeEra(pl, pl.timeFullLayout, t, 0)
}

// DateFullLayout returns the full date layout for 'pl'
func (pl *pl) DateFullLayout() string {
	return pl.dateFullLayout
}

// DateLongLayout returns the long date layout for 'pl'
func (pl *pl) DateLongLayout() string {
	return pl.dateLongLayout
}

// DateMediumLayout returns the medium date layout for 'pl'
func (pl *pl) DateMediumLayout() string {
	return pl.dateMediumLayout
}

// DateShortLayout returns the short date layout for 'pl'
func (pl *pl) DateShortLayout() string {
	return pl.dateShortLayout
}

// TimeFullLayout returns the full time layout for 'pl'
func (pl *pl) TimeFullLayout() string {
	return pl.timeFullLayout
}

// TimeLongLayout returns the full long layout for 'pl'
func (pl *pl) TimeLongLayout() string {
	return pl.timeLongLayout
}

// TimeMediumLayout returns the medium time layout for 'pl'
func (pl *pl) TimeMediumLayout() string {
	return pl.timeMediumLayout
}

// TimeShortLayout returns the short time layout for 'pl'
func (pl *pl) TimeShortLayout() string {
	return pl.timeShortLayout
}

func (pl *pl) ListPatterns() *locales.ListPatterns {
	return pl.listPatterns
}

func (pl *pl) GetDurationSpec() *locales.DurationSpec {
	return pl.DurationSpec
}

func (pl *pl) GetMiscPatterns() *locales.MiscPatterns {
	return &pl.miscPatterns
}