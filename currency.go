package fpmoney

import "errors"

// Currency is ISO 4217 without deprecated currencies.
// Zero value is undefined currency.
type Currency uint8

// Alpha returns the ISO 4217 three-letter alphabetic code.
func (c Currency) Alpha() string { return currencies[c].alpha }

// Exponent returns the decimal point location.
func (c Currency) Exponent() int { return currencies[c].exponent }

func (c Currency) String() string { return c.Alpha() }

func (c *Currency) UnmarshalText(text []byte) error {
	v, ok := fromAlpha[string(text)]
	if !ok {
		return errors.New("wrong text: " + string(text))
	}
	*c = v
	return nil
}

func (c Currency) MarshalText() (text []byte, err error) { return []byte(currencies[c].alpha), nil }

// CurrencyFromAlpha returns Currency for the three-letter alpha code.
// Or an error if it does not exist.
func CurrencyFromAlpha(alpha string) (Currency, error) {
	if c, ok := fromAlpha[alpha]; ok {
		return c, nil
	}
	return Currency(0), errors.New("no currency exists with alphabetic code " + alpha)
}

func (c Currency) scale() int64 {
	switch c.Exponent() {
	case 4:
		return 10000
	case 3:
		return 1000
	case 2:
		return 100
	default:
		return 1
	}
}

const numCurrencies = 181

const (
	_ Currency = iota
	AED
	AFN
	ALL
	AMD
	ANG
	AOA
	ARS
	AUD
	AWG
	AZN
	BAM
	BBD
	BDT
	BGN
	BHD
	BIF
	BMD
	BND
	BOB
	BOV
	BRL
	BSD
	BTN
	BWP
	BYN
	BZD
	CAD
	CDF
	CHE
	CHF
	CHW
	CLF
	CLP
	CNY
	COP
	COU
	CRC
	CUP
	CVE
	CZK
	DJF
	DKK
	DOP
	DZD
	EGP
	ERN
	ETB
	EUR
	FJD
	FKP
	GBP
	GEL
	GHS
	GIP
	GMD
	GNF
	GTQ
	GYD
	HKD
	HNL
	HRD
	HRK
	HTG
	HUF
	IDR
	ILS
	INR
	IQD
	IRR
	ISK
	JMD
	JOD
	JPY
	KES
	KGS
	KHR
	KMF
	KPW
	KRW
	KWD
	KYD
	KZT
	LAK
	LBP
	LKR
	LRD
	LSL
	LYD
	MAD
	MDL
	MGA
	MKD
	MMK
	MNT
	MOP
	MRU
	MUR
	MVR
	MWK
	MXN
	MXV
	MYR
	MZN
	NAD
	NGN
	NIO
	NOK
	NPR
	NZD
	OMR
	PAB
	PEN
	PGK
	PHP
	PKR
	PLN
	PYG
	QAR
	RON
	RSD
	RUB
	RWF
	SAR
	SBD
	SCR
	SDG
	SEK
	SGD
	SHP
	SLE
	SLL
	SOS
	SRD
	SSP
	STN
	SVC
	SYP
	SZL
	THB
	TJS
	TMT
	TND
	TOP
	TRY
	TTD
	TWD
	TZS
	UAH
	UGX
	USD
	USN
	UYI
	UYU
	UYW
	UZS
	VED
	VES
	VND
	VUV
	WST
	XAF
	XAG
	XAU
	XBA
	XBB
	XBC
	XBD
	XCD
	XDR
	XOF
	XPD
	XPF
	XPT
	XSU
	XTS
	XUA
	XXX
	YER
	ZAR
	ZMW
	ZWL
)

var currencies = [...]struct {
	alpha    string
	exponent int
}{
	AED: {alpha: "AED", exponent: 2},
	AFN: {alpha: "AFN", exponent: 2},
	ALL: {alpha: "ALL", exponent: 2},
	AMD: {alpha: "AMD", exponent: 2},
	ANG: {alpha: "ANG", exponent: 2},
	AOA: {alpha: "AOA", exponent: 2},
	ARS: {alpha: "ARS", exponent: 2},
	AUD: {alpha: "AUD", exponent: 2},
	AWG: {alpha: "AWG", exponent: 2},
	BAM: {alpha: "BAM", exponent: 2},
	BBD: {alpha: "BBD", exponent: 2},
	BDT: {alpha: "BDT", exponent: 2},
	BGN: {alpha: "BGN", exponent: 2},
	BHD: {alpha: "BHD", exponent: 3},
	BIF: {alpha: "BIF", exponent: 0},
	BMD: {alpha: "BMD", exponent: 2},
	BND: {alpha: "BND", exponent: 2},
	BOB: {alpha: "BOB", exponent: 2},
	BOV: {alpha: "BOV", exponent: 2},
	BRL: {alpha: "BRL", exponent: 2},
	BSD: {alpha: "BSD", exponent: 2},
	BTN: {alpha: "BTN", exponent: 2},
	BWP: {alpha: "BWP", exponent: 2},
	BYN: {alpha: "BYN", exponent: 2},
	BZD: {alpha: "BZD", exponent: 2},
	CAD: {alpha: "CAD", exponent: 2},
	CDF: {alpha: "CDF", exponent: 2},
	CHE: {alpha: "CHE", exponent: 2},
	CHF: {alpha: "CHF", exponent: 2},
	CHW: {alpha: "CHW", exponent: 2},
	CLF: {alpha: "CLF", exponent: 4},
	CLP: {alpha: "CLP", exponent: 0},
	CNY: {alpha: "CNY", exponent: 2},
	COP: {alpha: "COP", exponent: 2},
	COU: {alpha: "COU", exponent: 2},
	CRC: {alpha: "CRC", exponent: 2},
	CUP: {alpha: "CUP", exponent: 2},
	CVE: {alpha: "CVE", exponent: 2},
	CZK: {alpha: "CZK", exponent: 2},
	DJF: {alpha: "DJF", exponent: 0},
	DKK: {alpha: "DKK", exponent: 2},
	DOP: {alpha: "DOP", exponent: 2},
	DZD: {alpha: "DZD", exponent: 2},
	EGP: {alpha: "EGP", exponent: 2},
	ERN: {alpha: "ERN", exponent: 2},
	ETB: {alpha: "ETB", exponent: 2},
	EUR: {alpha: "EUR", exponent: 2},
	FJD: {alpha: "FJD", exponent: 2},
	FKP: {alpha: "FKP", exponent: 2},
	GBP: {alpha: "GBP", exponent: 2},
	GEL: {alpha: "GEL", exponent: 2},
	GHS: {alpha: "GHS", exponent: 2},
	GIP: {alpha: "GIP", exponent: 2},
	GMD: {alpha: "GMD", exponent: 2},
	GNF: {alpha: "GNF", exponent: 0},
	GTQ: {alpha: "GTQ", exponent: 2},
	GYD: {alpha: "GYD", exponent: 2},
	HKD: {alpha: "HKD", exponent: 2},
	HNL: {alpha: "HNL", exponent: 2},
	HRD: {alpha: "HRD", exponent: 0},
	HRK: {alpha: "HRK", exponent: 2},
	HTG: {alpha: "HTG", exponent: 2},
	HUF: {alpha: "HUF", exponent: 2},
	IDR: {alpha: "IDR", exponent: 2},
	ILS: {alpha: "ILS", exponent: 2},
	INR: {alpha: "INR", exponent: 2},
	IQD: {alpha: "IQD", exponent: 3},
	IRR: {alpha: "IRR", exponent: 2},
	ISK: {alpha: "ISK", exponent: 0},
	JMD: {alpha: "JMD", exponent: 2},
	JOD: {alpha: "JOD", exponent: 3},
	JPY: {alpha: "JPY", exponent: 0},
	KES: {alpha: "KES", exponent: 2},
	KGS: {alpha: "KGS", exponent: 2},
	KHR: {alpha: "KHR", exponent: 2},
	KMF: {alpha: "KMF", exponent: 0},
	KPW: {alpha: "KPW", exponent: 2},
	KRW: {alpha: "KRW", exponent: 0},
	KWD: {alpha: "KWD", exponent: 3},
	KYD: {alpha: "KYD", exponent: 2},
	KZT: {alpha: "KZT", exponent: 2},
	LAK: {alpha: "LAK", exponent: 2},
	LBP: {alpha: "LBP", exponent: 2},
	LKR: {alpha: "LKR", exponent: 2},
	LRD: {alpha: "LRD", exponent: 2},
	LSL: {alpha: "LSL", exponent: 2},
	LYD: {alpha: "LYD", exponent: 3},
	MAD: {alpha: "MAD", exponent: 2},
	MDL: {alpha: "MDL", exponent: 2},
	MGA: {alpha: "MGA", exponent: 2},
	MKD: {alpha: "MKD", exponent: 2},
	MMK: {alpha: "MMK", exponent: 2},
	MNT: {alpha: "MNT", exponent: 2},
	MOP: {alpha: "MOP", exponent: 2},
	MRU: {alpha: "MRU", exponent: 2},
	MVR: {alpha: "MVR", exponent: 2},
	MWK: {alpha: "MWK", exponent: 2},
	MXN: {alpha: "MXN", exponent: 2},
	MXV: {alpha: "MXV", exponent: 2},
	MYR: {alpha: "MYR", exponent: 2},
	MZN: {alpha: "MZN", exponent: 2},
	NAD: {alpha: "NAD", exponent: 2},
	NGN: {alpha: "NGN", exponent: 2},
	NIO: {alpha: "NIO", exponent: 2},
	NOK: {alpha: "NOK", exponent: 2},
	NPR: {alpha: "NPR", exponent: 2},
	NZD: {alpha: "NZD", exponent: 2},
	OMR: {alpha: "OMR", exponent: 3},
	PAB: {alpha: "PAB", exponent: 2},
	PEN: {alpha: "PEN", exponent: 2},
	PGK: {alpha: "PGK", exponent: 2},
	PHP: {alpha: "PHP", exponent: 2},
	PKR: {alpha: "PKR", exponent: 2},
	PLN: {alpha: "PLN", exponent: 2},
	PYG: {alpha: "PYG", exponent: 0},
	QAR: {alpha: "QAR", exponent: 2},
	RON: {alpha: "RON", exponent: 2},
	RSD: {alpha: "RSD", exponent: 2},
	RUB: {alpha: "RUB", exponent: 2},
	RWF: {alpha: "RWF", exponent: 0},
	SAR: {alpha: "SAR", exponent: 2},
	SBD: {alpha: "SBD", exponent: 2},
	SCR: {alpha: "SCR", exponent: 2},
	SDG: {alpha: "SDG", exponent: 2},
	SEK: {alpha: "SEK", exponent: 2},
	SGD: {alpha: "SGD", exponent: 2},
	SHP: {alpha: "SHP", exponent: 2},
	SLE: {alpha: "SLE", exponent: 2},
	SLL: {alpha: "SLL", exponent: 2},
	SOS: {alpha: "SOS", exponent: 2},
	SRD: {alpha: "SRD", exponent: 2},
	SSP: {alpha: "SSP", exponent: 2},
	STN: {alpha: "STN", exponent: 2},
	SVC: {alpha: "SVC", exponent: 2},
	SYP: {alpha: "SYP", exponent: 2},
	SZL: {alpha: "SZL", exponent: 2},
	THB: {alpha: "THB", exponent: 2},
	TJS: {alpha: "TJS", exponent: 2},
	TMT: {alpha: "TMT", exponent: 2},
	TND: {alpha: "TND", exponent: 3},
	TOP: {alpha: "TOP", exponent: 2},
	TRY: {alpha: "TRY", exponent: 2},
	TTD: {alpha: "TTD", exponent: 2},
	TWD: {alpha: "TWD", exponent: 2},
	TZS: {alpha: "TZS", exponent: 2},
	UAH: {alpha: "UAH", exponent: 2},
	UGX: {alpha: "UGX", exponent: 0},
	USD: {alpha: "USD", exponent: 2},
	USN: {alpha: "USN", exponent: 2},
	UYI: {alpha: "UYI", exponent: 0},
	UYU: {alpha: "UYU", exponent: 2},
	UYW: {alpha: "UYW", exponent: 4},
	UZS: {alpha: "UZS", exponent: 2},
	VED: {alpha: "VED", exponent: 2},
	VES: {alpha: "VES", exponent: 2},
	VND: {alpha: "VND", exponent: 0},
	VUV: {alpha: "VUV", exponent: 0},
	WST: {alpha: "WST", exponent: 2},
	XAF: {alpha: "XAF", exponent: 0},
	XAG: {alpha: "XAG", exponent: 0},
	XAU: {alpha: "XAU", exponent: 0},
	XBA: {alpha: "XBA", exponent: 0},
	XBB: {alpha: "XBB", exponent: 0},
	XBC: {alpha: "XBC", exponent: 0},
	XBD: {alpha: "XBD", exponent: 0},
	XCD: {alpha: "XCD", exponent: 2},
	XDR: {alpha: "XDR", exponent: 0},
	XOF: {alpha: "XOF", exponent: 0},
	XPD: {alpha: "XPD", exponent: 0},
	XPF: {alpha: "XPF", exponent: 0},
	XPT: {alpha: "XPT", exponent: 0},
	XSU: {alpha: "XSU", exponent: 0},
	XTS: {alpha: "XTS", exponent: 0},
	XUA: {alpha: "XUA", exponent: 0},
	XXX: {alpha: "XXX", exponent: 0},
	YER: {alpha: "YER", exponent: 2},
	ZAR: {alpha: "ZAR", exponent: 2},
	ZMW: {alpha: "ZMW", exponent: 2},
	ZWL: {alpha: "ZWL", exponent: 2},
}

var fromAlpha = map[string]Currency{
	"AED": AED,
	"AFN": AFN,
	"ALL": ALL,
	"AMD": AMD,
	"ANG": ANG,
	"AOA": AOA,
	"ARS": ARS,
	"AUD": AUD,
	"AWG": AWG,
	"AZN": AZN,
	"BAM": BAM,
	"BBD": BBD,
	"BDT": BDT,
	"BGN": BGN,
	"BHD": BHD,
	"BIF": BIF,
	"BMD": BMD,
	"BND": BND,
	"BOB": BOB,
	"BOV": BOV,
	"BRL": BRL,
	"BSD": BSD,
	"BTN": BTN,
	"BWP": BWP,
	"BYN": BYN,
	"BZD": BZD,
	"CAD": CAD,
	"CDF": CDF,
	"CHE": CHE,
	"CHF": CHF,
	"CHW": CHW,
	"CLF": CLF,
	"CLP": CLP,
	"CNY": CNY,
	"COP": COP,
	"COU": COU,
	"CRC": CRC,
	"CUP": CUP,
	"CVE": CVE,
	"CZK": CZK,
	"DJF": DJF,
	"DKK": DKK,
	"DOP": DOP,
	"DZD": DZD,
	"EGP": EGP,
	"ERN": ERN,
	"ETB": ETB,
	"EUR": EUR,
	"FJD": FJD,
	"FKP": FKP,
	"GBP": GBP,
	"GEL": GEL,
	"GHS": GHS,
	"GIP": GIP,
	"GMD": GMD,
	"GNF": GNF,
	"GTQ": GTQ,
	"GYD": GYD,
	"HKD": HKD,
	"HNL": HNL,
	"HRD": HRD,
	"HRK": HRK,
	"HTG": HTG,
	"HUF": HUF,
	"IDR": IDR,
	"ILS": ILS,
	"INR": INR,
	"IQD": IQD,
	"IRR": IRR,
	"ISK": ISK,
	"JMD": JMD,
	"JOD": JOD,
	"JPY": JPY,
	"KES": KES,
	"KGS": KGS,
	"KHR": KHR,
	"KMF": KMF,
	"KPW": KPW,
	"KRW": KRW,
	"KWD": KWD,
	"KYD": KYD,
	"KZT": KZT,
	"LAK": LAK,
	"LBP": LBP,
	"LKR": LKR,
	"LRD": LRD,
	"LSL": LSL,
	"LYD": LYD,
	"MAD": MAD,
	"MDL": MDL,
	"MGA": MGA,
	"MKD": MKD,
	"MMK": MMK,
	"MNT": MNT,
	"MOP": MOP,
	"MRU": MRU,
	"MUR": MUR,
	"MVR": MVR,
	"MWK": MWK,
	"MXN": MXN,
	"MXV": MXV,
	"MYR": MYR,
	"MZN": MZN,
	"NAD": NAD,
	"NGN": NGN,
	"NIO": NIO,
	"NOK": NOK,
	"NPR": NPR,
	"NZD": NZD,
	"OMR": OMR,
	"PAB": PAB,
	"PEN": PEN,
	"PGK": PGK,
	"PHP": PHP,
	"PKR": PKR,
	"PLN": PLN,
	"PYG": PYG,
	"QAR": QAR,
	"RON": RON,
	"RSD": RSD,
	"RUB": RUB,
	"RWF": RWF,
	"SAR": SAR,
	"SBD": SBD,
	"SCR": SCR,
	"SDG": SDG,
	"SEK": SEK,
	"SGD": SGD,
	"SHP": SHP,
	"SLE": SLE,
	"SLL": SLL,
	"SOS": SOS,
	"SRD": SRD,
	"SSP": SSP,
	"STN": STN,
	"SVC": SVC,
	"SYP": SYP,
	"SZL": SZL,
	"THB": THB,
	"TJS": TJS,
	"TMT": TMT,
	"TND": TND,
	"TOP": TOP,
	"TRY": TRY,
	"TTD": TTD,
	"TWD": TWD,
	"TZS": TZS,
	"UAH": UAH,
	"UGX": UGX,
	"USD": USD,
	"USN": USN,
	"UYI": UYI,
	"UYU": UYU,
	"UYW": UYW,
	"UZS": UZS,
	"VED": VED,
	"VES": VES,
	"VND": VND,
	"VUV": VUV,
	"WST": WST,
	"XAF": XAF,
	"XAG": XAG,
	"XAU": XAU,
	"XBA": XBA,
	"XBB": XBB,
	"XBC": XBC,
	"XBD": XBD,
	"XCD": XCD,
	"XDR": XDR,
	"XOF": XOF,
	"XPD": XPD,
	"XPF": XPF,
	"XPT": XPT,
	"XSU": XSU,
	"XTS": XTS,
	"XUA": XUA,
	"XXX": XXX,
	"YER": YER,
	"ZAR": ZAR,
	"ZMW": ZMW,
	"ZWL": ZWL,
}
