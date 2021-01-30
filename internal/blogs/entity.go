package blogs

// Blog is
type Blog struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// TableName is
func (Blog) TableName() string {
	return "blogs"
}
