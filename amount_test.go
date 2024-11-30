package fpmoney_test

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/nikolaydubina/fpmoney"
)

func ExampleAmount() {
	var BuySP500Price = fpmoney.FromInt(9000, fpmoney.SGD)

	input := []byte(`{"sp500": {"amount": 9000.02, "currency": "SGD"}}`)

	type Stonks struct {
		SP500 fpmoney.Amount `json:"sp500"`
	}
	var v Stonks
	if err := json.Unmarshal(input, &v); err != nil {
		log.Fatal(err)
	}

	amountToBuy := fpmoney.FromInt(0, fpmoney.SGD)
	if v.SP500.GreaterThan(BuySP500Price) {
		amountToBuy = amountToBuy.Add(v.SP500.Mul(2))
	}

	fmt.Println(amountToBuy)
	// Output: SGD18000.04
}

func ExampleAmount_DivMod_part() {
	x := fpmoney.FromInt(1, fpmoney.SGD)
	a, r := x.DivMod(3)
	fmt.Println(a, r)
	// Output: SGD0.33 SGD0.01
}

func ExampleAmount_DivMod_whole() {
	x := fpmoney.FromInt(1, fpmoney.SGD)
	a, r := x.DivMod(5)
	fmt.Println(a, r)
	// Output: SGD0.2 SGD0
}

func ExampleAmount_equality() {
	x := fpmoney.FromInt(3, fpmoney.SGD)
	y := fpmoney.FromInt(9, fpmoney.SGD)
	fmt.Println(y == x.Mul(3))
	// Output: true
}

func ExampleAmount_equality_same_currency() {
	x := fpmoney.FromInt(10, fpmoney.SGD)
	y := fpmoney.FromInt(10, fpmoney.SGD)
	fmt.Println(y == x)
	// Output: true
}

func ExampleAmount_equality_wrong_currency() {
	x := fpmoney.FromInt(10, fpmoney.USD)
	y := fpmoney.FromInt(10, fpmoney.SGD)
	fmt.Println(y == x)
	// Output: false
}

func ExampleFromFloat() {
	x := fpmoney.FromFloat(144.96, fpmoney.SGD)
	fmt.Println(x)
	// Output: SGD144.96
}

func ExampleAmount_Scaled_fractions() {
	v := fpmoney.FromFloat(42.23, fpmoney.EUR)
	fmt.Println(v.Scaled())
	// Output: 4223
}

func ExampleAmount_Scaled_many_fractions() {
	v := fpmoney.FromFloat(17.0, fpmoney.BHD)
	fmt.Println(v.Scaled())
	// Output: 17000
}

func ExampleAmount_Scaled_large() {
	v := fpmoney.FromFloat(8764534896.42, fpmoney.USD)
	fmt.Println(v.Scaled())
	// Output: 876453489642
}
func ExampleAmount_Scaled_whole() {
	v := fpmoney.FromFloat(23.0, fpmoney.EUR)
	fmt.Println(v.Scaled())
	// Output: 2300
}

func ExampleAmount_Scaled_from_scaled() {
	v := fpmoney.FromIntScaled(17, fpmoney.EUR)
	fmt.Println(v.Scaled())
	// Output: 17
}

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		s string
		v fpmoney.Amount
	}{
		// 2 cents
		{
			s: `{"currency": "SGD","amount": 9002.01}`,
			v: fpmoney.FromIntScaled(900201, fpmoney.SGD),
		},
		{
			s: `{"amount": 9002.01, "currency": "SGD"}`,
			v: fpmoney.FromIntScaled(900201, fpmoney.SGD),
		},
		{
			s: `{"amount": -9002.01, "currency": "SGD"}`,
			v: fpmoney.FromIntScaled(-900201, fpmoney.SGD),
		},
		{
			s: `{"amount": 0, "currency": "SGD"}`,
			v: fpmoney.FromIntScaled(0, fpmoney.SGD),
		},
		{
			s: `{"amount": 0.01, "currency": "SGD"}`,
			v: fpmoney.FromIntScaled(1, fpmoney.SGD),
		},
		{
			s: `{"amount": -0.01, "currency": "SGD"}`,
			v: fpmoney.FromIntScaled(-1, fpmoney.SGD),
		},
		// 0 cents
		{
			s: `{"amount":1,"currency":"KRW"}`,
			v: fpmoney.FromIntScaled(1, fpmoney.KRW),
		},
		{
			s: `{"amount":-1,"currency":"KRW"}`,
			v: fpmoney.FromIntScaled(-1, fpmoney.KRW),
		},
		{
			s: `{"amount":123,"currency":"KRW"}`,
			v: fpmoney.FromIntScaled(123, fpmoney.KRW),
		},
		{
			s: `{"amount":-123,"currency":"KRW"}`,
			v: fpmoney.FromIntScaled(-123, fpmoney.KRW),
		},
		// 2 cents strange valid input
		{
			s: `    {   "amount" : 9002.01
            
            , 
            
            "currency"
                : 
            
            "SGD"}   


             `,
			v: fpmoney.FromIntScaled(900201, fpmoney.SGD),
		},
	}
	for _, tc := range tests {
		t.Run(tc.s, func(t *testing.T) {
			var v fpmoney.Amount
			err := json.Unmarshal([]byte(tc.s), &v)
			if err != nil {
				t.Error(err)
			}
			if tc.v != v {
				t.Error(tc.v, v)
			}
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	tests := []struct {
		s string
		v fpmoney.Amount
	}{
		// 2 cents
		{
			s: `{"amount":9002.01,"currency":"SGD"}`,
			v: fpmoney.FromIntScaled(900201, fpmoney.SGD),
		},
		{
			s: `{"amount":0,"currency":"SGD"}`,
			v: fpmoney.FromIntScaled(0, fpmoney.SGD),
		},
		{
			s: `{"amount":0.01,"currency":"SGD"}`,
			v: fpmoney.FromIntScaled(1, fpmoney.SGD),
		},
		{
			s: `{"amount":-0.01,"currency":"SGD"}`,
			v: fpmoney.FromIntScaled(-1, fpmoney.SGD),
		},
		{
			s: `{"amount":1.01,"currency":"SGD"}`,
			v: fpmoney.FromIntScaled(101, fpmoney.SGD),
		},
		{
			s: `{"amount":-1.01,"currency":"SGD"}`,
			v: fpmoney.FromIntScaled(-101, fpmoney.SGD),
		},
		{
			s: `{"amount":1,"currency":"SGD"}`,
			v: fpmoney.FromInt(1, fpmoney.SGD),
		},
		{
			s: `{"amount":5,"currency":"SGD"}`,
			v: fpmoney.FromInt(5, fpmoney.SGD),
		},
		{
			s: `{"amount":-1,"currency":"SGD"}`,
			v: fpmoney.FromInt(-1, fpmoney.SGD),
		},
		{
			s: `{"amount":-5,"currency":"SGD"}`,
			v: fpmoney.FromInt(-5, fpmoney.SGD),
		},
		// 0 cents
		{
			s: `{"amount":1,"currency":"KRW"}`,
			v: fpmoney.FromIntScaled(1, fpmoney.KRW),
		},
		{
			s: `{"amount":-1,"currency":"KRW"}`,
			v: fpmoney.FromIntScaled(-1, fpmoney.KRW),
		},
		{
			s: `{"amount":123,"currency":"KRW"}`,
			v: fpmoney.FromIntScaled(123, fpmoney.KRW),
		},
		{
			s: `{"amount":-123,"currency":"KRW"}`,
			v: fpmoney.FromIntScaled(-123, fpmoney.KRW),
		},
	}
	for _, tc := range tests {
		t.Run(tc.s, func(t *testing.T) {
			s, err := json.Marshal(tc.v)
			if err != nil {
				t.Error(err)
			}
			if tc.s != string(s) {
				t.Error(tc.s, string(s))
			}
		})
	}
}

func FuzzJSON_MarshalUnmarshal(f *testing.F) {
	currencies := [...]fpmoney.Currency{
		fpmoney.KRW,
		fpmoney.SGD,
		fpmoney.BHD,
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
		q := fpmoney.FromIntScaled(v, currency)

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
	x := fpmoney.FromFloat(12312.31, fpmoney.SGD)
	y := fpmoney.FromFloat(12.02, fpmoney.SGD)

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

	if x == fpmoney.FromInt(0, fpmoney.SGD) {
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
				err := json.Unmarshal([]byte(tc.vals[n%len(tc.vals)]), &s)
				if err != nil || s == fpmoney.FromInt(0, fpmoney.SGD) {
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
			a: fpmoney.FromInt(10, fpmoney.SGD),
			b: fpmoney.FromInt(11, fpmoney.USD),
			f: func(a, b fpmoney.Amount) { a.Add(b) },
			e: &fpmoney.ErrCurrencyMismatch{A: fpmoney.SGD, B: fpmoney.USD},
		},
		{
			a: fpmoney.FromInt(10, fpmoney.SGD),
			b: fpmoney.FromInt(11, fpmoney.USD),
			f: func(a, b fpmoney.Amount) { a.Sub(b) },
			e: &fpmoney.ErrCurrencyMismatch{A: fpmoney.SGD, B: fpmoney.USD},
		},
		{
			a: fpmoney.FromInt(10, fpmoney.SGD),
			b: fpmoney.FromInt(11, fpmoney.USD),
			f: func(a, b fpmoney.Amount) { a.LessThan(b) },
			e: &fpmoney.ErrCurrencyMismatch{A: fpmoney.SGD, B: fpmoney.USD},
		},
		{
			a: fpmoney.FromInt(10, fpmoney.SGD),
			b: fpmoney.FromInt(11, fpmoney.USD),
			f: func(a, b fpmoney.Amount) { a.LessThanOrEqual(b) },
			e: &fpmoney.ErrCurrencyMismatch{A: fpmoney.SGD, B: fpmoney.USD},
		},
		{
			a: fpmoney.FromInt(10, fpmoney.SGD),
			b: fpmoney.FromInt(11, fpmoney.USD),
			f: func(a, b fpmoney.Amount) { a.GreaterThan(b) },
			e: &fpmoney.ErrCurrencyMismatch{A: fpmoney.SGD, B: fpmoney.USD},
		},
		{
			a: fpmoney.FromInt(10, fpmoney.SGD),
			b: fpmoney.FromInt(11, fpmoney.USD),
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
