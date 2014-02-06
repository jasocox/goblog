package view

import (
	l4g "code.google.com/p/log4go"
	"fmt"
	"github.com/jasocox/goblog/blog"
	"html/template"
	"net/http"
)

var (
	views      = "view/"
	header     = "header.html"
	footer     = "footer.html"
	index      = "index.html"
	soon       = "soon.html"
	ablog      = "blog.html"
	blog_list  = "blog_list.html"
	blog_links = "blog_links.html"
	templates  = template.Must(template.ParseFiles(
		views+header,
		views+footer,
		views+index,
		views+soon,
		views+ablog,
		views+blog_list,
		views+blog_links,
	))
)

type View struct {
	blogs *blog.Blogs
	log   l4g.Logger
}

func New(blogs *blog.Blogs, log l4g.Logger) View {
	return View{blogs, log}
}

func (v View) Blog(w http.ResponseWriter, b *blog.Blog) (err error) {
	if b == nil {
		v.log.Info("Requested blog that does not exist")
		fmt.Fprintln(w, "NOT FOUND! :D")
		return
	}

	v.log.Trace("Displaying blog: " + b.Title)

	v.log.Trace("Rendering the blog")
	err = templates.ExecuteTemplate(w, ablog, b)

	if err != nil {
		v.log.Error("Problems rendering template: " + err.Error())
		fmt.Fprintln(w, "Nope! "+err.Error())
	}

	v.log.Trace("Done rendering")
	return
}

func (v View) BlogList(w http.ResponseWriter) (err error) {
	v.log.Info("Blog list")

	v.log.Trace("Blogs given:")
	for _, b := range v.blogs.First() {
		v.log.Trace("\t%s", b.Title)
	}

	err = templates.ExecuteTemplate(w, header, nil)
	if err == nil {
		err = templates.ExecuteTemplate(w, blog_list, v.blogs.First())
	}
	if err == nil {
		err = templates.ExecuteTemplate(w, blog_links, v.blogs.Last())
	}
	if err == nil {
		err = templates.ExecuteTemplate(w, footer, nil)
	}

	if err != nil {
		v.log.Error("Problems rendering template: " + err.Error())
		fmt.Fprintln(w, "Nopes! "+err.Error())
	}

	return
}

func (v View) Index(w http.ResponseWriter) (err error) {
	v.log.Info("Index page")

	err = templates.ExecuteTemplate(w, index, nil)

	return
}

func (v View) Soon(w http.ResponseWriter) (err error) {
	v.log.Info("Coming soon page")

	err = templates.ExecuteTemplate(w, soon, nil)

	return
}
