package fpmoney

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
	alpha      string
	alphaBytes []byte
	exponent   int
}{
	AED: {alpha: "AED", alphaBytes: []byte("AED"), exponent: 2},
	AFN: {alpha: "AFN", alphaBytes: []byte("AFN"), exponent: 2},
	ALL: {alpha: "ALL", alphaBytes: []byte("ALL"), exponent: 2},
	AMD: {alpha: "AMD", alphaBytes: []byte("AMD"), exponent: 2},
	ANG: {alpha: "ANG", alphaBytes: []byte("ANG"), exponent: 2},
	AOA: {alpha: "AOA", alphaBytes: []byte("AOA"), exponent: 2},
	ARS: {alpha: "ARS", alphaBytes: []byte("ARS"), exponent: 2},
	AUD: {alpha: "AUD", alphaBytes: []byte("AUD"), exponent: 2},
	AWG: {alpha: "AWG", alphaBytes: []byte("AWG"), exponent: 2},
	BAM: {alpha: "BAM", alphaBytes: []byte("BAM"), exponent: 2},
	BBD: {alpha: "BBD", alphaBytes: []byte("BBD"), exponent: 2},
	BDT: {alpha: "BDT", alphaBytes: []byte("BDT"), exponent: 2},
	BGN: {alpha: "BGN", alphaBytes: []byte("BGN"), exponent: 2},
	BHD: {alpha: "BHD", alphaBytes: []byte("BHD"), exponent: 3},
	BIF: {alpha: "BIF", alphaBytes: []byte("BIF"), exponent: 0},
	BMD: {alpha: "BMD", alphaBytes: []byte("BMD"), exponent: 2},
	BND: {alpha: "BND", alphaBytes: []byte("BND"), exponent: 2},
	BOB: {alpha: "BOB", alphaBytes: []byte("BOB"), exponent: 2},
	BOV: {alpha: "BOV", alphaBytes: []byte("BOV"), exponent: 2},
	BRL: {alpha: "BRL", alphaBytes: []byte("BRL"), exponent: 2},
	BSD: {alpha: "BSD", alphaBytes: []byte("BSD"), exponent: 2},
	BTN: {alpha: "BTN", alphaBytes: []byte("BTN"), exponent: 2},
	BWP: {alpha: "BWP", alphaBytes: []byte("BWP"), exponent: 2},
	BYN: {alpha: "BYN", alphaBytes: []byte("BYN"), exponent: 2},
	BZD: {alpha: "BZD", alphaBytes: []byte("BZD"), exponent: 2},
	CAD: {alpha: "CAD", alphaBytes: []byte("CAD"), exponent: 2},
	CDF: {alpha: "CDF", alphaBytes: []byte("CDF"), exponent: 2},
	CHE: {alpha: "CHE", alphaBytes: []byte("CHE"), exponent: 2},
	CHF: {alpha: "CHF", alphaBytes: []byte("CHF"), exponent: 2},
	CHW: {alpha: "CHW", alphaBytes: []byte("CHW"), exponent: 2},
	CLF: {alpha: "CLF", alphaBytes: []byte("CLF"), exponent: 4},
	CLP: {alpha: "CLP", alphaBytes: []byte("CLP"), exponent: 0},
	CNY: {alpha: "CNY", alphaBytes: []byte("CNY"), exponent: 2},
	COP: {alpha: "COP", alphaBytes: []byte("COP"), exponent: 2},
	COU: {alpha: "COU", alphaBytes: []byte("COU"), exponent: 2},
	CRC: {alpha: "CRC", alphaBytes: []byte("CRC"), exponent: 2},
	CUP: {alpha: "CUP", alphaBytes: []byte("CUP"), exponent: 2},
	CVE: {alpha: "CVE", alphaBytes: []byte("CVE"), exponent: 2},
	CZK: {alpha: "CZK", alphaBytes: []byte("CZK"), exponent: 2},
	DJF: {alpha: "DJF", alphaBytes: []byte("DJF"), exponent: 0},
	DKK: {alpha: "DKK", alphaBytes: []byte("DKK"), exponent: 2},
	DOP: {alpha: "DOP", alphaBytes: []byte("DOP"), exponent: 2},
	DZD: {alpha: "DZD", alphaBytes: []byte("DZD"), exponent: 2},
	EGP: {alpha: "EGP", alphaBytes: []byte("EGP"), exponent: 2},
	ERN: {alpha: "ERN", alphaBytes: []byte("ERN"), exponent: 2},
	ETB: {alpha: "ETB", alphaBytes: []byte("ETB"), exponent: 2},
	EUR: {alpha: "EUR", alphaBytes: []byte("EUR"), exponent: 2},
	FJD: {alpha: "FJD", alphaBytes: []byte("FJD"), exponent: 2},
	FKP: {alpha: "FKP", alphaBytes: []byte("FKP"), exponent: 2},
	GBP: {alpha: "GBP", alphaBytes: []byte("GBP"), exponent: 2},
	GEL: {alpha: "GEL", alphaBytes: []byte("GEL"), exponent: 2},
	GHS: {alpha: "GHS", alphaBytes: []byte("GHS"), exponent: 2},
	GIP: {alpha: "GIP", alphaBytes: []byte("GIP"), exponent: 2},
	GMD: {alpha: "GMD", alphaBytes: []byte("GMD"), exponent: 2},
	GNF: {alpha: "GNF", alphaBytes: []byte("GNF"), exponent: 0},
	GTQ: {alpha: "GTQ", alphaBytes: []byte("GTQ"), exponent: 2},
	GYD: {alpha: "GYD", alphaBytes: []byte("GYD"), exponent: 2},
	HKD: {alpha: "HKD", alphaBytes: []byte("HKD"), exponent: 2},
	HNL: {alpha: "HNL", alphaBytes: []byte("HNL"), exponent: 2},
	HRD: {alpha: "HRD", alphaBytes: []byte("HRD"), exponent: 0},
	HRK: {alpha: "HRK", alphaBytes: []byte("HRK"), exponent: 2},
	HTG: {alpha: "HTG", alphaBytes: []byte("HTG"), exponent: 2},
	HUF: {alpha: "HUF", alphaBytes: []byte("HUF"), exponent: 2},
	IDR: {alpha: "IDR", alphaBytes: []byte("IDR"), exponent: 2},
	ILS: {alpha: "ILS", alphaBytes: []byte("ILS"), exponent: 2},
	INR: {alpha: "INR", alphaBytes: []byte("INR"), exponent: 2},
	IQD: {alpha: "IQD", alphaBytes: []byte("IQD"), exponent: 3},
	IRR: {alpha: "IRR", alphaBytes: []byte("IRR"), exponent: 2},
	ISK: {alpha: "ISK", alphaBytes: []byte("ISK"), exponent: 0},
	JMD: {alpha: "JMD", alphaBytes: []byte("JMD"), exponent: 2},
	JOD: {alpha: "JOD", alphaBytes: []byte("JOD"), exponent: 3},
	JPY: {alpha: "JPY", alphaBytes: []byte("JPY"), exponent: 0},
	KES: {alpha: "KES", alphaBytes: []byte("KES"), exponent: 2},
	KGS: {alpha: "KGS", alphaBytes: []byte("KGS"), exponent: 2},
	KHR: {alpha: "KHR", alphaBytes: []byte("KHR"), exponent: 2},
	KMF: {alpha: "KMF", alphaBytes: []byte("KMF"), exponent: 0},
	KPW: {alpha: "KPW", alphaBytes: []byte("KPW"), exponent: 2},
	KRW: {alpha: "KRW", alphaBytes: []byte("KRW"), exponent: 0},
	KWD: {alpha: "KWD", alphaBytes: []byte("KWD"), exponent: 3},
	KYD: {alpha: "KYD", alphaBytes: []byte("KYD"), exponent: 2},
	KZT: {alpha: "KZT", alphaBytes: []byte("KZT"), exponent: 2},
	LAK: {alpha: "LAK", alphaBytes: []byte("LAK"), exponent: 2},
	LBP: {alpha: "LBP", alphaBytes: []byte("LBP"), exponent: 2},
	LKR: {alpha: "LKR", alphaBytes: []byte("LKR"), exponent: 2},
	LRD: {alpha: "LRD", alphaBytes: []byte("LRD"), exponent: 2},
	LSL: {alpha: "LSL", alphaBytes: []byte("LSL"), exponent: 2},
	LYD: {alpha: "LYD", alphaBytes: []byte("LYD"), exponent: 3},
	MAD: {alpha: "MAD", alphaBytes: []byte("MAD"), exponent: 2},
	MDL: {alpha: "MDL", alphaBytes: []byte("MDL"), exponent: 2},
	MGA: {alpha: "MGA", alphaBytes: []byte("MGA"), exponent: 2},
	MKD: {alpha: "MKD", alphaBytes: []byte("MKD"), exponent: 2},
	MMK: {alpha: "MMK", alphaBytes: []byte("MMK"), exponent: 2},
	MNT: {alpha: "MNT", alphaBytes: []byte("MNT"), exponent: 2},
	MOP: {alpha: "MOP", alphaBytes: []byte("MOP"), exponent: 2},
	MRU: {alpha: "MRU", alphaBytes: []byte("MRU"), exponent: 2},
	MVR: {alpha: "MVR", alphaBytes: []byte("MVR"), exponent: 2},
	MWK: {alpha: "MWK", alphaBytes: []byte("MWK"), exponent: 2},
	MXN: {alpha: "MXN", alphaBytes: []byte("MXN"), exponent: 2},
	MXV: {alpha: "MXV", alphaBytes: []byte("MXV"), exponent: 2},
	MYR: {alpha: "MYR", alphaBytes: []byte("MYR"), exponent: 2},
	MZN: {alpha: "MZN", alphaBytes: []byte("MZN"), exponent: 2},
	NAD: {alpha: "NAD", alphaBytes: []byte("NAD"), exponent: 2},
	NGN: {alpha: "NGN", alphaBytes: []byte("NGN"), exponent: 2},
	NIO: {alpha: "NIO", alphaBytes: []byte("NIO"), exponent: 2},
	NOK: {alpha: "NOK", alphaBytes: []byte("NOK"), exponent: 2},
	NPR: {alpha: "NPR", alphaBytes: []byte("NPR"), exponent: 2},
	NZD: {alpha: "NZD", alphaBytes: []byte("NZD"), exponent: 2},
	OMR: {alpha: "OMR", alphaBytes: []byte("OMR"), exponent: 3},
	PAB: {alpha: "PAB", alphaBytes: []byte("PAB"), exponent: 2},
	PEN: {alpha: "PEN", alphaBytes: []byte("PEN"), exponent: 2},
	PGK: {alpha: "PGK", alphaBytes: []byte("PGK"), exponent: 2},
	PHP: {alpha: "PHP", alphaBytes: []byte("PHP"), exponent: 2},
	PKR: {alpha: "PKR", alphaBytes: []byte("PKR"), exponent: 2},
	PLN: {alpha: "PLN", alphaBytes: []byte("PLN"), exponent: 2},
	PYG: {alpha: "PYG", alphaBytes: []byte("PYG"), exponent: 0},
	QAR: {alpha: "QAR", alphaBytes: []byte("QAR"), exponent: 2},
	RON: {alpha: "RON", alphaBytes: []byte("RON"), exponent: 2},
	RSD: {alpha: "RSD", alphaBytes: []byte("RSD"), exponent: 2},
	RUB: {alpha: "RUB", alphaBytes: []byte("RUB"), exponent: 2},
	RWF: {alpha: "RWF", alphaBytes: []byte("RWF"), exponent: 0},
	SAR: {alpha: "SAR", alphaBytes: []byte("SAR"), exponent: 2},
	SBD: {alpha: "SBD", alphaBytes: []byte("SBD"), exponent: 2},
	SCR: {alpha: "SCR", alphaBytes: []byte("SCR"), exponent: 2},
	SDG: {alpha: "SDG", alphaBytes: []byte("SDG"), exponent: 2},
	SEK: {alpha: "SEK", alphaBytes: []byte("SEK"), exponent: 2},
	SGD: {alpha: "SGD", alphaBytes: []byte("SGD"), exponent: 2},
	SHP: {alpha: "SHP", alphaBytes: []byte("SHP"), exponent: 2},
	SLE: {alpha: "SLE", alphaBytes: []byte("SLE"), exponent: 2},
	SLL: {alpha: "SLL", alphaBytes: []byte("SLL"), exponent: 2},
	SOS: {alpha: "SOS", alphaBytes: []byte("SOS"), exponent: 2},
	SRD: {alpha: "SRD", alphaBytes: []byte("SRD"), exponent: 2},
	SSP: {alpha: "SSP", alphaBytes: []byte("SSP"), exponent: 2},
	STN: {alpha: "STN", alphaBytes: []byte("STN"), exponent: 2},
	SVC: {alpha: "SVC", alphaBytes: []byte("SVC"), exponent: 2},
	SYP: {alpha: "SYP", alphaBytes: []byte("SYP"), exponent: 2},
	SZL: {alpha: "SZL", alphaBytes: []byte("SZL"), exponent: 2},
	THB: {alpha: "THB", alphaBytes: []byte("THB"), exponent: 2},
	TJS: {alpha: "TJS", alphaBytes: []byte("TJS"), exponent: 2},
	TMT: {alpha: "TMT", alphaBytes: []byte("TMT"), exponent: 2},
	TND: {alpha: "TND", alphaBytes: []byte("TND"), exponent: 3},
	TOP: {alpha: "TOP", alphaBytes: []byte("TOP"), exponent: 2},
	TRY: {alpha: "TRY", alphaBytes: []byte("TRY"), exponent: 2},
	TTD: {alpha: "TTD", alphaBytes: []byte("TTD"), exponent: 2},
	TWD: {alpha: "TWD", alphaBytes: []byte("TWD"), exponent: 2},
	TZS: {alpha: "TZS", alphaBytes: []byte("TZS"), exponent: 2},
	UAH: {alpha: "UAH", alphaBytes: []byte("UAH"), exponent: 2},
	UGX: {alpha: "UGX", alphaBytes: []byte("UGX"), exponent: 0},
	USD: {alpha: "USD", alphaBytes: []byte("USD"), exponent: 2},
	USN: {alpha: "USN", alphaBytes: []byte("USN"), exponent: 2},
	UYI: {alpha: "UYI", alphaBytes: []byte("UYI"), exponent: 0},
	UYU: {alpha: "UYU", alphaBytes: []byte("UYU"), exponent: 2},
	UYW: {alpha: "UYW", alphaBytes: []byte("UYW"), exponent: 4},
	UZS: {alpha: "UZS", alphaBytes: []byte("UZS"), exponent: 2},
	VED: {alpha: "VED", alphaBytes: []byte("VED"), exponent: 2},
	VES: {alpha: "VES", alphaBytes: []byte("VES"), exponent: 2},
	VND: {alpha: "VND", alphaBytes: []byte("VND"), exponent: 0},
	VUV: {alpha: "VUV", alphaBytes: []byte("VUV"), exponent: 0},
	WST: {alpha: "WST", alphaBytes: []byte("WST"), exponent: 2},
	XAF: {alpha: "XAF", alphaBytes: []byte("XAF"), exponent: 0},
	XAG: {alpha: "XAG", alphaBytes: []byte("XAG"), exponent: 0},
	XAU: {alpha: "XAU", alphaBytes: []byte("XAU"), exponent: 0},
	XBA: {alpha: "XBA", alphaBytes: []byte("XBA"), exponent: 0},
	XBB: {alpha: "XBB", alphaBytes: []byte("XBB"), exponent: 0},
	XBC: {alpha: "XBC", alphaBytes: []byte("XBC"), exponent: 0},
	XBD: {alpha: "XBD", alphaBytes: []byte("XBD"), exponent: 0},
	XCD: {alpha: "XCD", alphaBytes: []byte("XCD"), exponent: 2},
	XDR: {alpha: "XDR", alphaBytes: []byte("XDR"), exponent: 0},
	XOF: {alpha: "XOF", alphaBytes: []byte("XOF"), exponent: 0},
	XPD: {alpha: "XPD", alphaBytes: []byte("XPD"), exponent: 0},
	XPF: {alpha: "XPF", alphaBytes: []byte("XPF"), exponent: 0},
	XPT: {alpha: "XPT", alphaBytes: []byte("XPT"), exponent: 0},
	XSU: {alpha: "XSU", alphaBytes: []byte("XSU"), exponent: 0},
	XTS: {alpha: "XTS", alphaBytes: []byte("XTS"), exponent: 0},
	XUA: {alpha: "XUA", alphaBytes: []byte("XUA"), exponent: 0},
	XXX: {alpha: "XXX", alphaBytes: []byte("XXX"), exponent: 0},
	YER: {alpha: "YER", alphaBytes: []byte("YER"), exponent: 2},
	ZAR: {alpha: "ZAR", alphaBytes: []byte("ZAR"), exponent: 2},
	ZMW: {alpha: "ZMW", alphaBytes: []byte("ZMW"), exponent: 2},
	ZWL: {alpha: "ZWL", alphaBytes: []byte("ZWL"), exponent: 2},
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
