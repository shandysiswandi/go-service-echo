package token

import (
	"errors"
	"go-service-echo/config"
	"go-service-echo/config/constant"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

// all errors token
var (
	ErrConfigTokenNil = errors.New("config token is nil")
	ErrKeyLengthToken = errors.New("bad key length")
	ErrTokenType      = errors.New("token type not support")
	ErrInvalidToken   = errors.New("token is invalid")
	ErrExpiredToken   = errors.New("token has expired")
)

type (
	// Token is
	Token interface {
		NewAccessToken(PayloadData, time.Duration) (string, error)
		NewRefreshToken(PayloadData, time.Duration) (string, error)
		VerifyAccessToken(string) (*Payload, error)
		VerifyRefreshToken(string) (*Payload, error)
	}

	// PayloadData is
	PayloadData struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}

	// Payload is
	Payload struct {
		ID        uuid.UUID   `json:"id"`
		Data      PayloadData `json:"data"`
		IssuedAt  time.Time   `json:"issued_at"`
		ExpiredAt time.Time   `json:"expired_at"`
	}

	token struct {
		tokenType  string
		accessKey  string
		refreshKey string
	}
)

// New is
func New(c *config.TokenConfig) (Token, error) {
	if c == nil {
		return nil, ErrConfigTokenNil
	}

	if len(c.AccessKey) != 32 || len(c.RefreshKey) != 32 {
		return nil, ErrKeyLengthToken
	}

	if c.TokenType == constant.JWT || c.TokenType == constant.Paseto {
		return &token{c.TokenType, c.AccessKey, c.RefreshKey}, nil
	}

	return nil, ErrTokenType
}

func (token *token) NewAccessToken(pd PayloadData, exp time.Duration) (string, error) {
	return token.generate(pd, exp, token.accessKey)
}

func (token *token) NewRefreshToken(pd PayloadData, exp time.Duration) (string, error) {
	return token.generate(pd, exp, token.refreshKey)
}

func (token *token) generate(pd PayloadData, exp time.Duration, key string) (string, error) {
	body := NewPayload(pd, exp)

	if token.tokenType == constant.JWT {
		return jwt.NewWithClaims(jwt.SigningMethodHS256, body).SignedString([]byte(key))
	}

	return paseto.NewV2().Encrypt([]byte(key), body, nil)
}

func (token *token) VerifyAccessToken(t string) (*Payload, error) {
	return token.verify(t, token.accessKey)
}

func (token *token) VerifyRefreshToken(t string) (*Payload, error) {
	return token.verify(t, token.refreshKey)
}

func (token *token) verify(t, k string) (*Payload, error) {
	payload := &Payload{}

	if token.tokenType == constant.JWT {
		keyFunc := func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrInvalidToken
			}

			return []byte(k), nil
		}

		jwtToken, err := jwt.ParseWithClaims(t, payload, keyFunc)
		if err != nil {
			verr, ok := err.(*jwt.ValidationError)
			if ok && errors.Is(verr.Inner, ErrExpiredToken) {
				return nil, ErrExpiredToken
			}

			return nil, ErrInvalidToken
		}

		payload, ok := jwtToken.Claims.(*Payload)
		if !ok {
			return nil, ErrInvalidToken
		}

		return payload, nil
	}

	if err := paseto.NewV2().Decrypt(t, []byte(k), payload, nil); err != nil {
		return nil, ErrInvalidToken
	}

	if err := payload.Valid(); err != nil {
		return nil, err
	}

	return payload, nil
}

// NewPayload is
func NewPayload(pd PayloadData, exp time.Duration) *Payload {
	return &Payload{
		ID:        uuid.New(),
		IssuedAt:  time.Now(),
		Data:      pd,
		ExpiredAt: time.Now().Add(exp),
	}
}

// Valid is
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
