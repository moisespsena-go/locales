package ccp

import (
	"time"

	"github.com/moisespsena-go/locales"
)

type ccp struct {
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

// New returns a new instance of translator for the 'ccp' locale
func New() locales.Translator {
	return &ccp{
		NumberSpec: &locales.NumberSpec{
			DecimalValue: ".",
		},
		CurrencySpec: &locales.CurrencySpec{
			CurrenciesValue: []locales.Currency{{Names: map[locales.PluralRule]string{0: "ADP"}, Symbols: locales.CurrencySymbols{Default: "ADP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AED"}, Symbols: locales.CurrencySymbols{Default: "AED", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AFA"}, Symbols: locales.CurrencySymbols{Default: "AFA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AFN"}, Symbols: locales.CurrencySymbols{Default: "AFN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ALK"}, Symbols: locales.CurrencySymbols{Default: "ALK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ALL"}, Symbols: locales.CurrencySymbols{Default: "ALL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AMD"}, Symbols: locales.CurrencySymbols{Default: "AMD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ANG"}, Symbols: locales.CurrencySymbols{Default: "ANG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AOA"}, Symbols: locales.CurrencySymbols{Default: "AOA", Narrow: "Kz"}}, {Names: map[locales.PluralRule]string{0: "AOK"}, Symbols: locales.CurrencySymbols{Default: "AOK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AON"}, Symbols: locales.CurrencySymbols{Default: "AON", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AOR"}, Symbols: locales.CurrencySymbols{Default: "AOR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARA"}, Symbols: locales.CurrencySymbols{Default: "ARA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARL"}, Symbols: locales.CurrencySymbols{Default: "ARL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARM"}, Symbols: locales.CurrencySymbols{Default: "ARM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARP"}, Symbols: locales.CurrencySymbols{Default: "ARP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ARS"}, Symbols: locales.CurrencySymbols{Default: "ARS", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "ATS"}, Symbols: locales.CurrencySymbols{Default: "ATS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AUD"}, Symbols: locales.CurrencySymbols{Default: "A$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "AWG"}, Symbols: locales.CurrencySymbols{Default: "AWG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AZM"}, Symbols: locales.CurrencySymbols{Default: "AZM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "AZN"}, Symbols: locales.CurrencySymbols{Default: "AZN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BAD"}, Symbols: locales.CurrencySymbols{Default: "BAD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BAM"}, Symbols: locales.CurrencySymbols{Default: "BAM", Narrow: "KM"}}, {Names: map[locales.PluralRule]string{0: "BAN"}, Symbols: locales.CurrencySymbols{Default: "BAN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BBD"}, Symbols: locales.CurrencySymbols{Default: "BBD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "BDT"}, Symbols: locales.CurrencySymbols{Default: "৳", Narrow: "৳"}}, {Names: map[locales.PluralRule]string{0: "BEC"}, Symbols: locales.CurrencySymbols{Default: "BEC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BEF"}, Symbols: locales.CurrencySymbols{Default: "BEF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BEL"}, Symbols: locales.CurrencySymbols{Default: "BEL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGL"}, Symbols: locales.CurrencySymbols{Default: "BGL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGM"}, Symbols: locales.CurrencySymbols{Default: "BGM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGN"}, Symbols: locales.CurrencySymbols{Default: "BGN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BGO"}, Symbols: locales.CurrencySymbols{Default: "BGO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BHD"}, Symbols: locales.CurrencySymbols{Default: "BHD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BIF"}, Symbols: locales.CurrencySymbols{Default: "BIF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BMD"}, Symbols: locales.CurrencySymbols{Default: "BMD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "BND"}, Symbols: locales.CurrencySymbols{Default: "BND", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "BOB"}, Symbols: locales.CurrencySymbols{Default: "BOB", Narrow: "Bs"}}, {Names: map[locales.PluralRule]string{0: "BOL"}, Symbols: locales.CurrencySymbols{Default: "BOL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BOP"}, Symbols: locales.CurrencySymbols{Default: "BOP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BOV"}, Symbols: locales.CurrencySymbols{Default: "BOV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRB"}, Symbols: locales.CurrencySymbols{Default: "BRB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRC"}, Symbols: locales.CurrencySymbols{Default: "BRC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRE"}, Symbols: locales.CurrencySymbols{Default: "BRE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRL"}, Symbols: locales.CurrencySymbols{Default: "R$", Narrow: "R$"}}, {Names: map[locales.PluralRule]string{0: "BRN"}, Symbols: locales.CurrencySymbols{Default: "BRN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRR"}, Symbols: locales.CurrencySymbols{Default: "BRR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BRZ"}, Symbols: locales.CurrencySymbols{Default: "BRZ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BSD"}, Symbols: locales.CurrencySymbols{Default: "BSD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "BTN"}, Symbols: locales.CurrencySymbols{Default: "BTN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BUK"}, Symbols: locales.CurrencySymbols{Default: "BUK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BWP"}, Symbols: locales.CurrencySymbols{Default: "BWP", Narrow: "P"}}, {Names: map[locales.PluralRule]string{0: "BYB"}, Symbols: locales.CurrencySymbols{Default: "BYB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BYN"}, Symbols: locales.CurrencySymbols{Default: "BYN", Narrow: "р."}}, {Names: map[locales.PluralRule]string{0: "BYR"}, Symbols: locales.CurrencySymbols{Default: "BYR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "BZD"}, Symbols: locales.CurrencySymbols{Default: "BZD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CAD"}, Symbols: locales.CurrencySymbols{Default: "CA$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CDF"}, Symbols: locales.CurrencySymbols{Default: "CDF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CHE"}, Symbols: locales.CurrencySymbols{Default: "CHE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CHF"}, Symbols: locales.CurrencySymbols{Default: "CHF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CHW"}, Symbols: locales.CurrencySymbols{Default: "CHW", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CLE"}, Symbols: locales.CurrencySymbols{Default: "CLE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CLF"}, Symbols: locales.CurrencySymbols{Default: "CLF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CLP"}, Symbols: locales.CurrencySymbols{Default: "CLP", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CNH"}, Symbols: locales.CurrencySymbols{Default: "CNH", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CNX"}, Symbols: locales.CurrencySymbols{Default: "CNX", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CNY"}, Symbols: locales.CurrencySymbols{Default: "CN¥", Narrow: "¥"}}, {Names: map[locales.PluralRule]string{0: "COP"}, Symbols: locales.CurrencySymbols{Default: "COP", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "COU"}, Symbols: locales.CurrencySymbols{Default: "COU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CRC"}, Symbols: locales.CurrencySymbols{Default: "CRC", Narrow: "₡"}}, {Names: map[locales.PluralRule]string{0: "CSD"}, Symbols: locales.CurrencySymbols{Default: "CSD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CSK"}, Symbols: locales.CurrencySymbols{Default: "CSK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CUC"}, Symbols: locales.CurrencySymbols{Default: "CUC", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CUP"}, Symbols: locales.CurrencySymbols{Default: "CUP", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "CVE"}, Symbols: locales.CurrencySymbols{Default: "CVE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CYP"}, Symbols: locales.CurrencySymbols{Default: "CYP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "CZK"}, Symbols: locales.CurrencySymbols{Default: "CZK", Narrow: "Kč"}}, {Names: map[locales.PluralRule]string{0: "DDM"}, Symbols: locales.CurrencySymbols{Default: "DDM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DEM"}, Symbols: locales.CurrencySymbols{Default: "DEM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DJF"}, Symbols: locales.CurrencySymbols{Default: "DJF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "DKK"}, Symbols: locales.CurrencySymbols{Default: "DKK", Narrow: "kr"}}, {Names: map[locales.PluralRule]string{0: "DOP"}, Symbols: locales.CurrencySymbols{Default: "DOP", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "DZD"}, Symbols: locales.CurrencySymbols{Default: "DZD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ECS"}, Symbols: locales.CurrencySymbols{Default: "ECS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ECV"}, Symbols: locales.CurrencySymbols{Default: "ECV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "EEK"}, Symbols: locales.CurrencySymbols{Default: "EEK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "EGP"}, Symbols: locales.CurrencySymbols{Default: "EGP", Narrow: "E£"}}, {Names: map[locales.PluralRule]string{0: "ERN"}, Symbols: locales.CurrencySymbols{Default: "ERN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ESA"}, Symbols: locales.CurrencySymbols{Default: "ESA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ESB"}, Symbols: locales.CurrencySymbols{Default: "ESB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ESP"}, Symbols: locales.CurrencySymbols{Default: "ESP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ETB"}, Symbols: locales.CurrencySymbols{Default: "ETB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "EUR"}, Symbols: locales.CurrencySymbols{Default: "€", Narrow: "€"}}, {Names: map[locales.PluralRule]string{0: "FIM"}, Symbols: locales.CurrencySymbols{Default: "FIM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "FJD"}, Symbols: locales.CurrencySymbols{Default: "FJD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "FKP"}, Symbols: locales.CurrencySymbols{Default: "FKP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "FRF"}, Symbols: locales.CurrencySymbols{Default: "FRF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GBP"}, Symbols: locales.CurrencySymbols{Default: "£", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "GEK"}, Symbols: locales.CurrencySymbols{Default: "GEK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GEL"}, Symbols: locales.CurrencySymbols{Default: "GEL", Narrow: "₾"}}, {Names: map[locales.PluralRule]string{0: "GHC"}, Symbols: locales.CurrencySymbols{Default: "GHC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GHS"}, Symbols: locales.CurrencySymbols{Default: "GHS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GIP"}, Symbols: locales.CurrencySymbols{Default: "GIP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "GMD"}, Symbols: locales.CurrencySymbols{Default: "GMD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GNF"}, Symbols: locales.CurrencySymbols{Default: "GNF", Narrow: "FG"}}, {Names: map[locales.PluralRule]string{0: "GNS"}, Symbols: locales.CurrencySymbols{Default: "GNS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GQE"}, Symbols: locales.CurrencySymbols{Default: "GQE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GRD"}, Symbols: locales.CurrencySymbols{Default: "GRD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GTQ"}, Symbols: locales.CurrencySymbols{Default: "GTQ", Narrow: "Q"}}, {Names: map[locales.PluralRule]string{0: "GWE"}, Symbols: locales.CurrencySymbols{Default: "GWE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GWP"}, Symbols: locales.CurrencySymbols{Default: "GWP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "GYD"}, Symbols: locales.CurrencySymbols{Default: "GYD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "HKD"}, Symbols: locales.CurrencySymbols{Default: "HK$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "HNL"}, Symbols: locales.CurrencySymbols{Default: "HNL", Narrow: "L"}}, {Names: map[locales.PluralRule]string{0: "HRD"}, Symbols: locales.CurrencySymbols{Default: "HRD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "HRK"}, Symbols: locales.CurrencySymbols{Default: "HRK", Narrow: "kn"}}, {Names: map[locales.PluralRule]string{0: "HTG"}, Symbols: locales.CurrencySymbols{Default: "HTG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "HUF"}, Symbols: locales.CurrencySymbols{Default: "HUF", Narrow: "Ft"}}, {Names: map[locales.PluralRule]string{0: "IDR"}, Symbols: locales.CurrencySymbols{Default: "IDR", Narrow: "Rp"}}, {Names: map[locales.PluralRule]string{0: "IEP"}, Symbols: locales.CurrencySymbols{Default: "IEP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ILP"}, Symbols: locales.CurrencySymbols{Default: "ILP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ILR"}, Symbols: locales.CurrencySymbols{Default: "ILR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ILS"}, Symbols: locales.CurrencySymbols{Default: "₪", Narrow: "₪"}}, {Names: map[locales.PluralRule]string{0: "INR"}, Symbols: locales.CurrencySymbols{Default: "₹", Narrow: "₹"}}, {Names: map[locales.PluralRule]string{0: "IQD"}, Symbols: locales.CurrencySymbols{Default: "IQD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "IRR"}, Symbols: locales.CurrencySymbols{Default: "IRR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ISJ"}, Symbols: locales.CurrencySymbols{Default: "ISJ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ISK"}, Symbols: locales.CurrencySymbols{Default: "ISK", Narrow: "kr"}}, {Names: map[locales.PluralRule]string{0: "ITL"}, Symbols: locales.CurrencySymbols{Default: "ITL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "JMD"}, Symbols: locales.CurrencySymbols{Default: "JMD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "JOD"}, Symbols: locales.CurrencySymbols{Default: "JOD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "JPY"}, Symbols: locales.CurrencySymbols{Default: "JP¥", Narrow: "¥"}}, {Names: map[locales.PluralRule]string{0: "KES"}, Symbols: locales.CurrencySymbols{Default: "KES", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KGS"}, Symbols: locales.CurrencySymbols{Default: "KGS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KHR"}, Symbols: locales.CurrencySymbols{Default: "KHR", Narrow: "៛"}}, {Names: map[locales.PluralRule]string{0: "KMF"}, Symbols: locales.CurrencySymbols{Default: "KMF", Narrow: "CF"}}, {Names: map[locales.PluralRule]string{0: "KPW"}, Symbols: locales.CurrencySymbols{Default: "KPW", Narrow: "₩"}}, {Names: map[locales.PluralRule]string{0: "KRH"}, Symbols: locales.CurrencySymbols{Default: "KRH", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KRO"}, Symbols: locales.CurrencySymbols{Default: "KRO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KRW"}, Symbols: locales.CurrencySymbols{Default: "₩", Narrow: "₩"}}, {Names: map[locales.PluralRule]string{0: "KWD"}, Symbols: locales.CurrencySymbols{Default: "KWD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "KYD"}, Symbols: locales.CurrencySymbols{Default: "KYD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "KZT"}, Symbols: locales.CurrencySymbols{Default: "KZT", Narrow: "₸"}}, {Names: map[locales.PluralRule]string{0: "LAK"}, Symbols: locales.CurrencySymbols{Default: "LAK", Narrow: "₭"}}, {Names: map[locales.PluralRule]string{0: "LBP"}, Symbols: locales.CurrencySymbols{Default: "LBP", Narrow: "L£"}}, {Names: map[locales.PluralRule]string{0: "LKR"}, Symbols: locales.CurrencySymbols{Default: "LKR", Narrow: "Rs"}}, {Names: map[locales.PluralRule]string{0: "LRD"}, Symbols: locales.CurrencySymbols{Default: "LRD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "LSL"}, Symbols: locales.CurrencySymbols{Default: "LSL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LTL"}, Symbols: locales.CurrencySymbols{Default: "LTL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LTT"}, Symbols: locales.CurrencySymbols{Default: "LTT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LUC"}, Symbols: locales.CurrencySymbols{Default: "LUC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LUF"}, Symbols: locales.CurrencySymbols{Default: "LUF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LUL"}, Symbols: locales.CurrencySymbols{Default: "LUL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LVL"}, Symbols: locales.CurrencySymbols{Default: "LVL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LVR"}, Symbols: locales.CurrencySymbols{Default: "LVR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "LYD"}, Symbols: locales.CurrencySymbols{Default: "LYD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MAD"}, Symbols: locales.CurrencySymbols{Default: "MAD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MAF"}, Symbols: locales.CurrencySymbols{Default: "MAF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MCF"}, Symbols: locales.CurrencySymbols{Default: "MCF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MDC"}, Symbols: locales.CurrencySymbols{Default: "MDC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MDL"}, Symbols: locales.CurrencySymbols{Default: "MDL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MGA"}, Symbols: locales.CurrencySymbols{Default: "MGA", Narrow: "Ar"}}, {Names: map[locales.PluralRule]string{0: "MGF"}, Symbols: locales.CurrencySymbols{Default: "MGF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MKD"}, Symbols: locales.CurrencySymbols{Default: "MKD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MKN"}, Symbols: locales.CurrencySymbols{Default: "MKN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MLF"}, Symbols: locales.CurrencySymbols{Default: "MLF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MMK"}, Symbols: locales.CurrencySymbols{Default: "MMK", Narrow: "K"}}, {Names: map[locales.PluralRule]string{0: "MNT"}, Symbols: locales.CurrencySymbols{Default: "MNT", Narrow: "₮"}}, {Names: map[locales.PluralRule]string{0: "MOP"}, Symbols: locales.CurrencySymbols{Default: "MOP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MRO"}, Symbols: locales.CurrencySymbols{Default: "MRO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MRU"}, Symbols: locales.CurrencySymbols{Default: "MRU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MTL"}, Symbols: locales.CurrencySymbols{Default: "MTL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MTP"}, Symbols: locales.CurrencySymbols{Default: "MTP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MUR"}, Symbols: locales.CurrencySymbols{Default: "MUR", Narrow: "Rs"}}, {Names: map[locales.PluralRule]string{0: "MVP"}, Symbols: locales.CurrencySymbols{Default: "MVP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MVR"}, Symbols: locales.CurrencySymbols{Default: "MVR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MWK"}, Symbols: locales.CurrencySymbols{Default: "MWK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MXN"}, Symbols: locales.CurrencySymbols{Default: "MX$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "MXP"}, Symbols: locales.CurrencySymbols{Default: "MXP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MXV"}, Symbols: locales.CurrencySymbols{Default: "MXV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MYR"}, Symbols: locales.CurrencySymbols{Default: "MYR", Narrow: "RM"}}, {Names: map[locales.PluralRule]string{0: "MZE"}, Symbols: locales.CurrencySymbols{Default: "MZE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MZM"}, Symbols: locales.CurrencySymbols{Default: "MZM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "MZN"}, Symbols: locales.CurrencySymbols{Default: "MZN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NAD"}, Symbols: locales.CurrencySymbols{Default: "NAD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "NGN"}, Symbols: locales.CurrencySymbols{Default: "NGN", Narrow: "₦"}}, {Names: map[locales.PluralRule]string{0: "NIC"}, Symbols: locales.CurrencySymbols{Default: "NIC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NIO"}, Symbols: locales.CurrencySymbols{Default: "NIO", Narrow: "C$"}}, {Names: map[locales.PluralRule]string{0: "NLG"}, Symbols: locales.CurrencySymbols{Default: "NLG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "NOK"}, Symbols: locales.CurrencySymbols{Default: "NOK", Narrow: "kr"}}, {Names: map[locales.PluralRule]string{0: "NPR"}, Symbols: locales.CurrencySymbols{Default: "NPR", Narrow: "Rs"}}, {Names: map[locales.PluralRule]string{0: "NZD"}, Symbols: locales.CurrencySymbols{Default: "NZ$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "OMR"}, Symbols: locales.CurrencySymbols{Default: "OMR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PAB"}, Symbols: locales.CurrencySymbols{Default: "PAB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PEI"}, Symbols: locales.CurrencySymbols{Default: "PEI", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PEN"}, Symbols: locales.CurrencySymbols{Default: "PEN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PES"}, Symbols: locales.CurrencySymbols{Default: "PES", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PGK"}, Symbols: locales.CurrencySymbols{Default: "PGK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PHP"}, Symbols: locales.CurrencySymbols{Default: "PHP", Narrow: "₱"}}, {Names: map[locales.PluralRule]string{0: "PKR"}, Symbols: locales.CurrencySymbols{Default: "PKR", Narrow: "Rs"}}, {Names: map[locales.PluralRule]string{0: "PLN"}, Symbols: locales.CurrencySymbols{Default: "PLN", Narrow: "zł"}}, {Names: map[locales.PluralRule]string{0: "PLZ"}, Symbols: locales.CurrencySymbols{Default: "PLZ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PTE"}, Symbols: locales.CurrencySymbols{Default: "PTE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "PYG"}, Symbols: locales.CurrencySymbols{Default: "PYG", Narrow: "₲"}}, {Names: map[locales.PluralRule]string{0: "QAR"}, Symbols: locales.CurrencySymbols{Default: "QAR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RHD"}, Symbols: locales.CurrencySymbols{Default: "RHD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ROL"}, Symbols: locales.CurrencySymbols{Default: "ROL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RON"}, Symbols: locales.CurrencySymbols{Default: "RON", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RSD"}, Symbols: locales.CurrencySymbols{Default: "RSD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RUB"}, Symbols: locales.CurrencySymbols{Default: "RUB", Narrow: "₽"}}, {Names: map[locales.PluralRule]string{0: "RUR"}, Symbols: locales.CurrencySymbols{Default: "RUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "RWF"}, Symbols: locales.CurrencySymbols{Default: "RWF", Narrow: "RF"}}, {Names: map[locales.PluralRule]string{0: "SAR"}, Symbols: locales.CurrencySymbols{Default: "SAR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SBD"}, Symbols: locales.CurrencySymbols{Default: "SBD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "SCR"}, Symbols: locales.CurrencySymbols{Default: "SCR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SDD"}, Symbols: locales.CurrencySymbols{Default: "SDD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SDG"}, Symbols: locales.CurrencySymbols{Default: "SDG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SDP"}, Symbols: locales.CurrencySymbols{Default: "SDP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SEK"}, Symbols: locales.CurrencySymbols{Default: "SEK", Narrow: "kr"}}, {Names: map[locales.PluralRule]string{0: "SGD"}, Symbols: locales.CurrencySymbols{Default: "SGD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "SHP"}, Symbols: locales.CurrencySymbols{Default: "SHP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "SIT"}, Symbols: locales.CurrencySymbols{Default: "SIT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SKK"}, Symbols: locales.CurrencySymbols{Default: "SKK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SLE"}, Symbols: locales.CurrencySymbols{Default: "SLE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SLL"}, Symbols: locales.CurrencySymbols{Default: "SLL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SOS"}, Symbols: locales.CurrencySymbols{Default: "SOS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SRD"}, Symbols: locales.CurrencySymbols{Default: "SRD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "SRG"}, Symbols: locales.CurrencySymbols{Default: "SRG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SSP"}, Symbols: locales.CurrencySymbols{Default: "SSP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "STD"}, Symbols: locales.CurrencySymbols{Default: "STD", Narrow: "Db"}}, {Names: map[locales.PluralRule]string{0: "STN"}, Symbols: locales.CurrencySymbols{Default: "STN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SUR"}, Symbols: locales.CurrencySymbols{Default: "SUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SVC"}, Symbols: locales.CurrencySymbols{Default: "SVC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "SYP"}, Symbols: locales.CurrencySymbols{Default: "SYP", Narrow: "£"}}, {Names: map[locales.PluralRule]string{0: "SZL"}, Symbols: locales.CurrencySymbols{Default: "SZL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "THB"}, Symbols: locales.CurrencySymbols{Default: "฿", Narrow: "฿"}}, {Names: map[locales.PluralRule]string{0: "TJR"}, Symbols: locales.CurrencySymbols{Default: "TJR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TJS"}, Symbols: locales.CurrencySymbols{Default: "TJS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TMM"}, Symbols: locales.CurrencySymbols{Default: "TMM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TMT"}, Symbols: locales.CurrencySymbols{Default: "TMT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TND"}, Symbols: locales.CurrencySymbols{Default: "TND", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TOP"}, Symbols: locales.CurrencySymbols{Default: "TOP", Narrow: "T$"}}, {Names: map[locales.PluralRule]string{0: "TPE"}, Symbols: locales.CurrencySymbols{Default: "TPE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TRL"}, Symbols: locales.CurrencySymbols{Default: "TRL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "TRY"}, Symbols: locales.CurrencySymbols{Default: "TRY", Narrow: "₺"}}, {Names: map[locales.PluralRule]string{0: "TTD"}, Symbols: locales.CurrencySymbols{Default: "TTD", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "TWD"}, Symbols: locales.CurrencySymbols{Default: "NT$", Narrow: "NT$"}}, {Names: map[locales.PluralRule]string{0: "TZS"}, Symbols: locales.CurrencySymbols{Default: "TZS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UAH"}, Symbols: locales.CurrencySymbols{Default: "UAH", Narrow: "₴"}}, {Names: map[locales.PluralRule]string{0: "UAK"}, Symbols: locales.CurrencySymbols{Default: "UAK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UGS"}, Symbols: locales.CurrencySymbols{Default: "UGS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UGX"}, Symbols: locales.CurrencySymbols{Default: "UGX", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "USD"}, Symbols: locales.CurrencySymbols{Default: "US$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "USN"}, Symbols: locales.CurrencySymbols{Default: "USN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "USS"}, Symbols: locales.CurrencySymbols{Default: "USS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UYI"}, Symbols: locales.CurrencySymbols{Default: "UYI", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UYP"}, Symbols: locales.CurrencySymbols{Default: "UYP", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UYU"}, Symbols: locales.CurrencySymbols{Default: "UYU", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "UYW"}, Symbols: locales.CurrencySymbols{Default: "UYW", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "UZS"}, Symbols: locales.CurrencySymbols{Default: "UZS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VEB"}, Symbols: locales.CurrencySymbols{Default: "VEB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VED"}, Symbols: locales.CurrencySymbols{Default: "VED", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VEF"}, Symbols: locales.CurrencySymbols{Default: "VEF", Narrow: "Bs"}}, {Names: map[locales.PluralRule]string{0: "VES"}, Symbols: locales.CurrencySymbols{Default: "VES", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VND"}, Symbols: locales.CurrencySymbols{Default: "₫", Narrow: "₫"}}, {Names: map[locales.PluralRule]string{0: "VNN"}, Symbols: locales.CurrencySymbols{Default: "VNN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "VUV"}, Symbols: locales.CurrencySymbols{Default: "VUV", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "WST"}, Symbols: locales.CurrencySymbols{Default: "WST", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XAF"}, Symbols: locales.CurrencySymbols{Default: "FCFA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XAG"}, Symbols: locales.CurrencySymbols{Default: "XAG", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XAU"}, Symbols: locales.CurrencySymbols{Default: "XAU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBA"}, Symbols: locales.CurrencySymbols{Default: "XBA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBB"}, Symbols: locales.CurrencySymbols{Default: "XBB", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBC"}, Symbols: locales.CurrencySymbols{Default: "XBC", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XBD"}, Symbols: locales.CurrencySymbols{Default: "XBD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XCD"}, Symbols: locales.CurrencySymbols{Default: "EC$", Narrow: "$"}}, {Names: map[locales.PluralRule]string{0: "XDR"}, Symbols: locales.CurrencySymbols{Default: "XDR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XEU"}, Symbols: locales.CurrencySymbols{Default: "XEU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XFO"}, Symbols: locales.CurrencySymbols{Default: "XFO", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XFU"}, Symbols: locales.CurrencySymbols{Default: "XFU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XOF"}, Symbols: locales.CurrencySymbols{Default: "F\u202fCFA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XPD"}, Symbols: locales.CurrencySymbols{Default: "XPD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XPF"}, Symbols: locales.CurrencySymbols{Default: "CFPF", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XPT"}, Symbols: locales.CurrencySymbols{Default: "XPT", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XRE"}, Symbols: locales.CurrencySymbols{Default: "XRE", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XSU"}, Symbols: locales.CurrencySymbols{Default: "XSU", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XTS"}, Symbols: locales.CurrencySymbols{Default: "XTS", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XUA"}, Symbols: locales.CurrencySymbols{Default: "XUA", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "XXX"}, Symbols: locales.CurrencySymbols{Default: "XXX", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YDD"}, Symbols: locales.CurrencySymbols{Default: "YDD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YER"}, Symbols: locales.CurrencySymbols{Default: "YER", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUD"}, Symbols: locales.CurrencySymbols{Default: "YUD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUM"}, Symbols: locales.CurrencySymbols{Default: "YUM", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUN"}, Symbols: locales.CurrencySymbols{Default: "YUN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "YUR"}, Symbols: locales.CurrencySymbols{Default: "YUR", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZAL"}, Symbols: locales.CurrencySymbols{Default: "ZAL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZAR"}, Symbols: locales.CurrencySymbols{Default: "ZAR", Narrow: "R"}}, {Names: map[locales.PluralRule]string{0: "ZMK"}, Symbols: locales.CurrencySymbols{Default: "ZMK", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZMW"}, Symbols: locales.CurrencySymbols{Default: "ZMW", Narrow: "ZK"}}, {Names: map[locales.PluralRule]string{0: "ZRN"}, Symbols: locales.CurrencySymbols{Default: "ZRN", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZRZ"}, Symbols: locales.CurrencySymbols{Default: "ZRZ", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZWD"}, Symbols: locales.CurrencySymbols{Default: "ZWD", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZWL"}, Symbols: locales.CurrencySymbols{Default: "ZWL", Narrow: ""}}, {Names: map[locales.PluralRule]string{0: "ZWR"}, Symbols: locales.CurrencySymbols{Default: "ZWR", Narrow: ""}}},
			Formatters: &locales.CurrencyFormatters{
				CurrencyFmt: locales.MustParseNumberFormatPatterns("#,##,##0.00¤"),
			},
		},
		TimeSpec: &locales.TimeSpec{
			Separator:                ":",
			MonthsAbbreviatedValue:   []string{"", "𑄎𑄚𑄪", "𑄜𑄬𑄛𑄴", "𑄟𑄢𑄴𑄌𑄧", "𑄃𑄬𑄛𑄳𑄢𑄨𑄣𑄴", "𑄟𑄬", "𑄎𑄪𑄚𑄴", "𑄎𑄪𑄣𑄭", "𑄃𑄉𑄧𑄌𑄴𑄑𑄴", "𑄥𑄬𑄛𑄴𑄑𑄬𑄟𑄴𑄝𑄧𑄢𑄴", "𑄃𑄧𑄇𑄴𑄑𑄮𑄝𑄧𑄢𑄴", "𑄚𑄧𑄞𑄬𑄟𑄴𑄝𑄧𑄢𑄴", "𑄓𑄨𑄥𑄬𑄟𑄴𑄝𑄢𑄴"},
			MonthsNarrowValue:        []string{"", "𑄎", "𑄜𑄬", "𑄟", "𑄃𑄬", "𑄟𑄬", "𑄎𑄪𑄚𑄴", "𑄎𑄪", "𑄃", "𑄥𑄬", "𑄃𑄧", "𑄚𑄧", "𑄓𑄨"},
			MonthsWideValue:          []string{"", "𑄎𑄚𑄪𑄠𑄢𑄨", "𑄜𑄬𑄛𑄴𑄝𑄳𑄢𑄪𑄠𑄢𑄨", "𑄟𑄢𑄴𑄌𑄧", "𑄃𑄬𑄛𑄳𑄢𑄨𑄣𑄴", "𑄟𑄬", "𑄎𑄪𑄚𑄴", "𑄎𑄪𑄣𑄭", "𑄃𑄉𑄧𑄌𑄴𑄑𑄴", "𑄥𑄬𑄛𑄴𑄑𑄬𑄟𑄴𑄝𑄧𑄢𑄴", "𑄃𑄧𑄇𑄴𑄑𑄬𑄝𑄧𑄢𑄴", "𑄚𑄧𑄞𑄬𑄟𑄴𑄝𑄧𑄢𑄴", "𑄓𑄨𑄥𑄬𑄟𑄴𑄝𑄧𑄢𑄴"},
			WeekdaysAbbreviatedValue: []string{"𑄢𑄧𑄝𑄨", "𑄥𑄧𑄟𑄴", "𑄟𑄧𑄁𑄉𑄧𑄣𑄴", "𑄝𑄪𑄖𑄴", "𑄝𑄳𑄢𑄨𑄥𑄪𑄛𑄴", "𑄥𑄪𑄇𑄴𑄇𑄮𑄢𑄴", "𑄥𑄧𑄚𑄨"},
			WeekdaysNarrowValue:      []string{"𑄢𑄧", "𑄥𑄧", "𑄟𑄧", "𑄝𑄪", "𑄝𑄳𑄢𑄨", "𑄥𑄪", "𑄥𑄧"},
			WeekdaysWideValue:        []string{"𑄢𑄧𑄝𑄨𑄝𑄢𑄴", "𑄥𑄧𑄟𑄴𑄝𑄢𑄴", "𑄟𑄧𑄁𑄉𑄧𑄣𑄴𑄝𑄢𑄴", "𑄝𑄪𑄖𑄴𑄝𑄢𑄴", "𑄝𑄳𑄢𑄨𑄥𑄪𑄛𑄴𑄝𑄢𑄴", "𑄥𑄪𑄇𑄴𑄇𑄮𑄢𑄴𑄝𑄢𑄴", "𑄥𑄧𑄚𑄨𑄝𑄢𑄴"},
			PeriodsAbbreviatedValue:  []string{"AM", "PM"},
			PeriodsNarrowValue:       []string{"AM", "PM"},
			PeriodsWideValue:         []string{"AM", "PM"},
			ErasAbbreviatedValue:     []string{"", ""},
			ErasNarrowValue:          []string{"", ""},
			ErasWideValue:            []string{"", ""},
			TimezonesValue:           map[string]string{"": "𑄛𑄳𑄠𑄢𑄉𑄪𑄠𑄬 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ACDT": "𑄃𑄧𑄌𑄴𑄑𑄳𑄢𑄬𑄣𑄨𑄠𑄧 𑄃𑄏𑄧𑄣𑄴 𑄉𑄧𑄢𑄳𑄦𑄢𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ACST": "𑄃𑄧𑄌𑄴𑄑𑄳𑄢𑄬𑄣𑄨𑄠𑄧 𑄃𑄏𑄧𑄣𑄴 𑄉𑄧𑄢𑄳𑄦𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ACT": "𑄃𑄬𑄉𑄧𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ACWDT": "ACWDT", "ACWST": "ACWST", "ADT": "𑄃𑄑𑄴𑄣𑄚𑄴𑄖𑄨𑄉𑄮𑄢𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AEDT": "𑄃𑄧𑄌𑄴𑄑𑄳𑄢𑄬𑄣𑄨𑄠𑄧 𑄛𑄪𑄉𑄬𑄘𑄨 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AEST": "𑄃𑄧𑄌𑄴𑄑𑄳𑄢𑄬𑄣𑄨𑄠𑄧 𑄛𑄪𑄉𑄬𑄘𑄨 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AFT": "𑄃𑄛𑄴𑄉𑄚𑄨𑄌𑄴𑄖𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AKDT": "𑄃𑄣𑄌𑄴𑄇 𑄘𑄨𑄚𑄮𑄃𑄣𑄮 𑄃𑄧𑄇𑄴𑄖𑄧", "AKST": "𑄃𑄣𑄌𑄴𑄇 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AMST": "𑄃𑄳𑄠𑄟𑄎𑄧𑄚𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧 𑄃𑄧𑄇𑄴𑄖𑄧", "AMT": "𑄃𑄳𑄠𑄟𑄎𑄧𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ARST": "𑄃𑄢𑄴𑄎𑄬𑄚𑄴𑄑𑄨𑄚 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ART": "𑄃𑄢𑄴𑄎𑄬𑄚𑄴𑄑𑄨𑄚 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AST": "𑄃𑄑𑄴𑄣𑄚𑄴𑄖𑄨𑄉𑄮𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AWDT": "𑄃𑄧𑄌𑄴𑄑𑄳𑄢𑄬𑄣𑄨𑄠𑄧 𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AWST": "𑄃𑄧𑄌𑄴𑄑𑄳𑄢𑄬𑄣𑄨𑄠𑄧 𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "BNT": "𑄝𑄳𑄢𑄪𑄚𑄬𑄭 𑄘𑄢𑄪𑄌𑄴𑄥𑄣𑄟𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "BOT": "𑄝𑄮𑄣𑄨𑄞𑄨𑄠 𑄃𑄧𑄇𑄴𑄖𑄧", "BRST": "𑄝𑄳𑄢𑄥𑄨𑄣𑄨𑄠 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "BRT": "𑄝𑄳𑄢𑄥𑄨𑄣𑄨𑄠 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "BST": "𑄝𑄁𑄣𑄘𑄬𑄌𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "BTT": "𑄞𑄪𑄑𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CAT": "𑄟𑄧𑄖𑄴𑄙𑄳𑄠 𑄃𑄜𑄳𑄢𑄨𑄇 𑄃𑄧𑄇𑄴𑄖𑄧", "CCT": "𑄇𑄮𑄇𑄮𑄌𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬𑄉𑄪𑄚𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CDT": "𑄃𑄏𑄧𑄣𑄴 𑄉𑄧𑄢𑄳𑄦 𑄘𑄨𑄚𑄮𑄃𑄣𑄮 𑄃𑄧𑄇𑄴𑄖𑄧", "CDT (Kuba)": "𑄇𑄨𑄃𑄪𑄝 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CHADT": "𑄌𑄳𑄠𑄗𑄟𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CHAST": "𑄌𑄳𑄠𑄗𑄟𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CLST": "𑄌𑄨𑄣𑄨 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CLT": "𑄌𑄨𑄣𑄨 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "COST": "𑄇𑄧𑄣𑄧𑄟𑄴𑄝𑄨𑄠 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "COT": "𑄇𑄧𑄣𑄧𑄟𑄴𑄝𑄨𑄠 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CST": "𑄃𑄏𑄧𑄣𑄴 𑄉𑄧𑄢𑄳𑄦 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CST (Kuba)": "𑄇𑄨𑄃𑄪𑄝 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CXT": "𑄇𑄳𑄢𑄨𑄌𑄴𑄟𑄥𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬 𑄃𑄧𑄇𑄴𑄖𑄧", "ChST": "𑄌𑄟𑄬𑄢𑄮 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "EASST": "𑄃𑄨𑄥𑄴𑄑𑄢𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "EAST": "𑄃𑄨𑄥𑄴𑄑𑄢𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "EAT": "𑄛𑄪𑄉𑄬𑄘𑄨 𑄃𑄜𑄳𑄢𑄨𑄇 𑄃𑄧𑄇𑄴𑄖𑄧", "ECT": "𑄃𑄨𑄇𑄪𑄠𑄬𑄓𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "EDT": "𑄛𑄪𑄉𑄮 𑄞𑄨𑄘𑄬𑄢𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "EGDT": "𑄛𑄪𑄉𑄬𑄘𑄨 𑄉𑄳𑄢𑄨𑄚𑄴𑄣𑄳𑄠𑄚𑄴𑄓𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "EGST": "𑄛𑄪𑄉𑄬𑄘𑄨 𑄉𑄳𑄢𑄨𑄚𑄴𑄣𑄳𑄠𑄚𑄴𑄓𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "EST": "𑄛𑄪𑄉𑄮 𑄞𑄨𑄘𑄬𑄢𑄴 𑄛𑄳𑄢𑄧𑄟𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "FKST": "𑄜𑄧𑄇𑄴𑄣𑄳𑄠𑄚𑄴𑄓𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬𑄉𑄪𑄚𑄮𑄢𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "GALT": "𑄉𑄣𑄛𑄉𑄮𑄌𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "GFT": "𑄜𑄧𑄢𑄥𑄨 𑄉𑄠𑄚 𑄃𑄧𑄇𑄴𑄖𑄧", "GMT": "𑄉𑄳𑄢𑄨𑄚𑄨𑄌𑄴 𑄟𑄨𑄚𑄴 𑄑𑄬𑄟𑄴", "GST": "𑄃𑄪𑄛𑄧𑄥𑄉𑄧𑄢𑄨𑄠𑄧 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "GYT": "𑄉𑄪𑄠𑄚 𑄃𑄧𑄇𑄴𑄖𑄧", "HADT": "𑄦𑄧𑄃𑄮𑄠𑄭-𑄃𑄣𑄬𑄃𑄪𑄖𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "HAST": "𑄦𑄧𑄃𑄮𑄠𑄭-𑄃𑄣𑄬𑄃𑄪𑄖𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "HENOMX": "𑄃𑄪𑄖𑄴𑄖𑄮𑄢𑄴 𑄛𑄧𑄏𑄨𑄟𑄴 𑄟𑄬𑄇𑄴𑄥𑄨𑄇𑄮𑄢𑄴 𑄘𑄨𑄚𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "HEOG": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄉𑄳𑄢𑄨𑄚𑄴𑄣𑄳𑄠𑄚𑄴𑄓𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "HEPMX": "𑄟𑄬𑄇𑄴𑄥𑄨𑄇𑄚𑄴 𑄛𑄳𑄢𑄧𑄥𑄚𑄴𑄖𑄧 𑄟𑄧𑄦𑄥𑄉𑄧𑄢𑄧𑄢𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "HKST": "𑄦𑄧𑄁 𑄇𑄧𑄁 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "HKT": "𑄦𑄧𑄁 𑄇𑄧𑄁 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "HNNOMX": "𑄃𑄪𑄖𑄴𑄖𑄮𑄢𑄴 𑄛𑄧𑄏𑄨𑄟𑄴 𑄟𑄬𑄇𑄴𑄥𑄨𑄇𑄮𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "HNOG": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄉𑄳𑄢𑄨𑄚𑄴𑄣𑄳𑄠𑄚𑄴𑄓𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "HNPMX": "𑄟𑄬𑄇𑄴𑄥𑄨𑄇𑄚𑄴 𑄛𑄳𑄢𑄧𑄥𑄚𑄴𑄖𑄧 𑄟𑄧𑄦𑄥𑄉𑄧𑄢𑄧𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ICT": "𑄃𑄨𑄚𑄴𑄘𑄮𑄌𑄩𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "IRDT": "𑄃𑄨𑄢𑄚𑄴 𑄘𑄨𑄚𑄬𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "IRST": "𑄃𑄨𑄢𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "IST": "𑄃𑄨𑄚𑄴𑄘𑄨𑄠𑄬 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "JDT": "𑄎𑄛𑄚𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "JST": "𑄎𑄛𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "LHDT": "𑄣𑄧𑄢𑄴𑄓𑄴 𑄦𑄤𑄬 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "LHST": "𑄣𑄧𑄢𑄴𑄓𑄴 𑄦𑄤𑄬 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MDT": "𑄦𑄨𑄣𑄧𑄧𑄱 𑄞𑄨𑄘𑄬𑄢𑄴 𑄘𑄨𑄚𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MESZ": "𑄟𑄧𑄖𑄴𑄙𑄳𑄠 𑄃𑄨𑄃𑄪𑄢𑄮𑄝𑄮𑄢𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MEZ": "𑄟𑄧𑄖𑄴𑄙𑄳𑄠 𑄃𑄨𑄃𑄪𑄢𑄮𑄝𑄮𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MST": "𑄦𑄨𑄣𑄧𑄧𑄱 𑄞𑄨𑄘𑄬𑄢𑄴 𑄛𑄳𑄢𑄧𑄟𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MVT": "𑄟𑄣𑄴𑄘𑄨𑄛𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MYT": "𑄟𑄣𑄴𑄠𑄬𑄥𑄨𑄠 𑄃𑄧𑄇𑄴𑄖𑄧", "NDT": "𑄚𑄨𑄃𑄪𑄜𑄃𑄪𑄚𑄴𑄣𑄳𑄠𑄚𑄴𑄓𑄨 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "NPT": "𑄚𑄬𑄛𑄣𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "NST": "𑄚𑄨𑄃𑄪𑄜𑄃𑄪𑄚𑄴𑄣𑄳𑄠𑄚𑄴𑄓𑄨 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "NZDT": "𑄚𑄨𑄃𑄪𑄎𑄨𑄣𑄳𑄠𑄚𑄴𑄓𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "NZST": "𑄚𑄨𑄃𑄪𑄎𑄨𑄣𑄳𑄠𑄚𑄴𑄓𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "OESZ": "𑄛𑄪𑄉𑄬𑄘𑄨 𑄃𑄨𑄃𑄪𑄢𑄮𑄝𑄮𑄢𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "OEZ": "𑄛𑄪𑄉𑄬𑄘𑄨 𑄃𑄨𑄃𑄪𑄢𑄮𑄝𑄮𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PDT": "𑄛𑄳𑄢𑄧𑄥𑄚𑄴𑄖𑄧 𑄟𑄧𑄦𑄥𑄉𑄧𑄢𑄧𑄢𑄴 𑄞𑄨𑄘𑄬𑄢𑄴 𑄘𑄨𑄚𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PKT": "𑄛𑄇𑄨𑄥𑄴𑄖𑄚𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PMDT": "𑄥𑄬𑄚𑄴𑄑𑄴 𑄛𑄨𑄠𑄬𑄢𑄴 𑄃𑄮 𑄟𑄨𑄇𑄬𑄣𑄧𑄚𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PMST": "𑄥𑄬𑄚𑄴𑄑𑄴 𑄛𑄨𑄠𑄬𑄢𑄴 𑄃𑄮 𑄟𑄨𑄇𑄬𑄣𑄧𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PST": "𑄛𑄳𑄢𑄧𑄥𑄚𑄴𑄖𑄧 𑄟𑄧𑄦𑄥𑄉𑄧𑄢𑄧𑄢𑄴 𑄞𑄨𑄘𑄬𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PYST": "𑄛𑄳𑄠𑄢𑄉𑄪𑄠𑄬 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "SAST": "𑄘𑄧𑄉𑄨𑄚𑄴 𑄃𑄜𑄳𑄢𑄨𑄇 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "SGT": "𑄥𑄨𑄁𑄉𑄛𑄪𑄢 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "SRT": "𑄥𑄪𑄢𑄨𑄚𑄟𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "TLT": "𑄛𑄪𑄉𑄬𑄘𑄨 𑄑𑄨𑄟𑄪𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "TMST": "𑄖𑄪𑄢𑄴𑄇𑄴𑄟𑄬𑄚𑄨𑄌𑄴𑄖𑄚𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "TMT": "𑄖𑄪𑄢𑄴𑄇𑄴𑄟𑄬𑄚𑄨𑄌𑄴𑄖𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "UYST": "𑄃𑄪𑄢𑄪𑄉𑄪𑄠𑄬 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "UYT": "𑄃𑄪𑄢𑄪𑄉𑄪𑄠𑄬 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "VET": "𑄞𑄬𑄚𑄬𑄎𑄪𑄠𑄬𑄣 𑄃𑄧𑄇𑄴𑄖𑄧", "WARST": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄃𑄢𑄴𑄎𑄬𑄚𑄴𑄑𑄨𑄚 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "WART": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄃𑄢𑄴𑄎𑄬𑄚𑄴𑄑𑄨𑄚 𑄛𑄳𑄢𑄧𑄟𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "WAST": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄃𑄜𑄳𑄢𑄨𑄇 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "WAT": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄃𑄜𑄳𑄢𑄨𑄇 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "WESZ": "𑄛𑄧𑄏𑄬𑄟𑄬𑄘𑄨 𑄃𑄨𑄃𑄪𑄢𑄮𑄝𑄮𑄢𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "WEZ": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄃𑄨𑄃𑄪𑄢𑄮𑄝𑄮𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "WIB": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄃𑄨𑄚𑄴𑄘𑄮𑄚𑄬𑄥𑄨𑄠 𑄃𑄧𑄇𑄴𑄖𑄧", "WIT": "𑄛𑄪𑄉𑄬𑄘𑄨 𑄃𑄨𑄚𑄴𑄘𑄮𑄚𑄬𑄥𑄨𑄠 𑄃𑄧𑄇𑄴𑄖𑄧", "WITA": "𑄃𑄏𑄧𑄣𑄴 𑄉𑄢𑄳𑄦 𑄃𑄨𑄚𑄴𑄘𑄮𑄚𑄬𑄥𑄨𑄠 𑄃𑄧𑄇𑄴𑄖𑄧", "∅∅∅": "𑄃𑄬𑄎𑄮𑄢𑄬𑄌𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧"},
		},

		locale:          "ccp",
		pluralsCardinal: nil,
		pluralsOrdinal:  nil,
		pluralsRange:    nil,
		percent:         "%",
		perMille:        "‰",
		inifinity:       "∞",

		dateFullLayout:   "EEEE, d MMMM, y",
		dateLongLayout:   "d MMMM, y",
		dateMediumLayout: "d MMM, y",
		dateShortLayout:  "d/M/yy",
		timeFullLayout:   "h:mm:ss a zzzz",
		timeLongLayout:   "h:mm:ss a z",
		timeMediumLayout: "h:mm:ss a",
		timeShortLayout:  "h:mm a",
		listPatterns:     locales.NewListPatternsFromSlice([][]string{{"{0} 𑄃𑄳𑄃 {1}", "{0}, {1}", "{0}, {1}", "{0} 𑄃𑄳𑄃 {1}"}, []string(nil)}),
		DurationSpec: &locales.DurationSpec{" ",
			"",
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"𑄥𑄧𑄖𑄧𑄇𑄴", "%v 𑄥𑄧𑄖𑄧𑄇𑄴", "%v 𑄥𑄧𑄖𑄧𑄇𑄴", ""},
				Long:   &locales.CounterFormat{"𑄃𑄬𑄇𑄴𑄥𑄧 𑄝𑄧𑄏𑄧𑄢𑄴", "%v 𑄃𑄬𑄇𑄴𑄥𑄧 𑄝𑄧𑄏𑄧𑄢𑄴", "%v 𑄃𑄬𑄇𑄴𑄥𑄧 𑄝𑄧𑄏𑄧𑄢𑄴", ""}},
			locales.DurationSpecPair{},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"𑄝𑄧𑄏𑄧𑄢𑄴", "%v 𑄝𑄧𑄏𑄧𑄢𑄴", "%v 𑄝𑄧𑄏𑄧𑄢𑄴", ""},
				Long:   &locales.CounterFormat{"𑄝𑄧𑄏𑄧𑄢𑄴", "%v 𑄝𑄧𑄏𑄧𑄢𑄴", "%v 𑄝𑄧𑄏𑄧𑄢𑄴", "%v 𑄛𑄳𑄢𑄧𑄖𑄨 𑄝𑄧𑄏𑄧𑄢𑄴"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"𑄟𑄌𑄴", "%v 𑄟𑄌𑄴", "%v 𑄟𑄌𑄴", ""},
				Long:   &locales.CounterFormat{"𑄟𑄌𑄴", "%v 𑄟𑄌𑄴", "%v 𑄟𑄌𑄴", "%v 𑄛𑄳𑄢𑄧𑄖𑄨 𑄟𑄌𑄴"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"𑄥𑄛𑄴𑄖", "%v 𑄥𑄛𑄴𑄖", "%v 𑄥𑄛𑄴𑄖", ""},
				Long:   &locales.CounterFormat{"𑄥𑄛𑄴𑄖", "%v 𑄥𑄛𑄴𑄖", "%v 𑄥𑄛𑄴𑄖", "%v 𑄛𑄳𑄢𑄧𑄖𑄨 𑄥𑄛𑄴𑄖"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"𑄘𑄨𑄚𑄴", "%v 𑄘𑄨𑄚𑄴", "%v 𑄘𑄨𑄚𑄴", ""}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"𑄊𑄧𑄚𑄴𑄑", "%v 𑄊𑄂", "%v 𑄊𑄂", ""},
				Long:   &locales.CounterFormat{"𑄊𑄧𑄚𑄴𑄘", "%v 𑄊𑄧𑄚𑄴𑄘", "%v 𑄊𑄧𑄚𑄴𑄘", "%v 𑄛𑄳𑄢𑄧𑄖𑄨 𑄊𑄧𑄚𑄴𑄘"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"𑄟𑄨𑄚𑄨𑄖𑄴", "%v 𑄟𑄨𑄂", "%v 𑄟𑄨𑄂", ""},
				Long:   &locales.CounterFormat{"𑄟𑄨𑄚𑄨𑄖𑄴", "%v 𑄟𑄨𑄚𑄨𑄖𑄴", "%v 𑄟𑄨𑄚𑄨𑄖𑄴", "%v 𑄛𑄳𑄢𑄧𑄖𑄨 𑄟𑄨𑄚𑄨𑄖𑄴"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"𑄥𑄬𑄇𑄬𑄚𑄳𑄓𑄴", "%v 𑄥𑄬𑄂", "%v 𑄥𑄬𑄂", ""},
				Long:   &locales.CounterFormat{"𑄥𑄬𑄇𑄬𑄚𑄳𑄓𑄴", "%v 𑄥𑄬𑄇𑄬𑄚𑄳𑄓𑄴", "%v 𑄥𑄬𑄇𑄬𑄚𑄳𑄓𑄴", "%v 𑄛𑄳𑄢𑄧𑄖𑄨 𑄥𑄬𑄇𑄬𑄚𑄳𑄓𑄴"}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"𑄟𑄨𑄣𑄨𑄥𑄬𑄇𑄬𑄚𑄳𑄓𑄴", "%v ms", "%v ms", ""},
				Long:   &locales.CounterFormat{"𑄟𑄨𑄣𑄨𑄥𑄬𑄇𑄬𑄚𑄳𑄓𑄴", "%v𑄟𑄨𑄣𑄨𑄥𑄬𑄇𑄬𑄚𑄳𑄓𑄴", "%v 𑄟𑄨𑄣𑄨𑄥𑄬𑄇𑄬𑄚𑄳𑄓𑄴", ""}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"μsecs", "%v μs", "%v μs", ""},
				Long:   &locales.CounterFormat{"𑄟𑄭𑄇𑄳𑄢𑄮𑄥𑄬𑄇𑄬𑄚𑄳𑄓𑄴", "%v 𑄟𑄭𑄇𑄳𑄢𑄮𑄥𑄬𑄇𑄬𑄚𑄳𑄓𑄴", "%v 𑄟𑄭𑄇𑄳𑄢𑄮𑄥𑄬𑄇𑄬𑄚𑄳𑄓𑄴", ""}},
			locales.DurationSpecPair{
				Gender: locales.Masculine,
				Short:  &locales.CounterFormat{"𑄚𑄳𑄠𑄚𑄮𑄥𑄬𑄇𑄬𑄚𑄳𑄓𑄴", "%v ns", "%v ns", ""},
				Long:   &locales.CounterFormat{"𑄚𑄳𑄠𑄚𑄮𑄥𑄬𑄇𑄬𑄚𑄳𑄓𑄴", "%v 𑄚𑄳𑄠𑄚𑄮𑄥𑄬𑄇𑄬𑄚𑄳𑄓𑄴", "%v 𑄚𑄳𑄠𑄚𑄮𑄥𑄬𑄇𑄬𑄚𑄳𑄓𑄴", ""}},
		},
		miscPatterns: locales.MiscPatterns{
			"",
			"%[1]v+",
			"",
			"%[1]v–%[2]v",
		},
	}
}

// Locale returns the current translators string locale
func (ccp *ccp) Locale() string {
	return ccp.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ccp'
func (ccp *ccp) PluralsCardinal() []locales.PluralRule {
	return ccp.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ccp'
func (ccp *ccp) PluralsOrdinal() []locales.PluralRule {
	return ccp.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ccp'
func (ccp *ccp) PluralsRange() []locales.PluralRule {
	return ccp.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ccp'
func (ccp *ccp) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ccp'
func (ccp *ccp) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ccp'
func (ccp *ccp) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// FmtDateShort returns the short date representation of 't' for 'ccp'
func (ccp *ccp) FmtDateShort(t time.Time) string {
	return locales.FormatTimeEra(ccp, ccp.dateShortLayout, t, 2)
}

// FmtDateMedium returns the medium date representation of 't' for 'ccp'
func (ccp *ccp) FmtDateMedium(t time.Time) string {
	return locales.FormatTimeEra(ccp, ccp.dateMediumLayout, t, 2)
}

// FmtDateLong returns the long date representation of 't' for 'ccp'
func (ccp *ccp) FmtDateLong(t time.Time) string {
	return locales.FormatTimeEra(ccp, ccp.dateLongLayout, t, 1)
}

// FmtDateFull returns the full date representation of 't' for 'ccp'
func (ccp *ccp) FmtDateFull(t time.Time) string {
	return locales.FormatTimeEra(ccp, ccp.dateFullLayout, t, 0)
}

// FmtTimeShort returns the short time representation of 't' for 'ccp'
func (ccp *ccp) FmtTimeShort(t time.Time) string {
	return locales.FormatTimeEra(ccp, ccp.timeShortLayout, t, 2)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ccp'
func (ccp *ccp) FmtTimeMedium(t time.Time) string {
	return locales.FormatTimeEra(ccp, ccp.timeMediumLayout, t, 2)
}

// FmtTimeLong returns the long time representation of 't' for 'ccp'
func (ccp *ccp) FmtTimeLong(t time.Time) string {
	return locales.FormatTimeEra(ccp, ccp.timeLongLayout, t, 1)
}

// FmtTimeFull returns the full time representation of 't' for 'ccp'
func (ccp *ccp) FmtTimeFull(t time.Time) string {
	return locales.FormatTimeEra(ccp, ccp.timeFullLayout, t, 0)
}

// DateFullLayout returns the full date layout for 'ccp'
func (ccp *ccp) DateFullLayout() string {
	return ccp.dateFullLayout
}

// DateLongLayout returns the long date layout for 'ccp'
func (ccp *ccp) DateLongLayout() string {
	return ccp.dateLongLayout
}

// DateMediumLayout returns the medium date layout for 'ccp'
func (ccp *ccp) DateMediumLayout() string {
	return ccp.dateMediumLayout
}

// DateShortLayout returns the short date layout for 'ccp'
func (ccp *ccp) DateShortLayout() string {
	return ccp.dateShortLayout
}

// TimeFullLayout returns the full time layout for 'ccp'
func (ccp *ccp) TimeFullLayout() string {
	return ccp.timeFullLayout
}

// TimeLongLayout returns the full long layout for 'ccp'
func (ccp *ccp) TimeLongLayout() string {
	return ccp.timeLongLayout
}

// TimeMediumLayout returns the medium time layout for 'ccp'
func (ccp *ccp) TimeMediumLayout() string {
	return ccp.timeMediumLayout
}

// TimeShortLayout returns the short time layout for 'ccp'
func (ccp *ccp) TimeShortLayout() string {
	return ccp.timeShortLayout
}

func (ccp *ccp) ListPatterns() *locales.ListPatterns {
	return ccp.listPatterns
}

func (ccp *ccp) GetDurationSpec() *locales.DurationSpec {
	return ccp.DurationSpec
}

func (ccp *ccp) GetMiscPatterns() *locales.MiscPatterns {
	return &ccp.miscPatterns
}
