package fpmoney_test

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/nikolaydubina/fpdecimal"
	"github.com/nikolaydubina/fpmoney"
)

func ExampleAmount() {
	var BuySP500Price = fpmoney.Amount{Amount: fpdecimal.FromInt(9000), Currency: fpmoney.SGD}

	input := []byte(`{"sp500": {"amount": 9000.02, "currency": "SGD"}}`)

	type Stonks struct {
		SP500 fpmoney.Amount `json:"sp500"`
	}
	var v Stonks
	if err := json.Unmarshal(input, &v); err != nil {
		log.Fatal(err)
	}

	amountToBuy := fpmoney.Amount{Amount: fpdecimal.Zero, Currency: fpmoney.SGD}
	if v.SP500.GreaterThan(BuySP500Price) {
		amountToBuy = amountToBuy.Add(v.SP500.Mul(2))
	}

	json.NewEncoder(os.Stdout).Encode(amountToBuy)
	// Output: {"amount":18000.04,"currency":"SGD"}
}

func ExampleAmount_Div_part() {
	x := fpmoney.Amount{Amount: fpdecimal.FromInt(1), Currency: fpmoney.SGD}
	a, r := x.Div(3)
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(a)
	enc.Encode(r)
	// Output:
	// {"amount":0.3333,"currency":"SGD"}
	// {"amount":0.0001,"currency":"SGD"}
}

func ExampleAmount_Div_whole() {
	x := fpmoney.Amount{Amount: fpdecimal.FromInt(1), Currency: fpmoney.SGD}
	a, r := x.Div(5)
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(a)
	enc.Encode(r)
	// Output:
	// {"amount":0.2,"currency":"SGD"}
	// {"amount":0,"currency":"SGD"}
}

func ExampleAmount_equality() {
	x := fpmoney.Amount{Amount: fpdecimal.FromInt(3), Currency: fpmoney.SGD}
	y := fpmoney.Amount{Amount: fpdecimal.FromInt(9), Currency: fpmoney.SGD}
	fmt.Println(y == x.Mul(3))
	// Output: true
}

func ExampleAmount_equality_same_currency() {
	x := fpmoney.Amount{Amount: fpdecimal.FromInt(10), Currency: fpmoney.SGD}
	y := fpmoney.Amount{Amount: fpdecimal.FromInt(10), Currency: fpmoney.SGD}
	fmt.Println(y == x)
	// Output: true
}

func ExampleAmount_equality_wrong_currency() {
	x := fpmoney.Amount{Amount: fpdecimal.FromInt(10), Currency: fpmoney.USD}
	y := fpmoney.Amount{Amount: fpdecimal.FromInt(10), Currency: fpmoney.SGD}
	fmt.Println(y == x)
	// Output: false
}

func ExampleFromFloat() {
	x := fpmoney.Amount{Amount: fpdecimal.FromFloat(144.96), Currency: fpmoney.SGD}
	json.NewEncoder(os.Stdout).Encode(x)
	// Output: {"amount":144.96,"currency":"SGD"}
}

func FuzzArithmetics(f *testing.F) {
	currencies := [...]fpmoney.Currency{
		fpmoney.KRW,
		fpmoney.SGD,
		fpmoney.BHD,
		fpmoney.CLF,
	}

	tests := [][2]int64{
		{1, 2},
		{1, -5},
		{1, 0},
		{1100, -2},
	}
	for _, tc := range tests {
		for i := range currencies {
			f.Add(tc[0], tc[1], i)
		}
	}
	f.Fuzz(func(t *testing.T, a, b int64, c int) {
		if c > len(currencies)-1 || c < 0 {
			t.Skip()
		}
		currency := currencies[c]

		fa := fpmoney.Amount{Amount: fpdecimal.FromIntScaled(a), Currency: currency}
		fb := fpmoney.Amount{Amount: fpdecimal.FromIntScaled(b), Currency: currency}

		zero := fpmoney.Amount{Amount: fpdecimal.Zero, Currency: currency}

		v := []bool{
			// sum commutativity
			fa.Add(fb) == fb.Add(fa),

			// sum associativity
			zero.Add(fa).Add(fb).Add(fa) == zero.Add(fb).Add(fa).Add(fa),

			// sum zero
			fa == fa.Add(fb).Sub(fb),
			fa == fa.Sub(fb).Add(fb),
			zero == zero.Add(fa).Sub(fa),

			// product identity
			fa == fa.Mul(1),

			// product zero
			zero == fa.Mul(0),

			// match number
			(a == b) == (fa == fb),
			a < b == fa.LessThan(fb),
			a > b == fa.GreaterThan(fb),
			a <= b == fa.LessThanOrEqual(fb),
			a >= b == fa.GreaterThanOrEqual(fb),

			// match number convert
			fpmoney.Amount{Amount: fpdecimal.FromIntScaled(a + b), Currency: currency} == fa.Add(fb),
			fpmoney.Amount{Amount: fpdecimal.FromIntScaled(a - b), Currency: currency} == fa.Sub(fb),
		}
		for i, q := range v {
			if !q {
				t.Error(i, a, b, fa, fb)
			}
		}

		if b != 0 {
			w, r := fa.Div(int(b))
			if w != (fpmoney.Amount{Amount: fpdecimal.FromIntScaled(a / b), Currency: currency}) {
				t.Error(w, a/b, a, b, fa)
			}
			if r != (fpmoney.Amount{Amount: fpdecimal.FromIntScaled(a % b), Currency: currency}) {
				t.Error(r, a%b, a, b, fa)
			}
		}
	})
}

func FuzzJSON_MarshalUnmarshal(f *testing.F) {
	currencies := [...]fpmoney.Currency{
		fpmoney.KRW,
		fpmoney.SGD,
		fpmoney.BHD,
		fpmoney.CLF,
	}

	tests := []int64{
		123456,
		0,
		1,
	}
	for _, tc := range tests {
		for i := range currencies {
			f.Add(tc, i)
			f.Add(-tc, i)
		}
	}
	f.Fuzz(func(t *testing.T, v int64, c int) {
		if c > len(currencies)-1 || c < 0 {
			t.Skip()
		}

		currency := currencies[c]
		q := fpmoney.Amount{Amount: fpdecimal.FromIntScaled(v), Currency: currency}

		s, err := json.Marshal(q)
		if err != nil {
			t.Error(err)
		}

		var x fpmoney.Amount
		if err := json.Unmarshal(s, &x); err != nil {
			t.Error(err, string(s))
		}

		if x != q {
			t.Error(x, q, v, c, s)
		}
	})
}

func BenchmarkArithmetic(b *testing.B) {
	x := fpmoney.Amount{Amount: fpdecimal.FromFloat(12312.31), Currency: fpmoney.SGD}
	y := fpmoney.Amount{Amount: fpdecimal.FromFloat(12.02), Currency: fpmoney.SGD}

	b.ResetTimer()
	b.Run("add_x1", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			x = x.Add(y)
		}
	})

	b.ResetTimer()
	b.Run("add_x100", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for i := 0; i < 100; i++ {
				x = x.Add(y)
			}
		}
	})

	if x == (fpmoney.Amount{Amount: fpdecimal.Zero, Currency: fpmoney.SGD}) {
		b.Error()
	}
}

//go:embed testdata/amount-float-large.jsonl
var amountFloatLargeJSONL string

//go:embed testdata/amount-float-small.jsonl
var amountFloatSmallJSONL string

var testsFloats = []struct {
	name string
	vals []string
}{
	{
		name: "small",
		vals: strings.Split(amountFloatSmallJSONL, "\n"),
	},
	{
		name: "large",
		vals: strings.Split(amountFloatLargeJSONL, "\n"),
	},
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	var s fpmoney.Amount
	for _, tc := range testsFloats {
		b.Run(tc.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				v := fpmoney.Amount{Amount: fpdecimal.Zero, Currency: fpmoney.SGD}
				if err := json.Unmarshal([]byte(tc.vals[n%len(tc.vals)]), &s); err != nil || s == v {
					b.Error(s, err)
				}
			}
		})
	}
}

func BenchmarkJSONMarshal(b *testing.B) {
	var s []byte
	var err error
	for _, tc := range testsFloats {
		tests := make([]fpmoney.Amount, 0, len(tc.vals))
		for _, q := range tc.vals {
			var x fpmoney.Amount
			if err := json.Unmarshal([]byte(q), &x); err != nil {
				b.Error(err)
			}
			tests = append(tests, x)
			tests = append(tests, x.Mul(-1))
		}

		b.ResetTimer()
		b.Run(tc.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				s, err = json.Marshal(tests[n%len(tc.vals)])
				if err != nil {
					b.Error(err)
				}
				if string(s) == "" {
					b.Error("empty str")
				}
			}
		})
	}
}

func BenchmarkJSONMarshal_Exact(b *testing.B) {
	var s []byte
	var err error
	for _, tc := range testsFloats {
		tests := make([]fpmoney.Amount, 0, len(tc.vals))
		for _, q := range tc.vals {
			var x fpmoney.Amount
			if err := json.Unmarshal([]byte(q), &x); err != nil {
				b.Error(err)
			}
			tests = append(tests, x)
			tests = append(tests, x.Mul(-1))
		}

		b.ResetTimer()
		b.Run(tc.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				s, err = json.Marshal(tests[n%len(tc.vals)])
				if err != nil {
					b.Error(err)
				}
				if string(s) == "" {
					b.Error("empty str")
				}
			}
		})
	}
}

func TestArithmetic_WrongCurrency(t *testing.T) {
	tests := []struct {
		a fpmoney.Amount
		b fpmoney.Amount
		f func(a, b fpmoney.Amount)
		e *fpmoney.ErrCurrencyMismatch
	}{
		{
			a: fpmoney.Amount{Amount: fpdecimal.FromInt(10), Currency: fpmoney.SGD},
			b: fpmoney.Amount{Amount: fpdecimal.FromInt(11), Currency: fpmoney.USD},
			f: func(a, b fpmoney.Amount) { a.Add(b) },
			e: &fpmoney.ErrCurrencyMismatch{A: fpmoney.SGD, B: fpmoney.USD},
		},
		{
			a: fpmoney.Amount{Amount: fpdecimal.FromInt(10), Currency: fpmoney.SGD},
			b: fpmoney.Amount{Amount: fpdecimal.FromInt(11), Currency: fpmoney.USD},
			f: func(a, b fpmoney.Amount) { a.Sub(b) },
			e: &fpmoney.ErrCurrencyMismatch{A: fpmoney.SGD, B: fpmoney.USD},
		},
		{
			a: fpmoney.Amount{Amount: fpdecimal.FromInt(10), Currency: fpmoney.SGD},
			b: fpmoney.Amount{Amount: fpdecimal.FromInt(11), Currency: fpmoney.USD},
			f: func(a, b fpmoney.Amount) { a.LessThan(b) },
			e: &fpmoney.ErrCurrencyMismatch{A: fpmoney.SGD, B: fpmoney.USD},
		},
		{
			a: fpmoney.Amount{Amount: fpdecimal.FromInt(10), Currency: fpmoney.SGD},
			b: fpmoney.Amount{Amount: fpdecimal.FromInt(11), Currency: fpmoney.USD},
			f: func(a, b fpmoney.Amount) { a.LessThanOrEqual(b) },
			e: &fpmoney.ErrCurrencyMismatch{A: fpmoney.SGD, B: fpmoney.USD},
		},
		{
			a: fpmoney.Amount{Amount: fpdecimal.FromInt(10), Currency: fpmoney.SGD},
			b: fpmoney.Amount{Amount: fpdecimal.FromInt(11), Currency: fpmoney.USD},
			f: func(a, b fpmoney.Amount) { a.GreaterThan(b) },
			e: &fpmoney.ErrCurrencyMismatch{A: fpmoney.SGD, B: fpmoney.USD},
		},
		{
			a: fpmoney.Amount{Amount: fpdecimal.FromInt(10), Currency: fpmoney.SGD},
			b: fpmoney.Amount{Amount: fpdecimal.FromInt(11), Currency: fpmoney.USD},
			f: func(a, b fpmoney.Amount) { a.GreaterThanOrEqual(b) },
			e: &fpmoney.ErrCurrencyMismatch{A: fpmoney.SGD, B: fpmoney.USD},
		},
	}
	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				re, ok := r.(error)
				if !ok {
					t.Error(r)
				}
				if err := fpmoney.NewErrCurrencyMismatch(); !errors.As(re, &err) || *err != *tc.e {
					t.Error(re, tc.e)
				}
			}()
			tc.f(tc.a, tc.b)
		})
	}
}

func TestErrCurrencyMismatch_Error(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		e := &fpmoney.ErrCurrencyMismatch{A: fpmoney.SGD, B: fpmoney.USD}
		if e.Error() != "SGD != USD" {
			t.Error(e)
		}
	})
}
