package delete

import (
	"go-rest-echo/entity"

	"gorm.io/gorm"
)

// Usecase is
func Usecase(user *entity.User, id string) error {
	count, err := Repository(user, id)
	if err != nil {
		return err
	}

	if count == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
