package blog

import (
	"bufio"
	l4g "code.google.com/p/log4go"
	"errors"
	"fmt"
	"os"
)

var (
	InvalidSection    = errors.New("Invalid Section")
	AlreadySetSection = errors.New("Already Set Section")
	InvalidStructure  = errors.New("Invalid Structure")
)

var log l4g.Logger

func init() {
	log = l4g.NewDefaultLogger(l4g.WARNING)
}

func MissingSection(s string) error {
	return errors.New(fmt.Sprintf("Missing Section: %s", s))
}

const BLOG_FILE_DELIM string = "-----"

func NewBlogFromFile(filename string) (blog *Blog, err error) {
	var (
		section           string
		subsection_header string
		body              string
		line_num          int
	)

	if _, err = os.Stat(filename); os.IsNotExist(err) {
		log.Error("Error finding blog file: %s", err.Error())
		return nil, err
	}

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Error("Error reading blog file: %s", err.Error())
		return nil, err
	}

	blog = new(Blog)
	scanner := bufio.NewScanner(file)
	for line_num = 1; scanner.Scan(); line_num++ {
		line := scanner.Text()

		if line == "" {
			continue
		}
		log.Trace("Line %d: %s", line_num, line)

		if section == "" && IsSection(line) {
			log.Trace("Setting the section to: %s", line)
			section = line
		} else if section == "" {
			err = InvalidSection
			break
		} else if line == BLOG_FILE_DELIM {
			log.Trace("Done parsing section: section=%s body=%s subsection_header=%s", section, body, subsection_header)
			err = addSection(blog, section, body, subsection_header)

			section = ""
			body = ""
			subsection_header = ""

			if err != nil {
				break
			}
		} else {
			if section == SUBSECTION && subsection_header == "" {
				subsection_header = line
			} else {
				body = line
			}
		}
	}

	if !(section == "") && err == nil {
		err = addSection(blog, section, body, subsection_header)
	}

	if err != nil {
		log.Error("Problems reading blog on line %d: %s", line_num, err.Error())
		return
	}

	err = scanner.Err()

	if blog.Title == "" {
		return nil, MissingSection("Title")
	} else if blog.Intro == "" {
		return nil, MissingSection("Intro")
	} else if len(blog.Tags) == 0 {
		return nil, MissingSection("Tags")
	}

	return
}

func addSection(blog *Blog, section string, text string, header string) error {
	if text == "" {
		return InvalidStructure
	}

	log.Trace("Add section: %s %s %s", section, text, header)
	switch section {
	case TITLE:
		if !(blog.Title == "") {
			return AlreadySetSection
		}

		blog.Title = text
	case INTRO:
		if !(blog.Intro == "") {
			return AlreadySetSection
		}

		blog.Intro = text
	case OUTRO:
		if !(blog.Outro == "") {
			return AlreadySetSection
		}

		blog.Outro = text
	case TAG:
		blog.Tags = append(blog.Tags, Tag{text})
	case SUBSECTION:
		if header == "" {
			return InvalidStructure
		}

		blog.Subsections = append(blog.Subsections, Subsection{header, text})
	}

	return nil
}
