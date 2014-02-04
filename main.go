package main

import (
	l4g "code.google.com/p/log4go"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jasocox/goblog/reader"
	"github.com/jasocox/goblog/view"
	"html"
	"net/http"
)

var (
	blog_dir   = flag.String("b", "", "directory where blogs a stored")
	blogReader reader.BlogReader
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	l4g.Trace("Handling request for " + html.EscapeString(r.URL.Path))

	err := view.Index(w)

	if err != nil {
		l4g.Error(err)
	}
}

func BlogListHandler(w http.ResponseWriter, r *http.Request) {
	l4g.Trace("List blog request " + html.EscapeString(r.URL.Path))

	err := view.BlogList(w, blogReader.First(), blogReader.Last())

	if err != nil {
		l4g.Error(err)
	}

	fmt.Fprintln(w, "Blogs")
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	l4g.Trace("Handling blog request " + html.EscapeString(r.URL.Path))

	err := view.Blog(w, blogReader.GetBlog(mux.Vars(r)["blog"]))

	if err != nil {
		l4g.Error(err)
	}
}

func ComingSoonHandler(w http.ResponseWriter, r *http.Request) {
	l4g.Trace("Request %s is listed as coming soon", html.EscapeString(r.URL.Path))

	err := view.Soon(w)

	if err != nil {
		l4g.Error(err)
	}
}

func main() {
	var err error
	l4g.Trace("Starting")

	flag.Parse()

	if *blog_dir == "" {
		l4g.Error("Must specify a directory where blogs are stored")
	}

	blogReader = reader.New(*blog_dir)

	err = blogReader.ReadBlogs()
	if err != nil {
		l4g.Error("Error creating blog reader: %s", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/blogs", BlogListHandler)
	router.HandleFunc("/blogs/{blog}", BlogHandler)

	router.HandleFunc("/contact_me", ComingSoonHandler)
	router.HandleFunc("/about_me", ComingSoonHandler)

	http.Handle("/", router)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	err = http.ListenAndServe(":2001", nil)
	if err != nil {
		l4g.Error("Problem with http server: %s", err)
	}

	l4g.Trace("Stopping")
}
