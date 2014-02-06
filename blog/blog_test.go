package blog

import (
	"testing"
	"time"
)

func Test_Blog(t *testing.T) {
	subs := []Subsection{Subsection{Header: "Header 1", Text: "Text 1"}, Subsection{Header: "Header 2", Text: "Text 2"}}

	tags := []Tag{Tag{Name: "Tag"}, Tag{Name: "Tag2"}}

	_ = Blog{Title: "Title", Date: time.Now(), Intro: "Intro text", Subsections: subs, Outro: "Outro text", Tags: tags}
}

func Test_HashFromTitle(t *testing.T) {
	blog := &Blog{Title: "Title 1"}

	if blog.HashTitle() != "title_1" {
		t.Error("Did not get expected title")
		return
	}
}
