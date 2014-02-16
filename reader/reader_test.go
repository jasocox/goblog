package reader

import (
	l4g "code.google.com/p/log4go"
	"github.com/jasocox/goblog/blog"
	"github.com/jasocox/goblog/reader"
	"testing"
)

var log l4g.Logger

func init() {
	log = l4g.NewLogger()
}

func Test_MakeBlogWithFile(t *testing.T) {
	b := blog.New()
	r := reader.New(b, "dir", log)
	err := r.NewBlogFromFile("example.txt")

	if err != nil {
		t.Error("Unexpected error parsing the blog file:", err.Error())
		return
	}

	if !(b.Blogs()[0].Title == "The Title") {
		t.Error("Did not set the proper title")
		return
	}

	if !(b.Blogs()[0].Intro == "The Intro") {
		t.Error("Did not set the proper intro")
		return
	}

	if !(b.Blogs()[0].Subsections[0].Header == "Sub Title 1") {
		t.Error("Did not set the proper subtitle", b.Blogs()[0].Subsections[0].Header)
		return
	}

	if !(b.Blogs()[0].Subsections[0].Text == "Sub Text 1") {
		t.Error("Did not set the proper sub text")
		return
	}

	if !(b.Blogs()[0].Subsections[1].Header == "Sub Title 2") {
		t.Error("Did not set the proper subtitle")
		return
	}

	if !(b.Blogs()[0].Subsections[1].Text == "Sub Text 2") {
		t.Error("Did not set the proper sub text")
		return
	}

	if !(b.Blogs()[0].Outro == "The Outro") {
		t.Error("Did not set the proper outro")
		return
	}

	if !(b.Blogs()[0].Tags[0].Name == "Tag 1") {
		t.Error("Did not set the proper tag")
		return
	}

	if !(b.Blogs()[0].Tags[1].Name == "Tag 2") {
		t.Error("Did not set the proper tag")
		return
	}
}

func Test_BlogRequiresTitle(t *testing.T) {
	r := reader.New(blog.New(), "dir", log)
	err := r.NewBlogFromFile("missing_title.txt")

	if err == nil {
		t.Error("Expected error")
		return
	}

	if !(err.Error() == "Missing Section: Title") {
		t.Error("Did not get the correct error")
		return
	}
}

func Test_BlogRequiresIntro(t *testing.T) {
	r := reader.New(blog.New(), "dir", log)
	err := r.NewBlogFromFile("missing_intro.txt")

	if err == nil {
		t.Error("Expected error")
		return
	}

	if !(err.Error() == "Missing Section: Intro") {
		t.Error("Did not get the correct error")
		return
	}
}

func Test_BlogRequiresTag(t *testing.T) {
	r := reader.New(blog.New(), "dir", log)
	err := r.NewBlogFromFile("missing_tag.txt")

	if err == nil {
		t.Error("Expected error")
		return
	}

	if !(err.Error() == "Missing Section: Tags") {
		t.Error("Did not get the correct error")
		return
	}
}

func Test_CanHaveMiltiLines(t *testing.T) {
	b := blog.New()
	r := reader.New(b, "dir", log)
	err := r.NewBlogFromFile("multiline_body.txt")

	if err != nil {
		t.Error("Unexpected error")
		return
	}

	if !(b.Blogs()[0].Intro == "The Intro, line 1\nThe Intro, line 2\nThe Intro, line 3") {
		t.Error("Unexpected Intro:", b.Blogs()[0].Intro)
		return
	}
}

func Test_CanReadAndGetBlogs(t *testing.T) {
	blogs := blog.New()
	reader := reader.New(blogs, "../blogs", log)

	reader.ReadBlogs()
	blog1 := blogs.Get("example_1")
	blog2 := blogs.Get("example_2")
	blog3 := blogs.Get("example_3")
	blog4 := blogs.Get("example_4")

	if blog1 == nil || blog2 == nil || blog3 == nil || blog4 == nil {
		t.Error("Did not properly read blogs")
		return
	}

	if blog1.Title != "Example 1" {
		t.Error("Did not receive expected blog")
		return
	}

	if blog2.Title != "Example 2" {
		t.Error("Did not receive expected blog")
		return
	}

	if blog3.Title != "Example 3" {
		t.Error("Did not receive expected blog")
		return
	}

	if blog4.Title != "Example 4" {
		t.Error("Did not receive expected blog")
		return
	}
}
