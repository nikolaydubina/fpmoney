package int_bench_test

import (
	_ "embed"
	"encoding/json"
	"strings"
	"testing"
)

//go:embed testdata/amount-int-large.jsonl
var amountIntLargeJSONL string

//go:embed testdata/amount-int-small.jsonl
var amountIntSmallJSONL string

var testsInts = []struct {
	name string
	vals []string
}{
	{
		name: "small",
		vals: strings.Split(amountIntSmallJSONL, "\n"),
	},
	{
		name: "large",
		vals: strings.Split(amountIntLargeJSONL, "\n"),
	},
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	type T struct {
		Amount   int    `json:"amount"`
		Currency string `json:"currency"`
	}
	var s T
	for _, tc := range testsInts {
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
		Amount   int    `json:"amount"`
		Currency string `json:"currency"`
	}
	var s []byte
	var err error
	for _, tc := range testsInts {
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
