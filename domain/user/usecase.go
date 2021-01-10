package user

import (
	"go-rest-echo/entity"

	"gorm.io/gorm"
)

type usecase struct {
	repository Repository
}

// Usecase is
type Usecase interface {
	Fetch(*[]entity.User) error
	Get(*entity.User, string) error
	Create(*entity.User) error
	Update(*entity.User, string) error
	Delete(*entity.User, string) error
}

// NewUsecase is
func NewUsecase() Usecase {
	return &usecase{repository: NewRepository()}
}

func (uc *usecase) Fetch(u *[]entity.User) (err error) {
	err = uc.repository.Fetch(u).Error
	if err != nil {
		return err
	}
	return nil
}

func (uc *usecase) Get(u *entity.User, id string) (err error) {
	err = uc.repository.Get(u, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (uc *usecase) Create(u *entity.User) error {
	err := uc.repository.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

func (uc *usecase) Update(u *entity.User, id string) error {
	repo := uc.repository.Update(u, id)
	if repo.Error != nil {
		return repo.Error
	}

	if repo.RowsAffected != 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (uc *usecase) Delete(u *entity.User, id string) error {
	repo := uc.repository.Delete(u, id)
	if repo.Error != nil {
		return repo.Error
	}

	if repo.RowsAffected != 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
