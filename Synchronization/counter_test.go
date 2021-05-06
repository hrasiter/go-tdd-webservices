package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment counter 3 times", func(t *testing.T) {
		counter := NewCounter()
		counter.Increment()
		counter.Increment()
		counter.Increment()
		assertCounter(t, 3, counter)
	})

	t.Run("counts safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Increment()
				w.Done()
			}(&wg)
		}

		wg.Wait()

		assertCounter(t, wantedCount, counter)
	})
}

func assertCounter(t testing.TB, want int, got *Counter) {
	if got.Value() != want {
		t.Errorf("got: %d, want: %d", got.Value(), want)
	}
}
