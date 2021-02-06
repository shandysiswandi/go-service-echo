package posts

import (
	"fmt"
	"go-rest-echo/external/jsonplaceholder"
	"net/http"
)

// Delete is
func Delete(ID int) error {
	link := fmt.Sprintf("%s/%d", jsonplaceholder.PostsPath, ID)
	client := new(http.Client)

	req, err := http.NewRequest(http.MethodDelete, link, nil)
	if err != nil {
		return jsonplaceholder.ErrPostDelete
	}

	response, err := client.Do(req)
	if err != nil {
		return jsonplaceholder.ErrPostDelete
	}
	defer response.Body.Close()

	if response.StatusCode == 200 || response.StatusCode == 204 {
		return nil
	}

	return jsonplaceholder.ErrPostDelete
}
