package reader

import "testing"

func Test_MakeBlogWithFile(t *testing.T) {
	blog, err := NewBlogFromFile("example.txt")

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
	_, err := NewBlogFromFile("missing_title.txt")

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
	_, err := NewBlogFromFile("missing_intro.txt")

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
	_, err := NewBlogFromFile("missing_tag.txt")

	if err == nil {
		t.Error("Expected error")
		return
	}

	if !(err.Error() == "Missing Section: Tags") {
		t.Error("Did not get the correct error")
		return
	}
}

func Test_CanAddAndGetBlogs(t *testing.T) {
	reader := New("dir")

	blog1 := &Blog{Title: "Title 1"}
	blog2 := &Blog{Title: "Title 2"}
	blog3 := &Blog{Title: "Title 3"}

	reader.addBlog(blog1)
	reader.addBlog(blog2)
	reader.addBlog(blog3)

	if reader.GetBlog("title_1") == nil {
		t.Error("Did not receive a blog")
		return
	}

	if reader.GetBlog("title_2") == nil {
		t.Error("Did not receive a blog")
		return
	}

	if reader.GetBlog("title_3") == nil {
		t.Error("Did not receive a blog")
		return
	}

	if reader.GetBlog("title_1") != blog1 {
		t.Error("Did not receive expected blog")
		return
	}

	if reader.GetBlog("title_2") != blog2 {
		t.Error("Did not receive expected blog")
		return
	}

	if reader.GetBlog("title_3") != blog3 {
		t.Error("Did not receive expected blog")
		return
	}
}

func Test_GetsNilIfDoesntExist(t *testing.T) {
	reader := New("dir")
	reader.addBlog(&Blog{Title: "dont care"})

	if reader.GetBlog("title_1") != nil {
		t.Error("Expected nil for non-existant blog")
		return
	}
}

func Test_FirstGivesFirstThree(t *testing.T) {
	var blogs []*Blog
	reader := New("dir")

	blogs = reader.First()
	if len(blogs) != 0 {
		t.Error("Shouldn't have returned a blog")
		return
	}

	reader.addBlog(&Blog{Title: "one"})

	blogs = reader.First()
	if len(blogs) != 1 {
		t.Error("Should have just one blog")
		return
	}

	if !(blogs[0].Title == "one") {
		t.Error("Wrong blog")
		return
	}
}
