<p align="center">
  <img width="300" height="300" src="https://github.com/nikolaydubina/fpmoney/assets/2933061/022c83e3-8a14-4e8f-b1b7-94ab262fa590">
</p>

## 🧧 Fixed-Point Decimal Money

[![codecov](https://codecov.io/gh/nikolaydubina/fpmoney/branch/master/graph/badge.svg?token=Eh52jhLERp)](https://codecov.io/gh/nikolaydubina/fpmoney)
[![Go Report Card](https://goreportcard.com/badge/github.com/nikolaydubina/fpmoney)](https://goreportcard.com/report/github.com/nikolaydubina/fpmoney)
[![Go Reference](https://pkg.go.dev/badge/github.com/nikolaydubina/fpmoney.svg)](https://pkg.go.dev/github.com/nikolaydubina/fpmoney)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/nikolaydubina/fpmoney/badge)](https://securityscorecards.dev/viewer/?uri=github.com/nikolaydubina/fpmoney)

> _**Be Precise:** using floats to represent currency is almost criminal. — Robert.C.Martin, "Clean Code" p.301_

* as fast as `int64`
* no `float` in parsing nor printing, does not leak precision
* `ISO 4217`[^1][^2] currency
* block mismatched currency arithmetics
* 100 LOC
* fuzz tests

```go
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
```

### Ultra Small Fractions

Some denominations have very low fractions.
Storing them `int64` you would get.

- `BTC` _satoshi_ is `1 BTC = 100,000,000 satoshi`, which is still enough for ~`92,233,720,368 BTC`.
- `ETH` _wei_ is `1 ETH = 1,000,000,000,000,000,000 wei`, which is ~`9 ETH`. If you deal with _wei_, you may consider `bigint` or multiple `int64`. In fact, official Ethereum code is in Go and it is using bigint ([code](https://github.com/ethereum/go-ethereum/blob/master/params/denomination.go)).

### Benchmarks

```bash
$ go test -bench=. -benchmem .
goos: darwin
goarch: arm64
pkg: github.com/nikolaydubina/fpmoney
cpu: Apple M3 Max
BenchmarkCurrency_UnmarshalText-16      711695404                1.610 ns/op           0 B/op          0 allocs/op
BenchmarkCurrency_AppendText-16         446232057                2.698 ns/op           0 B/op          0 allocs/op
BenchmarkCurrency_MarshalText-16        81956246                13.99 ns/op            8 B/op          1 allocs/op
BenchmarkCurrency_String-16             1000000000               1.064 ns/op           0 B/op          0 allocs/op
BenchmarkArithmetic/add-16              924924993                1.305 ns/op           0 B/op          0 allocs/op
BenchmarkJSON/small/encode-16            6004620               198.5 ns/op           160 B/op          3 allocs/op
BenchmarkJSON/small/decode-16            5047149               238.7 ns/op           152 B/op          2 allocs/op
BenchmarkJSON/large/encode-16            4739722               255.7 ns/op           176 B/op          3 allocs/op
BenchmarkJSON/large/decode-16            3737406               315.3 ns/op           152 B/op          2 allocs/op
BenchmarkBinary/small/encode-16         132380481                9.044 ns/op          16 B/op          1 allocs/op
BenchmarkBinary/small/decode-16         100000000               10.80 ns/op           16 B/op          1 allocs/op
BenchmarkBinary/large/encode-16         133549021                8.995 ns/op          16 B/op          1 allocs/op
BenchmarkBinary/large/decode-16         100000000               10.61 ns/op           16 B/op          1 allocs/op
PASS
ok      github.com/nikolaydubina/fpmoney        15.804s
```

## References and Related Work

- [ferdypruis/iso4217](https://github.com/ferdypruis/iso4217) was a good inspiration and reference material. it was used in early version as well. it is well maintained and fast library for currencies. 
- `github.com/shopspring/decimal`: fixed precision; faster printing/parsing/arithmetics; currency handling 
- `github.com/Rhymond/go-money`: does not use `float` or `interface{}` in parsing; currency is enum
- `github.com/ferdypruis/iso4217`: skipped deprecated currencies to fit into `uint8` and smaller struct size
- https://en.wikipedia.org/wiki/ISO_4217

[^1]: excluding currencies with 4+ minor units `CLF`, `UYW`
[^2]: excluding deprecated currencies `HRD`, `HRK`, `SLL`, `ZWL`
