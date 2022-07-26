package main

import (
	"fmt"
	"net/http"
	"project1/handlers"
)

const portNumber = ":4000"




func main() {
    http.HandleFunc("/", handlers.Home)
    http.HandleFunc("/about", handlers.AboutUs)

    fmt.Printf("Server starting at port %v ...", portNumber)
    http.ListenAndServe(portNumber, nil)
}