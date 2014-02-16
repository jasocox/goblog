package blog

import (
	"github.com/jasocox/goblog/blog"
	"testing"
)

func Test_AddABlog(t *testing.T) {
	blogs := blog.New()

	if len(blogs.Blogs()) != 0 {
		t.Error("Shouldn't have any blogs yet")
		return
	}

	blogs.Add(&blog.Blog{Title: "Example 1", Intro: "Intro"})

	if len(blogs.Blogs()) != 1 {
		t.Error("Should have a blog")
		return
	}

	if !(blogs.Blogs()[0].Title == "Example 1") {
		t.Error("Got the wrong blog title")
		return
	}

	blogs.Add(&blog.Blog{Title: "Example 2", Intro: "Intro"})

	if len(blogs.Blogs()) != 2 {
		t.Error("Should have two blogs")
		return
	}

	if !(blogs.Blogs()[0].Title == "Example 1") {
		t.Error("Got the wrong blog title")
		return
	}

	if !(blogs.Blogs()[1].Title == "Example 2") {
		t.Error("Got the wrong blog title")
		return
	}
}

func Test_CanGetBlogs(t *testing.T) {
	blogs := blog.New()

	blog1 := &blog.Blog{Title: "Title 1"}
	blog2 := &blog.Blog{Title: "Title 2"}
	blog3 := &blog.Blog{Title: "Title 3"}

	blogs.Add(blog1)
	blogs.Add(blog2)
	blogs.Add(blog3)

	if blogs.Get("title_1") == nil {
		t.Error("Did not receive a blog")
		return
	}

	if blogs.Get("title_2") == nil {
		t.Error("Did not receive a blog")
		return
	}

	if blogs.Get("title_3") == nil {
		t.Error("Did not receive a blog")
		return
	}

	if blogs.Get("title_1") != blog1 {
		t.Error("Did not receive expected blog")
		return
	}

	if blogs.Get("title_2") != blog2 {
		t.Error("Did not receive expected blog")
		return
	}

	if blogs.Get("title_3") != blog3 {
		t.Error("Did not receive expected blog")
		return
	}
}

func Test_GetsNilIfDoesntExist(t *testing.T) {
	blogs := blog.New()
	blogs.Add(&blog.Blog{Title: "dont care"})

	if blogs.Get("title_1") != nil {
		t.Error("Expected nil for non-existant blog")
		return
	}
}

func Test_FirstGivesFirstThree(t *testing.T) {
	blogs := blog.New()

	if len(blogs.First()) != 0 {
		t.Error("Shouldn't have returned a blog")
		return
	}

	blogs.Add(&blog.Blog{Title: "one"})

	if len(blogs.First()) != 1 {
		t.Error("Should have just one blog")
		return
	}

	if !(blogs.First()[0].Title == "one") {
		t.Error("Wrong blog")
		return
	}

	blogs.Add(&blog.Blog{Title: "two"})

	if len(blogs.First()) != 2 {
		t.Error("Should have two blogs")
		return
	}

	if !(blogs.First()[1].Title == "one") {
		t.Error("Wrong blog")
		return
	}

	if !(blogs.First()[0].Title == "two") {
		t.Error("Wrong blog")
		return
	}

	blogs.Add(&blog.Blog{Title: "three"})

	if len(blogs.First()) != 3 {
		t.Error("Should have three blogs")
		return
	}

	if !(blogs.First()[2].Title == "one") {
		t.Error("Wrong blog")
		return
	}

	if !(blogs.First()[1].Title == "two") {
		t.Error("Wrong blog")
		return
	}

	if !(blogs.First()[0].Title == "three") {
		t.Error("Wrong blog")
		return
	}

	blogs.Add(&blog.Blog{Title: "four"})

	if len(blogs.First()) != 3 {
		t.Error("Should have three blogs")
		return
	}

	if !(blogs.First()[2].Title == "two") {
		t.Error("Wrong blog")
		return
	}

	if !(blogs.First()[1].Title == "three") {
		t.Error("Wrong blog")
		return
	}

	if !(blogs.First()[0].Title == "four") {
		t.Error("Wrong blog")
		return
	}
}

func Test_LastSkipsFirstThree(t *testing.T) {
	blogs := blog.New()

	if len(blogs.Last()) != 0 {
		t.Error("Shouldn't have returned a blog")
		return
	}

	blogs.Add(&blog.Blog{Title: "one"})

	if len(blogs.Last()) != 0 {
		t.Error("Shouldn't have returned a blog")
		return
	}

	blogs.Add(&blog.Blog{Title: "two"})

	if len(blogs.Last()) != 0 {
		t.Error("Should have no blogs")
		return
	}

	blogs.Add(&blog.Blog{Title: "three"})

	if len(blogs.Last()) != 0 {
		t.Error("Should have no blogs")
		return
	}

	blogs.Add(&blog.Blog{Title: "four"})

	if len(blogs.Last()) != 1 {
		t.Error("Should have a blog")
		return
	}

	if !(blogs.Last()[0].Title == "one") {
		t.Error("Wrong blog")
		return
	}

	blogs.Add(&blog.Blog{Title: "five"})

	if len(blogs.Last()) != 2 {
		t.Error("Should have a blog")
		return
	}

	if !(blogs.Last()[0].Title == "two") {
		t.Error("Wrong blog")
		return
	}

	if !(blogs.Last()[1].Title == "one") {
		t.Error("Wrong blog")
		return
	}
}
