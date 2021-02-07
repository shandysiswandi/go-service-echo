package jsonplaceholder

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-rest-echo/config"
	"net/http"
)

// The URL from jsonplaceholder
const (
	BaseURL   string = "https://jsonplaceholder.typicode.com"
	PostsPath string = BaseURL + "/posts"
	UsersPath string = BaseURL + "/users"
)

type (
	// Instance is
	Instance struct {
		config *config.Config
	}

	// Post is
	Post struct {
		UserID int    `json:"userId"`
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}
)

// New is
func New(c *config.Config) *Instance {
	return &Instance{c}
}

// FetchPost is
func (j *Instance) FetchPost() ([]Post, error) {
	var posts []Post

	response, err := http.Get(PostsPath)
	if err != nil {
		return nil, errors.New("Failed Fetch From `posts`" + BaseURL)
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		return nil, errors.New("Failed Fetch From `posts`" + BaseURL)
	}

	return posts, nil
}

// GetPost is
func (j *Instance) GetPost(ID int) (*Post, error) {
	post := new(Post)
	link := fmt.Sprintf("%s/%d", PostsPath, ID)

	response, err := http.Get(link)
	if err != nil {
		return nil, errors.New("Failed Get From `posts`" + BaseURL)
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(post); err != nil {
		return nil, errors.New("Failed Get From `posts`" + BaseURL)
	}

	return post, nil
}

// CreatePost is
func (j *Instance) CreatePost(data Post) (*Post, error) {
	post := new(Post)

	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("Failed Create From `posts`" + BaseURL)
	}

	response, err := http.Post(PostsPath, "application/json; charset=utf-8", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, errors.New("Failed Create From `posts`" + BaseURL)
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(post); err != nil {
		return nil, errors.New("Failed Create From `posts`" + BaseURL)
	}

	return post, nil
}

// UpdatePost is
func (j *Instance) UpdatePost(data Post, ID int) (*Post, error) {
	post := new(Post)
	link := fmt.Sprintf("%s/%d", PostsPath, ID)

	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("Failed Update From `posts`" + BaseURL)
	}

	req, err := http.NewRequest(http.MethodPut, link, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, errors.New("Failed Update From `posts`" + BaseURL)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Failed Update From `posts`" + BaseURL)
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(post); err != nil {
		return nil, errors.New("Failed Update From `posts`" + BaseURL)
	}

	return post, nil
}

// DeletePost is
func (j *Instance) DeletePost(ID int) error {
	link := fmt.Sprintf("%s/%d", PostsPath, ID)
	client := new(http.Client)

	req, err := http.NewRequest(http.MethodDelete, link, nil)
	if err != nil {
		return errors.New("Failed Delete From `posts`" + BaseURL)
	}

	response, err := client.Do(req)
	if err != nil {
		return errors.New("Failed Delete From `posts`" + BaseURL)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 || response.StatusCode == 204 {
		return nil
	}

	return errors.New("Failed Delete From `posts`" + BaseURL)
}
