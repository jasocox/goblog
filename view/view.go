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

func (v View) Blog(w http.ResponseWriter, blog_name string) (err error) {
	b := v.blogs.Get(blog_name)
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

func (v View) BlogList(w http.ResponseWriter) error {
	v.log.Trace("Blogs given:")
	for _, b := range v.blogs.First() {
		v.log.Trace("\t%s", b.Title)
	}

	execers := make([]templateExecer, 4)
	execers[0] = templateExecer{header, nil}
	execers[1] = templateExecer{blog_list, v.blogs.First()}
	execers[2] = templateExecer{blog_links, v.blogs.Last()}
	execers[3] = templateExecer{footer, nil}

	return v.execTemplates("Blog List", execers, w)
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

type templateExecer struct {
	name string
	data interface{}
}

func (e templateExecer) executeTemplate(w http.ResponseWriter) error {
	return templates.ExecuteTemplate(w, e.name, e.data)
}

func executeTemplates(execers []templateExecer, w http.ResponseWriter) (err error) {
	for _, e := range execers {
		err = e.executeTemplate(w)

		if err != nil {
			break
		}
	}

	return
}

func (v View) execTemplates(name string, execers []templateExecer, w http.ResponseWriter) (err error) {
	v.log.Info("%s page", name)

	err = executeTemplates(execers, w)

	if err != nil {
		v.log.Error("Problems rendering template: " + err.Error())
		fmt.Fprintln(w, "Nopes! "+err.Error())
	}

	return
}
