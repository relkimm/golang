package webserver

import (
	"fmt"
	"log"
	"net/http"
)

func Execute() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
