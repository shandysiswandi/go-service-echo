package posts

import (
	"encoding/json"
	"fmt"
	"go-rest-echo/external/jsonplaceholder"
	"net/http"
)

// Get is
func Get(ID int) (*Post, error) {
	post := new(Post)
	link := fmt.Sprintf("%s/%d", jsonplaceholder.PostsPath, ID)

	response, err := http.Get(link)
	if err != nil {
		return nil, jsonplaceholder.ErrPostFetch
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(post); err != nil {
		return nil, jsonplaceholder.ErrPostFetch
	}

	return post, nil
}
