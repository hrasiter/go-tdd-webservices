package main

import (
	"context"
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

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.data
}

func TestStore(t *testing.T) {
	t.Run("storing data", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{data: data}
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

		store := &SpyStore{data: data}
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
