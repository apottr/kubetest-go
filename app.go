package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		h, e := os.Hostname()
		if e != nil {
			fmt.Fprintf(w, "Error occured! %q", html.EscapeString(e.Error()))
		} else {
			fmt.Fprintf(w, "My name is %q", html.EscapeString(h))
		}
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
