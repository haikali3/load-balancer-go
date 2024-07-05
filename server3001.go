package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from server on port %s!", r.Host)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on port 3001")
	http.ListenAndServe(":3001", nil)
}
