package view

import (
	l4g "code.google.com/p/log4go"
	"fmt"
	"github.com/jasocox/goblog/reader"
	"html/template"
	"net/http"
)

var (
	views     = "view/"
	header    = "header.html"
	footer    = "footer.html"
	blog      = "blog.html"
	templates = template.Must(template.ParseFiles(
		views+header,
		views+footer,
		views+blog,
	))
)

func Blog(w http.ResponseWriter, b *reader.Blog) (err error) {
	if b == nil {
		l4g.Info("Requested blog that does not exist")
		fmt.Fprintln(w, "NOT FOUND! :D")
		return
	}

	l4g.Trace("Displaying blog: " + b.Title)

	l4g.Trace("Rendering the header")
	err = renderHeader(w)
	if err == nil {
		l4g.Trace("Rendering the blog")
		err = templates.ExecuteTemplate(w, blog, b)
	}
	if err == nil {
		l4g.Trace("Rendering the footer")
		err = renderFooter(w)
	}

	if err != nil {
		l4g.Error("Problems rendering template: " + err.Error())
		fmt.Fprintln(w, "Nope! "+err.Error())
	}

	l4g.Trace("Done rendering")
	return
}

func renderHeader(w http.ResponseWriter) error {
	return templates.ExecuteTemplate(w, header, nil)
}

func renderFooter(w http.ResponseWriter) error {
	return templates.ExecuteTemplate(w, footer, nil)
}
