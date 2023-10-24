package fpmoney

import "fmt"

type ErrCurrencyMismatch[T comparable] struct {
	A, B T
}

func NewErrCurrencyMismatch[T comparable]() *ErrCurrencyMismatch[T] { return &ErrCurrencyMismatch[T]{} }

func (e *ErrCurrencyMismatch[T]) Error() string { return fmt.Sprintf("%v != %v", e.A, e.B) }

type ErrWrongCurrencyString struct{}

func NewErrWrongCurrencyString() *ErrWrongCurrencyString { return &ErrWrongCurrencyString{} }

func (e *ErrWrongCurrencyString) Error() string { return "wrong currency string" }
