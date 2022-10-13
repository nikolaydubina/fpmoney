package fpmoney

import "errors"

// Currency is ISO 4217 without deprecated currencies.
type Currency uint8

// Alpha returns the ISO 4217 three-letter alphabetic code.
func (c Currency) Alpha() string { return currencies[c].alpha }

// Exponent returns the decimal point location.
func (c Currency) Exponent() int { return currencies[c].exponent }

func (c Currency) Scale() int64 {
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

func (c Currency) append(b []byte) []byte { return append(b, currencies[c].alphaBytes...) }

// CurrencyFromAlpha returns Currency for the three-letter alpha code.
// Or an error if it does not exist.
func CurrencyFromAlpha(alpha string) (Currency, error) {
	if c, ok := fromAlpha[alpha]; ok {
		return c, nil
	}
	return Currency(0), errors.New("no currency exists with alphabetic code " + alpha)
}
