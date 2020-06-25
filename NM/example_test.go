package main

import (
	"testing"
)

func TestHex_String(t *testing.T) {
	expect := "a"
	actual := Hex(10).String()

	if actual != expect {
		t.Errorf(`expect=%s actual=%s`, expect, actual)
	}
}