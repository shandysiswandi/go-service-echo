package get

import (
	"go-rest-echo/entity"
)

// Usecase is
func Usecase(t *entity.Task, id string) (err error) {
	err = Repository(t, id).Error
	if err != nil {
		return err
	}
	return nil
}
