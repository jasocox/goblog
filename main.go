package main

import (
	l4g "code.google.com/p/log4go"
	"flag"
	"github.com/gorilla/mux"
	"github.com/jasocox/goblog/blog"
	"github.com/jasocox/goblog/reader"
	"github.com/jasocox/goblog/view"
	"html"
	"net/http"
)

var (
	blog_dir   = flag.String("b", "", "directory where blogs a stored")
	protocol   = flag.String("p", "2001", "protocal to run on")
	blogs      *blog.Blogs
	blogReader reader.BlogReader
	v          view.View
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	l4g.Trace("Handling request for " + html.EscapeString(r.URL.Path))

	err := v.Index(w)

	if err != nil {
		l4g.Error(err)
	}
}

func BlogListHandler(w http.ResponseWriter, r *http.Request) {
	l4g.Trace("List blog request " + html.EscapeString(r.URL.Path))

	err := v.BlogList(w)

	if err != nil {
		l4g.Error(err)
	}
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	l4g.Trace("Handling blog request " + html.EscapeString(r.URL.Path))

	err := v.Blog(w, blogs.Get(mux.Vars(r)["blog"]))

	if err != nil {
		l4g.Error(err)
	}
}

func ComingSoonHandler(w http.ResponseWriter, r *http.Request) {
	l4g.Trace("Request %s is listed as coming soon", html.EscapeString(r.URL.Path))

	err := v.Soon(w)

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

	blogs = blog.New()
	blogReader = reader.New(blogs, *blog_dir)
	v = view.New(blogs)

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
	err = http.ListenAndServe(":"+*protocol, nil)
	if err != nil {
		l4g.Error("Problem with http server: %s", err)
	}

	l4g.Trace("Stopping")
}
