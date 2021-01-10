package delete

import (
	"go-rest-echo/entity"

	"gorm.io/gorm"
)

// Usecase is
func Usecase(t *entity.Task, id string) error {
	count, err := Repository(t, id)
	if err != nil {
		return err
	}

	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
