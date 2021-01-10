package fetch

import (
	"go-rest-echo/entity"
)

// Usecase is
func Usecase(t *[]entity.Task) (err error) {
	err = Repository(t).Error
	if err != nil {
		return err
	}
	return nil
}
