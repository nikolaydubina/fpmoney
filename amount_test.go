package fpmoney_test

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"testing"
	"unsafe"

	"github.com/ferdypruis/iso4217"
	"github.com/nikolaydubina/fpmoney"
)

func ExampleAmount() {
	var BuySP500Price = fpmoney.FromInt(9000, iso4217.SGD)

	input := []byte(`{"sp500": {"amount": 9000.02, "currency": "SGD"}}`)

	type Stonks struct {
		SP500 fpmoney.Amount `json:"sp500"`
	}
	var v Stonks
	if err := json.Unmarshal(input, &v); err != nil {
		log.Fatal(err)
	}

	amountToBuy := fpmoney.FromInt(0, iso4217.SGD)
	if v.SP500.GreaterThan(BuySP500Price) {
		amountToBuy = amountToBuy.Add(v.SP500.Mul(2))
	}

	fmt.Println(amountToBuy)
	// Output: 18000.04 SGD
}

func ExampleDiv_remainder() {
	x := fpmoney.FromInt(1, iso4217.SGD)
	a, r := x.Div(3)
	fmt.Println(a, r)
	// Output: 0.33 SGD 0.01 SGD
}

func ExampleDiv_whole() {
	x := fpmoney.FromInt(1, iso4217.SGD)
	a, r := x.Div(5)
	fmt.Println(a, r)
	// Output: 0.20 SGD 0 SGD
}

func FuzzArithmetics(f *testing.F) {
	tests := [][2]int64{
		{1, 2},
		{1, -5},
		{1, 0},
		{1100, -2},
	}
	for _, tc := range tests {
		f.Add(tc[0], tc[1], uint16(iso4217.KRW))
		f.Add(tc[0], tc[1], uint16(iso4217.SGD))
		f.Add(tc[0], tc[1], uint16(iso4217.BHD))
		f.Add(tc[0], tc[1], uint16(iso4217.CLF))
	}
	f.Fuzz(func(t *testing.T, a, b int64, c uint16) {
		if c > 300 {
			t.Skip()
		}
		currency := iso4217.Currency(c)

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
			w, r := fa.Div(int(b))
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
		f.Add(tc, uint16(iso4217.KRW), uint8(5))
		f.Add(tc, uint16(iso4217.SGD), uint8(5))
		f.Add(tc, uint16(iso4217.BHD), uint8(5))
		f.Add(tc, uint16(iso4217.CLF), uint8(5))

		f.Add(-tc, uint16(iso4217.KRW), uint8(5))
		f.Add(-tc, uint16(iso4217.SGD), uint8(5))
		f.Add(-tc, uint16(iso4217.BHD), uint8(5))
		f.Add(-tc, uint16(iso4217.CLF), uint8(5))
	}
	f.Fuzz(func(t *testing.T, r float32, c uint16, nf uint8) {
		if c > 300 {
			t.Skip()
		}
		if nf > 10 {
			t.Skip()
		}
		var l float32 = 10000000
		if r > l || r < -l {
			t.Skip()
		}

		currency := iso4217.Currency(c)

		fs := `%.` + strconv.Itoa(int(nf)) + `f`
		rs := fmt.Sprintf(fs, r)
		s := fmt.Sprintf(`{"amount": %s, "currency": "%s"}`, rs, currency.Alpha())
		if _, err := fmt.Sscanf(rs, "%f", &r); err != nil {
			t.Error(err)
		}

		if r == -0 {
			t.Skip()
		}

		var x fpmoney.Amount
		if err := json.Unmarshal([]byte(s), &x); err != nil {
			t.Error(err)
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
		"sgd",
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
		}
	}
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
		a := fpmoney.FromIntScaled(v, iso4217.KRW)

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
			s: `{"amount": 9002.01, "currency": "SGD"}`,
			v: fpmoney.FromIntScaled(900201, iso4217.SGD),
		},
		{
			s: `{"amount": -9002.01, "currency": "SGD"}`,
			v: fpmoney.FromIntScaled(-900201, iso4217.SGD),
		},
		{
			s: `{"amount": 0, "currency": "SGD"}`,
			v: fpmoney.FromIntScaled(0, iso4217.SGD),
		},
		{
			s: `{"amount": 0.01, "currency": "SGD"}`,
			v: fpmoney.FromIntScaled(1, iso4217.SGD),
		},
		{
			s: `{"amount": -0.01, "currency": "SGD"}`,
			v: fpmoney.FromIntScaled(-1, iso4217.SGD),
		},
		// 0 cents
		{
			s: `{"amount":1,"currency":"KRW"}`,
			v: fpmoney.FromIntScaled(1, iso4217.KRW),
		},
		{
			s: `{"amount":-1,"currency":"KRW"}`,
			v: fpmoney.FromIntScaled(-1, iso4217.KRW),
		},
		{
			s: `{"amount":123,"currency":"KRW"}`,
			v: fpmoney.FromIntScaled(123, iso4217.KRW),
		},
		{
			s: `{"amount":-123,"currency":"KRW"}`,
			v: fpmoney.FromIntScaled(-123, iso4217.KRW),
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
			v: fpmoney.FromIntScaled(900201, iso4217.SGD),
		},
		{
			s: `{"amount":0,"currency":"SGD"}`,
			v: fpmoney.FromIntScaled(0, iso4217.SGD),
		},
		{
			s: `{"amount":0.01,"currency":"SGD"}`,
			v: fpmoney.FromIntScaled(1, iso4217.SGD),
		},
		{
			s: `{"amount":-0.01,"currency":"SGD"}`,
			v: fpmoney.FromIntScaled(-1, iso4217.SGD),
		},
		{
			s: `{"amount":1.01,"currency":"SGD"}`,
			v: fpmoney.FromIntScaled(101, iso4217.SGD),
		},
		{
			s: `{"amount":-1.01,"currency":"SGD"}`,
			v: fpmoney.FromIntScaled(-101, iso4217.SGD),
		},
		{
			s: `{"amount":1.00,"currency":"SGD"}`,
			v: fpmoney.FromInt(1, iso4217.SGD),
		},
		{
			s: `{"amount":5.00,"currency":"SGD"}`,
			v: fpmoney.FromInt(5, iso4217.SGD),
		},
		{
			s: `{"amount":-1.00,"currency":"SGD"}`,
			v: fpmoney.FromInt(-1, iso4217.SGD),
		},
		{
			s: `{"amount":-5.00,"currency":"SGD"}`,
			v: fpmoney.FromInt(-5, iso4217.SGD),
		},
		// 0 cents
		{
			s: `{"amount":1,"currency":"KRW"}`,
			v: fpmoney.FromIntScaled(1, iso4217.KRW),
		},
		{
			s: `{"amount":-1,"currency":"KRW"}`,
			v: fpmoney.FromIntScaled(-1, iso4217.KRW),
		},
		{
			s: `{"amount":123,"currency":"KRW"}`,
			v: fpmoney.FromIntScaled(123, iso4217.KRW),
		},
		{
			s: `{"amount":-123,"currency":"KRW"}`,
			v: fpmoney.FromIntScaled(-123, iso4217.KRW),
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
	tests := []int64{
		123456,
		0,
		1,
	}
	for _, tc := range tests {
		f.Add(tc, uint16(iso4217.KRW))
		f.Add(tc, uint16(iso4217.SGD))
		f.Add(tc, uint16(iso4217.BHD))
		f.Add(tc, uint16(iso4217.CLF))

		f.Add(-tc, uint16(iso4217.KRW))
		f.Add(-tc, uint16(iso4217.SGD))
		f.Add(-tc, uint16(iso4217.BHD))
		f.Add(-tc, uint16(iso4217.CLF))
	}
	f.Fuzz(func(t *testing.T, v int64, c uint16) {
		if c > 300 {
			t.Skip()
		}

		q := fpmoney.FromIntScaled(v, iso4217.Currency(c))

		s, err := json.Marshal(q)
		if err != nil {
			t.Error(err)
		}

		var x fpmoney.Amount
		if err := json.Unmarshal(s, &x); err != nil {
			t.Error(err, s)
		}

		if x != q {
			t.Error(x, q, v, c, s)
		}
	})
}

func BenchmarkArithmetic(b *testing.B) {
	x := fpmoney.FromFloat(12312.31, iso4217.SGD)
	y := fpmoney.FromFloat(12.02, iso4217.SGD)

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

	if x == fpmoney.FromInt(0, iso4217.SGD) {
		b.Error()
	}
}

var testsFloats = []struct {
	name string
	vals []string
}{
	{
		name: "small",
		vals: []string{
			`{"currency": "KRW", "amount": 123.456}`,
			`{"currency": "KRW", "amount": 0.123}`,
			`{"currency": "KRW", "amount": 0.012}`,
			`{"currency": "KRW", "amount": 0.001}`,
			`{"currency": "KRW", "amount": 0.982}`,
			`{"currency": "KRW", "amount": 0.101}`,
			`{"currency": "KRW", "amount": 10}`,
			`{"currency": "KRW", "amount": 11}`,
			`{"currency": "KRW", "amount": 1}`,
		},
	},
	{
		name: "large",
		vals: []string{
			`{"currency": "KRW", "amount": 123123123112312.1232}`,
			`{"currency": "KRW", "amount": 5341320482340234.123}`,
		},
	},
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	var s fpmoney.Amount
	for _, tc := range testsFloats {
		b.Run(tc.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				err := json.Unmarshal([]byte(tc.vals[n%len(tc.vals)]), &s)
				if err != nil || s == fpmoney.FromInt(0, iso4217.SGD) {
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

func BenchmarkJSONUnmarshal_float32(b *testing.B) {
	type T struct {
		Amount   float32 `json:"amount"`
		Currency string  `json:"currency"`
	}
	var s T
	for _, tc := range testsFloats {
		b.Run(tc.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				err := json.Unmarshal([]byte(tc.vals[n%len(tc.vals)]), &s)
				if err != nil || (s == T{}) {
					b.Error(s, err)
				}
			}
		})
	}
}

func BenchmarkJSONMarshal_float32(b *testing.B) {
	type T struct {
		Amount   float32 `json:"amount"`
		Currency string  `json:"currency"`
	}
	var s []byte
	var err error
	for _, tc := range testsFloats {
		tests := make([]T, 0, len(tc.vals))
		for _, q := range tc.vals {
			var x T
			if err := json.Unmarshal([]byte(q), &x); err != nil {
				b.Error(err)
			}
			tests = append(tests, x)
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

func TestMemoryLayout(t *testing.T) {
	a := fpmoney.FromFloat(-1000.123, iso4217.SGD)
	if v := unsafe.Sizeof(a); v != 16 {
		t.Error(a, v)
	}
}
