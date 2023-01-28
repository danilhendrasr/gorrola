package gorrola

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

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

func (b *BackendPool) CheckHealth() {
	for _, be := range b.backends {
		conn, err := net.DialTimeout("tcp", be.URL.Host, 1*time.Second)

		if err != nil {
			be.SetAlive(false)
		} else {
			be.SetAlive(true)
			conn.Close()
		}
	}
}

func (b *BackendPool) MarkAsDown(backendUrl *url.URL) {
	for _, backend := range b.backends {
		if backend.URL == backendUrl {
			backend.mux.Lock()
			backend.Alive = false
			backend.mux.Unlock()
		}
	}
}

func (b *BackendPool) GetNextAliveBackend() (*Backend, error) {
	nextIdx := int((backendPool.current + 1) % uint64(len(backendPool.backends)))
	i := nextIdx

	// Bail out of the loop if alive backend is not found
	// after 3 iterations.
	for i < (nextIdx+1)*3 {
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

func healthCheck() {
	for {
		time.Sleep(time.Second * 20)
		backendPool.CheckHealth()
	}
}

func Run(backendUrls []string, port uint) {
	for _, backendUrl := range backendUrls {
		u, _ := url.Parse(backendUrl)

		rp := httputil.NewSingleHostReverseProxy(u)

		rp.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			fmt.Printf("Error: %s\n", err.Error())
			retries := GetRetryFromContext(r)
			if retries < 3 {
				time.Sleep(10 * time.Millisecond)
				ctx := context.WithValue(r.Context(), Retry, retries+1)
				rp.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			backendPool.MarkAsDown(u)

			attempts := GetAttemptsFromContext(r)
			ctx := context.WithValue(r.Context(), Attempts, attempts+1)
			balanceLoad(w, r.WithContext(ctx))
		}

		newNode := &Backend{URL: u, ReverseProxy: rp, Alive: true}
		backendPool.backends = append(backendPool.backends, newNode)
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(balanceLoad),
	}

	go healthCheck()
	log.Println("Load balancer served at: ", server.Addr)
	log.Fatal(server.ListenAndServe())
}
