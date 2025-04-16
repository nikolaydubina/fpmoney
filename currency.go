package fpmoney

// Currency is ISO 4217 without deprecated currencies.
//   - excluding currencies with 4+ minor units `CLF`, `UYW`
//   - excluding deprecated currencies `HRD`, `HRK`, `SLL`, `ZWL`
type Currency uint8

// Exponent returns the decimal point location.
func (c Currency) Exponent() uint8 {
	switch c {
	case BIF, CLP, DJF, GNF, ISK, JPY, KMF, KRW, PYG, RWF, UGX, UYI, VND, VUV, XAF, XAG, XAU, XBA, XBB, XBC, XBD, XDR, XOF, XPD, XPF, XPT, XSU, XTS, XUA, XXX:
		return 0
	case BHD, IQD, JOD, KWD, LYD, OMR, TND:
		return 3
	default:
		return 2
	}
}

func (c Currency) scale() int64 {
	switch c.Exponent() {
	case 3:
		return 1000
	case 2:
		return 100
	default:
		return 1
	}
}

//go:generate go-enum-encoding -type=Currency -string
const (
	UndefinedCurrency Currency = iota //
	AED                               // json:"AED"
	AFN                               // json:"AFN"
	ALL                               // json:"ALL"
	AMD                               // json:"AMD"
	ANG                               // json:"ANG"
	AOA                               // json:"AOA"
	ARS                               // json:"ARS"
	AUD                               // json:"AUD"
	AWG                               // json:"AWG"
	AZN                               // json:"AZN"
	BAM                               // json:"BAM"
	BBD                               // json:"BBD"
	BDT                               // json:"BDT"
	BGN                               // json:"BGN"
	BHD                               // json:"BHD"
	BIF                               // json:"BIF"
	BMD                               // json:"BMD"
	BND                               // json:"BND"
	BOB                               // json:"BOB"
	BOV                               // json:"BOV"
	BRL                               // json:"BRL"
	BSD                               // json:"BSD"
	BTN                               // json:"BTN"
	BWP                               // json:"BWP"
	BYN                               // json:"BYN"
	BZD                               // json:"BZD"
	CAD                               // json:"CAD"
	CDF                               // json:"CDF"
	CHE                               // json:"CHE"
	CHF                               // json:"CHF"
	CHW                               // json:"CHW"
	CLP                               // json:"CLP"
	CNY                               // json:"CNY"
	COP                               // json:"COP"
	COU                               // json:"COU"
	CRC                               // json:"CRC"
	CUP                               // json:"CUP"
	CVE                               // json:"CVE"
	CZK                               // json:"CZK"
	DJF                               // json:"DJF"
	DKK                               // json:"DKK"
	DOP                               // json:"DOP"
	DZD                               // json:"DZD"
	EGP                               // json:"EGP"
	ERN                               // json:"ERN"
	ETB                               // json:"ETB"
	EUR                               // json:"EUR"
	FJD                               // json:"FJD"
	FKP                               // json:"FKP"
	GBP                               // json:"GBP"
	GEL                               // json:"GEL"
	GHS                               // json:"GHS"
	GIP                               // json:"GIP"
	GMD                               // json:"GMD"
	GNF                               // json:"GNF"
	GTQ                               // json:"GTQ"
	GYD                               // json:"GYD"
	HKD                               // json:"HKD"
	HNL                               // json:"HNL"
	HRD                               // json:"HRD"
	HTG                               // json:"HTG"
	HUF                               // json:"HUF"
	IDR                               // json:"IDR"
	ILS                               // json:"ILS"
	INR                               // json:"INR"
	IQD                               // json:"IQD"
	IRR                               // json:"IRR"
	ISK                               // json:"ISK"
	JMD                               // json:"JMD"
	JOD                               // json:"JOD"
	JPY                               // json:"JPY"
	KES                               // json:"KES"
	KGS                               // json:"KGS"
	KHR                               // json:"KHR"
	KMF                               // json:"KMF"
	KPW                               // json:"KPW"
	KRW                               // json:"KRW"
	KWD                               // json:"KWD"
	KYD                               // json:"KYD"
	KZT                               // json:"KZT"
	LAK                               // json:"LAK"
	LBP                               // json:"LBP"
	LKR                               // json:"LKR"
	LRD                               // json:"LRD"
	LSL                               // json:"LSL"
	LYD                               // json:"LYD"
	MAD                               // json:"MAD"
	MDL                               // json:"MDL"
	MGA                               // json:"MGA"
	MKD                               // json:"MKD"
	MMK                               // json:"MMK"
	MNT                               // json:"MNT"
	MOP                               // json:"MOP"
	MRU                               // json:"MRU"
	MUR                               // json:"MUR"
	MVR                               // json:"MVR"
	MWK                               // json:"MWK"
	MXN                               // json:"MXN"
	MXV                               // json:"MXV"
	MYR                               // json:"MYR"
	MZN                               // json:"MZN"
	NAD                               // json:"NAD"
	NGN                               // json:"NGN"
	NIO                               // json:"NIO"
	NOK                               // json:"NOK"
	NPR                               // json:"NPR"
	NZD                               // json:"NZD"
	OMR                               // json:"OMR"
	PAB                               // json:"PAB"
	PEN                               // json:"PEN"
	PGK                               // json:"PGK"
	PHP                               // json:"PHP"
	PKR                               // json:"PKR"
	PLN                               // json:"PLN"
	PYG                               // json:"PYG"
	QAR                               // json:"QAR"
	RON                               // json:"RON"
	RSD                               // json:"RSD"
	RUB                               // json:"RUB"
	RWF                               // json:"RWF"
	SAR                               // json:"SAR"
	SBD                               // json:"SBD"
	SCR                               // json:"SCR"
	SDG                               // json:"SDG"
	SEK                               // json:"SEK"
	SGD                               // json:"SGD"
	SHP                               // json:"SHP"
	SLE                               // json:"SLE"
	SOS                               // json:"SOS"
	SRD                               // json:"SRD"
	SSP                               // json:"SSP"
	STN                               // json:"STN"
	SVC                               // json:"SVC"
	SYP                               // json:"SYP"
	SZL                               // json:"SZL"
	THB                               // json:"THB"
	TJS                               // json:"TJS"
	TMT                               // json:"TMT"
	TND                               // json:"TND"
	TOP                               // json:"TOP"
	TRY                               // json:"TRY"
	TTD                               // json:"TTD"
	TWD                               // json:"TWD"
	TZS                               // json:"TZS"
	UAH                               // json:"UAH"
	UGX                               // json:"UGX"
	USD                               // json:"USD"
	USN                               // json:"USN"
	UYI                               // json:"UYI"
	UYU                               // json:"UYU"
	UZS                               // json:"UZS"
	VED                               // json:"VED"
	VES                               // json:"VES"
	VND                               // json:"VND"
	VUV                               // json:"VUV"
	WST                               // json:"WST"
	XAF                               // json:"XAF"
	XAG                               // json:"XAG"
	XAU                               // json:"XAU"
	XBA                               // json:"XBA"
	XBB                               // json:"XBB"
	XBC                               // json:"XBC"
	XBD                               // json:"XBD"
	XCD                               // json:"XCD"
	XDR                               // json:"XDR"
	XOF                               // json:"XOF"
	XPD                               // json:"XPD"
	XPF                               // json:"XPF"
	XPT                               // json:"XPT"
	XSU                               // json:"XSU"
	XTS                               // json:"XTS"
	XUA                               // json:"XUA"
	XXX                               // json:"XXX"
	YER                               // json:"YER"
	ZAR                               // json:"ZAR"
	ZMW                               // json:"ZMW"
	ZWG                               // json:"ZWG"
)

var Currencies = [...]Currency{
	AED,
	AFN,
	ALL,
	AMD,
	ANG,
	AOA,
	ARS,
	AUD,
	AWG,
	AZN,
	BAM,
	BBD,
	BDT,
	BGN,
	BHD,
	BIF,
	BMD,
	BND,
	BOB,
	BOV,
	BRL,
	BSD,
	BTN,
	BWP,
	BYN,
	BZD,
	CAD,
	CDF,
	CHE,
	CHF,
	CHW,
	CLP,
	CNY,
	COP,
	COU,
	CRC,
	CUP,
	CVE,
	CZK,
	DJF,
	DKK,
	DOP,
	DZD,
	EGP,
	ERN,
	ETB,
	EUR,
	FJD,
	FKP,
	GBP,
	GEL,
	GHS,
	GIP,
	GMD,
	GNF,
	GTQ,
	GYD,
	HKD,
	HNL,
	HRD,
	HTG,
	HUF,
	IDR,
	ILS,
	INR,
	IQD,
	IRR,
	ISK,
	JMD,
	JOD,
	JPY,
	KES,
	KGS,
	KHR,
	KMF,
	KPW,
	KRW,
	KWD,
	KYD,
	KZT,
	LAK,
	LBP,
	LKR,
	LRD,
	LSL,
	LYD,
	MAD,
	MDL,
	MGA,
	MKD,
	MMK,
	MNT,
	MOP,
	MRU,
	MUR,
	MVR,
	MWK,
	MXN,
	MXV,
	MYR,
	MZN,
	NAD,
	NGN,
	NIO,
	NOK,
	NPR,
	NZD,
	OMR,
	PAB,
	PEN,
	PGK,
	PHP,
	PKR,
	PLN,
	PYG,
	QAR,
	RON,
	RSD,
	RUB,
	RWF,
	SAR,
	SBD,
	SCR,
	SDG,
	SEK,
	SGD,
	SHP,
	SLE,
	SOS,
	SRD,
	SSP,
	STN,
	SVC,
	SYP,
	SZL,
	THB,
	TJS,
	TMT,
	TND,
	TOP,
	TRY,
	TTD,
	TWD,
	TZS,
	UAH,
	UGX,
	USD,
	USN,
	UYI,
	UYU,
	UZS,
	VED,
	VES,
	VND,
	VUV,
	WST,
	XAF,
	XAG,
	XAU,
	XBA,
	XBB,
	XBC,
	XBD,
	XCD,
	XDR,
	XOF,
	XPD,
	XPF,
	XPT,
	XSU,
	XTS,
	XUA,
	XXX,
	YER,
	ZAR,
	ZMW,
	ZWG,
}
