package reader

import (
	l4g "code.google.com/p/log4go"
	"github.com/jasocox/goblog/blog"
	"testing"
)

var log l4g.Logger

func init() {
	log = l4g.NewLogger()
}

func Test_MakeBlogWithFile(t *testing.T) {
	r := New(blog.New(), "dir", log)
	blog, err := r.NewBlogFromFile("example.txt")

	if err != nil {
		t.Error("Unexpected error parsing the blog file:", err.Error())
		return
	}

	if !(blog.Title == "The Title") {
		t.Error("Did not set the proper title")
		return
	}

	if !(blog.Intro == "The Intro") {
		t.Error("Did not set the proper intro")
		return
	}

	if !(blog.Subsections[0].Header == "Sub Title 1") {
		t.Error("Did not set the proper subtitle", blog.Subsections[0].Header)
		return
	}

	if !(blog.Subsections[0].Text == "Sub Text 1") {
		t.Error("Did not set the proper sub text")
		return
	}

	if !(blog.Subsections[1].Header == "Sub Title 2") {
		t.Error("Did not set the proper subtitle")
		return
	}

	if !(blog.Subsections[1].Text == "Sub Text 2") {
		t.Error("Did not set the proper sub text")
		return
	}

	if !(blog.Outro == "The Outro") {
		t.Error("Did not set the proper outro")
		return
	}

	if !(blog.Tags[0].Name == "Tag 1") {
		t.Error("Did not set the proper tag")
		return
	}

	if !(blog.Tags[1].Name == "Tag 2") {
		t.Error("Did not set the proper tag")
		return
	}
}

func Test_BlogRequiresTitle(t *testing.T) {
	r := New(blog.New(), "dir", log)
	_, err := r.NewBlogFromFile("missing_title.txt")

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
	r := New(blog.New(), "dir", log)
	_, err := r.NewBlogFromFile("missing_intro.txt")

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
	r := New(blog.New(), "dir", log)
	_, err := r.NewBlogFromFile("missing_tag.txt")

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
	r := New(blog.New(), "dir", log)
	blog, err := r.NewBlogFromFile("multiline_body.txt")

	if err != nil {
		t.Error("Unexpected error")
		return
	}

	if !(blog.Intro == "The Intro, line 1\nThe Intro, line 2\nThe Intro, line 3") {
		t.Error("Unexpected Intro:", blog.Intro)
		return
	}
}

func Test_CanAddAndGetBlogs(t *testing.T) {
	blogs := blog.New()
	reader := New(blogs, "dir", log)

	blog1 := &blog.Blog{Title: "Title 1"}
	blog2 := &blog.Blog{Title: "Title 2"}
	blog3 := &blog.Blog{Title: "Title 3"}

	reader.addBlog(blog1)
	reader.addBlog(blog2)
	reader.addBlog(blog3)

	if blogs.Blogs()[0] != blog1 {
		t.Error("Did not receive expected blog")
		return
	}

	if blogs.Blogs()[1] != blog2 {
		t.Error("Did not receive expected blog")
		return
	}

	if blogs.Blogs()[2] != blog3 {
		t.Error("Did not receive expected blog")
		return
	}
}
