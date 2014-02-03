package reader

import (
	"bufio"
	l4g "code.google.com/p/log4go"
	"errors"
	"fmt"
	"os"
)

var (
	NotDirectory      = errors.New("Not a Directory")
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

type BlogReader struct {
	blogs     map[string]*Blog
	blog_list []*Blog
	blog_dir  string
}

const BLOG_FILE_DELIM string = "-----"

func New(blog_dir string) (reader BlogReader) {
	reader.blogs = make(map[string]*Blog, 0)
	reader.blog_list = make([]*Blog, 0)
	reader.blog_dir = blog_dir

	return
}

func (r *BlogReader) ReadBlogs() error {
	stat, err := os.Stat(r.blog_dir)

	if os.IsNotExist(err) {
		log.Error("Specified blog directory does not exist: %s", r.blog_dir)
		return err
	}

	if !stat.IsDir() {
		err = NotDirectory
		log.Error("Specified blog location is not a directory: %s", r.blog_dir)
		return err
	}

	file, err := os.Open(r.blog_dir)
	if err != nil {
		log.Error("Error reading the blog directory: %s", err.Error())
		return err
	}

	file_list, err := file.Readdirnames(0)
	if err != nil {
		log.Error("Error getting list of file names in the blog directory: %s", err.Error())
		return err
	}

	for _, filename := range file_list {
		var (
			blog     *Blog
			errr     error
			filepath string
		)

		filepath = r.blog_dir + "/" + filename
		log.Trace("Reading blog file: %s", filepath)

		blog, errr = NewBlogFromFile(filepath)
		if errr != nil {
			log.Error("Problems reading a blog file: %s", errr.Error())
			err = errr
		} else {
			r.addBlog(blog)
		}
	}

	return err
}

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

func (reader *BlogReader) addBlog(b *Blog) {
	l4g.Trace("Adding blog: %s", b.Title)
	reader.blogs[b.HashTitle()] = b
	reader.blog_list = append(reader.blog_list, b)
}

func (reader *BlogReader) GetBlog(hashed_title string) *Blog {
	return reader.blogs[hashed_title]
}

func (reader *BlogReader) First() []*Blog {
	blog_len := len(reader.blog_list)
	size := blog_len

	if size > 3 {
		size = 3
	}

	blog_list := make([]*Blog, size)

	for i := 0; i < size; i++ {
		blog_list[i] = reader.blog_list[blog_len-1-i]
	}

	return blog_list
}
