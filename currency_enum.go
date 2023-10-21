package fpmoney

const (
	_ uint8 = iota
	iAED
	iAFN
	iALL
	iAMD
	iANG
	iAOA
	iARS
	iAUD
	iAWG
	iAZN
	iBAM
	iBBD
	iBDT
	iBGN
	iBHD
	iBIF
	iBMD
	iBND
	iBOB
	iBOV
	iBRL
	iBSD
	iBTN
	iBWP
	iBYN
	iBZD
	iCAD
	iCDF
	iCHE
	iCHF
	iCHW
	iCLF
	iCLP
	iCNY
	iCOP
	iCOU
	iCRC
	iCUP
	iCVE
	iCZK
	iDJF
	iDKK
	iDOP
	iDZD
	iEGP
	iERN
	iETB
	iEUR
	iFJD
	iFKP
	iGBP
	iGEL
	iGHS
	iGIP
	iGMD
	iGNF
	iGTQ
	iGYD
	iHKD
	iHNL
	iHRD
	iHRK
	iHTG
	iHUF
	iIDR
	iILS
	iINR
	iIQD
	iIRR
	iISK
	iJMD
	iJOD
	iJPY
	iKES
	iKGS
	iKHR
	iKMF
	iKPW
	iKRW
	iKWD
	iKYD
	iKZT
	iLAK
	iLBP
	iLKR
	iLRD
	iLSL
	iLYD
	iMAD
	iMDL
	iMGA
	iMKD
	iMMK
	iMNT
	iMOP
	iMRU
	iMUR
	iMVR
	iMWK
	iMXN
	iMXV
	iMYR
	iMZN
	iNAD
	iNGN
	iNIO
	iNOK
	iNPR
	iNZD
	iOMR
	iPAB
	iPEN
	iPGK
	iPHP
	iPKR
	iPLN
	iPYG
	iQAR
	iRON
	iRSD
	iRUB
	iRWF
	iSAR
	iSBD
	iSCR
	iSDG
	iSEK
	iSGD
	iSHP
	iSLE
	iSLL
	iSOS
	iSRD
	iSSP
	iSTN
	iSVC
	iSYP
	iSZL
	iTHB
	iTJS
	iTMT
	iTND
	iTOP
	iTRY
	iTTD
	iTWD
	iTZS
	iUAH
	iUGX
	iUSD
	iUSN
	iUYI
	iUYU
	iUYW
	iUZS
	iVED
	iVES
	iVND
	iVUV
	iWST
	iXAF
	iXAG
	iXAU
	iXBA
	iXBB
	iXBC
	iXBD
	iXCD
	iXDR
	iXOF
	iXPD
	iXPF
	iXPT
	iXSU
	iXTS
	iXUA
	iXXX
	iYER
	iZAR
	iZMW
	iZWL
)

var (
	AED = Currency{iAED}
	AFN = Currency{iAFN}
	ALL = Currency{iALL}
	AMD = Currency{iAMD}
	ANG = Currency{iANG}
	AOA = Currency{iAOA}
	ARS = Currency{iARS}
	AUD = Currency{iAUD}
	AWG = Currency{iAWG}
	AZN = Currency{iAZN}
	BAM = Currency{iBAM}
	BBD = Currency{iBBD}
	BDT = Currency{iBDT}
	BGN = Currency{iBGN}
	BHD = Currency{iBHD}
	BIF = Currency{iBIF}
	BMD = Currency{iBMD}
	BND = Currency{iBND}
	BOB = Currency{iBOB}
	BOV = Currency{iBOV}
	BRL = Currency{iBRL}
	BSD = Currency{iBSD}
	BTN = Currency{iBTN}
	BWP = Currency{iBWP}
	BYN = Currency{iBYN}
	BZD = Currency{iBZD}
	CAD = Currency{iCAD}
	CDF = Currency{iCDF}
	CHE = Currency{iCHE}
	CHF = Currency{iCHF}
	CHW = Currency{iCHW}
	CLF = Currency{iCLF}
	CLP = Currency{iCLP}
	CNY = Currency{iCNY}
	COP = Currency{iCOP}
	COU = Currency{iCOU}
	CRC = Currency{iCRC}
	CUP = Currency{iCUP}
	CVE = Currency{iCVE}
	CZK = Currency{iCZK}
	DJF = Currency{iDJF}
	DKK = Currency{iDKK}
	DOP = Currency{iDOP}
	DZD = Currency{iDZD}
	EGP = Currency{iEGP}
	ERN = Currency{iERN}
	ETB = Currency{iETB}
	EUR = Currency{iEUR}
	FJD = Currency{iFJD}
	FKP = Currency{iFKP}
	GBP = Currency{iGBP}
	GEL = Currency{iGEL}
	GHS = Currency{iGHS}
	GIP = Currency{iGIP}
	GMD = Currency{iGMD}
	GNF = Currency{iGNF}
	GTQ = Currency{iGTQ}
	GYD = Currency{iGYD}
	HKD = Currency{iHKD}
	HNL = Currency{iHNL}
	HRD = Currency{iHRD}
	HRK = Currency{iHRK}
	HTG = Currency{iHTG}
	HUF = Currency{iHUF}
	IDR = Currency{iIDR}
	ILS = Currency{iILS}
	INR = Currency{iINR}
	IQD = Currency{iIQD}
	IRR = Currency{iIRR}
	ISK = Currency{iISK}
	JMD = Currency{iJMD}
	JOD = Currency{iJOD}
	JPY = Currency{iJPY}
	KES = Currency{iKES}
	KGS = Currency{iKGS}
	KHR = Currency{iKHR}
	KMF = Currency{iKMF}
	KPW = Currency{iKPW}
	KRW = Currency{iKRW}
	KWD = Currency{iKWD}
	KYD = Currency{iKYD}
	KZT = Currency{iKZT}
	LAK = Currency{iLAK}
	LBP = Currency{iLBP}
	LKR = Currency{iLKR}
	LRD = Currency{iLRD}
	LSL = Currency{iLSL}
	LYD = Currency{iLYD}
	MAD = Currency{iMAD}
	MDL = Currency{iMDL}
	MGA = Currency{iMGA}
	MKD = Currency{iMKD}
	MMK = Currency{iMMK}
	MNT = Currency{iMNT}
	MOP = Currency{iMOP}
	MRU = Currency{iMRU}
	MUR = Currency{iMUR}
	MVR = Currency{iMVR}
	MWK = Currency{iMWK}
	MXN = Currency{iMXN}
	MXV = Currency{iMXV}
	MYR = Currency{iMYR}
	MZN = Currency{iMZN}
	NAD = Currency{iNAD}
	NGN = Currency{iNGN}
	NIO = Currency{iNIO}
	NOK = Currency{iNOK}
	NPR = Currency{iNPR}
	NZD = Currency{iNZD}
	OMR = Currency{iOMR}
	PAB = Currency{iPAB}
	PEN = Currency{iPEN}
	PGK = Currency{iPGK}
	PHP = Currency{iPHP}
	PKR = Currency{iPKR}
	PLN = Currency{iPLN}
	PYG = Currency{iPYG}
	QAR = Currency{iQAR}
	RON = Currency{iRON}
	RSD = Currency{iRSD}
	RUB = Currency{iRUB}
	RWF = Currency{iRWF}
	SAR = Currency{iSAR}
	SBD = Currency{iSBD}
	SCR = Currency{iSCR}
	SDG = Currency{iSDG}
	SEK = Currency{iSEK}
	SGD = Currency{iSGD}
	SHP = Currency{iSHP}
	SLE = Currency{iSLE}
	SLL = Currency{iSLL}
	SOS = Currency{iSOS}
	SRD = Currency{iSRD}
	SSP = Currency{iSSP}
	STN = Currency{iSTN}
	SVC = Currency{iSVC}
	SYP = Currency{iSYP}
	SZL = Currency{iSZL}
	THB = Currency{iTHB}
	TJS = Currency{iTJS}
	TMT = Currency{iTMT}
	TND = Currency{iTND}
	TOP = Currency{iTOP}
	TRY = Currency{iTRY}
	TTD = Currency{iTTD}
	TWD = Currency{iTWD}
	TZS = Currency{iTZS}
	UAH = Currency{iUAH}
	UGX = Currency{iUGX}
	USD = Currency{iUSD}
	USN = Currency{iUSN}
	UYI = Currency{iUYI}
	UYU = Currency{iUYU}
	UYW = Currency{iUYW}
	UZS = Currency{iUZS}
	VED = Currency{iVED}
	VES = Currency{iVES}
	VND = Currency{iVND}
	VUV = Currency{iVUV}
	WST = Currency{iWST}
	XAF = Currency{iXAF}
	XAG = Currency{iXAG}
	XAU = Currency{iXAU}
	XBA = Currency{iXBA}
	XBB = Currency{iXBB}
	XBC = Currency{iXBC}
	XBD = Currency{iXBD}
	XCD = Currency{iXCD}
	XDR = Currency{iXDR}
	XOF = Currency{iXOF}
	XPD = Currency{iXPD}
	XPF = Currency{iXPF}
	XPT = Currency{iXPT}
	XSU = Currency{iXSU}
	XTS = Currency{iXTS}
	XUA = Currency{iXUA}
	XXX = Currency{iXXX}
	YER = Currency{iYER}
	ZAR = Currency{iZAR}
	ZMW = Currency{iZMW}
	ZWL = Currency{iZWL}
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
