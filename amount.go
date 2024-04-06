package fpmoney

import (
	"github.com/nikolaydubina/fpdecimal"
)

func init() {
	fpdecimal.FractionDigits = 4
}

// Amount stores fixed-precision decimal money.
// Values fit in ~92 quadrillion for 2 decimal currencies.
// Does not use float in printing nor parsing.
// Rounds down fractional cents during parsing.
// Blocking arithmetic operations that result in loss of precision.
type Amount struct {
	Amount   fpdecimal.Decimal `json:"amount"`
	Currency Currency          `json:"currency"`
}

func (a Amount) Add(b Amount) Amount {
	checkCurrency(a.Currency, b.Currency)
	return Amount{Amount: a.Amount.Add(b.Amount), Currency: a.Currency}
}

func (a Amount) Sub(b Amount) Amount {
	checkCurrency(a.Currency, b.Currency)
	return Amount{Amount: a.Amount.Sub(b.Amount), Currency: a.Currency}
}

func (a Amount) GreaterThan(b Amount) bool {
	checkCurrency(a.Currency, b.Currency)
	return a.Amount.GreaterThan(b.Amount)
}

func (a Amount) LessThan(b Amount) bool {
	checkCurrency(a.Currency, b.Currency)
	return a.Amount.LessThan(b.Amount)
}

func (a Amount) GreaterThanOrEqual(b Amount) bool {
	checkCurrency(a.Currency, b.Currency)
	return a.Amount.GreaterThanOrEqual(b.Amount)
}

func (a Amount) LessThanOrEqual(b Amount) bool {
	checkCurrency(a.Currency, b.Currency)
	return a.Amount.LessThanOrEqual(b.Amount)
}

func checkCurrency(a, b Currency) {
	if a != b {
		panic(&ErrCurrencyMismatch{A: a, B: b})
	}
}

func (a Amount) Mul(b int) Amount {
	return Amount{Amount: a.Amount.Mul(fpdecimal.FromInt(b)), Currency: a.Currency}
}

func (a Amount) Div(b int) (part Amount, remainder Amount) {
	return Amount{Amount: a.Amount.Div(fpdecimal.FromInt(b)), Currency: a.Currency}, Amount{Amount: a.Amount.Mod(fpdecimal.FromInt(b)), Currency: a.Currency}
}

type ErrCurrencyMismatch struct {
	A, B Currency
}

func NewErrCurrencyMismatch() *ErrCurrencyMismatch { return &ErrCurrencyMismatch{} }

func (e *ErrCurrencyMismatch) Error() string {
	a, _ := e.A.MarshalText()
	b, _ := e.B.MarshalText()
	return string(a) + " != " + string(b)
}
