package float32_bench_test

import (
	_ "embed"
	"encoding/json"
	"strings"
	"testing"
)

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

func BenchmarkJSONMarshal(b *testing.B) {
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
