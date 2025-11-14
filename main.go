package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func addHandler(w http.ResponseWriter, r *http.Request) {
    aStr := r.URL.Query().Get("a")
    bStr := r.URL.Query().Get("b")
    a, err := strconv.Atoi(aStr)
    if err != nil {
        http.Error(w, "invalid or missing 'a' param", http.StatusBadRequest)
        return
    }
    b, err := strconv.Atoi(bStr)
    if err != nil {
        http.Error(w, "invalid or missing 'b' param", http.StatusBadRequest)
        return
    }
    fmt.Fprintf(w, "%d", Add(a, b))
}

func main() {
    http.HandleFunc("/add", addHandler)
    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}
