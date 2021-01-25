package posts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-rest-echo/external/jsonplaceholder"
	"net/http"
)

// Update is
func Update(data Post, ID int) (*Post, error) {
	post := new(Post)
	link := fmt.Sprintf("%s/%d", jsonplaceholder.PostsPath, ID)

	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, jsonplaceholder.ErrPostUpdate
	}

	req, err := http.NewRequest(http.MethodPut, link, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, jsonplaceholder.ErrPostUpdate
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return nil, jsonplaceholder.ErrPostUpdate
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(post); err != nil {
		return nil, jsonplaceholder.ErrPostUpdate
	}

	return post, nil
}
