package blog

import (
	"errors"
	"time"
)

var InvalidBlog = errors.New("Invalid blog")

func NewBlog(parsed []string) (Blog, error) {
	var (
		title       string
		intro       string
		outro       string
		subsections = make([]Subsection, 0)
		tags        = make([]Tag, 0)
	)

	for i := 0; i < len(parsed); i++ {
		switch parsed[i] {
		case TITLE:
			title = parsed[i+1]
		case INTRO:
			intro = parsed[i+1]
		case OUTRO:
			outro = parsed[i+1]
		case TAG:
			tags = append(tags, Tag{parsed[i+1]})
		case SUBSECTION:
			subsections = append(subsections, Subsection{parsed[i+1], parsed[i+2]})
			i++
		default:
			return Blog{}, InvalidBlog
		}

		i++
	}

	return Blog{title, time.Now(), intro, subsections, outro, tags}, nil
}
