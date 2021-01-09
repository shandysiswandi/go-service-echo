package delete

import (
	"go-rest-echo/domain/user/get"
	"go-rest-echo/entity"
)

// Usecase is
func Usecase(user *entity.User, id string) (err error) {
	err = get.Repository(user, id).Error
	if err != nil {
		return err
	}

	err = Repository(user, id).Error
	if err != nil {
		return err
	}
	return nil
}
