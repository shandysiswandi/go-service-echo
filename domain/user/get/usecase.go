package get

import (
	"go-rest-echo/entity"
)

// Usecase is
func Usecase(u *entity.User, id string) error {
	err := Repository(u, id).Error
	if err != nil {
		return err
	}
	return nil
}
