package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/danilhendrasr/gorrola/pkg/test_util"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		test_util.RunGorrola()
	}()

	resp, _ := http.Get("http://localhost:3000")
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s\n", body)
	wg.Done()
}
