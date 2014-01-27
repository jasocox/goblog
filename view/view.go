package view

import (
	l4g "code.google.com/p/log4go"
	"fmt"
	"github.com/jasocox/goblog/reader"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("view/blog.html"))

func Blog(w http.ResponseWriter, b *reader.Blog) {
	if b == nil {
		l4g.Info("Requested blog that does not exist")
		fmt.Fprintln(w, "NOT FOUND! :D")
		return
	}

	l4g.Trace("Displaying blog: " + b.Title)

	templates.ExecuteTemplate(w, "blog.html", b)
}
