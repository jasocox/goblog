package main

import (
  "fmt"
  "net/http"
)

func main() {
  fmt.Println("Starting")

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Blogs")
  })

  err := http.ListenAndServe(":2001", nil)
  if err != nil {
    fmt.Println("Something broke...")
    fmt.Println(err.Error())
  }

  fmt.Println("Stopping")
}
