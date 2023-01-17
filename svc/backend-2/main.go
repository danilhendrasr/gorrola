package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Backend 2 got it!")
	})

	http.ListenAndServe(":8081", nil)
}
