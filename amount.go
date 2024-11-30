package fpmoney

import (
	"github.com/nikolaydubina/fpdecimal/fp3"
)

type integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// Amount stores fixed-precision decimal money and ISO 4217 currency.
// Blocking arithmetic operations that result in loss of precision.
type Amount struct {
	Amount   fp3.Decimal `json:"amount"`
	Currency Currency    `json:"currency"`
}

func FromInt[T integer](amount T, currency Currency) Amount {
	return Amount{Amount: fp3.FromInt(amount), Currency: currency}
}

// FromIntScaled converts amount as minor units of currency.
func FromIntScaled[T integer](amount T, currency Currency) Amount {
	return Amount{Amount: fp3.FromIntScaled(amount * T(1000/currency.scale())), Currency: currency}
}

func FromFloat[T float32 | float64](amount T, currency Currency) Amount {
	return Amount{Amount: fp3.FromFloat(amount), Currency: currency}
}

// Scaled returns amount in minor units for currency
func (a Amount) Scaled() int64 { return a.Amount.Scaled() * (a.Currency.scale() / 1000) }

func (a Amount) String() string { return a.Currency.String() + a.Amount.String() }

func (a Amount) Add(b Amount) Amount {
	if a.Currency != b.Currency {
		panic(&ErrCurrencyMismatch{A: a.Currency, B: b.Currency})
	}
	return Amount{Amount: a.Amount.Add(b.Amount), Currency: a.Currency}
}

func (a Amount) Sub(b Amount) Amount {
	if a.Currency != b.Currency {
		panic(&ErrCurrencyMismatch{A: a.Currency, B: b.Currency})
	}
	return Amount{Amount: a.Amount.Sub(b.Amount), Currency: a.Currency}
}

func (a Amount) Mul(b int) Amount { return Amount{a.Amount.Mul(fp3.FromInt(b)), a.Currency} }

func (a Amount) Div(b int) Amount { return Amount{a.Amount.Div(fp3.FromInt(b)), a.Currency} }

func (a Amount) Mod(b int) Amount { return Amount{a.Amount.Mod(fp3.FromInt(b)), a.Currency} }

func (a Amount) DivMod(b int) (part, remainder Amount) { return a.Div(b), a.Mod(b) }

func (a Amount) GreaterThan(b Amount) bool {
	if a.Currency != b.Currency {
		panic(&ErrCurrencyMismatch{A: a.Currency, B: b.Currency})
	}
	return a.Amount.GreaterThan(b.Amount)
}

func (a Amount) LessThan(b Amount) bool {
	if a.Currency != b.Currency {
		panic(&ErrCurrencyMismatch{A: a.Currency, B: b.Currency})
	}
	return a.Amount.LessThan(b.Amount)
}

func (a Amount) GreaterThanOrEqual(b Amount) bool {
	if a.Currency != b.Currency {
		panic(&ErrCurrencyMismatch{A: a.Currency, B: b.Currency})
	}
	return a.Amount.GreaterThanOrEqual(b.Amount)
}

func (a Amount) LessThanOrEqual(b Amount) bool {
	if a.Currency != b.Currency {
		panic(&ErrCurrencyMismatch{A: a.Currency, B: b.Currency})
	}
	return a.Amount.LessThanOrEqual(b.Amount)
}

type ErrCurrencyMismatch struct {
	A, B Currency
}

func NewErrCurrencyMismatch() *ErrCurrencyMismatch { return &ErrCurrencyMismatch{} }

func (e ErrCurrencyMismatch) Error() string { return e.A.String() + " != " + e.B.String() }
