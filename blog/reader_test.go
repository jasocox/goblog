package blog

import "testing"

func Test_MakeBlogWithParsedStrings(t *testing.T) {
	parsed := []string{"Title", "The Title", "Intro", "The Intro", "Subsection", "Sub Title 1", "Sub Text 1", "Subsection", "Sub Title 2", "Sub Text 2", "Outro", "The Outro", "Tag", "Tag 1", "Tag", "Tag 2"}

	blog, err := NewBlog(parsed)

	if err != nil {
		t.Error("Unexpected error parsing the blog strings:", err.Error())
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
		t.Error("Did not set the proper subtitle")
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
	parsed := []string{"Intro", "The Intro", "Subsection", "Sub Title 1", "Sub Text 1", "Subsection", "Sub Title 2", "Sub Text 2", "Outro", "The Outro", "Tag", "Tag 1", "Tag", "Tag 2"}

	_, err := NewBlog(parsed)

	if err == nil {
		t.Error("Expected error")
		return
	}
}

func Test_BlogRequiresIntro(t *testing.T) {
	parsed := []string{"Title", "The Title", "Subsection", "Sub Title 1", "Sub Text 1", "Subsection", "Sub Title 2", "Sub Text 2", "Outro", "The Outro", "Tag", "Tag 1", "Tag", "Tag 2"}

	_, err := NewBlog(parsed)

	if err == nil {
		t.Error("Expected error")
		return
	}
}

func Test_BlogRequiresTag(t *testing.T) {
	parsed := []string{"Title", "The Title", "Intro", "The Intro", "Subsection", "Sub Title 1", "Sub Text 1", "Subsection", "Sub Title 2", "Sub Text 2", "Outro", "The Outro"}

	_, err := NewBlog(parsed)

	if err == nil {
		t.Error("Expected error")
		return
	}
}
