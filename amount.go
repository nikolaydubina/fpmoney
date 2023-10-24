package fpmoney

import (
	"github.com/nikolaydubina/fpdecimal"
)

// Amount stores fixed-precision decimal money.
// Stores integer number of cents for ISO 4217 currency.
// Values fit in ~92 quadrillion for 2 decimal currencies.
// Does not use float in printing nor parsing.
// Rounds down fractional cents during parsing.
// Blocking arithmetic operations that result in loss of precision.
type Amount[C comparable] struct {
	Amount   fpdecimal.Decimal `json:"amount"`
	Currency C                 `json:"currency"`
}

func FromIntScaled[T ~int | ~int8 | ~int16 | ~int32 | ~int64, C comparable](v T, currency C) Amount[C] {
	return Amount[C]{Amount: fpdecimal.FromIntScaled(v), Currency: currency}
}

func FromInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64, C comparable](v T, currency C) Amount[C] {
	return Amount[C]{Amount: fpdecimal.FromInt(v), Currency: currency}
}

func FromFloat[T ~float32 | ~float64, C comparable](v T, currency C) Amount[C] {
	return Amount[C]{Amount: fpdecimal.FromFloat(v), Currency: currency}
}

func (a Amount[C]) Float32() float32 { return a.Amount.Float32() }

func (a Amount[C]) Float64() float64 { return a.Amount.Float64() }

func (a Amount[C]) Add(b Amount[C]) Amount[C] {
	checkCurrency(a.Currency, b.Currency)
	return Amount[C]{Amount: a.Amount.Add(b.Amount), Currency: a.Currency}
}

func (a Amount[C]) Sub(b Amount[C]) Amount[C] {
	checkCurrency(a.Currency, b.Currency)
	return Amount[C]{Amount: a.Amount.Sub(b.Amount), Currency: a.Currency}
}

func (a Amount[C]) GreaterThan(b Amount[C]) bool {
	checkCurrency(a.Currency, b.Currency)
	return a.Amount.GreaterThan(b.Amount)
}

func (a Amount[C]) LessThan(b Amount[C]) bool {
	checkCurrency(a.Currency, b.Currency)
	return a.Amount.LessThan(b.Amount)
}

func (a Amount[C]) GreaterThanOrEqual(b Amount[C]) bool {
	checkCurrency(a.Currency, b.Currency)
	return a.Amount.GreaterThanOrEqual(b.Amount)
}

func (a Amount[C]) LessThanOrEqual(b Amount[C]) bool {
	checkCurrency(a.Currency, b.Currency)
	return a.Amount.LessThanOrEqual(b.Amount)
}

func checkCurrency[C comparable](a, b C) {
	if a != b {
		panic("currency mismatch")
	}
}

func (a Amount[C]) Mul(b int) Amount[C] {
	return Amount[C]{Amount: a.Amount.Mul(fpdecimal.FromInt(b)), Currency: a.Currency}
}

func (a Amount[C]) Div(b int) (part, remainder Amount[C]) {
	r, m := a.Amount.DivMod(fpdecimal.FromInt(b))
	return Amount[C]{Amount: r, Currency: a.Currency}, Amount[C]{Amount: m, Currency: a.Currency}
}
