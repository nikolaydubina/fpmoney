package fpmoney

import (
	"github.com/nikolaydubina/fpdecimal"
)

var (
	Currency           string = "USD"
	CurrencyMinorUnits uint8  = 2

	multiplier = [...]int{1, 10, 100, 1000, 10000}
)

// Amount stores fixed-precision decimal money.
// Values fit in ~92 quadrillion for 2 decimal currencies.
// Does not use float in printing nor parsing.
// Rounds down fractional cents during parsing.
// Blocking arithmetic operations that result in loss of precision.
type Amount struct{ v int64 }

func FromIntScaled[T ~int | ~int8 | ~int16 | ~int32 | ~int64](v T) Amount { return Amount{v: int64(v)} }

func FromInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](v T) Amount {
	return Amount{v: int64(v) * int64(multiplier[CurrencyMinorUnits])}
}

func FromFloat[T ~float32 | ~float64](v T) Amount {
	return Amount{v: int64(T(v) * T(multiplier[CurrencyMinorUnits]))}
}

func (a Amount) Float32() float32 { return float32(a.v) / float32(multiplier[CurrencyMinorUnits]) }

func (a Amount) Float64() float64 { return float64(a.v) / float64(multiplier[CurrencyMinorUnits]) }

func (a Amount) Add(b Amount) Amount { return Amount{v: a.v + b.v} }

func (a Amount) Sub(b Amount) Amount { return Amount{v: a.v - b.v} }

func (a Amount) GreaterThan(b Amount) bool { return a.v > b.v }

func (a Amount) LessThan(b Amount) bool { return a.v < b.v }

func (a Amount) GreaterThanOrEqual(b Amount) bool { return a.v >= b.v }

func (a Amount) LessThanOrEqual(b Amount) bool { return a.v <= b.v }

func (a Amount) Mul(b int) Amount { return Amount{v: a.v * int64(b)} }

func (a Amount) DivMod(b int) (part, remainder Amount) {
	return Amount{v: a.v / int64(b)}, Amount{v: a.v % int64(b)}
}

func (a Amount) String() string {
	return fpdecimal.FixedPointDecimalToString(a.v, CurrencyMinorUnits) + " " + Currency
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
				return &errorString{"wrong currency"}
			}
			i++ // opening `"`
			e = i + lenISO427Currency
			if e > len(b) {
				return &errorString{"wrong currency"}
			}

			if string(b[i:e]) != Currency {
				return &errorString{"wrong currency"}
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

	a.v, err = fpdecimal.ParseFixedPointDecimal(b[as:ae], CurrencyMinorUnits)

	return err
}

func (a Amount) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, 100)
	b = append(b, `{"`...)
	b = append(b, keyAmount...)
	b = append(b, `":`...)
	b = fpdecimal.AppendFixedPointDecimal(b, a.v, CurrencyMinorUnits)
	b = append(b, `,"`...)
	b = append(b, keyCurrency...)
	b = append(b, `":"`...)
	b = append(b, Currency...)
	b = append(b, `"}`...)
	return b, nil
}

type errorString struct{ v string }

func (e *errorString) Error() string { return e.v }
