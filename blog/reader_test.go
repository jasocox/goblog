package blog

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
