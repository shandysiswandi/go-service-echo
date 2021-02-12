package jsonplaceholder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-rest-echo/config"
	"log"
	"net/http"
)

type (
	// JSONPlaceHolder is
	JSONPlaceHolder struct {
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
func New(c *config.Config) *JSONPlaceHolder {
	return &JSONPlaceHolder{c}
}

// FetchPost is
func (j *JSONPlaceHolder) FetchPost() ([]Post, error) {
	var posts []Post

	response, err := http.Get(j.config.External.JsonplaceholderURL + "/posts")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		log.Println("2", err)
		return nil, err
	}

	return posts, nil
}

// GetPost is
func (j *JSONPlaceHolder) GetPost(ID int) (*Post, error) {
	var post Post

	response, err := http.Get(fmt.Sprintf("%s/posts/%d", j.config.External.JsonplaceholderURL, ID))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&post); err != nil {
		return nil, err
	}

	return &post, nil
}

// CreatePost is
func (j *JSONPlaceHolder) CreatePost(data Post) (*Post, error) {
	var post Post

	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(j.config.External.JsonplaceholderURL+"/posts", "application/json; charset=utf-8", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&post); err != nil {
		return nil, err
	}

	return &post, nil
}

// UpdatePost is
func (j *JSONPlaceHolder) UpdatePost(data Post, ID int) (*Post, error) {
	var post Post

	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/posts/%d", j.config.External.JsonplaceholderURL, ID), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&post); err != nil {
		return nil, err
	}

	return &post, nil
}

// DeletePost is
func (j *JSONPlaceHolder) DeletePost(ID int) error {
	var client = http.Client{}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/posts/%d", j.config.External.JsonplaceholderURL, ID), nil)
	if err != nil {
		return err
	}

	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode == 200 || response.StatusCode == 204 {
		return nil
	}

	return err
}
