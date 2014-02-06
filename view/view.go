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
}

func New(blogs *blog.Blogs) View {
	return View{blogs}
}

func (v View) Blog(w http.ResponseWriter, b *blog.Blog) (err error) {
	if b == nil {
		l4g.Info("Requested blog that does not exist")
		fmt.Fprintln(w, "NOT FOUND! :D")
		return
	}

	l4g.Trace("Displaying blog: " + b.Title)

	l4g.Trace("Rendering the blog")
	err = templates.ExecuteTemplate(w, ablog, b)

	if err != nil {
		l4g.Error("Problems rendering template: " + err.Error())
		fmt.Fprintln(w, "Nope! "+err.Error())
	}

	l4g.Trace("Done rendering")
	return
}

func (v View) BlogList(w http.ResponseWriter) (err error) {
	l4g.Info("Blog list")

	l4g.Trace("Blogs given:")
	for _, b := range v.blogs.First() {
		l4g.Trace("\t%s", b.Title)
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
		l4g.Error("Problems rendering template: " + err.Error())
		fmt.Fprintln(w, "Nopes! "+err.Error())
	}

	return
}

func (v View) Index(w http.ResponseWriter) (err error) {
	l4g.Info("Index page")

	err = templates.ExecuteTemplate(w, index, nil)

	return
}

func (v View) Soon(w http.ResponseWriter) (err error) {
	l4g.Info("Coming soon page")

	err = templates.ExecuteTemplate(w, soon, nil)

	return
}
