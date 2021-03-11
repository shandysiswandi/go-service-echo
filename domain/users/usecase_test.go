package users_test

import (
	"errors"
	"go-service-echo/domain/users"
	"go-service-echo/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserUsecase_Fetch_Sucess(t *testing.T) {
	now := time.Now()
	expected := users.Users{{
		ID:        "1",
		Name:      "name",
		Email:     "email",
		Password:  "password",
		CreatedAt: now,
		UpdatedAt: now,
	}}

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Fetch").Return(expected, nil)

	userUsecase := users.NewUserUsecase(mockUserRepo)
	us, err := userUsecase.Fetch()

	assert.NoError(t, err)
	assert.Equal(t, expected, us)
	mockUserRepo.AssertExpectations(t)
}

func TestUserUsecase_Fetch_Error(t *testing.T) {
	expected := errors.New("can't Fetch")

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Fetch").Return(nil, expected)

	userUsecase := users.NewUserUsecase(mockUserRepo)
	us, err := userUsecase.Fetch()

	assert.NotNil(t, err)
	assert.Equal(t, expected, err)
	assert.Nil(t, us)
	mockUserRepo.AssertExpectations(t)
}

func TestUserUsecase_Get_Success(t *testing.T) {
	now := time.Now()
	expected := &users.User{
		ID:        "1",
		Name:      "name",
		Email:     "email",
		Password:  "password",
		CreatedAt: now,
		UpdatedAt: now,
	}

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Get", "1").Return(expected, nil)

	userUsecase := users.NewUserUsecase(mockUserRepo)
	us, err := userUsecase.Get("1")

	assert.NoError(t, err)
	assert.Equal(t, expected, us)
	mockUserRepo.AssertExpectations(t)
}

func TestUserUsecase_Get_Error(t *testing.T) {
	expected := errors.New("can't Get")

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Get", "1").Return(nil, expected)

	userUsecase := users.NewUserUsecase(mockUserRepo)
	us, err := userUsecase.Get("1")

	assert.NotNil(t, err)
	assert.Equal(t, expected, err)
	assert.Nil(t, us)
	mockUserRepo.AssertExpectations(t)
}

func TestUserUsecase_GetByEmail_Success(t *testing.T) {
	now := time.Now()
	expected := &users.User{
		ID:        "1",
		Name:      "name",
		Email:     "email",
		Password:  "password",
		CreatedAt: now,
		UpdatedAt: now,
	}

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("GetByEmail", "email").Return(expected, nil)

	userUsecase := users.NewUserUsecase(mockUserRepo)
	us, err := userUsecase.GetByEmail("email")

	assert.NoError(t, err)
	assert.Equal(t, expected, us)
	mockUserRepo.AssertExpectations(t)
}

func TestUserUsecase_GetByEmail_Error(t *testing.T) {
	expected := errors.New("can't GetByEmail")

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("GetByEmail", "email").Return(nil, expected)

	userUsecase := users.NewUserUsecase(mockUserRepo)
	us, err := userUsecase.GetByEmail("email")

	assert.NotNil(t, err)
	assert.Equal(t, expected, err)
	assert.Nil(t, us)
	mockUserRepo.AssertExpectations(t)
}

func TestUserUsecase_Create_Success(t *testing.T) {
	now := time.Now()
	payload := &users.UserCreatePayload{
		ID:        "1",
		Name:      "name",
		Email:     "email",
		Password:  "password",
		CreatedAt: now,
		UpdatedAt: now,
	}

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Create", payload).Return(nil)

	userUsecase := users.NewUserUsecase(mockUserRepo)
	err := userUsecase.Create(payload)

	assert.NoError(t, err)
	assert.Equal(t, nil, err)
	mockUserRepo.AssertExpectations(t)
}

func TestUserUsecase_Create_Error(t *testing.T) {
	expected := errors.New("can't Create")
	now := time.Now()
	payload := &users.UserCreatePayload{
		ID:        "1",
		Name:      "name",
		Email:     "email",
		Password:  "password",
		CreatedAt: now,
		UpdatedAt: now,
	}

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Create", payload).Return(expected)

	userUsecase := users.NewUserUsecase(mockUserRepo)
	err := userUsecase.Create(payload)

	assert.NotNil(t, err)
	assert.Equal(t, expected, err)
	mockUserRepo.AssertExpectations(t)
}

func TestUserUsecase_Update_Success(t *testing.T) {
	now := time.Now()
	payload := &users.UserUpdatePayload{
		Name:      "name",
		Email:     "email",
		Password:  "password",
		UpdatedAt: now,
	}

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Update", payload, "id").Return(nil)

	userUsecase := users.NewUserUsecase(mockUserRepo)
	err := userUsecase.Update(payload, "id")

	assert.NoError(t, err)
	assert.Equal(t, nil, err)
	mockUserRepo.AssertExpectations(t)
}

func TestUserUsecase_Update_Error(t *testing.T) {
	expected := errors.New("can't Update")
	now := time.Now()
	payload := &users.UserUpdatePayload{
		Name:      "name",
		Email:     "email",
		Password:  "password",
		UpdatedAt: now,
	}

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Update", payload, "id").Return(expected)

	userUsecase := users.NewUserUsecase(mockUserRepo)
	err := userUsecase.Update(payload, "id")

	assert.NotNil(t, err)
	assert.Equal(t, expected, err)
	mockUserRepo.AssertExpectations(t)
}

func TestUserUsecase_Delete_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Delete", "id").Return(nil)

	userUsecase := users.NewUserUsecase(mockUserRepo)
	err := userUsecase.Delete("id")

	assert.NoError(t, err)
	assert.Equal(t, nil, err)
	mockUserRepo.AssertExpectations(t)
}

func TestUserUsecase_Delete_Error(t *testing.T) {
	expected := errors.New("can't Update")

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Delete", "id").Return(expected)

	userUsecase := users.NewUserUsecase(mockUserRepo)
	err := userUsecase.Delete("id")

	assert.NotNil(t, err)
	assert.Equal(t, expected, err)
	mockUserRepo.AssertExpectations(t)
}
