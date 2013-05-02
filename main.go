package main

import (
    "net"
    "net/http"
    "fmt"
    "io"
    "log"
)

func dotHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := net.Dial("tcp", "localhost:2004")
    if err != nil {
        fmt.Fprint(w, "Error connecting to olsrd dot_draw")
    }
    io.Copy(w, conn)
}

func main() {
    http.Handle("/", http.FileServer(http.Dir("html")))
    http.HandleFunc("/dot", dotHandler)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
