package delete

import (
	"go-rest-echo/database/mysql"
	"go-rest-echo/entity"
)

// Repository is
func Repository(user *entity.User, id string) (int64, error) {
	result := mysql.DB.Where("id = ?", id).Delete(user)
	return result.RowsAffected, result.Error
}
