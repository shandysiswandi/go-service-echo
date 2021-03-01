package jsonplaceholder

// JSONPlaceHolder is
type JSONPlaceHolder struct {
	postsEndpoint string
}

// Post is
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
