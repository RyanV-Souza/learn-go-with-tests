package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	ErrTimeout       = errors.New("request timed out")
	tenSecondTimeout = 10 * time.Second
)

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		fmt.Println("a")
		return a, nil
	case <-ping(b):
		fmt.Println("b")
		return b, nil
	case <-time.After(timeout):
		return "", ErrTimeout
	}
}

func ping(URL string) chan bool {
	ch := make(chan bool)

	go func() {
		http.Get(URL)
		ch <- true
	}()

	return ch
}
