package token_test

import (
	"go-service-echo/app/library/token"
	"go-service-echo/config"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNew_Config_Key_Length(t *testing.T) {
	os.Setenv("TOKEN_ACCESS_KEY", "access")
	os.Setenv("TOKEN_REFRESH_KEY", "refresh")

	theToken, err := token.New(config.New().Token)
	assert.Equal(t, token.ErrKeyLengthToken, err)
	assert.Nil(t, theToken)
}

func TestNew_Config_Token_Type_Not_Support(t *testing.T) {
	os.Setenv("TOKEN_TYPE", "none")
	os.Setenv("TOKEN_ACCESS_KEY", "eyJhbGciOiJub25lIiwidHlwIjoiSldU")
	os.Setenv("TOKEN_REFRESH_KEY", "48PX2KNDvmoPnF4BASrszNkwvw0l3Mgh")

	theToken, err := token.New(config.New().Token)
	assert.Equal(t, token.ErrTokenType, err)
	assert.Nil(t, theToken)
}

func TestNew_GetTokenType(t *testing.T) {
	os.Setenv("TOKEN_TYPE", "jwt")
	os.Setenv("TOKEN_ACCESS_KEY", "eyJhbGciOiJub25lIiwidHlwIjoiSldU")
	os.Setenv("TOKEN_REFRESH_KEY", "48PX2KNDvmoPnF4BASrszNkwvw0l3Mgh")

	theToken, err := token.New(config.New().Token)
	assert.Nil(t, err)
	assert.NotNil(t, theToken)
	assert.Equal(t, "jwt", theToken.GetTokenType())
}

func TestNew_Token_JWT(t *testing.T) {
	is := assert.New(t)
	exp := time.Minute
	pd := token.PayloadData{ID: "1", Email: "2"}

	os.Setenv("TOKEN_TYPE", "jwt")
	os.Setenv("TOKEN_ACCESS_KEY", "eyJhbGciOiJub25lIiwidHlwIjoiSldU")
	os.Setenv("TOKEN_REFRESH_KEY", "48PX2KNDvmoPnF4BASrszNkwvw0l3Mgh")

	/* test new token jwt */
	theToken, err := token.New(config.New().Token)
	is.Nil(err)
	is.NotNil(theToken)

	/* test new access token jwt */
	accJWT, err := theToken.NewAccessToken(pd, exp)
	is.Nil(err)
	is.NotEmpty(accJWT)

	/* test verify access token jwt */
	accVeri, err := theToken.VerifyAccessToken(accJWT)
	is.Nil(err)
	is.NotNil(accVeri)
	is.Equal("1", accVeri.Data.ID)
	is.Equal("2", accVeri.Data.Email)
	is.NotZero(accVeri.ID)
	is.NotEqual(uuid.New(), accVeri.ID)
	is.WithinDuration(time.Now(), accVeri.IssuedAt, time.Second)
	is.WithinDuration(time.Now().Add(exp), accVeri.ExpiredAt, time.Second)

	/* test new refresh token jwt */
	refJWT, err := theToken.NewRefreshToken(pd, exp)
	is.Nil(err)
	is.NotEmpty(refJWT)

	/* test verify access token jwt */
	refVeri, err := theToken.VerifyRefreshToken(refJWT)
	is.Nil(err)
	is.NotNil(refVeri)
	is.Equal("1", refVeri.Data.ID)
	is.Equal("2", refVeri.Data.Email)
	is.NotZero(refVeri.ID)
	is.NotEqual(uuid.New(), refVeri.ID)
	is.WithinDuration(time.Now(), refVeri.IssuedAt, time.Second)
	is.WithinDuration(time.Now().Add(exp), refVeri.ExpiredAt, time.Second)
}

func TestNew_Token_Paseto(t *testing.T) {
	is := assert.New(t)
	exp := time.Minute
	pd := token.PayloadData{ID: "1", Email: "2"}

	os.Setenv("TOKEN_TYPE", "paseto")
	os.Setenv("TOKEN_ACCESS_KEY", "eyJhbGciOiJub25lIiwidHlwIjoiSldU")
	os.Setenv("TOKEN_REFRESH_KEY", "48PX2KNDvmoPnF4BASrszNkwvw0l3Mgh")

	/* test new token jwt */
	theToken, err := token.New(config.New().Token)
	is.Nil(err)
	is.NotNil(theToken)

	/* test new token paseto */
	theToken, err = token.New(config.New().Token)
	is.Nil(err)
	is.NotNil(theToken)

	/* test new access token paseto */
	accPaseto, err := theToken.NewAccessToken(pd, exp)
	is.Nil(err)
	is.NotEmpty(accPaseto)

	/* test verify access token paseto */
	accVeriPaseto, err := theToken.VerifyAccessToken(accPaseto)
	is.Nil(err)
	is.NotNil(accVeriPaseto)
	is.Equal("1", accVeriPaseto.Data.ID)
	is.Equal("2", accVeriPaseto.Data.Email)
	is.NotZero(accVeriPaseto.ID)
	is.NotEqual(uuid.New(), accVeriPaseto.ID)
	is.WithinDuration(time.Now(), accVeriPaseto.IssuedAt, time.Second)
	is.WithinDuration(time.Now().Add(exp), accVeriPaseto.ExpiredAt, time.Second)

	/* test new refresh token paseto */
	refPaseto, err := theToken.NewAccessToken(pd, exp)
	is.Nil(err)
	is.NotEmpty(refPaseto)

	/* test verify refresh token paseto */
	refVeriPaseto, err := theToken.VerifyAccessToken(refPaseto)
	is.Nil(err)
	is.NotNil(refVeriPaseto)
	is.Equal("1", refVeriPaseto.Data.ID)
	is.Equal("2", refVeriPaseto.Data.Email)
	is.NotZero(refVeriPaseto.ID)
	is.NotEqual(uuid.New(), refVeriPaseto.ID)
	is.WithinDuration(time.Now(), refVeriPaseto.IssuedAt, time.Second)
	is.WithinDuration(time.Now().Add(exp), refVeriPaseto.ExpiredAt, time.Second)
}

func TestNew_Expired_Token_Verify_JWT(t *testing.T) {
	is := assert.New(t)
	exp := -time.Hour
	pd := token.PayloadData{ID: "1", Email: "2"}

	os.Setenv("TOKEN_TYPE", "jwt")
	os.Setenv("TOKEN_ACCESS_KEY", "eyJhbGciOiJub25lIiwidHlwIjoiSldU")
	os.Setenv("TOKEN_REFRESH_KEY", "48PX2KNDvmoPnF4BASrszNkwvw0l3Mgh")

	/* test new token jwt */
	theToken, err := token.New(config.New().Token)
	is.Nil(err)
	is.NotNil(theToken)

	/* test new access token jwt */
	accJWT, err := theToken.NewAccessToken(pd, exp)
	is.Nil(err)
	is.NotEmpty(accJWT)

	/* test verify access token jwt */
	accVeri, err := theToken.VerifyAccessToken(accJWT)
	is.Error(err)
	is.Equal(token.ErrExpiredToken, err)
	is.Nil(accVeri)

	/* test new refresh token jwt */
	refJWT, err := theToken.NewRefreshToken(pd, exp)
	is.Nil(err)
	is.NotEmpty(refJWT)

	/* test verify refresh token jwt */
	refVeri, err := theToken.VerifyRefreshToken(refJWT)
	is.Error(err)
	is.Equal(token.ErrExpiredToken, err)
	is.Nil(refVeri)
}

func TestNew_Invalid_Token_Verify_JWT(t *testing.T) {
	is := assert.New(t)
	tok := "eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJpZCI6IjZmY2I5YmY4LTI1NjEtNGE5ZC04ZjNmLThmOGNhY2E1NDcxOSIsImRhdGEiOnsiaWQiOiIxIiwiZW1haWwiOiIyIn0sImlzc3VlZF9hdCI6IjIwMjEtMDItMjdUMTE6MjQ6MTUuODM0MTY0MTUzKzA3OjAwIiwiZXhwaXJlZF9hdCI6IjIwMjEtMDItMjdUMTE6MjU6MTUuODM0MTY0MjI0KzA3OjAwIn0.48PX2KNDvmoPnF4BASrszNkwvw0l3MghKTS48bOQrK8"

	os.Setenv("TOKEN_TYPE", "jwt")
	os.Setenv("TOKEN_ACCESS_KEY", "eyJhbGciOiJub25lIiwidHlwIjoiSldU")
	os.Setenv("TOKEN_REFRESH_KEY", "48PX2KNDvmoPnF4BASrszNkwvw0l3Mgh")

	/* test new token jwt */
	theToken, err := token.New(config.New().Token)
	is.Nil(err)
	is.NotNil(theToken)

	/* test verify access token jwt */
	accVeri, err := theToken.VerifyAccessToken(tok)
	is.Error(err)
	is.Equal(token.ErrInvalidToken, err)
	is.Nil(accVeri)

	/* test verify refresh token jwt */
	refVeri, err := theToken.VerifyRefreshToken(tok)
	is.Error(err)
	is.Equal(token.ErrInvalidToken, err)
	is.Nil(refVeri)
}

func TestNew_Expired_Token_Verify_Paseto(t *testing.T) {
	is := assert.New(t)
	exp := -time.Hour
	pd := token.PayloadData{ID: "1", Email: "2"}

	os.Setenv("TOKEN_TYPE", "paseto")
	os.Setenv("TOKEN_ACCESS_KEY", "eyJhbGciOiJub25lIiwidHlwIjoiSldU")
	os.Setenv("TOKEN_REFRESH_KEY", "48PX2KNDvmoPnF4BASrszNkwvw0l3Mgh")

	/* test new token paseto */
	theToken, err := token.New(config.New().Token)
	is.Nil(err)
	is.NotNil(theToken)

	/* test new access token paseto */
	accJWT, err := theToken.NewAccessToken(pd, exp)
	is.Nil(err)
	is.NotEmpty(accJWT)

	/* test verify access token paseto */
	accVeri, err := theToken.VerifyAccessToken(accJWT)
	is.Error(err)
	is.Equal(token.ErrExpiredToken, err)
	is.Nil(accVeri)

	/* test new refresh token paseto */
	refJWT, err := theToken.NewRefreshToken(pd, exp)
	is.Nil(err)
	is.NotEmpty(refJWT)

	/* test verify refresh token paseto */
	refVeri, err := theToken.VerifyRefreshToken(refJWT)
	is.Error(err)
	is.Equal(token.ErrExpiredToken, err)
	is.Nil(refVeri)
}

func TestNew_Invalid_Token_Verify_Paseto(t *testing.T) {
	is := assert.New(t)
	tok := "eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJpZCI6IjZmY2I5YmY4LTI1NjEtNGE5ZC04ZjNmLThmOGNhY2E1NDcxOSIsImRhdGEiOnsiaWQiOiIxIiwiZW1haWwiOiIyIn0sImlzc3VlZF9hdCI6IjIwMjEtMDItMjdUMTE6MjQ6MTUuODM0MTY0MTUzKzA3OjAwIiwiZXhwaXJlZF9hdCI6IjIwMjEtMDItMjdUMTE6MjU6MTUuODM0MTY0MjI0KzA3OjAwIn0.48PX2KNDvmoPnF4BASrszNkwvw0l3MghKTS48bOQrK8"

	os.Setenv("TOKEN_TYPE", "paseto")
	os.Setenv("TOKEN_ACCESS_KEY", "eyJhbGciOiJub25lIiwidHlwIjoiSldU")
	os.Setenv("TOKEN_REFRESH_KEY", "48PX2KNDvmoPnF4BASrszNkwvw0l3Mgh")

	/* test new token paseto */
	theToken, err := token.New(config.New().Token)
	is.Nil(err)
	is.NotNil(theToken)

	/* test verify access token paseto */
	accVeri, err := theToken.VerifyAccessToken(tok)
	is.Error(err)
	is.Equal(token.ErrInvalidToken, err)
	is.Nil(accVeri)

	/* test verify refresh token paseto */
	refVeri, err := theToken.VerifyRefreshToken(tok)
	is.Error(err)
	is.Equal(token.ErrInvalidToken, err)
	is.Nil(refVeri)
}
