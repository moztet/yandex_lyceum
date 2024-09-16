package main

import (
	"errors"
	"testing"
)

func TestU(t *testing.T) {
	valid := []byte("Hello")
	invalid := []byte{0xff}
	val, er := GetUTFLength(valid)
	if val != 5 || er != nil {
		t.Fatalf(`GetUTFLength("%v") = %v %v, wait %v %v`, valid, val, er, 5, nil)
	}
	val, er = GetUTFLength(invalid)
	if val != 0 || er != errors.New("invalid utf8") {
		t.Fatalf(`GetUTFLength("%v") = %v %v, wait %v %v`, valid, val, er, 0, errors.New("invalid utf8"))
	}
}
