package main

import (
	"errors"
	"net/http"
	"time"
)

const defaultTimeout = 10 * time.Second

func Race(a, b string) (string, error) {
	return ConfigurableRace(a, b, defaultTimeout)
}

func ConfigurableRace(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", errors.New("timeout waiting for 10 seconds")
	}
}

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
