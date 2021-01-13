package task

import (
	"go-rest-echo/entity"
)

type usecase struct {
	repository Repository
}

// Usecase is
type Usecase interface {
	Fetch(*[]entity.Task) error
	Get(*entity.Task, string) error
	Create(*entity.Task) error
	Update(*entity.Task, string) error
	Delete(*entity.Task, string) error
}

// NewUsecase is
func NewUsecase() Usecase {
	return &usecase{repository: NewRepository()}
}

func (uc *usecase) Fetch(u *[]entity.Task) (err error) {
	err = uc.repository.Fetch(u)
	if err != nil {
		return err
	}

	return nil
}

func (uc *usecase) Get(u *entity.Task, id string) (err error) {
	err = uc.repository.Get(u, id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *usecase) Create(u *entity.Task) (err error) {
	err = uc.repository.Create(u)
	if err != nil {
		return err
	}

	return nil
}

func (uc *usecase) Update(u *entity.Task, id string) (err error) {
	err = uc.repository.Update(u, id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *usecase) Delete(u *entity.Task, id string) (err error) {
	err = uc.repository.Delete(u, id)
	if err != nil {
		return err
	}

	return nil
}
