package dto

type BookDto struct {
	ID         uint          `json:"id"`
	Title      string        `json:"title"`
	Pages      int64         `json:"pages"`
	CoverImage string        `json:"cover_image"`
	Notes      string        `json:"notes"`
	Finished   bool          `json:"finished"`
	Author     AuthorDto     `json:"author"`
	Categories []CategoryDto `json:"categories"`
}
