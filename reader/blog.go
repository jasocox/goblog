package reader

import "time"

type Blog struct {
	Title       string
	Date        time.Time
	Intro       string
	Subsections []Subsection
	Outro       string
	Tags        []Tag
}

type Subsection struct {
	Header string
	Text   string
}

type Tag struct {
	Name string
}

const (
	TITLE      string = "Title"
	INTRO      string = "Intro"
	SUBSECTION string = "Subsection"
	OUTRO      string = "Outro"
	TAG        string = "Tag"
)

func IsSection(s string) bool {
	switch s {
	case TITLE:
		return true
	case INTRO:
		return true
	case SUBSECTION:
		return true
	case OUTRO:
		return true
	case TAG:
		return true
	}

	return false
}
