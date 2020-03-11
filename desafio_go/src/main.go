package main

import (
    "fmt"
    "log"
    "net/http"
)

func Negrito(texto string) string {
    return "<b>" + texto + "</b>"
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, Negrito("Code.education rocks!"))
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}
