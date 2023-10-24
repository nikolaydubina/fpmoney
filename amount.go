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
type Amount[T comparable] struct {
	V fpdecimal.Decimal `json:"amount"`
	C T                 `json:"currency"`
}

func FromIntScaled[T ~int | ~int8 | ~int16 | ~int32 | ~int64, C comparable](v T, currency C) Amount[C] {
	return Amount[C]{V: fpdecimal.FromIntScaled(v), C: currency}
}

func FromInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64, C comparable](v T, currency C) Amount[C] {
	return Amount[C]{V: fpdecimal.FromInt(v), C: currency}
}

func FromFloat[T ~float32 | ~float64, C comparable](v T, currency C) Amount[C] {
	return Amount[C]{V: fpdecimal.FromFloat(v), C: currency}
}

func (a Amount[T]) Float32() float32 { return a.V.Float32() }

func (a Amount[T]) Float64() float64 { return a.V.Float64() }

func (a Amount[T]) Add(b Amount[T]) Amount[T] {
	checkCurrency(a.C, b.C)
	return Amount[T]{V: a.V.Add(b.V), C: a.C}
}

func (a Amount[T]) Sub(b Amount[T]) Amount[T] {
	checkCurrency(a.C, b.C)
	return Amount[T]{V: a.V.Sub(b.V), C: a.C}
}

func (a Amount[T]) GreaterThan(b Amount[T]) bool {
	checkCurrency(a.C, b.C)
	return a.V.GreaterThan(b.V)
}

func (a Amount[T]) LessThan(b Amount[T]) bool {
	checkCurrency(a.C, b.C)
	return a.V.LessThan(b.V)
}

func (a Amount[T]) GreaterThanOrEqual(b Amount[T]) bool {
	checkCurrency(a.C, b.C)
	return a.V.GreaterThanOrEqual(b.V)
}

func (a Amount[T]) LessThanOrEqual(b Amount[T]) bool {
	checkCurrency(a.C, b.C)
	return a.V.LessThanOrEqual(b.V)
}

func checkCurrency[T comparable](a, b T) {
	if a != b {
		panic("currency mismatch")
	}
}

func (a Amount[T]) Mul(b int) Amount[T] { return Amount[T]{V: a.V.Mul(fpdecimal.FromInt(b)), C: a.C} }

func (a Amount[T]) Div(b int) (part Amount[T], remainder Amount[T]) {
	r, m := a.V.DivMod(fpdecimal.FromInt(b))
	return Amount[T]{V: r, C: a.C}, Amount[T]{V: m, C: a.C}
}
