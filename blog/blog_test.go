package blog

import (
	"github.com/jasocox/goblog/blog"
	"testing"
	"time"
)

func Test_Blog(t *testing.T) {
	subs := []blog.Subsection{blog.Subsection{Header: "Header 1", Text: "Text 1"}, blog.Subsection{Header: "Header 2", Text: "Text 2"}}

	tags := []blog.Tag{blog.Tag{Name: "Tag"}, blog.Tag{Name: "Tag2"}}

	_ = blog.Blog{Title: "Title", Date: time.Now(), Intro: "Intro text", Subsections: subs, Outro: "Outro text", Tags: tags}
}

func Test_HashFromTitle(t *testing.T) {
	b := &blog.Blog{Title: "Title 1"}

	if b.HashTitle() != "title_1" {
		t.Error("Did not get expected title")
		return
	}
}
