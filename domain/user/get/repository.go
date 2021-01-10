package get

import (
	"go-rest-echo/database/mysql"
	"go-rest-echo/entity"

	"gorm.io/gorm"
)

// Repository is
func Repository(u *entity.User, id string) *gorm.DB {
	return mysql.DB.Where("id = ?", id).First(u)
}
