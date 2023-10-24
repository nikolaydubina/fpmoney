<p align="center">
  <img width="300" height="300" src="https://github.com/nikolaydubina/fpmoney/assets/2933061/022c83e3-8a14-4e8f-b1b7-94ab262fa590">
</p>

## 🧧 Fixed-Point Decimal Money

[![codecov](https://codecov.io/gh/nikolaydubina/fpmoney/branch/master/graph/badge.svg?token=Eh52jhLERp)](https://codecov.io/gh/nikolaydubina/fpmoney)
[![Go Report Card](https://goreportcard.com/badge/github.com/nikolaydubina/fpmoney)](https://goreportcard.com/report/github.com/nikolaydubina/fpmoney)
[![Go Reference](https://pkg.go.dev/badge/github.com/nikolaydubina/fpmoney.svg)](https://pkg.go.dev/github.com/nikolaydubina/fpmoney)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

> _**Be Precise:** using floats to represent currency is almost criminal. — Robert.C.Martin, "Clean Code" p.301_

* as fast as `int64`
* no `float` in parsing nor printing
* `ISO 4217` currency
* block mismatched currency arithmetics
* does not leak precision
* parsing faster than `int`, `float`, `string`
* Fuzz tests, Benchmarks, Generics
* 200 LOC

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

### Division

Division always returns remainder.
Fractional cents can never be reached.

```go
x := fpmoney.FromInt(1, fpmoney.SGD)
a, r := x.Div(3)
fmt.Println(a, r)
// Output: 0.33 SGD 0.01 SGD

a, r = x.Div(5)
fmt.Println(a, r)
// Output: 0.20 SGD 0 SGD
```

### Equality

Equality operator can be used to compare values.

```go
x := fpmoney.FromInt(3, fpmoney.SGD)
y := fpmoney.FromInt(9, fpmoney.SGD)
fmt.Println(y == x.Mul(3))
// Output: true
```

### Cross Currency Protection

Akin to integer division by 0 which panics in Go, mismatched currenices result in panic.

Arithmetics
```go
x := fpmoney.FromInt(10, fpmoney.USD)
y := fpmoney.FromInt(10, fpmoney.SGD)
c := x.Add(y) // panics
```

Equality
```go
x := fpmoney.FromInt(10, fpmoney.USD)
y := fpmoney.FromInt(10, fpmoney.SGD)
fmt.Println(y == x)
// Output: false
```

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
BenchmarkArithmetic/add_x1-10                   1000000000             0.5 ns/op           0 B/op           0 allocs/op
BenchmarkArithmetic/add_x100-10                   12525424            51.9 ns/op           0 B/op           0 allocs/op
BenchmarkJSONUnmarshal/small-10                    3610992           329.8 ns/op         198 B/op           3 allocs/op
BenchmarkJSONUnmarshal/large-10                    2901363           412.4 ns/op         216 B/op           3 allocs/op
BenchmarkJSONMarshal/small-10                      5032456           238.1 ns/op         160 B/op           3 allocs/op
BenchmarkJSONMarshal/large-10                      4072776           295.5 ns/op         176 B/op           3 allocs/op
BenchmarkJSONMarshal_Exact/small-10               40404832            29.6 ns/op         112 B/op           1 allocs/op
BenchmarkJSONMarshal_Exact/large-10               28532677            41.6 ns/op         112 B/op           1 allocs/op
PASS
ok      github.com/nikolaydubina/fpmoney    62.744s
```

`float32` (old) and `fpmoney` (new)
```
$ benchstat -split="XYZ" float32.bench fpmoney.bench
name                    old time/op    new time/op    delta
JSONUnmarshal/small-10     502ns ± 0%     338ns ± 1%   -32.63%  (p=0.008 n=5+5)
JSONUnmarshal/large-10     572ns ± 0%     419ns ± 1%   -26.79%  (p=0.008 n=5+5)
JSONMarshal/small-10       189ns ± 0%     245ns ± 1%   +29.12%  (p=0.008 n=5+5)
JSONMarshal/large-10       176ns ± 0%     305ns ± 1%   +73.07%  (p=0.008 n=5+5)

name                    old alloc/op   new alloc/op   delta
JSONUnmarshal/small-10      271B ± 0%      198B ± 0%   -26.94%  (p=0.008 n=5+5)
JSONUnmarshal/large-10      312B ± 0%      216B ± 0%   -30.77%  (p=0.008 n=5+5)
JSONMarshal/small-10       66.0B ± 0%    160.0B ± 0%  +142.42%  (p=0.008 n=5+5)
JSONMarshal/large-10       72.0B ± 0%    176.0B ± 0%  +144.44%  (p=0.008 n=5+5)

name                    old allocs/op  new allocs/op  delta
JSONUnmarshal/small-10      6.00 ± 0%      3.00 ± 0%   -50.00%  (p=0.008 n=5+5)
JSONUnmarshal/large-10      7.00 ± 0%      3.00 ± 0%   -57.14%  (p=0.008 n=5+5)
JSONMarshal/small-10        2.00 ± 0%      3.00 ± 0%   +50.00%  (p=0.008 n=5+5)
JSONMarshal/large-10        2.00 ± 0%      3.00 ± 0%   +50.00%  (p=0.008 n=5+5)
```

`int`, `float32`, `fpmoney`
```
$ benchstat -split="XYZ" int.bench float32.bench fpmoney.bench
name \ time/op              int.bench   float32.bench  fpmoney.bench
JSONUnmarshal/small-10      481ns ± 2%     502ns ± 0%     338ns ± 1%
JSONUnmarshal/large-10      530ns ± 1%     572ns ± 0%     419ns ± 1%
JSONMarshal/small-10        140ns ± 1%     189ns ± 0%     245ns ± 1%
JSONMarshal/large-10        145ns ± 0%     176ns ± 0%     305ns ± 1%

name \ alloc/op             int.bench   float32.bench  fpmoney.bench
JSONUnmarshal/small-10       269B ± 0%      271B ± 0%      198B ± 0%
JSONUnmarshal/large-10       288B ± 0%      312B ± 0%      216B ± 0%
JSONMarshal/small-10        57.0B ± 0%     66.0B ± 0%    160.0B ± 0%
JSONMarshal/large-10        72.0B ± 0%     72.0B ± 0%    176.0B ± 0%

name \ allocs/op            int.bench   float32.bench  fpmoney.bench
JSONUnmarshal/small-10       6.00 ± 0%      6.00 ± 0%      3.00 ± 0%
JSONUnmarshal/large-10       7.00 ± 0%      7.00 ± 0%      3.00 ± 0%
JSONMarshal/small-10         2.00 ± 0%      2.00 ± 0%      3.00 ± 0%
JSONMarshal/large-10         2.00 ± 0%      2.00 ± 0%      3.00 ± 0%
```

## References

- [ferdypruis/iso4217](https://github.com/ferdypruis/iso4217) was a good inspiration and reference material. it was used in early version as well. it is well maintained and fast library for currencies. 

## Appendix A: `json.Unmarshal` optimizations

Parsing is surprisingly slow. It is ~6x of `float32` + `string`.

Use `json.NewDecoder` and parse directly.
```
BenchmarkJSONUnmarshal/small-10           2030568          2977 ns/op        1599 B/op          38 allocs/op
BenchmarkJSONUnmarshal/large-10           1956444          3106 ns/op        1640 B/op          39 allocs/op

```

Make container struct and wrap int and ISO 4217 currency and copy values.
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

## Appendix B: Other Libraries

`github.com/shopspring/decimal`
* fixed precision
* faster printing/parsing/arithmetics
* currency handling 

`github.com/Rhymond/go-money`
* does not use `float` or `interface{}` in parsing
* currency is enum

`github.com/ferdypruis/iso4217`
* skipped deprecated currencies to fit into `uint8` and smaller struct size

## Appendix C: Extra malloc in Printing

Even though `MarshalJSON` does exactly one malloc, using it with `json.Marshall` package adds two more mallocs.
This looks like penalty of reflect nature of `json` package and is unavoidable.

```
BenchmarkJSONMarshal_Exact/small-10     40404832    29.6 ns/op      112 B/op        1 allocs/op
BenchmarkJSONMarshal_Exact/large-10     28532677    41.6 ns/op      112 B/op        1 allocs/op
```

## Appendix D: Strict Currency Enum

It is possible to rely on Go compiler to strictiy currency enum by wrapping into a struct.
There is no performance penalty.
Implementation is almost same.
API is the same, but much safer.

## Appendix E: comparable generics for currency

Using `comparable` generic constraint is attractive option, since it allows to plug-in any type for currencies, including from other packages.
Marshalling and Unamrshalling will be fully delegated to other packages too.
However, this inccurs penalty for performance.
Arithmetics is 3x slower, and due to `json` pacakge `mallocs` increase 2x. 
Using fpdecimal.Decimal directly, also fixes precision for whole binary, meaning other pacakges in binary should be using same number of decimal values in fpdecimal package.
For large scale systems, it is recommended to not access fpdecimal.Decimal directly, but define wrappers on top of it.

```go
type Amount[T comparable] struct {
	Amount   fpdecimal.Decimal `json:"amount"`
	Currency T                 `json:"currency"`
}
```
