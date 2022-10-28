package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// TODO: set a http client timeout
	timeout := 50 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://andcloud.io", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	// Close the response body on the return.
	defer resp.Body.Close()

	// Write the response to stdout.
	io.Copy(os.Stdout, resp.Body)

}
