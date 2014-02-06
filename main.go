package main

import (
	l4g "code.google.com/p/log4go"
	"flag"
	"github.com/jasocox/goblog/blog"
	"github.com/jasocox/goblog/reader"
	"github.com/jasocox/goblog/router"
	"github.com/jasocox/goblog/view"
	"net/http"
)

var (
	blog_dir = flag.String("b", "", "directory where blogs a stored")
	protocol = flag.String("p", "2001", "protocal to run on")
)

func main() {
	var err error

	log := l4g.NewDefaultLogger(l4g.WARNING)
	log.Trace("Starting")

	flag.Parse()

	if *blog_dir == "" {
		log.Error("Must specify a directory where blogs are stored")
	}

	blogs := blog.New()
	blogReader := reader.New(blogs, *blog_dir, log)
	v := view.New(blogs, log)
	router := router.New(v, log)

	err = blogReader.ReadBlogs()
	if err != nil {
		log.Error("Error creating blog reader: %s", err)
	}

	http.Handle("/", router)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	err = http.ListenAndServe(":"+*protocol, nil)
	if err != nil {
		log.Error("Problem with http server: %s", err)
	}

	log.Trace("Stopping")
}
