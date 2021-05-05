package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{}
	CountDown(buffer, sleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("want: %q, got: %q", want, got)
	}

	if sleeper.Calls != 4 {
		t.Errorf("sleeper should be called 4 times! but called: %d", sleeper.Calls)
	}
}
