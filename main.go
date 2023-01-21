package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

var serverUrls = []string{
	"http://localhost:8080",
	"http://localhost:8081",
	"http://localhost:8082",
}

type Backend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

func (b *Backend) IsAlive() bool {
	var alive bool
	b.mux.Lock()
	alive = b.Alive
	b.mux.Unlock()
	return alive
}

type BackendPool struct {
	backends []*Backend
	current  uint64
}

func (b *BackendPool) MarkAsDown(serverUrl string) {
	for _, backend := range b.backends {
		if backend.URL.String() == serverUrl {
			backend.mux.Lock()
			backend.Alive = false
			backend.mux.Unlock()
		}
	}
}

func (b *BackendPool) GetNextAliveBackend() (*Backend, error) {
	nextIdx := int((backendPool.current + 1) % uint64(len(backendPool.backends)))
	i := nextIdx

	for i < nextIdx*3 {
		idx := i % len(b.backends)

		if b.backends[idx].IsAlive() {
			b.current = uint64(idx)
			return b.backends[idx], nil
		}

		i++
	}

	return nil, errors.New("cannot find an alive backend")
}

var backendPool = BackendPool{
	[]*Backend{},
	0,
}

func balanceLoad(w http.ResponseWriter, r *http.Request) {
	nextBackend, err := backendPool.GetNextAliveBackend()
	if err != nil {
		http.Error(w, "Service is not available", http.StatusServiceUnavailable)
		return
	}

	nextBackend.ReverseProxy.ServeHTTP(w, r)
}

func main() {
	for idx, serverUrl := range serverUrls {
		u, _ := url.Parse(serverUrl)

		rp := httputil.NewSingleHostReverseProxy(u)

		rp.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			fmt.Printf("Error: %s\n", err.Error())
			retries := GetRetryFromContext(r)
			if retries < 3 {
				select {
				case <-time.After(10 * time.Millisecond):
					ctx := context.WithValue(r.Context(), Retry, retries+1)
					rp.ServeHTTP(w, r.WithContext(ctx))
				}
				return
			}

			backendPool.MarkAsDown(serverUrl)

			attempts := GetAttemptsFromContext(r)
			ctx := context.WithValue(r.Context(), Attempts, attempts+1)
			balanceLoad(w, r.WithContext(ctx))
		}

		newNode := &Backend{URL: u, ReverseProxy: rp, Alive: true}
		if idx == 2 {
			newNode.Alive = false
		}
		backendPool.backends = append(backendPool.backends, newNode)
	}

	server := http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(balanceLoad),
	}

	server.ListenAndServe()
}
