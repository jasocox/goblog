package main

import (
	l4g "code.google.com/p/log4go"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jasocox/goblog/reader"
	"html"
	"net/http"
)

var blog_dir = flag.String("b", "", "directory where blogs a stored")

func RootHandler(w http.ResponseWriter, r *http.Request) {
	l4g.Trace("Handling request for " + html.EscapeString(r.URL.Path))

	fmt.Fprintln(w, "Home")
}

func BlogListHandler(w http.ResponseWriter, r *http.Request) {
	l4g.Trace("List blog request " + html.EscapeString(r.URL.Path))

	fmt.Fprintln(w, "Blogs")
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	l4g.Trace("Handling blog request " + html.EscapeString(r.URL.Path))

	fmt.Fprintln(w, "A Blog")
}

func main() {
	var err error
	l4g.Trace("Starting")

	flag.Parse()

	if *blog_dir == "" {
		l4g.Error("Must specify a directory where blogs are stored")
	}

	blogReader := reader.New(*blog_dir)

	err = blogReader.ReadBlogs()
	if err != nil {
		l4g.Error("Error creating blog reader: %s", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/blogs", BlogListHandler)
	router.HandleFunc("/blogs/{blog}", BlogHandler)

	http.Handle("/", router)
	err = http.ListenAndServe(":2001", nil)
	if err != nil {
		l4g.Error("Problem with http server: %s", err)
	}

	l4g.Trace("Stopping")
}
