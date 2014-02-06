package blog

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

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

var replaceSpaces *strings.Replacer

var (
	InvalidSection    = errors.New("Invalid Section")
	AlreadySetSection = errors.New("Already Set Section")
	InvalidStructure  = errors.New("Invalid Structure")
)

func init() {
	replaceSpaces = strings.NewReplacer(" ", "_")
}

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

func (b *Blog) AddSection(section string, body string, subsection_header string) error {
	if !IsSection(section) {
		return InvalidSection
	}

	if body == "" {
		return InvalidStructure
	}

	switch section {
	case TITLE:
		if !(b.Title == "") {
			return AlreadySetSection
		}

		b.Title = body
	case INTRO:
		if !(b.Intro == "") {
			return AlreadySetSection
		}

		b.Intro = body
	case OUTRO:
		if !(b.Outro == "") {
			return AlreadySetSection
		}

		b.Outro = body
	case TAG:
		b.Tags = append(b.Tags, Tag{body})
	case SUBSECTION:
		if subsection_header == "" {
			return InvalidStructure
		}

		b.Subsections = append(b.Subsections, Subsection{subsection_header, body})
	}

	return nil

}

func MissingSection(s string) error {
	return errors.New(fmt.Sprintf("Missing Section: %s", s))
}

func (b *Blog) HashTitle() string {
	return replaceSpaces.Replace(strings.ToLower(strings.TrimSpace(b.Title)))
}

func (b Blog) Validate() error {
	if b.Title == "" {
		return MissingSection("Title")
	} else if b.Intro == "" {
		return MissingSection("Intro")
	} else if len(b.Tags) == 0 {
		return MissingSection("Tags")
	}

	return nil
}
