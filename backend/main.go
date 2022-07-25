package main

import (
	"fmt"
	"net/http"
)

func setupRoutes() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Simple Server")
    })
}

func main() {
    setupRoutes()
    fmt.Println("Server starting ...")
    http.ListenAndServe(":4000", nil)
}