package main

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}

	CountDown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("want: %q, got: %q", want, got)
	}
}
