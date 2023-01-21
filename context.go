package main

import "net/http"

const (
	Attempts int = iota
	Retry
)

func GetRetryFromContext(r *http.Request) int {
	if retry, ok := r.Context().Value(Retry).(int); ok {
		return retry
	}

	return 0
}

func GetAttemptsFromContext(r *http.Request) int {
	if retry, ok := r.Context().Value(Attempts).(int); ok {
		return retry
	}

	return 0
}
