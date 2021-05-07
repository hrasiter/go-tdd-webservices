package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	data      string
	cancelled bool
	t         *testing.T
}

func (s *SpyStore) assertWasCanceled() {
	s.t.Helper()

	if !s.cancelled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCanceled() {
	s.t.Helper()

	if s.cancelled {
		s.t.Error("it should not have cancelled the store")
	}
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	log.Println("Fetching...")
	go func() {
		var result string

		for _, c := range s.data {
			select {
			case <-ctx.Done():
				log.Println("spyStore got canceled")
				log.Println("CANCELED")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}

		data <- result
	}()

	select {
	case <-ctx.Done():
		log.Println("context canceled!!!")
		return "", ctx.Err()
	case res := <-data:
		log.Println("data received...")
		return res, nil
	}
}

func TestStore(t *testing.T) {
	t.Run("storing data", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{data: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		store.assertWasNotCanceled()
	})
	t.Run("canceling request", func(t *testing.T) {
		data := "hello, world"

		store := &SpyStore{data: data, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		store.assertWasCanceled()
	})
}
