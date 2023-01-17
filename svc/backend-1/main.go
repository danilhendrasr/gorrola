package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Backend 1 got it!")
	})

	http.ListenAndServe(":8080", nil)
}
