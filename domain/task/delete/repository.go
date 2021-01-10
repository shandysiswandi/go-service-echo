package delete

import (
	"go-rest-echo/database/mysql"
	"go-rest-echo/entity"
)

// Repository is
func Repository(t *entity.Task, id string) (int64, error) {
	result := mysql.DB.Where("id = ?", id).Delete(t)
	return result.RowsAffected, result.Error
}
