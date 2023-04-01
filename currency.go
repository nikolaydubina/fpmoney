package fpmoney

// Currency is ISO 4217 without deprecated currencies.
// Zero value is undefined currency.
type Currency struct{ v uint8 }

// Alpha returns the ISO 4217 three-letter alphabetic code.
func (c Currency) Alpha() string { return currencies[c].alpha }

// Exponent returns the decimal point location.
func (c Currency) Exponent() int { return currencies[c].exponent }

func (c Currency) String() string { return c.Alpha() }

func (c Currency) IsUndefined() bool { return c.v == 0 }

func (c *Currency) UnmarshalText(text []byte) error {
	v, ok := fromAlpha[string(text)]
	if !ok {
		return &ErrWrongCurrencyString{}
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
	return Currency{}, &ErrWrongCurrencyString{}
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
