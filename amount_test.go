package fpmoney_test

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/nikolaydubina/fpmoney"
)

func init() {
	fpmoney.Currency = "SGD"
	fpmoney.CurrencyMinorUnits = 2
}

func ExampleAmount() {
	var BuySP500Price = fpmoney.FromInt(9000)

	input := []byte(`{"sp500": {"amount": 9000.02, "currency": "SGD"}}`)

	type Stonks struct {
		SP500 fpmoney.Amount `json:"sp500"`
	}
	var v Stonks
	if err := json.Unmarshal(input, &v); err != nil {
		log.Fatal(err)
	}

	amountToBuy := fpmoney.FromInt(0)
	if v.SP500.GreaterThan(BuySP500Price) {
		amountToBuy = amountToBuy.Add(v.SP500.Mul(2))
	}

	fmt.Println(amountToBuy)
	// Output: 18000.04 SGD
}

func ExampleAmount_DivMod_part() {
	x := fpmoney.FromInt(1)
	a, r := x.DivMod(3)
	fmt.Println(a, r)
	// Output: 0.33 SGD 0.01 SGD
}

func ExampleAmount_Div_whole() {
	x := fpmoney.FromInt(1)
	a, r := x.DivMod(5)
	fmt.Println(a, r)
	// Output: 0.2 SGD 0 SGD
}

func ExampleAmount_equality() {
	x := fpmoney.FromInt(3)
	y := fpmoney.FromInt(9)
	fmt.Println(y == x.Mul(3))
	// Output: true
}

func ExampleAmount_equality_same_currency() {
	x := fpmoney.FromInt(10)
	y := fpmoney.FromInt(10)
	fmt.Println(y == x)
	// Output: true
}

func ExampleAmount_equality_wrong_currency() {
	x := fpmoney.FromInt(10)
	y := fpmoney.FromInt(10)
	fmt.Println(y == x)
	// Output: false
}

func ExampleFromFloat() {
	x := fpmoney.FromFloat(144.96)
	fmt.Println(x)
	// Output: 144.96 SGD
}

func FuzzArithmetics(f *testing.F) {
	tests := [][2]int64{
		{1, 2},
		{1, -5},
		{1, 0},
		{1100, -2},
	}
	for _, tc := range tests {
		f.Add(tc[0], tc[1])
	}
	f.Fuzz(func(t *testing.T, a, b int64) {
		fa := fpmoney.FromIntScaled(a)
		fb := fpmoney.FromIntScaled(b)

		v := []bool{
			// sum commutativity
			fa.Add(fb) == fb.Add(fa),

			// sum associativity
			fpmoney.FromInt(0).Add(fa).Add(fb).Add(fa) == fpmoney.FromInt(0).Add(fb).Add(fa).Add(fa),

			// sum zero
			fa == fa.Add(fb).Sub(fb),
			fa == fa.Sub(fb).Add(fb),
			fpmoney.FromInt(0) == fpmoney.FromInt(0).Add(fa).Sub(fa),

			// product identity
			fa == fa.Mul(1),

			// product zero
			fpmoney.FromInt(0) == fa.Mul(0),

			// match number
			(a == b) == (fa == fb),
			a < b == fa.LessThan(fb),
			a > b == fa.GreaterThan(fb),
			a <= b == fa.LessThanOrEqual(fb),
			a >= b == fa.GreaterThanOrEqual(fb),

			// match number convert
			fpmoney.FromIntScaled(a+b) == fa.Add(fb),
			fpmoney.FromIntScaled(a-b) == fa.Sub(fb),
		}
		for i, q := range v {
			if !q {
				t.Error(i, a, b, fa, fb)
			}
		}

		if b != 0 {
			w, r := fa.DivMod(int(b))
			if w != fpmoney.FromIntScaled(a/b) {
				t.Error(w, a/b, a, b, fa)
			}
			if r != fpmoney.FromIntScaled(a%b) {
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
		f.Add(tc, uint8(5))
		f.Add(-tc, uint8(5))
	}
	f.Fuzz(func(t *testing.T, r float32, nf uint8) {
		if nf > 10 {
			t.Skip()
		}
		var l float32 = 10000000
		if r > l || r < -l {
			t.Skip()
		}

		fs := `%.` + strconv.Itoa(int(nf)) + `f`
		rs := fmt.Sprintf(fs, r)
		s := fmt.Sprintf(`{"amount": %s, "currency": "%s"}`, rs, "USD")
		if _, err := fmt.Sscanf(rs, "%f", &r); err != nil {
			t.Error(err)
		}

		if r == -0 {
			t.Skip()
		}

		var x fpmoney.Amount
		if err := json.Unmarshal([]byte(s), &x); err != nil {
			t.Error(rs, "USD", err)
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
		a := fpmoney.FromIntScaled(v)

		if float32(v) != a.Float32() {
			t.Error(a, a.Float32(), float32(v))
		}

		if float64(v) != a.Float64() {
			t.Error(a, a.Float64(), v)
		}
	})
}

func TestUnmarshalJSON(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		fpmoney.Currency = "SGD"
		fpmoney.CurrencyMinorUnits = 2

		tests := []struct {
			s string
			v fpmoney.Amount
		}{
			{
				s: `{"currency": "SGD","amount": 9002.01}`,
				v: fpmoney.FromIntScaled(900201),
			},
			{
				s: `{"amount": 9002.01, "currency": "SGD"}`,
				v: fpmoney.FromIntScaled(900201),
			},
			{
				s: `{"amount": -9002.01, "currency": "SGD"}`,
				v: fpmoney.FromIntScaled(-900201),
			},
			{
				s: `{"amount": 0, "currency": "SGD"}`,
				v: fpmoney.FromIntScaled(0),
			},
			{
				s: `{"amount": 0.01, "currency": "SGD"}`,
				v: fpmoney.FromIntScaled(1),
			},
			{
				s: `{"amount": -0.01, "currency": "SGD"}`,
				v: fpmoney.FromIntScaled(-1),
			},
			{
				s: `    {   "amount" : 9002.01
            
            , 
            
            "currency"
                : 
            
            "SGD"}   


             `,
				v: fpmoney.FromIntScaled(900201),
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
	})

	t.Run("zero cents", func(t *testing.T) {
		fpmoney.Currency = "KRW"
		fpmoney.CurrencyMinorUnits = 3

		tests := []struct {
			s string
			v fpmoney.Amount
		}{
			{
				s: `{"amount":1,"currency":"KRW"}`,
				v: fpmoney.FromIntScaled(1),
			},
			{
				s: `{"amount":-1,"currency":"KRW"}`,
				v: fpmoney.FromIntScaled(-1),
			},
			{
				s: `{"amount":123,"currency":"KRW"}`,
				v: fpmoney.FromIntScaled(123),
			},
			{
				s: `{"amount":-123,"currency":"KRW"}`,
				v: fpmoney.FromIntScaled(-123),
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
	})
}

func TestMarshalJSON(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		fpmoney.Currency = "SGD"
		fpmoney.CurrencyMinorUnits = 2

		tests := []struct {
			s string
			v fpmoney.Amount
		}{
			{
				s: `{"amount":9002.01,"currency":"SGD"}`,
				v: fpmoney.FromIntScaled(900201),
			},
			{
				s: `{"amount":0,"currency":"SGD"}`,
				v: fpmoney.FromIntScaled(0),
			},
			{
				s: `{"amount":0.01,"currency":"SGD"}`,
				v: fpmoney.FromIntScaled(1),
			},
			{
				s: `{"amount":-0.01,"currency":"SGD"}`,
				v: fpmoney.FromIntScaled(-1),
			},
			{
				s: `{"amount":1.01,"currency":"SGD"}`,
				v: fpmoney.FromIntScaled(101),
			},
			{
				s: `{"amount":-1.01,"currency":"SGD"}`,
				v: fpmoney.FromIntScaled(-101),
			},
			{
				s: `{"amount":1,"currency":"SGD"}`,
				v: fpmoney.FromInt(1),
			},
			{
				s: `{"amount":5,"currency":"SGD"}`,
				v: fpmoney.FromInt(5),
			},
			{
				s: `{"amount":-1,"currency":"SGD"}`,
				v: fpmoney.FromInt(-1),
			},
			{
				s: `{"amount":-5,"currency":"SGD"}`,
				v: fpmoney.FromInt(-5),
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
	})

	t.Run("zero cents", func(t *testing.T) {
		fpmoney.Currency = "KRW"
		fpmoney.CurrencyMinorUnits = 3

		tests := []struct {
			s string
			v fpmoney.Amount
		}{
			{
				s: `{"amount":1,"currency":"KRW"}`,
				v: fpmoney.FromIntScaled(1),
			},
			{
				s: `{"amount":-1,"currency":"KRW"}`,
				v: fpmoney.FromIntScaled(-1),
			},
			{
				s: `{"amount":123,"currency":"KRW"}`,
				v: fpmoney.FromIntScaled(123),
			},
			{
				s: `{"amount":-123,"currency":"KRW"}`,
				v: fpmoney.FromIntScaled(-123),
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

	})
}

func FuzzJSON_MarshalUnmarshal(f *testing.F) {
	currencies := [...]struct {
		currency   string
		minorUnits int
	}{
		{currency: "KRW", minorUnits: 3},
		{currency: "SGD", minorUnits: 2},
		{currency: "BHD", minorUnits: 2},
		{currency: "CLF", minorUnits: 2},
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

		q := fpmoney.FromIntScaled(v)

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
	x := fpmoney.FromFloat(12312.31)
	y := fpmoney.FromFloat(12.02)

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

	if x == fpmoney.FromInt(0) {
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
				if err != nil || s == fpmoney.FromInt(0) {
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
