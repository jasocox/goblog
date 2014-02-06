package router

import (
	l4g "code.google.com/p/log4go"
	"github.com/gorilla/mux"
	"github.com/jasocox/goblog/view"
	"html"
	"net/http"
)

type Router struct {
	log l4g.Logger
	v   view.View
}

func New(v view.View, log l4g.Logger) *mux.Router {
	r := Router{log, v}
	router := mux.NewRouter()
	router.HandleFunc("/", r.RootHandler)
	router.HandleFunc("/blogs", r.BlogListHandler)
	router.HandleFunc("/blogs/{blog}", r.BlogHandler)

	router.HandleFunc("/contact_me", r.ComingSoonHandler)
	router.HandleFunc("/about_me", r.ComingSoonHandler)

	return router
}

func (router Router) RootHandler(w http.ResponseWriter, r *http.Request) {
	router.log.Trace("Handling request for " + html.EscapeString(r.URL.Path))

	err := router.v.Index(w)

	if err != nil {
		router.log.Error(err)
	}
}

func (router Router) BlogListHandler(w http.ResponseWriter, r *http.Request) {
	router.log.Trace("List blog request " + html.EscapeString(r.URL.Path))

	err := router.v.BlogList(w)

	if err != nil {
		router.log.Error(err)
	}
}

func (router Router) BlogHandler(w http.ResponseWriter, r *http.Request) {
	router.log.Trace("Handling blog request " + html.EscapeString(r.URL.Path))

	err := router.v.Blog(w, mux.Vars(r)["blog"])

	if err != nil {
		router.log.Error(err)
	}
}

func (router Router) ComingSoonHandler(w http.ResponseWriter, r *http.Request) {
	router.log.Trace("Request %s is listed as coming soon", html.EscapeString(r.URL.Path))

	err := router.v.Soon(w)

	if err != nil {
		router.log.Error(err)
	}
}
