package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	u, _ := url.Parse("http://localhost:8080")

	rp := httputil.NewSingleHostReverseProxy(u)
	handler := http.HandlerFunc(rp.ServeHTTP)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
