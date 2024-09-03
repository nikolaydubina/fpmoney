package fpmoney

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
)

func ExampleAmount() {
	var BuySP500Price = FromInt(9000, SGD)

	input := []byte(`{"sp500": {"amount": 9000.02, "currency": "SGD"}}`)

	type Stonks struct {
		SP500 Amount `json:"sp500"`
	}
	var v Stonks
	if err := json.Unmarshal(input, &v); err != nil {
		log.Fatal(err)
	}

	amountToBuy := FromInt(0, SGD)
	if v.SP500.GreaterThan(BuySP500Price) {
		amountToBuy = amountToBuy.Add(v.SP500.Mul(2))
	}

	fmt.Println(amountToBuy)
	// Output: 18000.04 SGD
}

func ExampleAmount_Div_part() {
	x := FromInt(1, SGD)
	a, r := x.Div(3)
	fmt.Println(a, r)
	// Output: 0.33 SGD 0.01 SGD
}

func ExampleAmount_Div_whole() {
	x := FromInt(1, SGD)
	a, r := x.Div(5)
	fmt.Println(a, r)
	// Output: 0.2 SGD 0 SGD
}

func ExampleAmount_equality() {
	x := FromInt(3, SGD)
	y := FromInt(9, SGD)
	fmt.Println(y == x.Mul(3))
	// Output: true
}

func ExampleAmount_equality_same_currency() {
	x := FromInt(10, SGD)
	y := FromInt(10, SGD)
	fmt.Println(y == x)
	// Output: true
}

func ExampleAmount_equality_wrong_currency() {
	x := FromInt(10, USD)
	y := FromInt(10, SGD)
	fmt.Println(y == x)
	// Output: false
}

func ExampleFromFloat() {
	x := FromFloat(144.96, SGD)
	fmt.Println(x)
	// Output: 144.96 SGD
}

func TestScaled(t *testing.T) {
	data := []struct {
		a        Amount
		expected int64
	}{
		{FromIntScaled(17, EUR), 17},
		{FromFloat(23.0, EUR), 2300},
		{fromJSONOrDie(`{"amount": 42.23, currency: "EUR"}`), 4223},
		{fromJSONOrDie(`{"amount": 17, currency: "CLF"}`), 170000},
		{fromJSONOrDie(`{"amount": 8764534896.42, currency: "USD"}`), 876453489642},
	}

	for _, d := range data {
		res := d.a.Scaled()
		if res != d.expected {
			t.Errorf("%v: expected: %d, got: %d", d, d.expected, res)
		}
	}
}

func FuzzArithmetics(f *testing.F) {
	currencies := [...]Currency{
		KRW,
		SGD,
		BHD,
		CLF,
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

		fa := FromIntScaled(a, currency)
		fb := FromIntScaled(b, currency)

		v := []bool{
			// sum commutativity
			fa.Add(fb) == fb.Add(fa),

			// sum associativity
			FromInt(0, currency).Add(fa).Add(fb).Add(fa) == FromInt(0, currency).Add(fb).Add(fa).Add(fa),

			// sum zero
			fa == fa.Add(fb).Sub(fb),
			fa == fa.Sub(fb).Add(fb),
			FromInt(0, currency) == FromInt(0, currency).Add(fa).Sub(fa),

			// product identity
			fa == fa.Mul(1),

			// product zero
			FromInt(0, currency) == fa.Mul(0),

			// match number
			(a == b) == (fa == fb),
			a < b == fa.LessThan(fb),
			a > b == fa.GreaterThan(fb),
			a <= b == fa.LessThanOrEqual(fb),
			a >= b == fa.GreaterThanOrEqual(fb),

			// match number convert
			FromIntScaled(a+b, currency) == fa.Add(fb),
			FromIntScaled(a-b, currency) == fa.Sub(fb),
		}
		for i, q := range v {
			if !q {
				t.Error(i, a, b, fa, fb)
			}
		}

		if b != 0 {
			w, r := fa.Div(int(b))
			if w != FromIntScaled(a/b, currency) {
				t.Error(w, a/b, a, b, fa)
			}
			if r != FromIntScaled(a%b, currency) {
				t.Error(r, a%b, a, b, fa)
			}
		}
	})
}

func FuzzJSONUnmarshal_Float(f *testing.F) {
	currencies := [...]Currency{
		KRW,
		SGD,
		BHD,
		CLF,
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

		var x Amount
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
		var x Amount
		err := json.Unmarshal([]byte(s), &x)
		if err != nil {
			if (x != Amount{}) {
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
		a := FromIntScaled(v, KRW)

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
		v Amount
	}{
		// 2 cents
		{
			s: `{"currency": "SGD","amount": 9002.01}`,
			v: FromIntScaled(900201, SGD),
		},
		{
			s: `{"amount": 9002.01, "currency": "SGD"}`,
			v: FromIntScaled(900201, SGD),
		},
		{
			s: `{"amount": -9002.01, "currency": "SGD"}`,
			v: FromIntScaled(-900201, SGD),
		},
		{
			s: `{"amount": 0, "currency": "SGD"}`,
			v: FromIntScaled(0, SGD),
		},
		{
			s: `{"amount": 0.01, "currency": "SGD"}`,
			v: FromIntScaled(1, SGD),
		},
		{
			s: `{"amount": -0.01, "currency": "SGD"}`,
			v: FromIntScaled(-1, SGD),
		},
		// 0 cents
		{
			s: `{"amount":1,"currency":"KRW"}`,
			v: FromIntScaled(1, KRW),
		},
		{
			s: `{"amount":-1,"currency":"KRW"}`,
			v: FromIntScaled(-1, KRW),
		},
		{
			s: `{"amount":123,"currency":"KRW"}`,
			v: FromIntScaled(123, KRW),
		},
		{
			s: `{"amount":-123,"currency":"KRW"}`,
			v: FromIntScaled(-123, KRW),
		},
		// 2 cents strange valid input
		{
			s: `    {   "amount" : 9002.01
            
            , 
            
            "currency"
                : 
            
            "SGD"}   


             `,
			v: FromIntScaled(900201, SGD),
		},
	}
	for _, tc := range tests {
		t.Run(tc.s, func(t *testing.T) {
			var v Amount
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
		v Amount
	}{
		// 2 cents
		{
			s: `{"amount":9002.01,"currency":"SGD"}`,
			v: FromIntScaled(900201, SGD),
		},
		{
			s: `{"amount":0,"currency":"SGD"}`,
			v: FromIntScaled(0, SGD),
		},
		{
			s: `{"amount":0.01,"currency":"SGD"}`,
			v: FromIntScaled(1, SGD),
		},
		{
			s: `{"amount":-0.01,"currency":"SGD"}`,
			v: FromIntScaled(-1, SGD),
		},
		{
			s: `{"amount":1.01,"currency":"SGD"}`,
			v: FromIntScaled(101, SGD),
		},
		{
			s: `{"amount":-1.01,"currency":"SGD"}`,
			v: FromIntScaled(-101, SGD),
		},
		{
			s: `{"amount":1,"currency":"SGD"}`,
			v: FromInt(1, SGD),
		},
		{
			s: `{"amount":5,"currency":"SGD"}`,
			v: FromInt(5, SGD),
		},
		{
			s: `{"amount":-1,"currency":"SGD"}`,
			v: FromInt(-1, SGD),
		},
		{
			s: `{"amount":-5,"currency":"SGD"}`,
			v: FromInt(-5, SGD),
		},
		// 0 cents
		{
			s: `{"amount":1,"currency":"KRW"}`,
			v: FromIntScaled(1, KRW),
		},
		{
			s: `{"amount":-1,"currency":"KRW"}`,
			v: FromIntScaled(-1, KRW),
		},
		{
			s: `{"amount":123,"currency":"KRW"}`,
			v: FromIntScaled(123, KRW),
		},
		{
			s: `{"amount":-123,"currency":"KRW"}`,
			v: FromIntScaled(-123, KRW),
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
	currencies := [...]Currency{
		KRW,
		SGD,
		BHD,
		CLF,
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
		q := FromIntScaled(v, currency)

		s, err := json.Marshal(q)
		if err != nil {
			t.Error(err)
		}

		var x Amount
		if err := json.Unmarshal(s, &x); err != nil {
			t.Error(err, string(s))
		}

		if x != q {
			t.Error(x, q, v, c, s)
		}
	})
}

func BenchmarkArithmetic(b *testing.B) {
	x := FromFloat(12312.31, SGD)
	y := FromFloat(12.02, SGD)

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

	if x == FromInt(0, SGD) {
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
	var s Amount
	for _, tc := range testsFloats {
		b.Run(tc.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				err := json.Unmarshal([]byte(tc.vals[n%len(tc.vals)]), &s)
				if err != nil || s == FromInt(0, SGD) {
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
		tests := make([]Amount, 0, len(tc.vals))
		for _, q := range tc.vals {
			var x Amount
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
		tests := make([]Amount, 0, len(tc.vals))
		for _, q := range tc.vals {
			var x Amount
			if err := json.Unmarshal([]byte(q), &x); err != nil {
				b.Error(err)
			}
			tests = append(tests, x)
			tests = append(tests, x.Mul(-1))
		}

		b.ResetTimer()
		b.Run(tc.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				s, err = tests[n%len(tc.vals)].MarshalJSON()
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
		a Amount
		b Amount
		f func(a, b Amount)
		e *ErrCurrencyMismatch
	}{
		{
			a: FromInt(10, SGD),
			b: FromInt(11, USD),
			f: func(a, b Amount) { a.Add(b) },
			e: &ErrCurrencyMismatch{A: SGD, B: USD},
		},
		{
			a: FromInt(10, SGD),
			b: FromInt(11, USD),
			f: func(a, b Amount) { a.Sub(b) },
			e: &ErrCurrencyMismatch{A: SGD, B: USD},
		},
		{
			a: FromInt(10, SGD),
			b: FromInt(11, USD),
			f: func(a, b Amount) { a.LessThan(b) },
			e: &ErrCurrencyMismatch{A: SGD, B: USD},
		},
		{
			a: FromInt(10, SGD),
			b: FromInt(11, USD),
			f: func(a, b Amount) { a.LessThanOrEqual(b) },
			e: &ErrCurrencyMismatch{A: SGD, B: USD},
		},
		{
			a: FromInt(10, SGD),
			b: FromInt(11, USD),
			f: func(a, b Amount) { a.GreaterThan(b) },
			e: &ErrCurrencyMismatch{A: SGD, B: USD},
		},
		{
			a: FromInt(10, SGD),
			b: FromInt(11, USD),
			f: func(a, b Amount) { a.GreaterThanOrEqual(b) },
			e: &ErrCurrencyMismatch{A: SGD, B: USD},
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
				if err := NewErrCurrencyMismatch(); !errors.As(re, &err) || *err != *tc.e {
					t.Error(re, tc.e)
				}
			}()
			tc.f(tc.a, tc.b)
		})
	}
}

func TestErrCurrencyMismatch_Error(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		e := &ErrCurrencyMismatch{A: SGD, B: USD}
		if e.Error() != "SGD != USD" {
			t.Error(e)
		}
	})
}

func fromJSONOrDie(json string) Amount {
	var a Amount
	if err := a.UnmarshalJSON([]byte(json)); err != nil {
		panic(err)
	}
	return a
}
