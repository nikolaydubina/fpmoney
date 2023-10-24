package fpmoney

import (
	"fmt"
	"testing"
)

func TestCurrencyBasic(t *testing.T) {
	c := SGD
	if c.String() != "SGD" {
		t.Error("wrong stringer")
	}
}

func TestCurrencyTextEncoding(t *testing.T) {
	tests := []struct {
		c Currency
		s []byte
	}{
		{
			SGD,
			[]byte("SGD"),
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%#v <-> %s", tc.c, tc.s), func(t *testing.T) {
			var v Currency
			if err := v.UnmarshalText(tc.s); err != nil {
				t.Error(err)
			}
			if v != tc.c {
				t.Error("wrong decode")
			}
			b, err := tc.c.MarshalText()
			if err != nil {
				t.Error(err)
			}
			if string(tc.s) != string(b) {
				t.Error(b, "!=", tc.s)
			}
		})
	}
}

func TestCurrencyTextEncoding_Error(t *testing.T) {
	tests := []struct {
		s []byte
	}{
		{s: []byte("asdf")},
		{s: []byte("sgd")},
		{s: []byte(" SGD")},
		{s: []byte("SGD ")},
		{s: []byte("")},
		{s: []byte("\n")},
		{s: nil},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			var v Currency
			err := v.UnmarshalText(tc.s)
			if err == nil {
				t.Error("expected error")
			}
			var c Currency
			if v != c {
				t.Error("wrong zero currency")
			}
		})
	}
}
