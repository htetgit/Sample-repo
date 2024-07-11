package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })
    http.Handle("/", r)
    fmt.Println("Server started at http://localhost:8082")
    if err := http.ListenAndServe(":8082", nil); err != nil {
        fmt.Println("Server failed:", err)
    }
}
