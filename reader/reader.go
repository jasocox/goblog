package reader

import (
	"bufio"
	l4g "code.google.com/p/log4go"
	"errors"
	"github.com/jasocox/goblog/blog"
	"os"
)

var NotDirectory = errors.New("Not a Directory")

type BlogReader struct {
	blogs    *blog.Blogs
	blog_dir string
	log      l4g.Logger
}

const BLOG_FILE_DELIM string = "-----"

func New(blogs *blog.Blogs, blog_dir string, log l4g.Logger) (reader BlogReader) {
	reader.blogs = blogs
	reader.blog_dir = blog_dir
	reader.log = log

	return
}

func (r *BlogReader) ReadBlogs() error {
	stat, err := os.Stat(r.blog_dir)

	if os.IsNotExist(err) {
		r.log.Error("Specified blog directory does not exist: %s", r.blog_dir)
		return err
	}

	if !stat.IsDir() {
		err = NotDirectory
		r.log.Error("Specified blog location is not a directory: %s", r.blog_dir)
		return err
	}

	file, err := os.Open(r.blog_dir)
	if err != nil {
		r.log.Error("Error reading the blog directory: %s", err.Error())
		return err
	}

	file_list, err := file.Readdirnames(0)
	if err != nil {
		r.log.Error("Error getting list of file names in the blog directory: %s", err.Error())
		return err
	}

	for _, filename := range file_list {
		var (
			errr     error
			filepath string
		)

		filepath = r.blog_dir + "/" + filename
		r.log.Trace("Reading blog file: %s", filepath)

		errr = r.NewBlogFromFile(filepath)
		if errr != nil {
			r.log.Error("Problems reading a blog file: %s", errr.Error())
			err = errr
		}
	}

	return err
}

func (r BlogReader) NewBlogFromFile(filename string) (err error) {
	b, err := r.readBlog(filename)

	if err == nil {
		r.addBlog(b)
	}

	return
}

func (r BlogReader) readBlog(filename string) (b *blog.Blog, err error) {
	var (
		section           string
		subsection_header string
		body              string
		line_num          int
	)

	if _, err = os.Stat(filename); os.IsNotExist(err) {
		r.log.Error("Error finding blog file: %s", err.Error())
		return nil, err
	}

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		r.log.Error("Error reading blog file: %s", err.Error())
		return nil, err
	}

	b = blog.NewBlog()
	scanner := bufio.NewScanner(file)
	for line_num = 1; scanner.Scan(); line_num++ {
		line := scanner.Text()

		if line == "" {
			continue
		}
		r.log.Trace("Line %d: %s", line_num, line)

		if section == "" {
			r.log.Trace("Setting the section to: %s", line)
			section = line
		} else if line == BLOG_FILE_DELIM {
			r.log.Trace("Done parsing section: section=%s body=%s subsection_header=%s", section, body, subsection_header)
			err = b.AddSection(section, body, subsection_header)

			section = ""
			body = ""
			subsection_header = ""

			if err != nil {
				break
			}
		} else {
			if section == blog.SUBSECTION && subsection_header == "" {
				subsection_header = line
			} else {
				if !(body == "") {
					body += "\n"
				}
				body += line
			}
		}
	}

	if !(section == "") && err == nil {
		r.log.Trace("Add section: %s %s %s", section, body, subsection_header)
		err = b.AddSection(section, body, subsection_header)
	}

	if err != nil {
		r.log.Error("Problems reading blog on line %d: %s", line_num, err.Error())
		return
	}

	err = scanner.Err()
	if err == nil {
		err = b.Validate()
	}

	return
}

func (reader *BlogReader) addBlog(b *blog.Blog) {
	reader.log.Trace("Adding blog: %s", b.Title)
	reader.log.Trace("Blogs:")
	reader.blogs.Add(b)
	reader.log.Trace(reader.blogs.Blogs())
}
