## ✈️ Small Fixed-Point Decimal Money

[![Go Report Card](https://goreportcard.com/badge/github.com/nikolaydubina/fpmoney)](https://goreportcard.com/report/github.com/nikolaydubina/fpmoney)

* `ISO 4217`
* as fast as `int64`
* no `float` in parsing nor printing
* zero overhead arithmetics
* block mismatched currency arithmetics
* does not leak precision
* Fuzz tests

```go
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
```

### Division

Division always returns remainder.
Fractional cents can never be reached.

```go
x := fpmoney.FromInt(1, iso4217.SGD)
a, r := x.Div(3)
fmt.Println(a, r)
// Output: 0.33 SGD 0.01 SGD

a, r = x.Div(5)
fmt.Println(a, r)
// Output: 0.20 SGD 0 SGD
```

### Benchmarks

```
nikolaydubina@Nikolays-MacBook-Pro fpmoney % go test -bench=. -benchtime=5s -benchmem ./...
goos: darwin
goarch: arm64
pkg: github.com/nikolaydubina/fpmoney
BenchmarkArithmetic/add_x1-10        	1000000000	         0.5440 ns/op	       0 B/op	       0 allocs/op
BenchmarkArithmetic/add_x100-10      	100000000	        52.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkJSONUnmarshal/small-10      	 2030568	      2977 ns/op	    1599 B/op	      38 allocs/op
BenchmarkJSONUnmarshal/large-10      	 1956444	      3106 ns/op	    1640 B/op	      39 allocs/op
BenchmarkJSONMarshal/small-10        	20263528	       295.2 ns/op	     337 B/op	       4 allocs/op
BenchmarkJSONMarshal/large-10        	16540860	       366.8 ns/op	     384 B/op	       5 allocs/op
BenchmarkJSONUnmarshal_float32/small-10         	11995531	       500.0 ns/op	     271 B/op	       6 allocs/op
BenchmarkJSONUnmarshal_float32/large-10         	10550324	       570.9 ns/op	     312 B/op	       7 allocs/op
BenchmarkJSONMarshal_float32/small-10           	31520827	       192.1 ns/op	      66 B/op	       2 allocs/op
BenchmarkJSONMarshal_float32/large-10           	33842808	       177.2 ns/op	      72 B/op	       2 allocs/op
PASS
ok  	github.com/nikolaydubina/fpmoney	62.744s
```

## Appendix A: json.Unmarshal optimizations

Parsing is surprisingly slow. It is ~6x of `float32` + `string`.

Use `json.NewDecoder` and parse directly.
```
BenchmarkJSONUnmarshal/small-10      	 2030568	      2977 ns/op	    1599 B/op	      38 allocs/op
BenchmarkJSONUnmarshal/large-10      	 1956444	      3106 ns/op	    1640 B/op	      39 allocs/op

```

Make container struct and wrap int and iso4217 currency and copy values.
```
BenchmarkJSONUnmarshal/small-10      	 2776969	      2160 ns/op	     430 B/op	       8 allocs/op
BenchmarkJSONUnmarshal/large-10      	 2649692	      2263 ns/op	     448 B/op	       8 allocs/op
```
