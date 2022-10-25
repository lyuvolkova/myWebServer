package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "hello DIMA :) !")

}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%s: %s\n", name, h)
		}
	}

}
