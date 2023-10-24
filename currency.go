package fpmoney

import (
	_ "embed"
	"errors"
	"strings"
)

// Currency is ISO 4217 without deprecated currencies.
// Zero value is undefined currency.
type Currency struct{ v uint8 }

func (c Currency) String() string { return currencies[c.v+1] }

func (c Currency) MarshalText() (text []byte, err error) { return []byte(c.String()), nil }

func (c *Currency) UnmarshalText(text []byte) error {
	v, ok := fromAlpha[string(text)]
	if !ok {
		return errors.New("unexpected alpha: " + string(text))
	}
	*c = v
	return nil
}

//go:embed currency.csv
var currencyCSV string

var (
	currencies []string
	fromAlpha  = make(map[string]Currency)
)

func init() {
	currencies = append(currencies, "")
	for i, s := range strings.Split(currencyCSV, "\n") {
		fromAlpha[s] = Currency{v: uint8(i)}
		currencies = append(currencies, s)
	}
}
