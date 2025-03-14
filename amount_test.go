package fpmoney_test

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
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
	// Output: 18000.04 SGD
}

func ExampleAmount_DivMod_part() {
	x := fpmoney.FromInt(1, fpmoney.SGD)
	a, r := x.DivMod(3)
	fmt.Println(a, r)
	// Output: 0.33 SGD 0.01 SGD
}

func ExampleAmount_DivMod_whole() {
	x := fpmoney.FromInt(1, fpmoney.SGD)
	a, r := x.DivMod(5)
	fmt.Println(a, r)
	// Output: 0.2 SGD 0 SGD
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
	// Output: 144.96 SGD
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

func FuzzArithmetics(f *testing.F) {
	currencies := [...]fpmoney.Currency{
		fpmoney.KRW,
		fpmoney.SGD,
		fpmoney.BHD,
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

		fa := fpmoney.FromIntScaled(a, currency)
		fb := fpmoney.FromIntScaled(b, currency)

		v := []bool{
			// sum commutativity
			fa.Add(fb) == fb.Add(fa),

			// sum associativity
			fpmoney.FromInt(0, currency).Add(fa).Add(fb).Add(fa) == fpmoney.FromInt(0, currency).Add(fb).Add(fa).Add(fa),

			// sum zero
			fa == fa.Add(fb).Sub(fb),
			fa == fa.Sub(fb).Add(fb),
			fpmoney.FromInt(0, currency) == fpmoney.FromInt(0, currency).Add(fa).Sub(fa),

			// product identity
			fa == fa.Mul(1),

			// product zero
			fpmoney.FromInt(0, currency) == fa.Mul(0),

			// match number
			(a == b) == (fa == fb),
			a < b == fa.LessThan(fb),
			a > b == fa.GreaterThan(fb),
			a <= b == fa.LessThanOrEqual(fb),
			a >= b == fa.GreaterThanOrEqual(fb),

			// match number convert
			fpmoney.FromIntScaled(a+b, currency) == fa.Add(fb),
			fpmoney.FromIntScaled(a-b, currency) == fa.Sub(fb),
		}
		for i, q := range v {
			if !q {
				t.Error(i, a, b, fa, fb)
			}
		}

		if b != 0 {
			w, r := fa.DivMod(int(b))
			if w != fpmoney.FromIntScaled(a/b, currency) {
				t.Error(w, a/b, a, b, fa)
			}
			if r != fpmoney.FromIntScaled(a%b, currency) {
				t.Error(r, a%b, a, b, fa)
			}
		}
	})
}

func FuzzJSONUnmarshal_Float(f *testing.F) {
	currencies := [...]fpmoney.Currency{
		fpmoney.KRW,
		fpmoney.SGD,
		fpmoney.BHD,
	}

	tests := []float32{
		0,
		0.100,
		0.101,
		0.010,
		0.001,
		0.0001,
		0.123,
		0.103,
		0.100001,
		12.001,
		12.010,
		12.345,
		1,
		2,
		10,
		12345678,
	}
	for _, tc := range tests {
		for i := range currencies {
			f.Add(tc, i, uint8(5))
			f.Add(-tc, i, uint8(5))
		}
	}
	f.Fuzz(func(t *testing.T, r float32, c int, nf uint8) {
		if c > len(currencies)-1 || c < 0 {
			t.Skip()
		}
		if nf > 10 {
			t.Skip()
		}
		var l float32 = 10000000
		if r > l || r < -l {
			t.Skip()
		}
		if c == 0 {
			t.Skip()
		}

		currency := currencies[c]

		fs := `%.` + strconv.Itoa(int(nf)) + `f`
		rs := fmt.Sprintf(fs, r)
		s := fmt.Sprintf(`{"amount": %s, "currency": "%s"}`, rs, currency.String())
		if _, err := fmt.Sscanf(rs, "%f", &r); err != nil {
			t.Error(err)
		}

		if r == -0 {
			t.Skip()
		}

		var x fpmoney.Amount
		if err := json.Unmarshal([]byte(s), &x); err != nil {
			t.Error(rs, currency, err)
		}

		if x.Currency() != currency {
			t.Error(x, currency)
		}
	})
}

func FuzzJSONUnmarshal_NoPanic(f *testing.F) {
	amounts := []string{
		"123.456",
		"0.123",
		"0.1",
		"0.01",
		"0.001",
		"0.000",
		"0.123.2",
		"0..1",
		"0.1.2",
		"123.1o2",
		"--123",
		"00000.123",
		"-",
		"",
		"123456",
	}
	currencies := []string{
		"SGD",
		"SGDSGD",
		"",
		"123",
		"'SGD'",
		`"TDF"`,
	}
	for _, a := range amounts {
		for _, c := range currencies {
			f.Add(a)
			f.Add("-" + a)
			f.Add(fmt.Sprintf(`{"amount": %s, "currency": %s}`, a, c))
			f.Add(fmt.Sprintf(`{"amount": -%s, "currency": %s}`, a, c))
			f.Add(fmt.Sprintf(`{"amount": -%s, "currency": %s}`, a, c))
			f.Add(fmt.Sprintf(`"amount": -%s, "currency": %s`, a, c))
			f.Add(fmt.Sprintf(`{"amount": -%s}`, a))
			f.Add(fmt.Sprintf(`{"currency": %s}`, c))
			f.Add(fmt.Sprintf(`{"amount": %s, "currency": %s}`, c, a))
			f.Add(fmt.Sprintf(`"amount": %s, "currency": %s}`, a, c))
			f.Add(fmt.Sprintf(`{"amount": %s, "currency": %s`, c, a))
			f.Add(fmt.Sprintf(`{"amount": %s,,""""currency": %s}`, a, c))
		}
	}

	f.Add(`{"amount": 123.32, "currency":""}`)
	f.Add(`{"amount": , "currency":""}`)
	f.Add(`{"amount":,"currency":""}`)

	f.Fuzz(func(t *testing.T, s string) {
		var x fpmoney.Amount
		err := json.Unmarshal([]byte(s), &x)
		if err != nil {
			if (x != fpmoney.Amount{}) {
				t.Errorf("has to be 0 on error")
			}
			return
		}
	})
}

func FuzzToFloat(f *testing.F) {
	tests := []int64{
		0,
		1,
		123456,
	}
	for _, tc := range tests {
		f.Add(tc)
		f.Add(-tc)
	}
	f.Fuzz(func(t *testing.T, v int64) {
		a := fpmoney.FromIntScaled(v, fpmoney.KRW)

		if float32(v) != a.Float32() {
			t.Error(a, a.Float32(), float32(v))
		}

		if float64(v) != a.Float64() {
			t.Error(a, a.Float64(), v)
		}
	})
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
	x := fpmoney.FromFloat(rand.Float32(), fpmoney.SGD)
	y := fpmoney.FromFloat(rand.Float32(), fpmoney.SGD)

	b.Run("add", func(b *testing.B) {
		for b.Loop() {
			x.Add(y)
		}
	})
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

func BenchmarkJSON(b *testing.B) {
	for _, tc := range testsFloats {
		b.Run(tc.name, func(b *testing.B) {
			tests := make([]fpmoney.Amount, 0, len(tc.vals))
			for _, q := range tc.vals {
				var x fpmoney.Amount
				if err := json.Unmarshal([]byte(q), &x); err != nil {
					b.Error(err)
				}
				tests = append(tests, x)
				tests = append(tests, x.Mul(-1))
			}

			b.Run("encode", func(b *testing.B) {
				v := tests[rand.Intn(len(tests))]

				for b.Loop() {
					_, _ = json.Marshal(v)
				}
			})

			b.Run("decode", func(b *testing.B) {
				var s fpmoney.Amount
				v := []byte(tc.vals[rand.Intn(len(tc.vals))])

				for b.Loop() {
					json.Unmarshal(v, &s)
				}
			})
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
