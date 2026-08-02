package fpmoney

import (
	"encoding/binary"
	"errors"

	"github.com/nikolaydubina/fpdecimal"
)

var ErrWrongCurrencyString = errors.New("wrong currency string")

type ErrCurrencyMismatch struct{ A, B Currency }

func (e *ErrCurrencyMismatch) Error() string { return e.A.String() + " != " + e.B.String() }

// Amount stores fixed-precision decimal money.
// Values fit in ~92 quadrillion for 2 decimal currencies.
// Does not use float in printing nor parsing.
// Rounds down fractional cents during parsing.
// Blocks arithmetics that result in loss of precision or mismatched currency.
type Amount struct {
	Minor    int64
	Currency Currency
}

func FromMinor[T ~int | ~int8 | ~int16 | ~int32 | ~int64](v T, currency Currency) Amount {
	return Amount{Minor: int64(v), Currency: currency}
}

func FromInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](v T, currency Currency) Amount {
	return Amount{Minor: int64(v) * int64(currency.scale()), Currency: currency}
}

func FromFloat[T ~float32 | ~float64](v T, currency Currency) Amount {
	return Amount{Minor: int64(T(v) * T(currency.scale())), Currency: currency}
}

func (a Amount) Float32() float32 { return float32(a.Minor) / float32(a.Currency.scale()) }

func (a Amount) Float64() float64 { return float64(a.Minor) / float64(a.Currency.scale()) }

func (a Amount) Add(b Amount) Amount {
	checkCurrency(a.Currency, b.Currency)
	return Amount{Minor: a.Minor + b.Minor, Currency: a.Currency}
}

func (a Amount) Sub(b Amount) Amount {
	checkCurrency(a.Currency, b.Currency)
	return Amount{Minor: a.Minor - b.Minor, Currency: a.Currency}
}

func (a Amount) GreaterThan(b Amount) bool {
	checkCurrency(a.Currency, b.Currency)
	return a.Minor > b.Minor
}

func (a Amount) LessThan(b Amount) bool {
	checkCurrency(a.Currency, b.Currency)
	return a.Minor < b.Minor
}

func (a Amount) GreaterThanOrEqual(b Amount) bool {
	checkCurrency(a.Currency, b.Currency)
	return a.Minor >= b.Minor
}

func (a Amount) LessThanOrEqual(b Amount) bool {
	checkCurrency(a.Currency, b.Currency)
	return a.Minor <= b.Minor
}

func checkCurrency(a, b Currency) {
	if a != b {
		panic(&ErrCurrencyMismatch{A: a, B: b})
	}
}

func (a Amount) Mul(b int) Amount { return Amount{Minor: a.Minor * int64(b), Currency: a.Currency} }

func (a Amount) DivMod(b int) (part Amount, remainder Amount) {
	return Amount{Minor: a.Minor / int64(b), Currency: a.Currency}, Amount{Minor: a.Minor % int64(b), Currency: a.Currency}
}

func (a Amount) MulFraction(b float64) Amount {
	if b == 0 {
		return Amount{Minor: 0, Currency: a.Currency}
	}
	if b == 1 {
		return a
	}
	if b == -1 {
		return Amount{Minor: -a.Minor, Currency: a.Currency}
	}
	if float64(int(b)) == b {
		return a.Mul(int(b))
	}
	// skip scale roundtrip
	return Amount{Minor: int64(float64(a.Minor) * b), Currency: a.Currency}
}

// Convert with rate=to/from
func (a Amount) Convert(to Currency, rate float64) Amount {
	if a.Currency == to {
		return a
	}
	// skip scale roundtrip
	return Amount{Minor: int64(float64(a.Minor) * rate), Currency: to}
}

func (a Amount) StringPair() (amount, currency string) {
	return fpdecimal.FixedPointDecimalToString(a.Minor, a.Currency.Exponent()), a.Currency.String()
}

func (a Amount) String() string {
	return fpdecimal.FixedPointDecimalToString(a.Minor, a.Currency.Exponent()) + " " + a.Currency.String()
}

const (
	keyCurrency       = "currency"
	keyAmount         = "amount"
	lenISO427Currency = 3
)

func (a *Amount) UnmarshalJSON(b []byte) (err error) {
	// This is implemented directly for speed.
	// Avoiding json.Decoder, interface{}, reflect, tags, temporary structs.
	// Avoiding mallocs.
	//
	// Go json package provides:
	// - check that pointer method receiver is not nil;
	// - removes whitespace in b []bytes
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

			if err := a.Currency.UnmarshalText(b[i:e]); err != nil {
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

	a.Minor, err = fpdecimal.ParseFixedPointDecimal(b[as:ae], a.Currency.Exponent())

	return err
}

func (a Amount) AppendJSON(b []byte) ([]byte, error) {
	b = append(b, `{"`...)
	b = append(b, keyAmount...)
	b = append(b, `":`...)

	b = fpdecimal.AppendFixedPointDecimal(b, a.Minor, a.Currency.Exponent())

	b = append(b, `,"`...)
	b = append(b, keyCurrency...)
	b = append(b, `":"`...)

	var err error
	b, err = a.Currency.AppendText(b)

	b = append(b, `"}`...)
	return b, err
}

func (a Amount) MarshalJSON() ([]byte, error) { return a.AppendJSON(make([]byte, 0, 100)) }

const NumBytes = 8 + 1

func (s Amount) AppendBinary(b []byte) ([]byte, error) {
	b = append(b, byte(s.Currency))
	b = binary.LittleEndian.AppendUint64(b, uint64(s.Minor))
	return b, nil
}

func (s Amount) MarshalBinary() ([]byte, error) { return s.AppendBinary(make([]byte, 0, NumBytes)) }

func (s *Amount) UnmarshalBinary(b []byte) error {
	if len(b) != NumBytes {
		return errors.New("invalid length")
	}
	s.Currency = Currency(b[0])
	s.Minor = int64(binary.LittleEndian.Uint64(b[1:9]))
	return nil
}
