package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		//Exercise 2: Insert http:// prefix
		if !strings.HasPrefix(strings.ToLower(url), "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		//Exercise 3: Print status code
		fmt.Fprintf(os.Stderr, "Status code: %s \n", resp.Status)

		//Exercise 1: Using io.Copy
		_, err = io.Copy(os.Stdout, resp.Body)

		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
	}
}
