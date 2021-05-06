package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("get the faster url", func(t *testing.T) {
		fastServer := makeDelayedServer(20 * time.Millisecond)

		slowServer := makeDelayedServer(40 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		fasturl := fastServer.URL
		slowurl := slowServer.URL

		want := fasturl

		got, _ := Race(slowurl, fasturl)

		if got != want {
			t.Errorf("want: %q, got: %q", want, got)
		}
	})

	t.Run("timeout error after 10 seconds", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		_, err := Race(serverA.URL, serverB.URL, 1*time.Second)

		if err == nil {
			t.Error("should respond an error")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		rw.WriteHeader(http.StatusOK)
	}))
}
