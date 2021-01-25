package jsonplaceholder

import "errors"

// The URL from jsonplaceholder
const (
	BaseURL   string = "https://jsonplaceholder.typicode.com"
	PostsPath string = BaseURL + "/posts"
	UsersPath string = BaseURL + "/users"
)

// The common error from post
var (
	ErrPostFetch  = errors.New("Failed Fetch From `posts`" + BaseURL)
	ErrPostGet    = errors.New("Failed Get From `posts`" + BaseURL)
	ErrPostCreate = errors.New("Failed Create From `posts`" + BaseURL)
	ErrPostUpdate = errors.New("Failed Update From `posts`" + BaseURL)
	ErrPostDelete = errors.New("Failed Delete From `posts`" + BaseURL)
)
