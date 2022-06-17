## 🧧 Fixed-Point Decimal Money

[![Go Report Card](https://goreportcard.com/badge/github.com/nikolaydubina/fpmoney)](https://goreportcard.com/report/github.com/nikolaydubina/fpmoney)
[![Go Reference](https://pkg.go.dev/badge/github.com/nikolaydubina/fpmoney.svg)](https://pkg.go.dev/github.com/nikolaydubina/fpmoney)

* `ISO 4217`
* as fast as `int64`
* no `float` in parsing nor printing
* zero overhead arithmetics
* block mismatched currency arithmetics
* does not leak precision
* Fuzz tests
* parsing is faster than `int`, `float`, `string`
* 200 LOC
* `int64` can fit enough of `BTC` _satoshi_ and even few `ETH` _wei_

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

### Cross Currency Protection

Akin to integer division by 0, which panics in Go, arithmetic operations on differnet currenices result in panic.
Returning error in arithmetic operation would prohibit chaning of method calls, which is not convenient.
It is better to stop execution, rather then corrupt value.
Mismatched or missing currencies must be caught at testing or QA of your code.

Two mechanisms to reduce panics are planned for future versions:
1. package level var for enable/disable currency check
2. package level var for fallback currency

### Ultra Small Fractions

Some denominatinos have very low fractions. 
Storing them `int64` you would get. 

- `BTC` _satoshi_ is `1 BTC = 100,000,000 satoshi`, which is still enough for ~`92,233,720,368 BTC`.
- `ETH` _wei_ is `1 ETH = 1,000,000,000,000,000,000 wei`, which is ~`9 ETH`. If you deal with _wei_, you may consider `bigint` or multiple `int64`. In fact, official Ethereum code is in Go and it is using bigint ([code](https://github.com/ethereum/go-ethereum/blob/master/params/denomination.go)).

Given that currency enumn still takes at least 1B in separate storage from `int64` in struct and Go allocates 16B of memory for struct regardless, current implementation reserved padding bytes.
It is sensible to use extra space our ot 16B to support long integer arithmetics.
Implementing this is area of furthter research.

### Benchmarks

```
$ go test -bench=. -benchmem ./...
goos: darwin
goarch: arm64
pkg: github.com/nikolaydubina/fpmoney
BenchmarkArithmetic/add_x1-10                     1000000000             0.5492 ns/op        0 B/op           0 allocs/op
BenchmarkArithmetic/add_x100-10                     18430124            64.64 ns/op          0 B/op           0 allocs/op
BenchmarkJSONUnmarshal/small-10                      3531835           340.7 ns/op         198 B/op           3 allocs/op
BenchmarkJSONUnmarshal/large-10                      2791712           426.9 ns/op         216 B/op           3 allocs/op
BenchmarkJSONUnmarshal_int/small-10                  2504600           478.5 ns/op         269 B/op           6 allocs/op
BenchmarkJSONUnmarshal_int/large-10                  2294034           522.2 ns/op         288 B/op           7 allocs/op
BenchmarkJSONUnmarshal_float32/small-10              2405636           496.9 ns/op         271 B/op           6 allocs/op
BenchmarkJSONUnmarshal_float32/large-10              2122207           567.6 ns/op         312 B/op           7 allocs/op
BenchmarkJSONMarshal/small-10                        4379685           274.4 ns/op         144 B/op           4 allocs/op
BenchmarkJSONMarshal/large-10                        3321205           345.8 ns/op         192 B/op           5 allocs/op
BenchmarkJSONMarshal_int/small-10                    8629840           138.6 ns/op          57 B/op           2 allocs/op
BenchmarkJSONMarshal_int/large-10                    8318066           143.2 ns/op          72 B/op           2 allocs/op
BenchmarkJSONMarshal_float32/small-10                6289126           189.9 ns/op          66 B/op           2 allocs/op
BenchmarkJSONMarshal_float32/large-10                6819679           175.7 ns/op          72 B/op           2 allocs/op
PASS
ok      github.com/nikolaydubina/fpmoney    62.744s
```

## Appendix A: json.Unmarshal optimizations

Parsing is surprisingly slow. It is ~6x of `float32` + `string`.

Use `json.NewDecoder` and parse directly.
```
BenchmarkJSONUnmarshal/small-10           2030568          2977 ns/op        1599 B/op          38 allocs/op
BenchmarkJSONUnmarshal/large-10           1956444          3106 ns/op        1640 B/op          39 allocs/op

```

Make container struct and wrap int and iso4217 currency and copy values.
```
BenchmarkJSONUnmarshal/small-10           2776969          2160 ns/op         430 B/op           8 allocs/op
BenchmarkJSONUnmarshal/large-10           2649692          2263 ns/op         448 B/op           8 allocs/op
```

Two passes over string, find `amount` and find `currency`.
```
BenchmarkJSONUnmarshal/small-10            686832          1732 ns/op         198 B/op           3 allocs/op
BenchmarkJSONUnmarshal/large-10            657272          1820 ns/op         216 B/op           3 allocs/op
```

Parsing just amount takes 400ns.
```
BenchmarkJSONUnmarshal/small-10              3339529           344.5 ns/op         198 B/op           3 allocs/op
BenchmarkJSONUnmarshal/large-10              2686135           443.2 ns/op         216 B/op           3 allocs/op
```

Package `github.com/ferdypruis/iso4217@v1.2.0` does cast of string to currency through loop.
But we have predefined currencies, we can rely on compiler for that.
Optimizing this cast by avoiding mallocs and loops.

As of `2022-06-17`, package `github.com/ferdypruis/iso4217@v1.2.1` uses map to cast currency.
It is as efficient as switch case.
Thanks @ferdypruis for the update!
