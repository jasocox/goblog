package main

import (
  "fmt"
  "html"
  "net/http"
  l4g "code.google.com/p/log4go"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
  l4g.Trace("Handling request for %s", html.EscapeString(r.URL.Path))

  fmt.Fprintln(w, "Blogs")
}

func main() {
  l4g.Trace("Starting")

  http.HandleFunc("/", HandleRoot)

  err := http.ListenAndServe(":2001", nil)
  if err != nil {
    fmt.Println("Something broke...")
    fmt.Println(err.Error())
  }

  l4g.Trace("Stopping")
}
