package fpmoney

import (
	"errors"

	"github.com/nikolaydubina/fpdecimal"
)

// Amount stores fixed-precision decimal money.
// Stores integer number of cents for ISO 4217 currency.
// Values fit in ~92 quadrillion for 2 decimal currencies.
// Does not use float in printing nor parsing.
// Rounds down fractional cents during parsing.
// Blocking arithmetic operations that result in loss of precision.
type Amount struct {
	v int64
	c Currency
	_ uint8  // padding for improved line alignment, 2x improves arithmetics
	_ uint32 // padding
}

func FromIntScaled[T ~int | ~int8 | ~int16 | ~int32 | ~int64](v T, currency Currency) Amount {
	return Amount{v: int64(v), c: currency}
}

func FromInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](v T, currency Currency) Amount {
	return Amount{v: int64(v) * int64(currency.scale()), c: currency}
}

func FromFloat[T ~float32 | ~float64](v T, currency Currency) Amount {
	return Amount{v: int64(T(v) * T(currency.scale())), c: currency}
}

func (a Amount) Float32() float32 { return float32(a.v) / float32(a.c.scale()) }

func (a Amount) Float64() float64 { return float64(a.v) / float64(a.c.scale()) }

func (a Amount) Currency() Currency { return a.c }

func (a Amount) Scaled() int64 { return a.v }

func (a Amount) Add(b Amount) Amount {
	checkCurrency(a.c, b.c)
	return Amount{v: a.v + b.v, c: a.c}
}

func (a Amount) Sub(b Amount) Amount {
	checkCurrency(a.c, b.c)
	return Amount{v: a.v - b.v, c: a.c}
}

func (a Amount) GreaterThan(b Amount) bool {
	checkCurrency(a.c, b.c)
	return a.v > b.v
}

func (a Amount) LessThan(b Amount) bool {
	checkCurrency(a.c, b.c)
	return a.v < b.v
}

func (a Amount) GreaterThanOrEqual(b Amount) bool {
	checkCurrency(a.c, b.c)
	return a.v >= b.v
}

func (a Amount) LessThanOrEqual(b Amount) bool {
	checkCurrency(a.c, b.c)
	return a.v <= b.v
}

func checkCurrency(a, b Currency) {
	if a != b {
		panic(&ErrCurrencyMismatch{A: a, B: b})
	}
}

func (a Amount) Mul(b int) Amount { return Amount{v: a.v * int64(b), c: a.c} }

func (a Amount) DivMod(b int) (part Amount, remainder Amount) {
	return Amount{v: a.v / int64(b), c: a.c}, Amount{v: a.v % int64(b), c: a.c}
}

func (a Amount) String() string {
	return fpdecimal.FixedPointDecimalToString(a.v, a.c.Exponent()) + " " + a.c.String()
}

const (
	keyCurrency       = "currency"
	keyAmount         = "amount"
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
		if b[i] == keyCurrency[0] && (i+len(keyCurrency)) <= len(b) && string(b[i:i+len(keyCurrency)]) == keyCurrency {
			i += len(keyCurrency) + 2 // `":`

			// find opening quote.
			for ; i < len(b) && b[i] != '"'; i++ {
			}
			if i == len(b) {
				return ErrWrongCurrencyString
			}
			i++ // opening `"`
			e = i + lenISO427Currency
			if e > len(b) {
				return ErrWrongCurrencyString
			}

			if err := a.c.UnmarshalText(b[i:e]); err != nil {
				return err
			}
			i = e
		}

		// amount
		if b[i] == keyAmount[0] && (i+len(keyCurrency)) <= len(b) && string(b[i:i+len(keyAmount)]) == keyAmount {
			i += len(keyAmount) + 2 // `":`
			// go until find either number or + or -, which is a start of simple number.
			for ; i < len(b) && !(b[i] >= '0' && b[i] <= '9') && b[i] != '-' && b[i] != '+'; i++ {
			}
			as = i
			// find end of number
			for ae = i; ae < len(b) && ((b[ae] >= '0' && b[ae] <= '9') || b[ae] == '-' || b[ae] == '+' || b[ae] == '.'); ae++ {
			}
			i = ae
		}
	}

	a.v, err = fpdecimal.ParseFixedPointDecimal(b[as:ae], a.c.Exponent())

	return err
}

func (a Amount) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, 100)
	b = append(b, `{"`...)
	b = append(b, keyAmount...)
	b = append(b, `":`...)
	b = fpdecimal.AppendFixedPointDecimal(b, a.v, a.c.Exponent())
	b = append(b, `,"`...)
	b = append(b, keyCurrency...)
	b = append(b, `":"`...)
	b = append(b, a.c.String()...)
	b = append(b, `"}`...)
	return b, nil
}

var ErrWrongCurrencyString = errors.New("wrong currency string")

type ErrCurrencyMismatch struct {
	A, B Currency
}

func NewErrCurrencyMismatch() *ErrCurrencyMismatch { return &ErrCurrencyMismatch{} }

func (e *ErrCurrencyMismatch) Error() string { return e.A.String() + " != " + e.B.String() }
