package fpmoney

import (
	_ "embed"
	"errors"
	"fmt"
)

// Currency is ISO 4217 without deprecated currencies.
// Zero value is undefined currency.
type Currency uint8

func (c Currency) MarshalText() (text []byte, err error) { return []byte(c.String()), nil }

func (c *Currency) UnmarshalText(text []byte) error {
	v, ok := fromAlpha[string(text)]
	if !ok {
		return errors.New("unexpected alpha: " + string(text))
	}
	*c = v
	return nil
}

var (
	fromAlpha = make(map[string]Currency)
)

func init() {
	for i := 0; i < len(_Currency_index); i++ {
		fromAlpha[Currency(_Currency_index[i]).String()] = Currency(_Currency_index[i])
	}
	fmt.Printf("%#v\n", fromAlpha)
}

//go:generate stringer -type=Currency

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
