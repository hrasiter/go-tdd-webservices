package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	}

	t.Run("Hello, with name", func(t *testing.T) {
		got := Hello("Chriss")
		want := "Hello, Chriss"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello without name", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})

}
