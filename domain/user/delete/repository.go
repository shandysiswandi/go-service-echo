package delete

import (
	"go-rest-echo/database/mysql"
	"go-rest-echo/entity"
)

// Repository is
func Repository(u *entity.User, id string) (int64, error) {
	result := mysql.DB.Where("id = ?", id).Delete(u)
	return result.RowsAffected, result.Error
}
