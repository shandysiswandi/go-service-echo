package get

import (
	"go-rest-echo/entity"
)

// Usecase is
func Usecase(user *entity.User, id string) error {
	err := Repository(user, id).Error
	if err != nil {
		return err
	}
	return nil
}
