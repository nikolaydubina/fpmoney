package fpmoney

const (
	_ uint8 = iota
	_AED
	_AFN
	_ALL
	_AMD
	_ANG
	_AOA
	_ARS
	_AUD
	_AWG
	_AZN
	_BAM
	_BBD
	_BDT
	_BGN
	_BHD
	_BIF
	_BMD
	_BND
	_BOB
	_BOV
	_BRL
	_BSD
	_BTN
	_BWP
	_BYN
	_BZD
	_CAD
	_CDF
	_CHE
	_CHF
	_CHW
	_CLF
	_CLP
	_CNY
	_COP
	_COU
	_CRC
	_CUP
	_CVE
	_CZK
	_DJF
	_DKK
	_DOP
	_DZD
	_EGP
	_ERN
	_ETB
	_EUR
	_FJD
	_FKP
	_GBP
	_GEL
	_GHS
	_GIP
	_GMD
	_GNF
	_GTQ
	_GYD
	_HKD
	_HNL
	_HRD
	_HRK
	_HTG
	_HUF
	_IDR
	_ILS
	_INR
	_IQD
	_IRR
	_ISK
	_JMD
	_JOD
	_JPY
	_KES
	_KGS
	_KHR
	_KMF
	_KPW
	_KRW
	_KWD
	_KYD
	_KZT
	_LAK
	_LBP
	_LKR
	_LRD
	_LSL
	_LYD
	_MAD
	_MDL
	_MGA
	_MKD
	_MMK
	_MNT
	_MOP
	_MRU
	_MUR
	_MVR
	_MWK
	_MXN
	_MXV
	_MYR
	_MZN
	_NAD
	_NGN
	_NIO
	_NOK
	_NPR
	_NZD
	_OMR
	_PAB
	_PEN
	_PGK
	_PHP
	_PKR
	_PLN
	_PYG
	_QAR
	_RON
	_RSD
	_RUB
	_RWF
	_SAR
	_SBD
	_SCR
	_SDG
	_SEK
	_SGD
	_SHP
	_SLE
	_SLL
	_SOS
	_SRD
	_SSP
	_STN
	_SVC
	_SYP
	_SZL
	_THB
	_TJS
	_TMT
	_TND
	_TOP
	_TRY
	_TTD
	_TWD
	_TZS
	_UAH
	_UGX
	_USD
	_USN
	_UYI
	_UYU
	_UYW
	_UZS
	_VED
	_VES
	_VND
	_VUV
	_WST
	_XAF
	_XAG
	_XAU
	_XBA
	_XBB
	_XBC
	_XBD
	_XCD
	_XDR
	_XOF
	_XPD
	_XPF
	_XPT
	_XSU
	_XTS
	_XUA
	_XXX
	_YER
	_ZAR
	_ZMW
	_ZWL
)

var (
	AED = Currency{_AED}
	AFN = Currency{_AFN}
	ALL = Currency{_ALL}
	AMD = Currency{_AMD}
	ANG = Currency{_ANG}
	AOA = Currency{_AOA}
	ARS = Currency{_ARS}
	AUD = Currency{_AUD}
	AWG = Currency{_AWG}
	AZN = Currency{_AZN}
	BAM = Currency{_BAM}
	BBD = Currency{_BBD}
	BDT = Currency{_BDT}
	BGN = Currency{_BGN}
	BHD = Currency{_BHD}
	BIF = Currency{_BIF}
	BMD = Currency{_BMD}
	BND = Currency{_BND}
	BOB = Currency{_BOB}
	BOV = Currency{_BOV}
	BRL = Currency{_BRL}
	BSD = Currency{_BSD}
	BTN = Currency{_BTN}
	BWP = Currency{_BWP}
	BYN = Currency{_BYN}
	BZD = Currency{_BZD}
	CAD = Currency{_CAD}
	CDF = Currency{_CDF}
	CHE = Currency{_CHE}
	CHF = Currency{_CHF}
	CHW = Currency{_CHW}
	CLF = Currency{_CLF}
	CLP = Currency{_CLP}
	CNY = Currency{_CNY}
	COP = Currency{_COP}
	COU = Currency{_COU}
	CRC = Currency{_CRC}
	CUP = Currency{_CUP}
	CVE = Currency{_CVE}
	CZK = Currency{_CZK}
	DJF = Currency{_DJF}
	DKK = Currency{_DKK}
	DOP = Currency{_DOP}
	DZD = Currency{_DZD}
	EGP = Currency{_EGP}
	ERN = Currency{_ERN}
	ETB = Currency{_ETB}
	EUR = Currency{_EUR}
	FJD = Currency{_FJD}
	FKP = Currency{_FKP}
	GBP = Currency{_GBP}
	GEL = Currency{_GEL}
	GHS = Currency{_GHS}
	GIP = Currency{_GIP}
	GMD = Currency{_GMD}
	GNF = Currency{_GNF}
	GTQ = Currency{_GTQ}
	GYD = Currency{_GYD}
	HKD = Currency{_HKD}
	HNL = Currency{_HNL}
	HRD = Currency{_HRD}
	HRK = Currency{_HRK}
	HTG = Currency{_HTG}
	HUF = Currency{_HUF}
	IDR = Currency{_IDR}
	ILS = Currency{_ILS}
	INR = Currency{_INR}
	IQD = Currency{_IQD}
	IRR = Currency{_IRR}
	ISK = Currency{_ISK}
	JMD = Currency{_JMD}
	JOD = Currency{_JOD}
	JPY = Currency{_JPY}
	KES = Currency{_KES}
	KGS = Currency{_KGS}
	KHR = Currency{_KHR}
	KMF = Currency{_KMF}
	KPW = Currency{_KPW}
	KRW = Currency{_KRW}
	KWD = Currency{_KWD}
	KYD = Currency{_KYD}
	KZT = Currency{_KZT}
	LAK = Currency{_LAK}
	LBP = Currency{_LBP}
	LKR = Currency{_LKR}
	LRD = Currency{_LRD}
	LSL = Currency{_LSL}
	LYD = Currency{_LYD}
	MAD = Currency{_MAD}
	MDL = Currency{_MDL}
	MGA = Currency{_MGA}
	MKD = Currency{_MKD}
	MMK = Currency{_MMK}
	MNT = Currency{_MNT}
	MOP = Currency{_MOP}
	MRU = Currency{_MRU}
	MUR = Currency{_MUR}
	MVR = Currency{_MVR}
	MWK = Currency{_MWK}
	MXN = Currency{_MXN}
	MXV = Currency{_MXV}
	MYR = Currency{_MYR}
	MZN = Currency{_MZN}
	NAD = Currency{_NAD}
	NGN = Currency{_NGN}
	NIO = Currency{_NIO}
	NOK = Currency{_NOK}
	NPR = Currency{_NPR}
	NZD = Currency{_NZD}
	OMR = Currency{_OMR}
	PAB = Currency{_PAB}
	PEN = Currency{_PEN}
	PGK = Currency{_PGK}
	PHP = Currency{_PHP}
	PKR = Currency{_PKR}
	PLN = Currency{_PLN}
	PYG = Currency{_PYG}
	QAR = Currency{_QAR}
	RON = Currency{_RON}
	RSD = Currency{_RSD}
	RUB = Currency{_RUB}
	RWF = Currency{_RWF}
	SAR = Currency{_SAR}
	SBD = Currency{_SBD}
	SCR = Currency{_SCR}
	SDG = Currency{_SDG}
	SEK = Currency{_SEK}
	SGD = Currency{_SGD}
	SHP = Currency{_SHP}
	SLE = Currency{_SLE}
	SLL = Currency{_SLL}
	SOS = Currency{_SOS}
	SRD = Currency{_SRD}
	SSP = Currency{_SSP}
	STN = Currency{_STN}
	SVC = Currency{_SVC}
	SYP = Currency{_SYP}
	SZL = Currency{_SZL}
	THB = Currency{_THB}
	TJS = Currency{_TJS}
	TMT = Currency{_TMT}
	TND = Currency{_TND}
	TOP = Currency{_TOP}
	TRY = Currency{_TRY}
	TTD = Currency{_TTD}
	TWD = Currency{_TWD}
	TZS = Currency{_TZS}
	UAH = Currency{_UAH}
	UGX = Currency{_UGX}
	USD = Currency{_USD}
	USN = Currency{_USN}
	UYI = Currency{_UYI}
	UYU = Currency{_UYU}
	UYW = Currency{_UYW}
	UZS = Currency{_UZS}
	VED = Currency{_VED}
	VES = Currency{_VES}
	VND = Currency{_VND}
	VUV = Currency{_VUV}
	WST = Currency{_WST}
	XAF = Currency{_XAF}
	XAG = Currency{_XAG}
	XAU = Currency{_XAU}
	XBA = Currency{_XBA}
	XBB = Currency{_XBB}
	XBC = Currency{_XBC}
	XBD = Currency{_XBD}
	XCD = Currency{_XCD}
	XDR = Currency{_XDR}
	XOF = Currency{_XOF}
	XPD = Currency{_XPD}
	XPF = Currency{_XPF}
	XPT = Currency{_XPT}
	XSU = Currency{_XSU}
	XTS = Currency{_XTS}
	XUA = Currency{_XUA}
	XXX = Currency{_XXX}
	YER = Currency{_YER}
	ZAR = Currency{_ZAR}
	ZMW = Currency{_ZMW}
	ZWL = Currency{_ZWL}
)

var currencies = map[Currency]struct {
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
	AZN: {alpha: "AZN", exponent: 2},
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
	MUR: {alpha: "MUR", exponent: 2},
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

var fromAlpha = make(map[string]Currency, len(currencies))

func init() {
	for c, v := range currencies {
		fromAlpha[v.alpha] = c
	}
}
