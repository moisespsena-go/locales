package lo_LA

import (
	"math"
	"time"

	"github.com/moisespsena-go/locales"
)

type lo_LA struct {
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

// New returns a new instance of translator for the 'lo_LA' locale
func New() locales.Translator {
	return &lo_LA{
		NumberSpec: &locales.NumberSpec{
			GroupValue:   ".",
			DecimalValue: ",",
		},
		CurrencySpec: &locales.CurrencySpec{
			CurrenciesValue: []locales.Currency{{Names: map[locales.PluralRule]string{0: "ADP"}, Symbols: locales.CurrencySymbols{Default: "ADP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AED"}, Symbols: locales.CurrencySymbols{Default: "AED", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AFA"}, Symbols: locales.CurrencySymbols{Default: "AFA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AFN"}, Symbols: locales.CurrencySymbols{Default: "AFN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ALK"}, Symbols: locales.CurrencySymbols{Default: "ALK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ALL"}, Symbols: locales.CurrencySymbols{Default: "ALL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AMD"}, Symbols: locales.CurrencySymbols{Default: "AMD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ANG"}, Symbols: locales.CurrencySymbols{Default: "ANG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AOA"}, Symbols: locales.CurrencySymbols{Default: "AOA", Narrow: "Kz"}}, {Names: map[locales.PluralRule]string{0: "AOK"}, Symbols: locales.CurrencySymbols{Default: "AOK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AON"}, Symbols: locales.CurrencySymbols{Default: "AON", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AOR"}, Symbols: locales.CurrencySymbols{Default: "AOR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARA"}, Symbols: locales.CurrencySymbols{Default: "ARA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARL"}, Symbols: locales.CurrencySymbols{Default: "ARL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARM"}, Symbols: locales.CurrencySymbols{Default: "ARM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARP"}, Symbols: locales.CurrencySymbols{Default: "ARP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARS"}, Symbols: locales.CurrencySymbols{Default: "ARS", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "ATS"}, Symbols: locales.CurrencySymbols{Default: "ATS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AUD"}, Symbols: locales.CurrencySymbols{Default: "A$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "AWG"}, Symbols: locales.CurrencySymbols{Default: "AWG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AZM"}, Symbols: locales.CurrencySymbols{Default: "AZM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AZN"}, Symbols: locales.CurrencySymbols{Default: "AZN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BAD"}, Symbols: locales.CurrencySymbols{Default: "BAD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BAM"}, Symbols: locales.CurrencySymbols{Default: "BAM", Narrow: "KM"}}, {Names: map[locales.PluralRule]string{0: "BAN"}, Symbols: locales.CurrencySymbols{Default: "BAN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BBD"}, Symbols: locales.CurrencySymbols{Default: "BBD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "BDT"}, Symbols: locales.CurrencySymbols{Default: "BDT", Narrow: "৳"}}, {Names: map[locales.PluralRule]string{0: "BEC"}, Symbols: locales.CurrencySymbols{Default: "BEC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BEF"}, Symbols: locales.CurrencySymbols{Default: "BEF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BEL"}, Symbols: locales.CurrencySymbols{Default: "BEL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGL"}, Symbols: locales.CurrencySymbols{Default: "BGL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGM"}, Symbols: locales.CurrencySymbols{Default: "BGM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGN"}, Symbols: locales.CurrencySymbols{Default: "BGN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGO"}, Symbols: locales.CurrencySymbols{Default: "BGO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BHD"}, Symbols: locales.CurrencySymbols{Default: "BHD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BIF"}, Symbols: locales.CurrencySymbols{Default: "BIF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BMD"}, Symbols: locales.CurrencySymbols{Default: "BMD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "BND"}, Symbols: locales.CurrencySymbols{Default: "BND", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "BOB"}, Symbols: locales.CurrencySymbols{Default: "BOB", Narrow: "Bs"}}, {Names: map[locales.PluralRule]string{0: "BOL"}, Symbols: locales.CurrencySymbols{Default: "BOL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BOP"}, Symbols: locales.CurrencySymbols{Default: "BOP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BOV"}, Symbols: locales.CurrencySymbols{Default: "BOV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRB"}, Symbols: locales.CurrencySymbols{Default: "BRB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRC"}, Symbols: locales.CurrencySymbols{Default: "BRC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRE"}, Symbols: locales.CurrencySymbols{Default: "BRE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRL"}, Symbols: locales.CurrencySymbols{Default: "R$", Narrow: "R$"}}, {Names: map[locales.PluralRule]string{0: "BRN"}, Symbols: locales.CurrencySymbols{Default: "BRN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRR"}, Symbols: locales.CurrencySymbols{Default: "BRR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRZ"}, Symbols: locales.CurrencySymbols{Default: "BRZ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BSD"}, Symbols: locales.CurrencySymbols{Default: "BSD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "BTN"}, Symbols: locales.CurrencySymbols{Default: "BTN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BUK"}, Symbols: locales.CurrencySymbols{Default: "BUK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BWP"}, Symbols: locales.CurrencySymbols{Default: "BWP", Narrow: "P"}}, {Names: map[locales.PluralRule]string{0: "BYB"}, Symbols: locales.CurrencySymbols{Default: "BYB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BYN"}, Symbols: locales.CurrencySymbols{Default: "BYN", Narrow: "р."}}, {Names: map[locales.PluralRule]string{0: "BYR"}, Symbols: locales.CurrencySymbols{Default: "BYR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BZD"}, Symbols: locales.CurrencySymbols{Default: "BZD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CAD"}, Symbols: locales.CurrencySymbols{Default: "CA$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CDF"}, Symbols: locales.CurrencySymbols{Default: "CDF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CHE"}, Symbols: locales.CurrencySymbols{Default: "CHE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CHF"}, Symbols: locales.CurrencySymbols{Default: "CHF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CHW"}, Symbols: locales.CurrencySymbols{Default: "CHW", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CLE"}, Symbols: locales.CurrencySymbols{Default: "CLE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CLF"}, Symbols: locales.CurrencySymbols{Default: "CLF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CLP"}, Symbols: locales.CurrencySymbols{Default: "CLP", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CNH"}, Symbols: locales.CurrencySymbols{Default: "CNH", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CNX"}, Symbols: locales.CurrencySymbols{Default: "CNX", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CNY"}, Symbols: locales.CurrencySymbols{Default: "CN¥", Narrow: "¥"}}, {Names: map[locales.PluralRule]string{0: "COP"}, Symbols: locales.CurrencySymbols{Default: "COP", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "COU"}, Symbols: locales.CurrencySymbols{Default: "COU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CRC"}, Symbols: locales.CurrencySymbols{Default: "CRC", Narrow: "₡"}}, {Names: map[locales.PluralRule]string{0: "CSD"}, Symbols: locales.CurrencySymbols{Default: "CSD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CSK"}, Symbols: locales.CurrencySymbols{Default: "CSK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CUC"}, Symbols: locales.CurrencySymbols{Default: "CUC", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CUP"}, Symbols: locales.CurrencySymbols{Default: "CUP", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CVE"}, Symbols: locales.CurrencySymbols{Default: "CVE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CYP"}, Symbols: locales.CurrencySymbols{Default: "CYP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CZK"}, Symbols: locales.CurrencySymbols{Default: "CZK", Narrow: "Kč"}}, {Names: map[locales.PluralRule]string{0: "DDM"}, Symbols: locales.CurrencySymbols{Default: "DDM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DEM"}, Symbols: locales.CurrencySymbols{Default: "DEM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DJF"}, Symbols: locales.CurrencySymbols{Default: "DJF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DKK"}, Symbols: locales.CurrencySymbols{Default: "DKK", Narrow: "kr"}}, {Names: map[locales.PluralRule]string{0: "DOP"}, Symbols: locales.CurrencySymbols{Default: "DOP", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "DZD"}, Symbols: locales.CurrencySymbols{Default: "DZD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ECS"}, Symbols: locales.CurrencySymbols{Default: "ECS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ECV"}, Symbols: locales.CurrencySymbols{Default: "ECV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "EEK"}, Symbols: locales.CurrencySymbols{Default: "EEK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "EGP"}, Symbols: locales.CurrencySymbols{Default: "EGP", Narrow: "E£"}}, {Names: map[locales.PluralRule]string{0: "ERN"}, Symbols: locales.CurrencySymbols{Default: "ERN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ESA"}, Symbols: locales.CurrencySymbols{Default: "ESA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ESB"}, Symbols: locales.CurrencySymbols{Default: "ESB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ESP"}, Symbols: locales.CurrencySymbols{Default: "ESP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ETB"}, Symbols: locales.CurrencySymbols{Default: "ETB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "EUR"}, Symbols: locales.CurrencySymbols{Default: "€", Narrow: "€"}}, {Names: map[locales.PluralRule]string{0: "FIM"}, Symbols: locales.CurrencySymbols{Default: "FIM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "FJD"}, Symbols: locales.CurrencySymbols{Default: "FJD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "FKP"}, Symbols: locales.CurrencySymbols{Default: "FKP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "FRF"}, Symbols: locales.CurrencySymbols{Default: "FRF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GBP"}, Symbols: locales.CurrencySymbols{Default: "£", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "GEK"}, Symbols: locales.CurrencySymbols{Default: "GEK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GEL"}, Symbols: locales.CurrencySymbols{Default: "GEL", Narrow: "₾"}}, {Names: map[locales.PluralRule]string{0: "GHC"}, Symbols: locales.CurrencySymbols{Default: "GHC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GHS"}, Symbols: locales.CurrencySymbols{Default: "GHS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GIP"}, Symbols: locales.CurrencySymbols{Default: "GIP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "GMD"}, Symbols: locales.CurrencySymbols{Default: "GMD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GNF"}, Symbols: locales.CurrencySymbols{Default: "GNF", Narrow: "FG"}}, {Names: map[locales.PluralRule]string{0: "GNS"}, Symbols: locales.CurrencySymbols{Default: "GNS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GQE"}, Symbols: locales.CurrencySymbols{Default: "GQE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GRD"}, Symbols: locales.CurrencySymbols{Default: "GRD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GTQ"}, Symbols: locales.CurrencySymbols{Default: "GTQ", Narrow: "Q"}}, {Names: map[locales.PluralRule]string{0: "GWE"}, Symbols: locales.CurrencySymbols{Default: "GWE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GWP"}, Symbols: locales.CurrencySymbols{Default: "GWP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GYD"}, Symbols: locales.CurrencySymbols{Default: "GYD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "HKD"}, Symbols: locales.CurrencySymbols{Default: "HK$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "HNL"}, Symbols: locales.CurrencySymbols{Default: "HNL", Narrow: "L"}}, {Names: map[locales.PluralRule]string{0: "HRD"}, Symbols: locales.CurrencySymbols{Default: "HRD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "HRK"}, Symbols: locales.CurrencySymbols{Default: "HRK", Narrow: "kn"}}, {Names: map[locales.PluralRule]string{0: "HTG"}, Symbols: locales.CurrencySymbols{Default: "HTG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "HUF"}, Symbols: locales.CurrencySymbols{Default: "HUF", Narrow: "Ft"}}, {Names: map[locales.PluralRule]string{0: "IDR"}, Symbols: locales.CurrencySymbols{Default: "IDR", Narrow: "Rp"}}, {Names: map[locales.PluralRule]string{0: "IEP"}, Symbols: locales.CurrencySymbols{Default: "IEP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ILP"}, Symbols: locales.CurrencySymbols{Default: "ILP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ILR"}, Symbols: locales.CurrencySymbols{Default: "ILR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ILS"}, Symbols: locales.CurrencySymbols{Default: "₪", Narrow: "₪"}}, {Names: map[locales.PluralRule]string{0: "INR"}, Symbols: locales.CurrencySymbols{Default: "₹", Narrow: "₹"}}, {Names: map[locales.PluralRule]string{0: "IQD"}, Symbols: locales.CurrencySymbols{Default: "IQD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "IRR"}, Symbols: locales.CurrencySymbols{Default: "IRR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ISJ"}, Symbols: locales.CurrencySymbols{Default: "ISJ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ISK"}, Symbols: locales.CurrencySymbols{Default: "ISK", Narrow: "kr"}}, {Names: map[locales.PluralRule]string{0: "ITL"}, Symbols: locales.CurrencySymbols{Default: "ITL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "JMD"}, Symbols: locales.CurrencySymbols{Default: "JMD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "JOD"}, Symbols: locales.CurrencySymbols{Default: "JOD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "JPY"}, Symbols: locales.CurrencySymbols{Default: "JP¥", Narrow: "¥"}}, {Names: map[locales.PluralRule]string{0: "KES"}, Symbols: locales.CurrencySymbols{Default: "KES", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KGS"}, Symbols: locales.CurrencySymbols{Default: "KGS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KHR"}, Symbols: locales.CurrencySymbols{Default: "KHR", Narrow: "៛"}}, {Names: map[locales.PluralRule]string{0: "KMF"}, Symbols: locales.CurrencySymbols{Default: "KMF", Narrow: "CF"}}, {Names: map[locales.PluralRule]string{0: "KPW"}, Symbols: locales.CurrencySymbols{Default: "KPW", Narrow: "₩"}}, {Names: map[locales.PluralRule]string{0: "KRH"}, Symbols: locales.CurrencySymbols{Default: "KRH", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KRO"}, Symbols: locales.CurrencySymbols{Default: "KRO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KRW"}, Symbols: locales.CurrencySymbols{Default: "₩", Narrow: "₩"}}, {Names: map[locales.PluralRule]string{0: "KWD"}, Symbols: locales.CurrencySymbols{Default: "KWD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KYD"}, Symbols: locales.CurrencySymbols{Default: "KYD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "KZT"}, Symbols: locales.CurrencySymbols{Default: "KZT", Narrow: "₸"}}, {Names: map[locales.PluralRule]string{0: "LAK"}, Symbols: locales.CurrencySymbols{Default: "₭", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LBP"}, Symbols: locales.CurrencySymbols{Default: "LBP", Narrow: "L£"}}, {Names: map[locales.PluralRule]string{0: "LKR"}, Symbols: locales.CurrencySymbols{Default: "LKR", Narrow: "Rs"}}, {Names: map[locales.PluralRule]string{0: "LRD"}, Symbols: locales.CurrencySymbols{Default: "LRD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "LSL"}, Symbols: locales.CurrencySymbols{Default: "LSL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LTL"}, Symbols: locales.CurrencySymbols{Default: "LTL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LTT"}, Symbols: locales.CurrencySymbols{Default: "LTT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LUC"}, Symbols: locales.CurrencySymbols{Default: "LUC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LUF"}, Symbols: locales.CurrencySymbols{Default: "LUF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LUL"}, Symbols: locales.CurrencySymbols{Default: "LUL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LVL"}, Symbols: locales.CurrencySymbols{Default: "LVL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LVR"}, Symbols: locales.CurrencySymbols{Default: "LVR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LYD"}, Symbols: locales.CurrencySymbols{Default: "LYD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MAD"}, Symbols: locales.CurrencySymbols{Default: "MAD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MAF"}, Symbols: locales.CurrencySymbols{Default: "MAF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MCF"}, Symbols: locales.CurrencySymbols{Default: "MCF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MDC"}, Symbols: locales.CurrencySymbols{Default: "MDC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MDL"}, Symbols: locales.CurrencySymbols{Default: "MDL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MGA"}, Symbols: locales.CurrencySymbols{Default: "MGA", Narrow: "Ar"}}, {Names: map[locales.PluralRule]string{0: "MGF"}, Symbols: locales.CurrencySymbols{Default: "MGF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MKD"}, Symbols: locales.CurrencySymbols{Default: "MKD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MKN"}, Symbols: locales.CurrencySymbols{Default: "MKN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MLF"}, Symbols: locales.CurrencySymbols{Default: "MLF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MMK"}, Symbols: locales.CurrencySymbols{Default: "MMK", Narrow: "K"}}, {Names: map[locales.PluralRule]string{0: "MNT"}, Symbols: locales.CurrencySymbols{Default: "MNT", Narrow: "₮"}}, {Names: map[locales.PluralRule]string{0: "MOP"}, Symbols: locales.CurrencySymbols{Default: "MOP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MRO"}, Symbols: locales.CurrencySymbols{Default: "MRO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MRU"}, Symbols: locales.CurrencySymbols{Default: "MRU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MTL"}, Symbols: locales.CurrencySymbols{Default: "MTL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MTP"}, Symbols: locales.CurrencySymbols{Default: "MTP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MUR"}, Symbols: locales.CurrencySymbols{Default: "MUR", Narrow: "Rs"}}, {Names: map[locales.PluralRule]string{0: "MVP"}, Symbols: locales.CurrencySymbols{Default: "MVP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MVR"}, Symbols: locales.CurrencySymbols{Default: "MVR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MWK"}, Symbols: locales.CurrencySymbols{Default: "MWK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MXN"}, Symbols: locales.CurrencySymbols{Default: "MX$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "MXP"}, Symbols: locales.CurrencySymbols{Default: "MXP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MXV"}, Symbols: locales.CurrencySymbols{Default: "MXV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MYR"}, Symbols: locales.CurrencySymbols{Default: "MYR", Narrow: "RM"}}, {Names: map[locales.PluralRule]string{0: "MZE"}, Symbols: locales.CurrencySymbols{Default: "MZE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MZM"}, Symbols: locales.CurrencySymbols{Default: "MZM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MZN"}, Symbols: locales.CurrencySymbols{Default: "MZN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NAD"}, Symbols: locales.CurrencySymbols{Default: "NAD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "NGN"}, Symbols: locales.CurrencySymbols{Default: "NGN", Narrow: "₦"}}, {Names: map[locales.PluralRule]string{0: "NIC"}, Symbols: locales.CurrencySymbols{Default: "NIC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NIO"}, Symbols: locales.CurrencySymbols{Default: "NIO", Narrow: "C$"}}, {Names: map[locales.PluralRule]string{0: "NLG"}, Symbols: locales.CurrencySymbols{Default: "NLG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NOK"}, Symbols: locales.CurrencySymbols{Default: "NOK", Narrow: "kr"}}, {Names: map[locales.PluralRule]string{0: "NPR"}, Symbols: locales.CurrencySymbols{Default: "NPR", Narrow: "Rs"}}, {Names: map[locales.PluralRule]string{0: "NZD"}, Symbols: locales.CurrencySymbols{Default: "NZ$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "OMR"}, Symbols: locales.CurrencySymbols{Default: "OMR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PAB"}, Symbols: locales.CurrencySymbols{Default: "PAB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PEI"}, Symbols: locales.CurrencySymbols{Default: "PEI", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PEN"}, Symbols: locales.CurrencySymbols{Default: "PEN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PES"}, Symbols: locales.CurrencySymbols{Default: "PES", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PGK"}, Symbols: locales.CurrencySymbols{Default: "PGK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PHP"}, Symbols: locales.CurrencySymbols{Default: "PHP", Narrow: "₱"}}, {Names: map[locales.PluralRule]string{0: "PKR"}, Symbols: locales.CurrencySymbols{Default: "PKR", Narrow: "Rs"}}, {Names: map[locales.PluralRule]string{0: "PLN"}, Symbols: locales.CurrencySymbols{Default: "PLN", Narrow: "zł"}}, {Names: map[locales.PluralRule]string{0: "PLZ"}, Symbols: locales.CurrencySymbols{Default: "PLZ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PTE"}, Symbols: locales.CurrencySymbols{Default: "PTE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PYG"}, Symbols: locales.CurrencySymbols{Default: "PYG", Narrow: "₲"}}, {Names: map[locales.PluralRule]string{0: "QAR"}, Symbols: locales.CurrencySymbols{Default: "QAR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RHD"}, Symbols: locales.CurrencySymbols{Default: "RHD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ROL"}, Symbols: locales.CurrencySymbols{Default: "ROL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RON"}, Symbols: locales.CurrencySymbols{Default: "RON", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RSD"}, Symbols: locales.CurrencySymbols{Default: "RSD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RUB"}, Symbols: locales.CurrencySymbols{Default: "RUB", Narrow: "₽"}}, {Names: map[locales.PluralRule]string{0: "RUR"}, Symbols: locales.CurrencySymbols{Default: "RUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RWF"}, Symbols: locales.CurrencySymbols{Default: "RWF", Narrow: "RF"}}, {Names: map[locales.PluralRule]string{0: "SAR"}, Symbols: locales.CurrencySymbols{Default: "SAR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SBD"}, Symbols: locales.CurrencySymbols{Default: "SBD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "SCR"}, Symbols: locales.CurrencySymbols{Default: "SCR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SDD"}, Symbols: locales.CurrencySymbols{Default: "SDD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SDG"}, Symbols: locales.CurrencySymbols{Default: "SDG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SDP"}, Symbols: locales.CurrencySymbols{Default: "SDP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SEK"}, Symbols: locales.CurrencySymbols{Default: "SEK", Narrow: "kr"}}, {Names: map[locales.PluralRule]string{0: "SGD"}, Symbols: locales.CurrencySymbols{Default: "SGD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "SHP"}, Symbols: locales.CurrencySymbols{Default: "SHP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "SIT"}, Symbols: locales.CurrencySymbols{Default: "SIT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SKK"}, Symbols: locales.CurrencySymbols{Default: "SKK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SLE"}, Symbols: locales.CurrencySymbols{Default: "SLE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SLL"}, Symbols: locales.CurrencySymbols{Default: "SLL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SOS"}, Symbols: locales.CurrencySymbols{Default: "SOS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SRD"}, Symbols: locales.CurrencySymbols{Default: "SRD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "SRG"}, Symbols: locales.CurrencySymbols{Default: "SRG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SSP"}, Symbols: locales.CurrencySymbols{Default: "SSP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "STD"}, Symbols: locales.CurrencySymbols{Default: "STD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "STN"}, Symbols: locales.CurrencySymbols{Default: "STN", Narrow: "Db"}}, {Names: map[locales.PluralRule]string{0: "SUR"}, Symbols: locales.CurrencySymbols{Default: "SUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SVC"}, Symbols: locales.CurrencySymbols{Default: "SVC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SYP"}, Symbols: locales.CurrencySymbols{Default: "SYP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "SZL"}, Symbols: locales.CurrencySymbols{Default: "SZL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "THB"}, Symbols: locales.CurrencySymbols{Default: "฿", Narrow: "฿"}}, {Names: map[locales.PluralRule]string{0: "TJR"}, Symbols: locales.CurrencySymbols{Default: "TJR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TJS"}, Symbols: locales.CurrencySymbols{Default: "TJS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TMM"}, Symbols: locales.CurrencySymbols{Default: "TMM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TMT"}, Symbols: locales.CurrencySymbols{Default: "TMT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TND"}, Symbols: locales.CurrencySymbols{Default: "TND", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TOP"}, Symbols: locales.CurrencySymbols{Default: "TOP", Narrow: "T$"}}, {Names: map[locales.PluralRule]string{0: "TPE"}, Symbols: locales.CurrencySymbols{Default: "TPE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TRL"}, Symbols: locales.CurrencySymbols{Default: "TRL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TRY"}, Symbols: locales.CurrencySymbols{Default: "TRY", Narrow: "₺"}}, {Names: map[locales.PluralRule]string{0: "TTD"}, Symbols: locales.CurrencySymbols{Default: "TTD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "TWD"}, Symbols: locales.CurrencySymbols{Default: "NT$", Narrow: "NT$"}}, {Names: map[locales.PluralRule]string{0: "TZS"}, Symbols: locales.CurrencySymbols{Default: "TZS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UAH"}, Symbols: locales.CurrencySymbols{Default: "UAH", Narrow: "₴"}}, {Names: map[locales.PluralRule]string{0: "UAK"}, Symbols: locales.CurrencySymbols{Default: "UAK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UGS"}, Symbols: locales.CurrencySymbols{Default: "UGS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UGX"}, Symbols: locales.CurrencySymbols{Default: "UGX", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "USD"}, Symbols: locales.CurrencySymbols{Default: "US$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "USN"}, Symbols: locales.CurrencySymbols{Default: "USN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "USS"}, Symbols: locales.CurrencySymbols{Default: "USS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UYI"}, Symbols: locales.CurrencySymbols{Default: "UYI", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UYP"}, Symbols: locales.CurrencySymbols{Default: "UYP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UYU"}, Symbols: locales.CurrencySymbols{Default: "UYU", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "UYW"}, Symbols: locales.CurrencySymbols{Default: "UYW", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UZS"}, Symbols: locales.CurrencySymbols{Default: "UZS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VEB"}, Symbols: locales.CurrencySymbols{Default: "VEB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VED"}, Symbols: locales.CurrencySymbols{Default: "VED", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VEF"}, Symbols: locales.CurrencySymbols{Default: "VEF", Narrow: "Bs"}}, {Names: map[locales.PluralRule]string{0: "VES"}, Symbols: locales.CurrencySymbols{Default: "VES", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VND"}, Symbols: locales.CurrencySymbols{Default: "₫", Narrow: "₫"}}, {Names: map[locales.PluralRule]string{0: "VNN"}, Symbols: locales.CurrencySymbols{Default: "VNN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VUV"}, Symbols: locales.CurrencySymbols{Default: "VUV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "WST"}, Symbols: locales.CurrencySymbols{Default: "WST", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XAF"}, Symbols: locales.CurrencySymbols{Default: "FCFA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XAG"}, Symbols: locales.CurrencySymbols{Default: "XAG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XAU"}, Symbols: locales.CurrencySymbols{Default: "XAU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBA"}, Symbols: locales.CurrencySymbols{Default: "XBA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBB"}, Symbols: locales.CurrencySymbols{Default: "XBB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBC"}, Symbols: locales.CurrencySymbols{Default: "XBC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBD"}, Symbols: locales.CurrencySymbols{Default: "XBD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XCD"}, Symbols: locales.CurrencySymbols{Default: "EC$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "XDR"}, Symbols: locales.CurrencySymbols{Default: "XDR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XEU"}, Symbols: locales.CurrencySymbols{Default: "XEU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XFO"}, Symbols: locales.CurrencySymbols{Default: "XFO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XFU"}, Symbols: locales.CurrencySymbols{Default: "XFU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XOF"}, Symbols: locales.CurrencySymbols{Default: "F\u202fCFA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XPD"}, Symbols: locales.CurrencySymbols{Default: "XPD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XPF"}, Symbols: locales.CurrencySymbols{Default: "CFPF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XPT"}, Symbols: locales.CurrencySymbols{Default: "XPT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XRE"}, Symbols: locales.CurrencySymbols{Default: "XRE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XSU"}, Symbols: locales.CurrencySymbols{Default: "XSU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XTS"}, Symbols: locales.CurrencySymbols{Default: "XTS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XUA"}, Symbols: locales.CurrencySymbols{Default: "XUA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XXX"}, Symbols: locales.CurrencySymbols{Default: "XXX", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YDD"}, Symbols: locales.CurrencySymbols{Default: "YDD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YER"}, Symbols: locales.CurrencySymbols{Default: "YER", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUD"}, Symbols: locales.CurrencySymbols{Default: "YUD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUM"}, Symbols: locales.CurrencySymbols{Default: "YUM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUN"}, Symbols: locales.CurrencySymbols{Default: "YUN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUR"}, Symbols: locales.CurrencySymbols{Default: "YUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZAL"}, Symbols: locales.CurrencySymbols{Default: "ZAL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZAR"}, Symbols: locales.CurrencySymbols{Default: "ZAR", Narrow: "R"}}, {Names: map[locales.PluralRule]string{0: "ZMK"}, Symbols: locales.CurrencySymbols{Default: "ZMK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZMW"}, Symbols: locales.CurrencySymbols{Default: "ZMW", Narrow: "ZK"}}, {Names: map[locales.PluralRule]string{0: "ZRN"}, Symbols: locales.CurrencySymbols{Default: "ZRN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZRZ"}, Symbols: locales.CurrencySymbols{Default: "ZRZ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZWD"}, Symbols: locales.CurrencySymbols{Default: "ZWD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZWL"}, Symbols: locales.CurrencySymbols{Default: "ZWL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZWR"}, Symbols: locales.CurrencySymbols{Default: "ZWR", Narrow: ""}}},
			Formatters: &locales.CurrencyFormatters{
				CurrencyFmt:   locales.MustParseNumberFormatPatterns("¤#,##0.00"),
				AccountingFmt: locales.MustParseNumberFormatPatterns("¤#,##0.00"),
				Short: locales.CurrencyAccountingFormatterByExp{
					CurrencyFmt: map[uint8]*locales.CurrencyAccountingFormatterByExpPlural{
						3: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤0 ພັນ"),
						}},
						4: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤00 ພັນ"),
						}},
						5: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤000 ກີບ"),
						}},
						6: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤0 ລ້ານ"),
						}},
						7: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤00 ລ້ານ"),
						}},
						8: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤000 ລ້ານ"),
						}},
						9: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤0 ຕື້"),
						}},
						10: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤00 ຕື້"),
						}},
						11: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤000 ຕື້"),
						}},
						12: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤0 ລ້ານລ້ານ"),
						}},
						13: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤00 ລ້ານລ້ານ"),
						}},
						14: {Rules: map[locales.PluralRule]*locales.NumberFormatProperties{
							locales.PluralRuleOther: locales.MustParseNumberFormatPatterns("¤000 ລ້ານລ້ານ"),
						}},
					},
					AccountingFmt: map[uint8]*locales.CurrencyAccountingFormatterByExpPlural{},
				},
			},
		},
		TimeSpec: &locales.TimeSpec{
			Separator:                ":",
			MonthsAbbreviatedValue:   []string{"", "ມ.ກ.", "ກ.ພ.", "ມ.ນ.", "ມ.ສ.", "ພ.ພ.", "ມິ.ຖ.", "ກ.ລ.", "ສ.ຫ.", "ກ.ຍ.", "ຕ.ລ.", "ພ.ຈ.", "ທ.ວ."},
			MonthsNarrowValue:        []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
			MonthsWideValue:          []string{"", "ມັງກອນ", "ກຸມພາ", "ມີນາ", "ເມສາ", "ພຶດສະພາ", "ມິຖຸນາ", "ກໍລະກົດ", "ສິງຫາ", "ກັນຍາ", "ຕຸລາ", "ພະຈິກ", "ທັນວາ"},
			WeekdaysAbbreviatedValue: []string{"ອາທິດ", "ຈັນ", "ອັງຄານ", "ພຸດ", "ພະຫັດ", "ສຸກ", "ເສົາ"},
			WeekdaysNarrowValue:      []string{"ອາ", "ຈ", "ອ", "ພ", "ພຫ", "ສຸ", "ສ"},
			WeekdaysShortValue:       []string{"ອາ.", "ຈ.", "ອ.", "ພ.", "ພຫ.", "ສຸ.", "ສ."},
			WeekdaysWideValue:        []string{"ວັນອາທິດ", "ວັນຈັນ", "ວັນອັງຄານ", "ວັນພຸດ", "ວັນພະຫັດ", "ວັນສຸກ", "ວັນເສົາ"},
			PeriodsAbbreviatedValue:  []string{"ກ່ອນທ່ຽງ", "ຫຼັງທ່ຽງ"},
			PeriodsNarrowValue:       []string{"", ""},
			PeriodsWideValue:         []string{"ກ່ອນທ່ຽງ", "ຫຼັງທ່ຽງ"},
			ErasAbbreviatedValue:     []string{"ກ່ອນ ຄ.ສ.", "ຄ.ສ."},
			ErasNarrowValue:          []string{"", ""},
			ErasWideValue:            []string{"ກ່ອນຄຣິດສັກກະລາດ", "ຄຣິດສັກກະລາດ"},
			TimezonesValue:           map[string]string{"": "\u200bເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານໝູ່\u200bເກາະ\u200bຟອ\u200bລ໌ກ\u200bແລນ", "ACDT": "ເວ\u200bລາ\u200bຕອນ\u200bທ່ຽງ\u200bອອສ\u200bເຕຣ\u200bເລຍ\u200bກາງ", "ACST": "ເວ\u200bລາມາດ\u200bຕະ\u200bຖານອອ\u200bສ\u200bເຕຣ\u200bເລຍ\u200bກ\u200bາງ", "ACT": "ເວລາມາດຕະຖານຂອງອາເກຣ", "ACWDT": "ເວ\u200bລາ\u200bຕອນ\u200bທ່ຽງ\u200bອອສ\u200bເຕຣ\u200bລຽນ\u200bກາງ\u200bຕາ\u200bເວັນ\u200bຕົກ", "ACWST": "ເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານອອສ\u200bເຕຣ\u200bລຽນ\u200bກາງ\u200bຕາ\u200bເວັນ\u200bຕົກ", "ADT": "ເວລາກາງເວັນຂອງອາແລນຕິກ", "AEDT": "ເວ\u200bລາ\u200bຕອນ\u200bທ່ຽງ\u200bອອສ\u200bເຕຣ\u200bລຽນ\u200bຕາ\u200bເວັນ\u200bອອກ", "AEST": "ເວ\u200bລາ\u200bມາດຕະຖານ\u200b\u200b\u200bອອສ\u200bເຕຣ\u200bລຽນ\u200bຕາ\u200bເວັນ\u200bອອກ", "AFT": "ເວລາ ອັຟການິສຖານ", "AKDT": "ເວລາກາງເວັນອະລັສກາ", "AKST": "ເວລາມາດຕະຖານອະລັສກາ", "AMST": "ເວ\u200bລາ\u200bລະ\u200bດູ\u200bຮ້ອນອາ\u200bເມ\u200bຊອນ", "AMT": "ເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານອາ\u200bເມ\u200bຊອນ", "ARST": "\u200bເວ\u200bລາ\u200bລະ\u200bດູ\u200bຮ້ອນ\u200bອາ\u200bເຈນ\u200bທິ\u200bນາ", "ART": "\u200bເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານອາ\u200bເຈນ\u200bທິ\u200bນາ", "AST": "ເວລາມາດຕະຖານຂອງອາແລນຕິກ", "AWDT": "ເວ\u200bລາ\u200bຕອນ\u200bທ່ຽງ\u200bອອສ\u200bເຕຣ\u200bລຽນ\u200bຕາ\u200bເວັນ\u200bຕົກ", "AWST": "ເວ\u200bລາ\u200bມາ\u200bດ\u200bຕະ\u200bຖານອອສ\u200bເຕຣ\u200bລຽນ\u200bຕາ\u200bເວັນ\u200bຕົກ", "BNT": "\u200bເວ\u200bລາບຣູ\u200bໄນດາ\u200bຣຸສ\u200bຊາ\u200bລາມ", "BOT": "ເວ\u200bລາ\u200bໂບ\u200bລິ\u200bເວຍ", "BRST": "ເວລາຕາມເຂດລະດູຮ້ອນຕາມເຂດບຣາຊີເລຍ", "BRT": "ເວລາມາດຕາຖານເບຣຊີເລຍ", "BST": "ເວລາ ລະດູຮ້ອນ ບັງກະລາເທດ", "BTT": "ເວ\u200bລາ\u200bພູ\u200bຖານ", "CAT": "ເວ\u200bລາ\u200bອາ\u200bຟຣິ\u200bກາ\u200bກາງ", "CCT": "ເວລາຫມູ່ເກາະໂກໂກສ", "CDT": "ເວລາກາງເວັນກາງ", "CDT (Kuba)": "ເວລາກາງເວັນຄິວບາ", "CHADT": "ເວ\u200bລາ\u200bຕອນ\u200bທ່ຽງ\u200bຊາ\u200bທາມ", "CHAST": "ເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານ\u200bຊາ\u200bທາມ", "CLST": "ເວ\u200bລາ\u200bລະ\u200bດູ\u200bຮ້ອນຊິ\u200bລີ", "CLT": "ເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານຊິ\u200bລີ", "COST": "ເວລາລະດູຮ້ອນໂຄລໍາເບຍ", "COT": "ເວລາມາດຕະຖານໂຄລຳເບຍ", "CST": "ເວລາມາດຕະຖານກາງ", "CST (Kuba)": "ເວລາມາດຕະຖານຂອງຄິວບາ", "CXT": "ເວ\u200bລາ\u200bເກາະ\u200bຄ\u200bຣິສ\u200bມາສ", "ChST": "ເວ\u200bລາ\u200bຈາ\u200bໂມ\u200bໂຣ", "EASST": "ເວ\u200bລາ\u200bລະ\u200bດູ\u200bຮ້ອນເກາະ\u200bອີ\u200bສ\u200bເຕີ", "EAST": "ເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານເກາະ\u200bອີ\u200bສ\u200bເຕີ", "EAT": "ເວ\u200bລາ\u200bອາ\u200bຟຣິ\u200bກາ\u200bຕາ\u200bເວັນ\u200bອອກ", "ECT": "ເວ\u200bລາ\u200bເອ\u200bກົວ\u200bດໍ", "EDT": "ເວລາກາງເວັນຕາເວັນອອກ", "EGDT": "ເວລາລະດູຮ້ອນກຣີນແລນຕາເວັນອອກ", "EGST": "ເວລາມາດຕະຖານຕາເວັນອອກກຣີນແລນ", "EST": "ເວລາມາດຕະຖານຕາເວັນອອກ", "FKST": "\u200bເວ\u200bລາ\u200bລະ\u200bດູ\u200bຮ້ອນໝູ່\u200bເກາະ\u200bຟອ\u200bລ໌ກ\u200bແລນ", "GALT": "ເວ\u200bລາ\u200bກາ\u200bລາ\u200bປາ\u200bກອ\u200bສ", "GFT": "ເວ\u200bລາ\u200bເຟ\u200bຣນ\u200bຊ໌\u200bເກຍ\u200bນາ", "GMT": "ເວ\u200bລາກຣີນ\u200bວິ\u200bຊ", "GST": "ເວ\u200bລາ\u200bກູ\u200bລ\u200b໌ຟ", "GYT": "ເວລາກາຍອານາ", "HADT": "ເວລາຕອນທ່ຽງຮາວາຍ-ເອລູທຽນ", "HAST": "ເວລາມາດຕະຖານຮາວາຍ-ເອລູທຽນ", "HENOMX": "ເວລາກາງເວັນເມັກຊິກັນນອດເວສ", "HEOG": "ເວລາຕອນທ່ຽງກຣີນແລນຕາເວັນຕົກ", "HEPMX": "ເວລາກາງເວັນແປຊິຟິກເມັກຊິກັນ", "HKST": "\u200bເວ\u200bລາ\u200bລະ\u200bດູ\u200bຮ້ອນ\u200bຮອງ\u200bກົງ", "HKT": "ເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານ\u200bຮອງ\u200bກົງ", "HNNOMX": "\u200bເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານນອດ\u200bເວ\u200bສ\u200bເມັກ\u200bຊິ\u200bໂກ", "HNOG": "ເວລາມາດຕະຖານກຣີນແລນຕາເວັນຕົກ", "HNPMX": "ເວລາມາດຕະຖານແປຊິຟິກເມັກຊິກັນ", "ICT": "ເວລາອິນດູຈີນ", "IRDT": "ເວ\u200bລາ\u200bຕອນ\u200bທ່ຽງ\u200bອີ\u200bຣາ\u200bນ", "IRST": "ເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານອີ\u200bຣານ", "IST": "ເວລາ ອິນເດຍ", "JDT": "ເວ\u200bລາ\u200bຕອນ\u200bທ່ຽງ\u200bຍີ່\u200bປຸ່ນ", "JST": "ເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານ\u200bຍີ່\u200bປຸ່ນ", "LHDT": "\u200bເວ\u200bລ\u200bສາ\u200bຕອນ\u200b\u200bທ່ຽງ\u200bລອດ\u200bເຮົາ\u200b", "LHST": "ເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານ\u200bລອດ\u200bເຮົາ", "MDT": "ເວລາລະດູຮ້ອນມາເກົາ", "MESZ": "\u200bເວ\u200bລາ\u200bລະ\u200bດູ\u200bຮ້ອນ\u200bຢູ\u200bໂຣບ\u200bກາງ", "MEZ": "ເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານ\u200bຢູ\u200bໂຣບກາງ", "MST": "ເວລາມາດຕະຖານມາເກົາ", "MVT": "ເວລາມັລດີຟ", "MYT": "ເວ\u200bລາ\u200bມາ\u200bເລ\u200bເຊຍ", "NDT": "ເວລາກາງເວັນນິວຟາວແລນ", "NPT": "\u200bເວ\u200bລາ\u200bເນ\u200bປານ", "NST": "ເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານ\u200bນິວ\u200bຟາວ\u200bແລນ", "NZDT": "ເວ\u200bລາ\u200bຕອນ\u200bທ່ຽງ\u200bນິວ\u200bຊີ\u200bແລນ", "NZST": "ເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານນິວ\u200bຊີ\u200bແລນ", "OESZ": "ເວ\u200bລາ\u200bລະ\u200bດູ\u200bຮ້ອນຢູ\u200bໂຣບ\u200bຕາ\u200bເວັນ\u200bອອກ", "OEZ": "ເວ\u200bລາ\u200bມາ\u200bດ\u200bຕະ\u200bຖານ\u200bຢູ\u200bໂຣບ\u200bຕາ\u200bເວັນ\u200bອອກ", "PDT": "ເວລາກາງເວັນແປຊິຟິກ", "PKT": "ເວ\u200bລາ\u200bລະ\u200bດູ\u200bຮ້ອນ\u200bປາ\u200bກີ\u200bສ\u200bຖານ", "PMDT": "\u200bເວ\u200bລາຕອນ\u200bທ່ຽງເຊນ\u200bປີ\u200bແອ ແລະ\u200bມິ\u200bກົວ\u200bລອນ", "PMST": "\u200bເວ\u200bລາມາດ\u200bຕະ\u200bຖານເຊນ\u200bປີ\u200bແອ ແລະ\u200bມິ\u200bກົວ\u200bລອນ", "PST": "ເວລາມາດຕະຖານແປຊິຟິກ", "PYST": "ເວ\u200bລາ\u200bລະ\u200bດູ\u200bຮ້ອນ\u200bປາ\u200bຣາ\u200bກວຍ", "SAST": "ເວ\u200bລາ\u200bອາ\u200bຟຣິ\u200bກາ\u200bໃຕ້", "SGT": "ເວ\u200bລາ\u200bສິງ\u200bກະ\u200bໂປ", "SRT": "ເວ\u200bລາ\u200bຊຸ\u200bຣິ\u200bນາມ", "TLT": "ເວລາຕີມໍຕາເວັນອອກ", "TMST": "ເວລາລະດູຮ້ອນຕວກເມນິສຖານ", "TMT": "ເວລາມາດຕະຖານຕວກເມນິສຖານ", "UYST": "ເວ\u200bລາ\u200bລະ\u200bດູ\u200bຮ້ອນ\u200bອູ\u200bຣູ\u200bກວຍ", "UYT": "ເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານ\u200bອູ\u200bຣູ\u200bກວຍ", "VET": "ເວ\u200bລາ\u200bເວ\u200bເນ\u200bຊູ\u200bເອ\u200bລາ", "WARST": "ເວ\u200bລາ\u200bລະ\u200bດູ\u200bຮ້ອນເວ\u200bສ\u200bເທິນອາ\u200bເຈນ\u200bທິ\u200bນາ", "WART": "ເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານເວ\u200bສ\u200bເທິນອາ\u200bເຈນ\u200bທິ\u200bນາ", "WAST": "ເວ\u200bລາ\u200bລະ\u200bດູ\u200bຮ້ອນ\u200bອາ\u200bຟຣິ\u200bກາ\u200bຕາ\u200bເວັນ\u200bຕົກ", "WAT": "ເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານ\u200bອາ\u200bຟຣິ\u200bກາ\u200bຕາ\u200bເວັນ\u200bຕົກ", "WESZ": "ເວ\u200bລາ\u200bລະ\u200bດູ\u200bຮ້ອນຢູ\u200bໂຣບ\u200bຕາ\u200bເວັນ\u200bຕົກ", "WEZ": "ເວ\u200bລາ\u200bມາດ\u200bຕະ\u200bຖານຢູ\u200bໂຣບ\u200bຕາ\u200bເວັນ\u200bຕົກ", "WIB": "ເວ\u200bລາ\u200bອິນ\u200bໂດ\u200bເນ\u200bເຊຍ\u200bຕາ\u200bເວັນ\u200bຕົກ", "WIT": "ເວ\u200bລາ\u200bອິນ\u200bໂດ\u200bເນ\u200bເຊຍ\u200bຕາ\u200bເວັນ\u200bອອກ", "WITA": "ເວ\u200bລາ\u200bອິນ\u200bໂດ\u200bເນ\u200bເຊຍ\u200bກາງ", "∅∅∅": "ເວ\u200bລາ\u200bລະ\u200bດູ\u200bຮ້ອນ\u200bເປ\u200bຣູ"},
		},

		locale:          "lo_LA",
		pluralsCardinal: []locales.PluralRule{6},
		pluralsOrdinal:  []locales.PluralRule{2, 6},
		pluralsRange:    []locales.PluralRule{6},
		percent:         "%",
		perMille:        "‰",
		inifinity:       "∞",

		dateFullLayout:   "EEEE ທີ d MMMM G y",
		dateLongLayout:   "d MMMM y",
		dateMediumLayout: "d MMM y",
		dateShortLayout:  "d/M/y",
		timeFullLayout:   "H ໂມງ m ນາທີ ss ວິນາທີ zzzz",
		timeLongLayout:   "H ໂມງ m ນາທີ ss ວິນາທີ z",
		timeMediumLayout: "H:mm:ss",
		timeShortLayout:  "H:mm",
		listPatterns:     locales.NewListPatternsFromSlice([][]string{{"{0} ແລະ {1}", "{0}, {1}", "{0}, {1}", "{0}, {1}"}, {"{0} ຫຼື {1}", "{0}, {1}", "{0}, {1}", "{0} ຫຼື {1}"}}),
		miscPatterns: locales.MiscPatterns{
			"~%[1]v",
			"%[1]v+",
			"≤%[1]v",
			"%[1]v–%[2]v",
		},
	}
}

// Locale returns the current translators string locale
func (lo *lo_LA) Locale() string {
	return lo.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'lo_LA'
func (lo *lo_LA) PluralsCardinal() []locales.PluralRule {
	return lo.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'lo_LA'
func (lo *lo_LA) PluralsOrdinal() []locales.PluralRule {
	return lo.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'lo_LA'
func (lo *lo_LA) PluralsRange() []locales.PluralRule {
	return lo.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'lo_LA'
func (lo *lo_LA) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'lo_LA'
func (lo *lo_LA) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'lo_LA'
func (lo *lo_LA) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// FmtDateShort returns the short date representation of 't' for 'lo_LA'
func (lo *lo_LA) FmtDateShort(t time.Time) string {
	return locales.FormatTimeEra(lo, lo.dateShortLayout, t, 2)
}

// FmtDateMedium returns the medium date representation of 't' for 'lo_LA'
func (lo *lo_LA) FmtDateMedium(t time.Time) string {
	return locales.FormatTimeEra(lo, lo.dateMediumLayout, t, 2)
}

// FmtDateLong returns the long date representation of 't' for 'lo_LA'
func (lo *lo_LA) FmtDateLong(t time.Time) string {
	return locales.FormatTimeEra(lo, lo.dateLongLayout, t, 1)
}

// FmtDateFull returns the full date representation of 't' for 'lo_LA'
func (lo *lo_LA) FmtDateFull(t time.Time) string {
	return locales.FormatTimeEra(lo, lo.dateFullLayout, t, 0)
}

// FmtTimeShort returns the short time representation of 't' for 'lo_LA'
func (lo *lo_LA) FmtTimeShort(t time.Time) string {
	return locales.FormatTimeEra(lo, lo.timeShortLayout, t, 2)
}

// FmtTimeMedium returns the medium time representation of 't' for 'lo_LA'
func (lo *lo_LA) FmtTimeMedium(t time.Time) string {
	return locales.FormatTimeEra(lo, lo.timeMediumLayout, t, 2)
}

// FmtTimeLong returns the long time representation of 't' for 'lo_LA'
func (lo *lo_LA) FmtTimeLong(t time.Time) string {
	return locales.FormatTimeEra(lo, lo.timeLongLayout, t, 1)
}

// FmtTimeFull returns the full time representation of 't' for 'lo_LA'
func (lo *lo_LA) FmtTimeFull(t time.Time) string {
	return locales.FormatTimeEra(lo, lo.timeFullLayout, t, 0)
}

// DateFullLayout returns the full date layout for 'lo_LA'
func (lo *lo_LA) DateFullLayout() string {
	return lo.dateFullLayout
}

// DateLongLayout returns the long date layout for 'lo_LA'
func (lo *lo_LA) DateLongLayout() string {
	return lo.dateLongLayout
}

// DateMediumLayout returns the medium date layout for 'lo_LA'
func (lo *lo_LA) DateMediumLayout() string {
	return lo.dateMediumLayout
}

// DateShortLayout returns the short date layout for 'lo_LA'
func (lo *lo_LA) DateShortLayout() string {
	return lo.dateShortLayout
}

// TimeFullLayout returns the full time layout for 'lo_LA'
func (lo *lo_LA) TimeFullLayout() string {
	return lo.timeFullLayout
}

// TimeLongLayout returns the full long layout for 'lo_LA'
func (lo *lo_LA) TimeLongLayout() string {
	return lo.timeLongLayout
}

// TimeMediumLayout returns the medium time layout for 'lo_LA'
func (lo *lo_LA) TimeMediumLayout() string {
	return lo.timeMediumLayout
}

// TimeShortLayout returns the short time layout for 'lo_LA'
func (lo *lo_LA) TimeShortLayout() string {
	return lo.timeShortLayout
}

func (lo *lo_LA) ListPatterns() *locales.ListPatterns {
	return lo.listPatterns
}

func (lo *lo_LA) GetDurationSpec() *locales.DurationSpec {
	return lo.DurationSpec
}

func (lo *lo_LA) GetMiscPatterns() *locales.MiscPatterns {
	return &lo.miscPatterns
}
