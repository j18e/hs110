package hs110

import (
	"reflect"
	"testing"
)

func TestEncrypt(t *testing.T) {
	input := []byte(`{"key":"value"}`)
	exp := []byte("\x00\x00\x00\x0f\xd0\xf2\x99\xfc\x85\xa7\x9d\xbfɨı\xd4\xf6\x8b")

	got := encrypt(input)
	if !reflect.DeepEqual(got, exp) {
		t.Errorf("expected %q, got %q", exp, got)
	}
}

func TestDecrypt(t *testing.T) {
	input := []byte("\x00\x00\x00\x0f\xd0\xf2\x99\xfc\x85\xa7\x9d\xbfɨı\xd4\xf6\x8b")
	exp := []byte(`{"key":"value"}`)

	got := decrypt(input)
	if !reflect.DeepEqual(got, exp) {
		t.Errorf("expected %q, got %q", exp, got)
	}
}
