package create

import (
	"go-rest-echo/entity"
)

// Usecase is
func Usecase(u *entity.User) error {
	err := Repository(u).Error
	if err != nil {
		return err
	}
	return nil
}
