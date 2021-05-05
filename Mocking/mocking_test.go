package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type CountDownOperationsSpy struct {
	Calls []string
}

func (c *CountDownOperationsSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountDownOperationsSpy) Write(p []byte) (n int, er error) {
	c.Calls = append(c.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"

func TestCountDown(t *testing.T) {

	t.Run("counts Sleep calls", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		CountDown(buffer, &CountDownOperationsSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("want: %q, got: %q", want, got)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepCounter := &CountDownOperationsSpy{}
		CountDown(spySleepCounter, spySleepCounter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepCounter.Calls) {
			t.Errorf("wanted calls: %v, got %v", want, spySleepCounter.Calls)
		}
	})

}
