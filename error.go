package fpmoney

type ErrCurrencyMismatch struct {
	A, B Currency
}

func NewErrCurrencyMismatch() *ErrCurrencyMismatch { return &ErrCurrencyMismatch{} }

func (e *ErrCurrencyMismatch) Error() string { return e.A.Alpha() + " != " + e.B.Alpha() }

type ErrWrongCurrencyString struct{}

func NewErrWrongCurrencyString() *ErrWrongCurrencyString { return &ErrWrongCurrencyString{} }

func (e *ErrWrongCurrencyString) Error() string { return "wrong currency string" }
