package fetch

import (
	"go-rest-echo/database/mysql"
	"go-rest-echo/entity"

	"gorm.io/gorm"
)

// Repository is
func Repository(t *[]entity.Task) *gorm.DB {
	return mysql.DB.Find(t)
}
