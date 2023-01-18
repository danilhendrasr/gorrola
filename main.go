package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

var serverUrls = []string{"http://localhost:8080", "http://localhost:8081", "http://localhost:8082"}

type BackendPool struct {
	backends []*httputil.ReverseProxy
	current  uint64
}

var pp = BackendPool{
	[]*httputil.ReverseProxy{},
	0,
}

func balanceLoad(w http.ResponseWriter, r *http.Request) {
	pp.backends[pp.current].ServeHTTP(w, r)
	pp.current = (pp.current + 1) % uint64(len(pp.backends))
}

func main() {
	for _, serverUrl := range serverUrls {
		u, _ := url.Parse(string(serverUrl))
		rp := httputil.NewSingleHostReverseProxy(u)
		pp.backends = append(pp.backends, rp)
	}

	server := http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(balanceLoad),
	}

	server.ListenAndServe()
}
