package users_test

import (
	"errors"
	"go-service-echo/domain/users"
	"go-service-echo/util/tester"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_GormRepository_Fetch_Success(t *testing.T) {
	db, mock := tester.MockMysqlGormConnection()
	userRepository := users.NewGormRepository(db)
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow("uui-1", "name 1", "email_1@email.com", "password", now, now, nil).
		AddRow("uuid-2", "name 2", "email_2@email.com", "password", now, now, nil)

	query := "SELECT * FROM `users`"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	results, err := userRepository.Fetch()
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 2, len(results))

	assert.Nil(t, mock.ExpectationsWereMet())
}

func Test_GormRepository_Fetch_Error(t *testing.T) {
	db, mock := tester.MockMysqlGormConnection()
	userRepository := users.NewGormRepository(db)

	query := "SELECT * FROM `users`"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("can't fetch"))

	results, err := userRepository.Fetch()
	assert.Error(t, err)
	assert.Nil(t, results)
	assert.Equal(t, 0, len(results))

	assert.Nil(t, mock.ExpectationsWereMet())
}

func Test_GormRepository_Get_Success(t *testing.T) {
	db, mock := tester.MockMysqlGormConnection()
	userRepository := users.NewGormRepository(db)
	now := time.Now()
	user := users.User{
		ID:        "uuid",
		Name:      "name",
		Email:     "email",
		Password:  "password",
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(user.ID, user.Name, user.Email, user.Password, now, now, nil)

	query := "SELECT * FROM `users` WHERE `id` = ? ORDER BY `users`.`id` LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(user.ID).
		WillReturnRows(rows)

	result, err := userRepository.Get(user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.ID, result.ID)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func Test_GormRepository_Get_Error(t *testing.T) {
	db, mock := tester.MockMysqlGormConnection()
	userRepository := users.NewGormRepository(db)

	t.Run("Something Error", func(t *testing.T) {
		query := "SELECT * FROM `users` WHERE `id` = ? ORDER BY `users`.`id` LIMIT 1"
		mock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs("nil").
			WillReturnError(errors.New("can't get"))

		result, err := userRepository.Get("nil")
		assert.Error(t, err)
		assert.Equal(t, errors.New("can't get"), err)
		assert.Nil(t, result)

		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("Gorm Not Found", func(t *testing.T) {
		mock.ExpectQuery(`.+`).WillReturnRows(sqlmock.NewRows(nil))

		result, err := userRepository.Get("nil")
		assert.Error(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.Nil(t, result)

		assert.Nil(t, mock.ExpectationsWereMet())
	})
}

func Test_GormRepository_GetByEmail_Success(t *testing.T) {
	db, mock := tester.MockMysqlGormConnection()
	userRepository := users.NewGormRepository(db)
	now := time.Now()
	user := users.User{
		ID:        "uuid",
		Name:      "name",
		Email:     "email",
		Password:  "password",
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(user.ID, user.Name, user.Email, user.Password, now, now, nil)

	query := "SELECT * FROM `users` WHERE `email` = ? ORDER BY `users`.`id` LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(user.Email).
		WillReturnRows(rows)

	result, err := userRepository.GetByEmail(user.Email)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.Email, result.Email)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func Test_GormRepository_GetByEmail_Error(t *testing.T) {
	db, mock := tester.MockMysqlGormConnection()
	userRepository := users.NewGormRepository(db)

	t.Run("Something Error", func(t *testing.T) {
		query := "SELECT * FROM `users` WHERE `email` = ? ORDER BY `users`.`id` LIMIT 1"
		mock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs("nil").
			WillReturnError(errors.New("can't getByEmail"))

		result, err := userRepository.GetByEmail("nil")
		assert.Error(t, err)
		assert.Equal(t, errors.New("can't getByEmail"), err)
		assert.Nil(t, result)

		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("Gorm Not Found", func(t *testing.T) {
		mock.ExpectQuery(`.+`).WillReturnRows(sqlmock.NewRows(nil))

		result, err := userRepository.GetByEmail("nil")
		assert.Error(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.Nil(t, result)

		assert.Nil(t, mock.ExpectationsWereMet())
	})
}

func Test_GormRepository_Create_Success(t *testing.T) {
	db, mock := tester.MockMysqlGormConnection()
	userRepository := users.NewGormRepository(db)
	now := time.Now()
	user := users.UserCreatePayload{
		ID:        "1234-3454-678",
		Name:      "thomas name",
		Email:     "wak@gmail.com",
		Password:  "123456tghbsasji",
		CreatedAt: now,
		UpdatedAt: now,
	}

	query := "INSERT INTO `users`"
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := userRepository.Create(&user)
	assert.NoError(t, err)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func Test_GormRepository_Create_Error(t *testing.T) {
	db, mock := tester.MockMysqlGormConnection()
	userRepository := users.NewGormRepository(db)
	now := time.Now()
	user := users.UserCreatePayload{
		ID:        "1234-3454-678",
		Name:      "thomas name",
		Email:     "wak@gmail.com",
		Password:  "123456tghbsasji",
		CreatedAt: now,
		UpdatedAt: now,
	}

	query := "INSERT INTO `users`"
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).
		WillReturnError(errors.New("can't create"))

	err := userRepository.Create(&user)
	assert.Error(t, err)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func Test_GormRepository_Update_Success(t *testing.T) {
	db, mock := tester.MockMysqlGormConnection()
	userRepository := users.NewGormRepository(db)
	user := users.UserUpdatePayload{
		Name:     "thomas name update",
		Email:    "wak-update@gmail.com",
		Password: "123456",
	}

	query := "UPDATE `users`"
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(user.Name, user.Email, user.Password, time.Now(), "12345").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := userRepository.Update(&user, "12345")
	assert.NoError(t, err)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func Test_GormRepository_Update_Error(t *testing.T) {
	//
}

func Test_GormRepository_Delete_Success(t *testing.T) {
	//
}

func Test_GormRepository_Delete_Error(t *testing.T) {
	//
}
