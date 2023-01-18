package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

var serverUrls = []string{
	"http://localhost:8080",
	"http://localhost:8081",
	"http://localhost:8082",
}

type Backend struct {
	URL          *url.URL
	Alive        bool
	ReverseProxy *httputil.ReverseProxy
}

type BackendPool struct {
	backends []*Backend
	current  uint64
}

var bp = BackendPool{
	[]*Backend{},
	0,
}

func balanceLoad(w http.ResponseWriter, r *http.Request) {
	bp.backends[bp.current].ReverseProxy.ServeHTTP(w, r)
	bp.current = (bp.current + 1) % uint64(len(bp.backends))
}

func main() {
	for _, serverUrl := range serverUrls {
		u, _ := url.Parse(serverUrl)
		rp := httputil.NewSingleHostReverseProxy(u)
		backend := &Backend{URL: u, ReverseProxy: rp}
		bp.backends = append(bp.backends, backend)
	}

	server := http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(balanceLoad),
	}

	server.ListenAndServe()
}
