package create

import (
	"go-rest-echo/entity"
)

// Usecase is
func Usecase(user *entity.User) error {
	err := Repository(user).Error
	if err != nil {
		return err
	}
	return nil
}
