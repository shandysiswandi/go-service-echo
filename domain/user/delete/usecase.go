package delete

import (
	"go-rest-echo/entity"

	"gorm.io/gorm"
)

// Usecase is
func Usecase(u *entity.User, id string) error {
	count, err := Repository(u, id)
	if err != nil {
		return err
	}

	if count == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
