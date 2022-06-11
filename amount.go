package fpmoney

import (
	"strings"

	"github.com/ferdypruis/iso4217"
	"github.com/nikolaydubina/fpdecimal"
	"github.com/nikolaydubina/fpmoney/currency"
	"golang.org/x/exp/constraints"
)

// Amount stores fixed-precision decimal money.
// Stores integer number of cents for ISO 4217 currency.
// Values fit in ~92 quadrillion for 2 decimal currencies.
// Does not use float in printing nor parsing.
// Rounds down fractional cents during parsing.
// Blocking arithmetic operations that result in loss of precision.
type Amount struct {
	v int64
	c iso4217.Currency
	_ uint16 // padding
	_ uint32 // padding
}

func FromInt[T constraints.Integer](v T, currency iso4217.Currency) Amount {
	return Amount{v: int64(v) * scale(currency), c: currency}
}

func FromFloat[T constraints.Float](v T, currency iso4217.Currency) Amount {
	return Amount{v: int64(v) * scale(currency), c: currency}
}

func FromIntScaled[T constraints.Integer](v T, currency iso4217.Currency) Amount {
	return Amount{v: int64(v), c: currency}
}

func (a Amount) Float32() float32 { return float32(a.v) / float32(scale(a.c)) }

func (a Amount) Float64() float64 { return float64(a.v) / float64(scale(a.c)) }

func (a Amount) Currency() iso4217.Currency { return a.c }

func (a Amount) Add(b Amount) Amount {
	if a.c != b.c {
		panic(NewErrCurrencyMismatch(a.c, b.c))
	}
	return Amount{v: a.v + b.v, c: a.c}
}

func (a Amount) Sub(b Amount) Amount {
	if a.c != b.c {
		panic(NewErrCurrencyMismatch(a.c, b.c))
	}
	return Amount{v: a.v - b.v, c: a.c}
}

func (a Amount) Mul(b int) Amount { return Amount{v: a.v * int64(b), c: a.c} }

func (a Amount) Div(b int) (part Amount, remainder Amount) {
	return Amount{v: a.v / int64(b), c: a.c}, Amount{v: a.v % int64(b), c: a.c}
}

func (a Amount) GreaterThan(b Amount) bool {
	if a.c != b.c {
		panic(NewErrCurrencyMismatch(a.c, b.c))
	}
	return a.v > b.v
}

func (a Amount) LessThan(b Amount) bool {
	if a.c != b.c {
		panic(NewErrCurrencyMismatch(a.c, b.c))
	}
	return a.v < b.v
}

func (a Amount) GreaterThanOrEqual(b Amount) bool {
	if a.c != b.c {
		panic(NewErrCurrencyMismatch(a.c, b.c))
	}
	return a.v >= b.v
}

func (a Amount) LessThanOrEqual(b Amount) bool {
	if a.c != b.c {
		panic(NewErrCurrencyMismatch(a.c, b.c))
	}
	return a.v <= b.v
}

func (a Amount) String() string {
	return fpdecimal.FixedPointDecimalToString(int64(a.v), a.c.Exponent()) + " " + a.c.Alpha()
}

const (
	keyCurrency = "currency"
	keyAmount   = "amount"

	lenISO427Currency = 3
)

// UnmarshalJSON parses string.
// This is implemented directly for speed.
// Avoiding json.Decoder, interface{}, reflect, tags, temporary structs.
// Avoiding mallocs.
// Go json package provides:
// - check that pointer method receiver is not nil;
// - removes whitespace in b []bytes
func (a *Amount) UnmarshalJSON(b []byte) (err error) {
	var as, ae, e int

	for i := 0; i < len(b); i++ {
		// currency
		if b[i] == 'c' && (i+len(keyCurrency)) <= len(b) && string(b[i:i+len(keyCurrency)]) == keyCurrency {
			i += len(keyCurrency) + 2 // `":`

			// find opening quote.
			for ; i < len(b) && b[i] != '"'; i++ {
			}
			if i == len(b) {
				return NewErrUnmarshalJSONWrongCurrency("missing json value")
			}
			i++ // opening `"`
			e = i + lenISO427Currency
			if e > len(b) {
				return NewErrUnmarshalJSONWrongCurrency(string(b[i:]))
			}

			a.c = currency.CastCurrency(b[i:e])
			i = e
		}

		// amount
		if b[i] == 'a' && (i+len(keyCurrency)) <= len(b) && string(b[i:i+len(keyAmount)]) == keyAmount {
			i += len(keyAmount) + 2 // `":`
			// go until find either number or + or -, which is a start of simple number.
			for ; i < len(b) && (b[i] < '0' || b[i] > '9') && b[i] != '-' && b[i] != '+'; i++ {
			}
			as = i
			// find end of number
			for ae = i; ae < len(b) && ((b[ae] >= '0' && b[ae] <= '9') || b[ae] == '-' || b[ae] == '+' || b[ae] == '.'); ae++ {
			}
			i = ae
		}
	}

	if a.c == iso4217.Currency(0) {
		return NewErrUnmarshalJSONWrongCurrency("not recognized")
	}

	a.v, err = fpdecimal.ParseFixedPointDecimal(string(b[as:ae]), int8(a.c.Exponent()))

	return err
}

func (a Amount) MarshalJSON() ([]byte, error) {
	var b strings.Builder
	b.Grow(50)
	b.WriteRune('{')
	b.WriteString(`"amount":`)
	b.WriteString(fpdecimal.FixedPointDecimalToString(a.v, a.c.Exponent()))
	b.WriteString(`,`)
	b.WriteString(`"currency":`)
	b.WriteRune('"')
	b.WriteString(a.c.Alpha())
	b.WriteRune('"')
	b.WriteRune('}')
	return []byte(b.String()), nil
}

// ErrCurrencyMismatch is a lazy error for mismatched ISO 4217 currencies.
type ErrCurrencyMismatch struct {
	a iso4217.Currency
	b iso4217.Currency
}

func NewErrCurrencyMismatch(a, b iso4217.Currency) error {
	return &ErrCurrencyMismatch{a: a, b: b}
}

func (e *ErrCurrencyMismatch) Error() string {
	if e == nil {
		return ""
	}
	return e.a.Alpha() + " != " + e.b.Alpha()
}

type ErrUnmarshalJSONWrongCurrency struct {
	s string
}

func NewErrUnmarshalJSONWrongCurrency(s string) *ErrUnmarshalJSONWrongCurrency {
	return &ErrUnmarshalJSONWrongCurrency{s: s}
}

func (e *ErrUnmarshalJSONWrongCurrency) Error() string {
	if e == nil {
		return ""
	}
	return "wrong currency: " + e.s
}

func scale(currency iso4217.Currency) int64 {
	switch currency.Exponent() {
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
