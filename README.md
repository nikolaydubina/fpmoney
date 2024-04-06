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
* no `float` in parsing nor printing
* `ISO 4217` currency
* block mismatched currency arithmetics
* does not leak precision
* parsing faster than `int`, `float`, `string`
* 100 LOC

```go
var BuySP500Price = fpmoney.Amount{Amount: fpdecimal.FromInt(9000), Currency: fpmoney.SGD}

input := []byte(`{"sp500": {"amount": 9000.02, "currency": "SGD"}}`)

type Stonks struct {
    SP500 fpmoney.Amount `json:"sp500"`
}
var v Stonks
if err := json.Unmarshal(input, &v); err != nil {
    log.Fatal(err)
}

amountToBuy := fpmoney.Amount{Amount: fpdecimal.Zero, Currency: fpmoney.SGD}
if v.SP500.GreaterThan(BuySP500Price) {
    amountToBuy = amountToBuy.Add(v.SP500.Mul(2))
}

json.NewEncoder(os.Stdout).Encode(amountToBuy)
// Output: {"amount":18000.04,"currency":"SGD"}
```

### Division

Division always returns remainder.
Fractional cents can never be reached.

```go
x := fpmoney.Amount{Amount: fpdecimal.FromInt(1), Currency: fpmoney.SGD}
a, r := x.Div(3)
enc := json.NewEncoder(os.Stdout)
enc.Encode(a)
enc.Encode(r)
// Output:
// {"amount":0.3333,"currency":"SGD"}
// {"amount":0.0001,"currency":"SGD"}
```

### Ultra Small Fractions

Some denominations have very low fractions.
Storing them `int64` you would get.

- `BTC` _satoshi_ is `1 BTC = 100,000,000 satoshi`, which is still enough for ~`92,233,720,368 BTC`.
- `ETH` _wei_ is `1 ETH = 1,000,000,000,000,000,000 wei`, which is ~`9 ETH`. If you deal with _wei_, you may consider `bigint` or multiple `int64`. In fact, official Ethereum code is in Go and it is using bigint ([code](https://github.com/ethereum/go-ethereum/blob/master/params/denomination.go)).

Given that currency enumn still takes at least 1B in separate storage from `int64` in struct and Go allocates 16B of memory for struct regardless, current implementation reserved padding bytes.
It is sensible to use extra space our ot 16B to support long integer arithmetics.
Implementing this is area of furthter research.

### Benchmarks

```bash
$ go test -bench=. -benchmem . > fpmoney.bench
$ go test -bench=. -benchmem ./internal/bench/float32 > float32.bench
$ go test -bench=. -benchmem ./internal/bench/int > int.bench
$ benchstat -split="XYZ" int.bench float32.bench fpmoney.bench
name \ time/op              int.bench   float32.bench  fpmoney.bench
JSONUnmarshal/small-16      382ns ± 0%     429ns ± 0%     419ns ± 0%
JSONUnmarshal/large-16      429ns ± 0%     503ns ± 0%     464ns ± 0%
JSONMarshal/small-16        112ns ± 0%     158ns ± 0%     187ns ± 0%
JSONMarshal/large-16        112ns ± 0%     144ns ± 0%     231ns ± 0%

name \ alloc/op             int.bench   float32.bench  fpmoney.bench
JSONUnmarshal/small-16       268B ± 0%      270B ± 0%      262B ± 0%
JSONUnmarshal/large-16       272B ± 0%      288B ± 0%      280B ± 0%
JSONMarshal/small-16        57.0B ± 0%     66.0B ± 0%     72.0B ± 0%
JSONMarshal/large-16        72.0B ± 0%     72.0B ± 0%     88.0B ± 0%

name \ allocs/op            int.bench   float32.bench  fpmoney.bench
JSONUnmarshal/small-16       6.00 ± 0%      6.00 ± 0%      5.00 ± 0%
JSONUnmarshal/large-16       6.00 ± 0%      6.00 ± 0%      5.00 ± 0%
JSONMarshal/small-16         2.00 ± 0%      2.00 ± 0%      3.00 ± 0%
JSONMarshal/large-16         2.00 ± 0%      2.00 ± 0%      3.00 ± 0%
```

## References and Related Work

- [ferdypruis/iso4217](https://github.com/ferdypruis/iso4217) was a good inspiration and reference material. it was used in early version as well. it is well maintained and fast library for currencies. 
- `github.com/shopspring/decimal`: fixed precision; faster printing/parsing/arithmetics; currency handling 
- `github.com/Rhymond/go-money`: does not use `float` or `interface{}` in parsing; currency is enum
- `github.com/ferdypruis/iso4217`: skipped deprecated currencies to fit into `uint8` and smaller struct size
