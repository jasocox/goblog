package main

import (
	l4g "code.google.com/p/log4go"
	"flag"
	"fmt"
	"html"
	"net/http"
)

var blog_dir = flag.String("b", "", "directory where blogs a stored")

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	l4g.Trace("Handling request for %s", html.EscapeString(r.URL.Path))

	fmt.Fprintln(w, "Blogs")
}

func main() {
	l4g.Trace("Starting")

	flag.Parse()

	if *blog_dir == "" {
		l4g.Error("Must specify a directory where blogs are stored")
	}

	http.HandleFunc("/", HandleRoot)

	err := http.ListenAndServe(":2001", nil)
	if err != nil {
		l4g.Error("Problem with http server: %s", err)
	}

	l4g.Trace("Stopping")
}
