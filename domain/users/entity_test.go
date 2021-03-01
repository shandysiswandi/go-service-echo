package users_test

import (
	"go-service-echo/internal/users"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTableName(t *testing.T) {
	u := &users.User{}
	assert.Equal(t, "users", u.TableName())
}

func TestSetID(t *testing.T) {
	u := &users.User{}
	assert.Equal(t, "", u.ID)
	u.SetID()
	assert.NotEqual(t, "", u.ID)
}

func TestTransform(t *testing.T) {
	us := users.Users{}
	assert.Equal(t, 0, len(us))

	now := time.Now()

	us = append(us, &users.User{
		ID:        "1",
		Name:      "2",
		Email:     "3",
		Password:  "4",
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	})
	assert.Equal(t, 1, len(us))

	temp := []users.UserResponse{{
		ID:        "1",
		Name:      "2",
		Email:     "3",
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	}}
	ust := us.Transform()
	assert.Equal(t, temp, ust)
}
