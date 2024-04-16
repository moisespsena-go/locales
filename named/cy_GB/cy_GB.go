package cy_GB

import (
	"math"
	"time"

	"github.com/moisespsena-go/locales"
)

type cy_GB struct {
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

// New returns a new instance of translator for the 'cy_GB' locale
func New() locales.Translator {
	return &cy_GB{
		NumberSpec: &locales.NumberSpec{
			MinusValue:   "-",
			GroupValue:   ",",
			DecimalValue: ".",
		},
		CurrencySpec: &locales.CurrencySpec{
			CurrenciesValue: []locales.Currency{{Names: map[locales.PluralRule]string{0: "ADP"}, Symbols: locales.CurrencySymbols{Default: "ADP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AED"}, Symbols: locales.CurrencySymbols{Default: "AED", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AFA"}, Symbols: locales.CurrencySymbols{Default: "AFA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AFN"}, Symbols: locales.CurrencySymbols{Default: "AFN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ALK"}, Symbols: locales.CurrencySymbols{Default: "ALK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ALL"}, Symbols: locales.CurrencySymbols{Default: "ALL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AMD"}, Symbols: locales.CurrencySymbols{Default: "AMD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ANG"}, Symbols: locales.CurrencySymbols{Default: "ANG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AOA"}, Symbols: locales.CurrencySymbols{Default: "AOA", Narrow: "Kz"}}, {Names: map[locales.PluralRule]string{0: "AOK"}, Symbols: locales.CurrencySymbols{Default: "AOK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AON"}, Symbols: locales.CurrencySymbols{Default: "AON", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AOR"}, Symbols: locales.CurrencySymbols{Default: "AOR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARA"}, Symbols: locales.CurrencySymbols{Default: "ARA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARL"}, Symbols: locales.CurrencySymbols{Default: "ARL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARM"}, Symbols: locales.CurrencySymbols{Default: "ARM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARP"}, Symbols: locales.CurrencySymbols{Default: "ARP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARS"}, Symbols: locales.CurrencySymbols{Default: "ARS", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "ATS"}, Symbols: locales.CurrencySymbols{Default: "ATS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AUD"}, Symbols: locales.CurrencySymbols{Default: "A$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "AWG"}, Symbols: locales.CurrencySymbols{Default: "AWG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AZM"}, Symbols: locales.CurrencySymbols{Default: "AZM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AZN"}, Symbols: locales.CurrencySymbols{Default: "AZN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BAD"}, Symbols: locales.CurrencySymbols{Default: "BAD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BAM"}, Symbols: locales.CurrencySymbols{Default: "BAM", Narrow: "KM"}}, {Names: map[locales.PluralRule]string{0: "BAN"}, Symbols: locales.CurrencySymbols{Default: "BAN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BBD"}, Symbols: locales.CurrencySymbols{Default: "BBD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "BDT"}, Symbols: locales.CurrencySymbols{Default: "BDT", Narrow: "TK"}}, {Names: map[locales.PluralRule]string{0: "BEC"}, Symbols: locales.CurrencySymbols{Default: "BEC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BEF"}, Symbols: locales.CurrencySymbols{Default: "BEF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BEL"}, Symbols: locales.CurrencySymbols{Default: "BEL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGL"}, Symbols: locales.CurrencySymbols{Default: "BGL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGM"}, Symbols: locales.CurrencySymbols{Default: "BGM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGN"}, Symbols: locales.CurrencySymbols{Default: "BGN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGO"}, Symbols: locales.CurrencySymbols{Default: "BGO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BHD"}, Symbols: locales.CurrencySymbols{Default: "BHD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BIF"}, Symbols: locales.CurrencySymbols{Default: "BIF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BMD"}, Symbols: locales.CurrencySymbols{Default: "BMD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "BND"}, Symbols: locales.CurrencySymbols{Default: "BND", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "BOB"}, Symbols: locales.CurrencySymbols{Default: "BOB", Narrow: "Bs"}}, {Names: map[locales.PluralRule]string{0: "BOL"}, Symbols: locales.CurrencySymbols{Default: "BOL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BOP"}, Symbols: locales.CurrencySymbols{Default: "BOP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BOV"}, Symbols: locales.CurrencySymbols{Default: "BOV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRB"}, Symbols: locales.CurrencySymbols{Default: "BRB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRC"}, Symbols: locales.CurrencySymbols{Default: "BRC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRE"}, Symbols: locales.CurrencySymbols{Default: "BRE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRL"}, Symbols: locales.CurrencySymbols{Default: "R$", Narrow: "R$"}}, {Names: map[locales.PluralRule]string{0: "BRN"}, Symbols: locales.CurrencySymbols{Default: "BRN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRR"}, Symbols: locales.CurrencySymbols{Default: "BRR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRZ"}, Symbols: locales.CurrencySymbols{Default: "BRZ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BSD"}, Symbols: locales.CurrencySymbols{Default: "BSD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "BTN"}, Symbols: locales.CurrencySymbols{Default: "BTN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BUK"}, Symbols: locales.CurrencySymbols{Default: "BUK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BWP"}, Symbols: locales.CurrencySymbols{Default: "BWP", Narrow: "BWP"}}, {Names: map[locales.PluralRule]string{0: "BYB"}, Symbols: locales.CurrencySymbols{Default: "BYB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BYN"}, Symbols: locales.CurrencySymbols{Default: "BYN", Narrow: "р."}}, {Names: map[locales.PluralRule]string{0: "BYR"}, Symbols: locales.CurrencySymbols{Default: "BYR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BZD"}, Symbols: locales.CurrencySymbols{Default: "BZD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CAD"}, Symbols: locales.CurrencySymbols{Default: "CA$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CDF"}, Symbols: locales.CurrencySymbols{Default: "CDF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CHE"}, Symbols: locales.CurrencySymbols{Default: "CHE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CHF"}, Symbols: locales.CurrencySymbols{Default: "CHF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CHW"}, Symbols: locales.CurrencySymbols{Default: "CHW", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CLE"}, Symbols: locales.CurrencySymbols{Default: "CLE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CLF"}, Symbols: locales.CurrencySymbols{Default: "CLF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CLP"}, Symbols: locales.CurrencySymbols{Default: "CLP", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CNH"}, Symbols: locales.CurrencySymbols{Default: "CNH", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CNX"}, Symbols: locales.CurrencySymbols{Default: "CNX", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CNY"}, Symbols: locales.CurrencySymbols{Default: "CN¥", Narrow: "¥"}}, {Names: map[locales.PluralRule]string{0: "COP"}, Symbols: locales.CurrencySymbols{Default: "COP", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "COU"}, Symbols: locales.CurrencySymbols{Default: "COU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CRC"}, Symbols: locales.CurrencySymbols{Default: "CRC", Narrow: "₡"}}, {Names: map[locales.PluralRule]string{0: "CSD"}, Symbols: locales.CurrencySymbols{Default: "CSD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CSK"}, Symbols: locales.CurrencySymbols{Default: "CSK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CUC"}, Symbols: locales.CurrencySymbols{Default: "CUC", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CUP"}, Symbols: locales.CurrencySymbols{Default: "CUP", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CVE"}, Symbols: locales.CurrencySymbols{Default: "CVE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CYP"}, Symbols: locales.CurrencySymbols{Default: "CYP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CZK"}, Symbols: locales.CurrencySymbols{Default: "CZK", Narrow: "Kč"}}, {Names: map[locales.PluralRule]string{0: "DDM"}, Symbols: locales.CurrencySymbols{Default: "DDM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DEM"}, Symbols: locales.CurrencySymbols{Default: "DEM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DJF"}, Symbols: locales.CurrencySymbols{Default: "DJF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DKK"}, Symbols: locales.CurrencySymbols{Default: "DKK", Narrow: "kr"}}, {Names: map[locales.PluralRule]string{0: "DOP"}, Symbols: locales.CurrencySymbols{Default: "DOP", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "DZD"}, Symbols: locales.CurrencySymbols{Default: "DZD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ECS"}, Symbols: locales.CurrencySymbols{Default: "ECS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ECV"}, Symbols: locales.CurrencySymbols{Default: "ECV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "EEK"}, Symbols: locales.CurrencySymbols{Default: "EEK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "EGP"}, Symbols: locales.CurrencySymbols{Default: "EGP", Narrow: "E£"}}, {Names: map[locales.PluralRule]string{0: "ERN"}, Symbols: locales.CurrencySymbols{Default: "ERN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ESA"}, Symbols: locales.CurrencySymbols{Default: "ESA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ESB"}, Symbols: locales.CurrencySymbols{Default: "ESB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ESP"}, Symbols: locales.CurrencySymbols{Default: "ESP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ETB"}, Symbols: locales.CurrencySymbols{Default: "ETB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "EUR"}, Symbols: locales.CurrencySymbols{Default: "€", Narrow: "€"}}, {Names: map[locales.PluralRule]string{0: "FIM"}, Symbols: locales.CurrencySymbols{Default: "FIM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "FJD"}, Symbols: locales.CurrencySymbols{Default: "FJD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "FKP"}, Symbols: locales.CurrencySymbols{Default: "FKP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "FRF"}, Symbols: locales.CurrencySymbols{Default: "FRF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GBP"}, Symbols: locales.CurrencySymbols{Default: "£", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "GEK"}, Symbols: locales.CurrencySymbols{Default: "GEK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GEL"}, Symbols: locales.CurrencySymbols{Default: "GEL", Narrow: "₾"}}, {Names: map[locales.PluralRule]string{0: "GHC"}, Symbols: locales.CurrencySymbols{Default: "GHC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GHS"}, Symbols: locales.CurrencySymbols{Default: "GHS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GIP"}, Symbols: locales.CurrencySymbols{Default: "GIP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "GMD"}, Symbols: locales.CurrencySymbols{Default: "GMD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GNF"}, Symbols: locales.CurrencySymbols{Default: "GNF", Narrow: "FG"}}, {Names: map[locales.PluralRule]string{0: "GNS"}, Symbols: locales.CurrencySymbols{Default: "GNS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GQE"}, Symbols: locales.CurrencySymbols{Default: "GQE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GRD"}, Symbols: locales.CurrencySymbols{Default: "GRD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GTQ"}, Symbols: locales.CurrencySymbols{Default: "GTQ", Narrow: "Q"}}, {Names: map[locales.PluralRule]string{0: "GWE"}, Symbols: locales.CurrencySymbols{Default: "GWE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GWP"}, Symbols: locales.CurrencySymbols{Default: "GWP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GYD"}, Symbols: locales.CurrencySymbols{Default: "GYD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "HKD"}, Symbols: locales.CurrencySymbols{Default: "HK$", Narrow: "HK$"}}, {Names: map[locales.PluralRule]string{0: "HNL"}, Symbols: locales.CurrencySymbols{Default: "HNL", Narrow: "L"}}, {Names: map[locales.PluralRule]string{0: "HRD"}, Symbols: locales.CurrencySymbols{Default: "HRD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "HRK"}, Symbols: locales.CurrencySymbols{Default: "HRK", Narrow: "kn"}}, {Names: map[locales.PluralRule]string{0: "HTG"}, Symbols: locales.CurrencySymbols{Default: "HTG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "HUF"}, Symbols: locales.CurrencySymbols{Default: "HUF", Narrow: "Ft"}}, {Names: map[locales.PluralRule]string{0: "IDR"}, Symbols: locales.CurrencySymbols{Default: "IDR", Narrow: "Rp"}}, {Names: map[locales.PluralRule]string{0: "IEP"}, Symbols: locales.CurrencySymbols{Default: "IEP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ILP"}, Symbols: locales.CurrencySymbols{Default: "ILP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ILR"}, Symbols: locales.CurrencySymbols{Default: "ILR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ILS"}, Symbols: locales.CurrencySymbols{Default: "₪", Narrow: "₪"}}, {Names: map[locales.PluralRule]string{0: "INR"}, Symbols: locales.CurrencySymbols{Default: "₹", Narrow: "₹"}}, {Names: map[locales.PluralRule]string{0: "IQD"}, Symbols: locales.CurrencySymbols{Default: "IQD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "IRR"}, Symbols: locales.CurrencySymbols{Default: "IRR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ISJ"}, Symbols: locales.CurrencySymbols{Default: "ISJ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ISK"}, Symbols: locales.CurrencySymbols{Default: "ISK", Narrow: "kr"}}, {Names: map[locales.PluralRule]string{0: "ITL"}, Symbols: locales.CurrencySymbols{Default: "ITL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "JMD"}, Symbols: locales.CurrencySymbols{Default: "JMD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "JOD"}, Symbols: locales.CurrencySymbols{Default: "JOD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "JPY"}, Symbols: locales.CurrencySymbols{Default: "JP¥", Narrow: "¥"}}, {Names: map[locales.PluralRule]string{0: "KES"}, Symbols: locales.CurrencySymbols{Default: "KES", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KGS"}, Symbols: locales.CurrencySymbols{Default: "KGS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KHR"}, Symbols: locales.CurrencySymbols{Default: "KHR", Narrow: "៛"}}, {Names: map[locales.PluralRule]string{0: "KMF"}, Symbols: locales.CurrencySymbols{Default: "KMF", Narrow: "CF"}}, {Names: map[locales.PluralRule]string{0: "KPW"}, Symbols: locales.CurrencySymbols{Default: "KPW", Narrow: "₩"}}, {Names: map[locales.PluralRule]string{0: "KRH"}, Symbols: locales.CurrencySymbols{Default: "KRH", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KRO"}, Symbols: locales.CurrencySymbols{Default: "KRO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KRW"}, Symbols: locales.CurrencySymbols{Default: "KRW", Narrow: "₩"}}, {Names: map[locales.PluralRule]string{0: "KWD"}, Symbols: locales.CurrencySymbols{Default: "KWD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KYD"}, Symbols: locales.CurrencySymbols{Default: "KYD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "KZT"}, Symbols: locales.CurrencySymbols{Default: "KZT", Narrow: "₸"}}, {Names: map[locales.PluralRule]string{0: "LAK"}, Symbols: locales.CurrencySymbols{Default: "LAK", Narrow: "₭"}}, {Names: map[locales.PluralRule]string{0: "LBP"}, Symbols: locales.CurrencySymbols{Default: "LBP", Narrow: "L£"}}, {Names: map[locales.PluralRule]string{0: "LKR"}, Symbols: locales.CurrencySymbols{Default: "LKR", Narrow: "Rs"}}, {Names: map[locales.PluralRule]string{0: "LRD"}, Symbols: locales.CurrencySymbols{Default: "LRD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "LSL"}, Symbols: locales.CurrencySymbols{Default: "LSL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LTL"}, Symbols: locales.CurrencySymbols{Default: "LTL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LTT"}, Symbols: locales.CurrencySymbols{Default: "LTT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LUC"}, Symbols: locales.CurrencySymbols{Default: "LUC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LUF"}, Symbols: locales.CurrencySymbols{Default: "LUF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LUL"}, Symbols: locales.CurrencySymbols{Default: "LUL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LVL"}, Symbols: locales.CurrencySymbols{Default: "LVL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LVR"}, Symbols: locales.CurrencySymbols{Default: "LVR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LYD"}, Symbols: locales.CurrencySymbols{Default: "LYD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MAD"}, Symbols: locales.CurrencySymbols{Default: "MAD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MAF"}, Symbols: locales.CurrencySymbols{Default: "MAF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MCF"}, Symbols: locales.CurrencySymbols{Default: "MCF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MDC"}, Symbols: locales.CurrencySymbols{Default: "MDC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MDL"}, Symbols: locales.CurrencySymbols{Default: "MDL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MGA"}, Symbols: locales.CurrencySymbols{Default: "MGA", Narrow: "Ar"}}, {Names: map[locales.PluralRule]string{0: "MGF"}, Symbols: locales.CurrencySymbols{Default: "MGF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MKD"}, Symbols: locales.CurrencySymbols{Default: "MKD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MKN"}, Symbols: locales.CurrencySymbols{Default: "MKN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MLF"}, Symbols: locales.CurrencySymbols{Default: "MLF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MMK"}, Symbols: locales.CurrencySymbols{Default: "MMK", Narrow: "K"}}, {Names: map[locales.PluralRule]string{0: "MNT"}, Symbols: locales.CurrencySymbols{Default: "MNT", Narrow: "₮"}}, {Names: map[locales.PluralRule]string{0: "MOP"}, Symbols: locales.CurrencySymbols{Default: "MOP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MRO"}, Symbols: locales.CurrencySymbols{Default: "MRO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MRU"}, Symbols: locales.CurrencySymbols{Default: "MRU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MTL"}, Symbols: locales.CurrencySymbols{Default: "MTL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MTP"}, Symbols: locales.CurrencySymbols{Default: "MTP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MUR"}, Symbols: locales.CurrencySymbols{Default: "MUR", Narrow: "Rs"}}, {Names: map[locales.PluralRule]string{0: "MVP"}, Symbols: locales.CurrencySymbols{Default: "MVP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MVR"}, Symbols: locales.CurrencySymbols{Default: "MVR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MWK"}, Symbols: locales.CurrencySymbols{Default: "MWK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MXN"}, Symbols: locales.CurrencySymbols{Default: "MX$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "MXP"}, Symbols: locales.CurrencySymbols{Default: "MXP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MXV"}, Symbols: locales.CurrencySymbols{Default: "MXV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MYR"}, Symbols: locales.CurrencySymbols{Default: "MYR", Narrow: "RM"}}, {Names: map[locales.PluralRule]string{0: "MZE"}, Symbols: locales.CurrencySymbols{Default: "MZE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MZM"}, Symbols: locales.CurrencySymbols{Default: "MZM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MZN"}, Symbols: locales.CurrencySymbols{Default: "MZN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NAD"}, Symbols: locales.CurrencySymbols{Default: "NAD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "NGN"}, Symbols: locales.CurrencySymbols{Default: "NGN", Narrow: "₦"}}, {Names: map[locales.PluralRule]string{0: "NIC"}, Symbols: locales.CurrencySymbols{Default: "NIC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NIO"}, Symbols: locales.CurrencySymbols{Default: "NIO", Narrow: "C$"}}, {Names: map[locales.PluralRule]string{0: "NLG"}, Symbols: locales.CurrencySymbols{Default: "NLG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NOK"}, Symbols: locales.CurrencySymbols{Default: "NOK", Narrow: "kr"}}, {Names: map[locales.PluralRule]string{0: "NPR"}, Symbols: locales.CurrencySymbols{Default: "NPR", Narrow: "Rs"}}, {Names: map[locales.PluralRule]string{0: "NZD"}, Symbols: locales.CurrencySymbols{Default: "NZ$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "OMR"}, Symbols: locales.CurrencySymbols{Default: "OMR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PAB"}, Symbols: locales.CurrencySymbols{Default: "PAB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PEI"}, Symbols: locales.CurrencySymbols{Default: "PEI", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PEN"}, Symbols: locales.CurrencySymbols{Default: "PEN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PES"}, Symbols: locales.CurrencySymbols{Default: "PES", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PGK"}, Symbols: locales.CurrencySymbols{Default: "PGK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PHP"}, Symbols: locales.CurrencySymbols{Default: "PHP", Narrow: "₱"}}, {Names: map[locales.PluralRule]string{0: "PKR"}, Symbols: locales.CurrencySymbols{Default: "PKR", Narrow: "Rs"}}, {Names: map[locales.PluralRule]string{0: "PLN"}, Symbols: locales.CurrencySymbols{Default: "PLN", Narrow: "zł"}}, {Names: map[locales.PluralRule]string{0: "PLZ"}, Symbols: locales.CurrencySymbols{Default: "PLZ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PTE"}, Symbols: locales.CurrencySymbols{Default: "PTE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PYG"}, Symbols: locales.CurrencySymbols{Default: "PYG", Narrow: "₲"}}, {Names: map[locales.PluralRule]string{0: "QAR"}, Symbols: locales.CurrencySymbols{Default: "QAR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RHD"}, Symbols: locales.CurrencySymbols{Default: "RHD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ROL"}, Symbols: locales.CurrencySymbols{Default: "ROL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RON"}, Symbols: locales.CurrencySymbols{Default: "RON", Narrow: "lei"}}, {Names: map[locales.PluralRule]string{0: "RSD"}, Symbols: locales.CurrencySymbols{Default: "RSD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RUB"}, Symbols: locales.CurrencySymbols{Default: "RUB", Narrow: "₽"}}, {Names: map[locales.PluralRule]string{0: "RUR"}, Symbols: locales.CurrencySymbols{Default: "RUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RWF"}, Symbols: locales.CurrencySymbols{Default: "RWF", Narrow: "RF"}}, {Names: map[locales.PluralRule]string{0: "SAR"}, Symbols: locales.CurrencySymbols{Default: "SAR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SBD"}, Symbols: locales.CurrencySymbols{Default: "SBD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "SCR"}, Symbols: locales.CurrencySymbols{Default: "SCR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SDD"}, Symbols: locales.CurrencySymbols{Default: "SDD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SDG"}, Symbols: locales.CurrencySymbols{Default: "SDG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SDP"}, Symbols: locales.CurrencySymbols{Default: "SDP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SEK"}, Symbols: locales.CurrencySymbols{Default: "SEK", Narrow: "kr"}}, {Names: map[locales.PluralRule]string{0: "SGD"}, Symbols: locales.CurrencySymbols{Default: "SGD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "SHP"}, Symbols: locales.CurrencySymbols{Default: "SHP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "SIT"}, Symbols: locales.CurrencySymbols{Default: "SIT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SKK"}, Symbols: locales.CurrencySymbols{Default: "SKK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SLE"}, Symbols: locales.CurrencySymbols{Default: "SLE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SLL"}, Symbols: locales.CurrencySymbols{Default: "SLL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SOS"}, Symbols: locales.CurrencySymbols{Default: "SOS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SRD"}, Symbols: locales.CurrencySymbols{Default: "SRD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "SRG"}, Symbols: locales.CurrencySymbols{Default: "SRG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SSP"}, Symbols: locales.CurrencySymbols{Default: "SSP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "STD"}, Symbols: locales.CurrencySymbols{Default: "STD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "STN"}, Symbols: locales.CurrencySymbols{Default: "STN", Narrow: "Db"}}, {Names: map[locales.PluralRule]string{0: "SUR"}, Symbols: locales.CurrencySymbols{Default: "SUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SVC"}, Symbols: locales.CurrencySymbols{Default: "SVC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SYP"}, Symbols: locales.CurrencySymbols{Default: "SYP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "SZL"}, Symbols: locales.CurrencySymbols{Default: "SZL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "THB"}, Symbols: locales.CurrencySymbols{Default: "฿", Narrow: "฿"}}, {Names: map[locales.PluralRule]string{0: "TJR"}, Symbols: locales.CurrencySymbols{Default: "TJR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TJS"}, Symbols: locales.CurrencySymbols{Default: "TJS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TMM"}, Symbols: locales.CurrencySymbols{Default: "TMM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TMT"}, Symbols: locales.CurrencySymbols{Default: "TMT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TND"}, Symbols: locales.CurrencySymbols{Default: "TND", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TOP"}, Symbols: locales.CurrencySymbols{Default: "TOP", Narrow: "T$"}}, {Names: map[locales.PluralRule]string{0: "TPE"}, Symbols: locales.CurrencySymbols{Default: "TPE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TRL"}, Symbols: locales.CurrencySymbols{Default: "TRL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TRY"}, Symbols: locales.CurrencySymbols{Default: "TRY", Narrow: "₺"}}, {Names: map[locales.PluralRule]string{0: "TTD"}, Symbols: locales.CurrencySymbols{Default: "TTD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "TWD"}, Symbols: locales.CurrencySymbols{Default: "NT$", Narrow: "NT$"}}, {Names: map[locales.PluralRule]string{0: "TZS"}, Symbols: locales.CurrencySymbols{Default: "TZS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UAH"}, Symbols: locales.CurrencySymbols{Default: "UAH", Narrow: "₴"}}, {Names: map[locales.PluralRule]string{0: "UAK"}, Symbols: locales.CurrencySymbols{Default: "UAK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UGS"}, Symbols: locales.CurrencySymbols{Default: "UGS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UGX"}, Symbols: locales.CurrencySymbols{Default: "UGX", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "USD"}, Symbols: locales.CurrencySymbols{Default: "US$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "USN"}, Symbols: locales.CurrencySymbols{Default: "USN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "USS"}, Symbols: locales.CurrencySymbols{Default: "USS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UYI"}, Symbols: locales.CurrencySymbols{Default: "UYI", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UYP"}, Symbols: locales.CurrencySymbols{Default: "UYP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UYU"}, Symbols: locales.CurrencySymbols{Default: "UYU", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "UYW"}, Symbols: locales.CurrencySymbols{Default: "UYW", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UZS"}, Symbols: locales.CurrencySymbols{Default: "UZS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VEB"}, Symbols: locales.CurrencySymbols{Default: "VEB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VED"}, Symbols: locales.CurrencySymbols{Default: "VED", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VEF"}, Symbols: locales.CurrencySymbols{Default: "VEF", Narrow: "Bs"}}, {Names: map[locales.PluralRule]string{0: "VES"}, Symbols: locales.CurrencySymbols{Default: "VES", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VND"}, Symbols: locales.CurrencySymbols{Default: "₫", Narrow: "₫"}}, {Names: map[locales.PluralRule]string{0: "VNN"}, Symbols: locales.CurrencySymbols{Default: "VNN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VUV"}, Symbols: locales.CurrencySymbols{Default: "VUV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "WST"}, Symbols: locales.CurrencySymbols{Default: "WST", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XAF"}, Symbols: locales.CurrencySymbols{Default: "FCFA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XAG"}, Symbols: locales.CurrencySymbols{Default: "XAG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XAU"}, Symbols: locales.CurrencySymbols{Default: "XAU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBA"}, Symbols: locales.CurrencySymbols{Default: "XBA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBB"}, Symbols: locales.CurrencySymbols{Default: "XBB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBC"}, Symbols: locales.CurrencySymbols{Default: "XBC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBD"}, Symbols: locales.CurrencySymbols{Default: "XBD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XCD"}, Symbols: locales.CurrencySymbols{Default: "EC$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "XDR"}, Symbols: locales.CurrencySymbols{Default: "XDR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XEU"}, Symbols: locales.CurrencySymbols{Default: "XEU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XFO"}, Symbols: locales.CurrencySymbols{Default: "XFO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XFU"}, Symbols: locales.CurrencySymbols{Default: "XFU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XOF"}, Symbols: locales.CurrencySymbols{Default: "F\u202fCFA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XPD"}, Symbols: locales.CurrencySymbols{Default: "XPD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XPF"}, Symbols: locales.CurrencySymbols{Default: "CFPF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XPT"}, Symbols: locales.CurrencySymbols{Default: "XPT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XRE"}, Symbols: locales.CurrencySymbols{Default: "XRE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XSU"}, Symbols: locales.CurrencySymbols{Default: "XSU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XTS"}, Symbols: locales.CurrencySymbols{Default: "XTS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XUA"}, Symbols: locales.CurrencySymbols{Default: "XUA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XXX"}, Symbols: locales.CurrencySymbols{Default: "XXX", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YDD"}, Symbols: locales.CurrencySymbols{Default: "YDD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YER"}, Symbols: locales.CurrencySymbols{Default: "YER", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUD"}, Symbols: locales.CurrencySymbols{Default: "YUD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUM"}, Symbols: locales.CurrencySymbols{Default: "YUM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUN"}, Symbols: locales.CurrencySymbols{Default: "YUN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUR"}, Symbols: locales.CurrencySymbols{Default: "YUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZAL"}, Symbols: locales.CurrencySymbols{Default: "ZAL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZAR"}, Symbols: locales.CurrencySymbols{Default: "ZAR", Narrow: "ZAR"}}, {Names: map[locales.PluralRule]string{0: "ZMK"}, Symbols: locales.CurrencySymbols{Default: "ZMK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZMW"}, Symbols: locales.CurrencySymbols{Default: "ZMW", Narrow: "ZMW"}}, {Names: map[locales.PluralRule]string{0: "ZRN"}, Symbols: locales.CurrencySymbols{Default: "ZRN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZRZ"}, Symbols: locales.CurrencySymbols{Default: "ZRZ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZWD"}, Symbols: locales.CurrencySymbols{Default: "ZWD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZWL"}, Symbols: locales.CurrencySymbols{Default: "ZWL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZWR"}, Symbols: locales.CurrencySymbols{Default: "ZWR", Narrow: ""}}},
			Formatters: &locales.CurrencyFormatters{
				CurrencyFmt:   locales.MustParseNumberFormatPatterns("¤#,##0.00"),
				AccountingFmt: locales.MustParseNumberFormatPatterns("¤#,##0.00;(¤#,##0.00)"),
				Short: locales.CurrencyAccountingFormatterByExp{
					CurrencyFmt: map[uint8]*locales.CurrencyAccountingFormatterByExpPlural{
						3: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("¤0K"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤0K"),
						}},
						4: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("¤00K"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤00K"),
						}},
						5: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("¤000K"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤000K"),
						}},
						6: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("¤0M"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤0M"),
						}},
						7: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("¤00M"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤00M"),
						}},
						8: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("¤000M"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤000M"),
						}},
						9: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("¤0B"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤0B"),
						}},
						10: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("¤00B"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤00B"),
						}},
						11: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("¤000B"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤000B"),
						}},
						12: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("¤0T"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤0T"),
						}},
						13: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("¤00T"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤00T"),
						}},
						14: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOne:   locales.MustParseNumberFormatPatterns("¤000T"),
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤000T"),
						}},
					},
					AccountingFmt: map[uint8]*locales.CurrencyAccountingFormatterByExpPlural{},
				},
			},
		},
		TimeSpec: &locales.TimeSpec{
			Separator:                ":",
			MonthsAbbreviatedValue:   []string{"", "Ion", "Chwef", "Maw", "Ebr", "Mai", "Meh", "Gorff", "Awst", "Medi", "Hyd", "Tach", "Rhag"},
			MonthsNarrowValue:        []string{"", "I", "Ch", "M", "E", "M", "M", "G", "A", "M", "H", "T", "Rh"},
			MonthsWideValue:          []string{"", "Ionawr", "Chwefror", "Mawrth", "Ebrill", "Mai", "Mehefin", "Gorffennaf", "Awst", "Medi", "Hydref", "Tachwedd", "Rhagfyr"},
			WeekdaysAbbreviatedValue: []string{"Sul", "Llun", "Maw", "Mer", "Iau", "Gwen", "Sad"},
			WeekdaysNarrowValue:      []string{"S", "Ll", "M", "M", "I", "G", "S"},
			WeekdaysShortValue:       []string{"Su", "Ll", "Ma", "Me", "Ia", "Gw", "Sa"},
			WeekdaysWideValue:        []string{"Dydd Sul", "Dydd Llun", "Dydd Mawrth", "Dydd Mercher", "Dydd Iau", "Dydd Gwener", "Dydd Sadwrn"},
			PeriodsAbbreviatedValue:  []string{"", ""},
			PeriodsNarrowValue:       []string{"b", "h"},
			PeriodsWideValue:         []string{"yb", "yh"},
			ErasAbbreviatedValue:     []string{"CC", "OC"},
			ErasNarrowValue:          []string{"C", "O"},
			ErasWideValue:            []string{"Cyn Crist", "Oed Crist"},
			TimezonesValue:           map[string]string{"": "Amser Safonol Paraguay", "ACDT": "Amser Haf Canolbarth Awstralia", "ACST": "Amser Safonol Canolbarth Awstralia", "ACT": "ACT", "ACWDT": "Amser Haf Canolbarth Gorllewin Awstralia", "ACWST": "Amser Safonol Canolbarth Gorllewin Awstralia", "ADT": "Amser Haf Cefnfor yr Iwerydd", "AEDT": "Amser Haf Dwyrain Awstralia", "AEST": "Amser Safonol Dwyrain Awstralia", "AFT": "Amser Afghanistan", "AKDT": "Amser Haf Alaska", "AKST": "Amser Safonol Alaska", "AMST": "Amser Haf Amazonas", "AMT": "Amser Safonol Amazonas", "ARST": "Amser Haf Ariannin", "ART": "Amser Safonol Ariannin", "AST": "Amser Safonol Cefnfor yr Iwerydd", "AWDT": "Amser Haf Gorllewin Awstralia", "AWST": "Amser Safonol Gorllewin Awstralia", "BNT": "Amser Brunei Darussalam", "BOT": "Amser Bolifia", "BRST": "Amser Haf Brasília", "BRT": "Amser Safonol Brasília", "BST": "Amser Haf Bangladesh", "BTT": "Amser Bhutan", "CAT": "Amser Canolbarth Affrica", "CCT": "Amser Ynysoedd Cocos", "CDT": "Amser Haf Canolbarth Gogledd America", "CDT (Kuba)": "Amser Haf Ciwa", "CHADT": "Amser Haf Chatham", "CHAST": "Amser Safonol Chatham", "CLST": "Amser Haf Chile", "CLT": "Amser Safonol Chile", "COST": "Amser Haf Colombia", "COT": "Amser Safonol Colombia", "CST": "Amser Safonol Canolbarth Gogledd America", "CST (Kuba)": "Amser Safonol Ciwba", "CXT": "Amser Ynys Y Nadolig", "ChST": "Amser Chamorro", "EASST": "Amser Haf Ynys y Pasg", "EAST": "Amser Safonol Ynys y Pasg", "EAT": "Amser Dwyrain Affrica", "ECT": "Amser Ecuador", "EDT": "Amser Haf Dwyrain Gogledd America", "EGDT": "Amser Haf Dwyrain yr Ynys Las", "EGST": "Amser Safonol Dwyrain yr Ynys Las", "EST": "Amser Safonol Dwyrain Gogledd America", "FKST": "Amser Haf Ynysoedd Falklands/Malvinas", "GALT": "Amser Galapagos", "GFT": "Amser Guyane Ffrengig", "GMT": "Amser Safonol Greenwich", "GST": "Amser Safonol y Gwlff", "GYT": "Amser Guyana", "HADT": "Amser Haf Hawaii-Aleutian", "HAST": "Amser Safonol Hawaii-Aleutian", "HENOMX": "Amser Haf Gogledd Orllewin Mecsico", "HEOG": "Amser Haf Gorllewin yr Ynys Las", "HEPMX": "Amser Haf Pasiffig Mecsico", "HKST": "Amser Haf Hong Kong", "HKT": "Amser Safonol Hong Kong", "HNNOMX": "Amser Safonol Gogledd Orllewin Mecsico", "HNOG": "Amser Safonol Gorllewin yr Ynys Las", "HNPMX": "Amser Safonol Pasiffig Mecsico", "ICT": "Amser Indo-Tsieina", "IRDT": "Amser Haf Iran", "IRST": "Amser Safonol Iran", "IST": "Amser India", "JDT": "Amser Haf Japan", "JST": "Amser Safonol Japan", "LHDT": "Amser Haf yr Arglwydd Howe", "LHST": "Amser Safonol yr Arglwydd Howe", "MDT": "Amser Haf Mynyddoedd Gogledd America", "MESZ": "Amser Haf Canolbarth Ewrop", "MEZ": "Amser Safonol Canolbarth Ewrop", "MST": "Amser Safonol Mynyddoedd Gogledd America", "MVT": "Amser Y Maldives", "MYT": "Amser Malaysia", "NDT": "Amser Haf Newfoundland", "NPT": "Amser Nepal", "NST": "Amser Safonol Newfoundland", "NZDT": "Amser Haf Seland Newydd", "NZST": "Amser Safonol Seland Newydd", "OESZ": "Amser Haf Dwyrain Ewrop", "OEZ": "Amser Safonol Dwyrain Ewrop", "PDT": "Amser Haf Cefnfor Tawel Gogledd America", "PKT": "Amser Haf Pakistan", "PMDT": "Amser Haf Saint-Pierre-et-Miquelon", "PMST": "Amser Safonol Saint-Pierre-et-Miquelon", "PST": "Amser Safonol Cefnfor Tawel Gogledd America", "PYST": "Amser Haf Paraguay", "SAST": "Amser Safonol De Affrica", "SGT": "Amser Singapore", "SRT": "Amser Suriname", "TLT": "Amser Dwyrain Timor", "TMST": "Amser Haf Tyrcmenistan", "TMT": "Amser Safonol Tyrcmenistan", "UYST": "Amser Haf Uruguay", "UYT": "Amser Safonol Uruguay", "VET": "Amser Venezuela", "WARST": "Amser Haf Gorllewin Ariannin", "WART": "Amser Safonol Gorllewin Ariannin", "WAST": "Amser Haf Gorllewin Affrica", "WAT": "Amser Safonol Gorllewin Affrica", "WESZ": "Amser Haf Gorllewin Ewrop", "WEZ": "Amser Safonol Gorllewin Ewrop", "WIB": "Amser Gorllewin Indonesia", "WIT": "Amser Dwyrain Indonesia", "WITA": "Amser Canolbarth Indonesia", "∅∅∅": "Amser Haf yr Azores"},
		},

		locale:          "cy_GB",
		pluralsCardinal: []locales.PluralRule{1, 2, 3, 4, 5, 6},
		pluralsOrdinal:  []locales.PluralRule{1, 2, 3, 4, 5, 6},
		pluralsRange:    []locales.PluralRule{2, 3, 4, 5, 6},
		percent:         "%",
		perMille:        "‰",
		inifinity:       "∞",

		dateFullLayout:   "EEEE, d MMMM y",
		dateLongLayout:   "d MMMM y",
		dateMediumLayout: "d MMM y",
		dateShortLayout:  "dd/MM/yy",
		timeFullLayout:   "HH:mm:ss zzzz",
		timeLongLayout:   "HH:mm:ss z",
		timeMediumLayout: "HH:mm:ss",
		timeShortLayout:  "HH:mm",
		listPatterns:     locales.NewListPatternsFromSlice([][]string{{"{0} a(c) {1}", "{0}, {1}", "{0}, {1}", "{0}, a(c) {1}"}, {"{0} neu {1}", "{0}, {1}", "{0}, {1}", "{0} neu {1}"}}),
		miscPatterns: locales.MiscPatterns{
			"~%[1]v",
			"%[1]v+",
			"≤%[1]v",
			"%[1]v–%[2]v",
		},
	}
}

// Locale returns the current translators string locale
func (cy *cy_GB) Locale() string {
	return cy.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'cy_GB'
func (cy *cy_GB) PluralsCardinal() []locales.PluralRule {
	return cy.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'cy_GB'
func (cy *cy_GB) PluralsOrdinal() []locales.PluralRule {
	return cy.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'cy_GB'
func (cy *cy_GB) PluralsRange() []locales.PluralRule {
	return cy.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'cy_GB'
func (cy *cy_GB) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 0 {
		return locales.PluralRuleZero
	} else if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n == 3 {
		return locales.PluralRuleFew
	} else if n == 6 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'cy_GB'
func (cy *cy_GB) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 0 || n == 7 || n == 8 || n == 9 {
		return locales.PluralRuleZero
	} else if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n == 3 || n == 4 {
		return locales.PluralRuleFew
	} else if n == 5 || n == 6 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'cy_GB'
func (cy *cy_GB) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := cy.CardinalPluralRule(num1, v1)
	end := cy.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleZero && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleTwo {
		return locales.PluralRuleTwo
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleTwo {
		return locales.PluralRuleTwo
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleTwo {
		return locales.PluralRuleTwo
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther

}

// FmtDateShort returns the short date representation of 't' for 'cy_GB'
func (cy *cy_GB) FmtDateShort(t time.Time) string {
	return locales.FormatTimeEra(cy, cy.dateShortLayout, t, 2)
}

// FmtDateMedium returns the medium date representation of 't' for 'cy_GB'
func (cy *cy_GB) FmtDateMedium(t time.Time) string {
	return locales.FormatTimeEra(cy, cy.dateMediumLayout, t, 2)
}

// FmtDateLong returns the long date representation of 't' for 'cy_GB'
func (cy *cy_GB) FmtDateLong(t time.Time) string {
	return locales.FormatTimeEra(cy, cy.dateLongLayout, t, 1)
}

// FmtDateFull returns the full date representation of 't' for 'cy_GB'
func (cy *cy_GB) FmtDateFull(t time.Time) string {
	return locales.FormatTimeEra(cy, cy.dateFullLayout, t, 0)
}

// FmtTimeShort returns the short time representation of 't' for 'cy_GB'
func (cy *cy_GB) FmtTimeShort(t time.Time) string {
	return locales.FormatTimeEra(cy, cy.timeShortLayout, t, 2)
}

// FmtTimeMedium returns the medium time representation of 't' for 'cy_GB'
func (cy *cy_GB) FmtTimeMedium(t time.Time) string {
	return locales.FormatTimeEra(cy, cy.timeMediumLayout, t, 2)
}

// FmtTimeLong returns the long time representation of 't' for 'cy_GB'
func (cy *cy_GB) FmtTimeLong(t time.Time) string {
	return locales.FormatTimeEra(cy, cy.timeLongLayout, t, 1)
}

// FmtTimeFull returns the full time representation of 't' for 'cy_GB'
func (cy *cy_GB) FmtTimeFull(t time.Time) string {
	return locales.FormatTimeEra(cy, cy.timeFullLayout, t, 0)
}

// DateFullLayout returns the full date layout for 'cy_GB'
func (cy *cy_GB) DateFullLayout() string {
	return cy.dateFullLayout
}

// DateLongLayout returns the long date layout for 'cy_GB'
func (cy *cy_GB) DateLongLayout() string {
	return cy.dateLongLayout
}

// DateMediumLayout returns the medium date layout for 'cy_GB'
func (cy *cy_GB) DateMediumLayout() string {
	return cy.dateMediumLayout
}

// DateShortLayout returns the short date layout for 'cy_GB'
func (cy *cy_GB) DateShortLayout() string {
	return cy.dateShortLayout
}

// TimeFullLayout returns the full time layout for 'cy_GB'
func (cy *cy_GB) TimeFullLayout() string {
	return cy.timeFullLayout
}

// TimeLongLayout returns the full long layout for 'cy_GB'
func (cy *cy_GB) TimeLongLayout() string {
	return cy.timeLongLayout
}

// TimeMediumLayout returns the medium time layout for 'cy_GB'
func (cy *cy_GB) TimeMediumLayout() string {
	return cy.timeMediumLayout
}

// TimeShortLayout returns the short time layout for 'cy_GB'
func (cy *cy_GB) TimeShortLayout() string {
	return cy.timeShortLayout
}

func (cy *cy_GB) ListPatterns() *locales.ListPatterns {
	return cy.listPatterns
}

func (cy *cy_GB) GetDurationSpec() *locales.DurationSpec {
	return cy.DurationSpec
}

func (cy *cy_GB) GetMiscPatterns() *locales.MiscPatterns {
	return &cy.miscPatterns
}
