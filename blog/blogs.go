package blog

type Blogs struct {
	blogs     map[string]*Blog
	blog_list []*Blog
}

func New() *Blogs {
	return &Blogs{make(map[string]*Blog, 0), make([]*Blog, 0)}
}

func NewBlog() *Blog {
	return &Blog{}
}

func (blogs *Blogs) Add(b *Blog) {
	blogs.blogs[b.HashTitle()] = b
	blogs.blog_list = append(blogs.blog_list, b)
}

func (blogs *Blogs) Blogs() []*Blog {
	return blogs.blog_list
}

func (blogs *Blogs) Get(title string) *Blog {
	return blogs.blogs[title]
}

func (blogs *Blogs) First() []*Blog {
	blog_len := len(blogs.blog_list)
	size := blog_len

	if size > 3 {
		size = 3
	}

	blog_list := make([]*Blog, size)

	for i := 0; i < size; i++ {
		blog_list[i] = blogs.blog_list[blog_len-1-i]
	}

	return blog_list
}

func (blogs *Blogs) Last() []*Blog {
	blog_len := len(blogs.blog_list)
	size := blog_len - 3

	if size <= 0 {
		return nil
	}

	blog_list := make([]*Blog, size)

	for i := 0; i < size; i++ {
		blog_list[i] = blogs.blog_list[size-1-i]
	}

	return blog_list
}
