package create

import (
	"go-rest-echo/domain/user/get"
	"go-rest-echo/entity"
)

// Usecase is
func Usecase(task *entity.Task, uID string) (err error) {
	// check user
	var user entity.User
	err = get.Repository(&user, uID).Error
	if err != nil {
		return err
	}

	err = Repository(task).Error
	if err != nil {
		return err
	}
	return nil
}
