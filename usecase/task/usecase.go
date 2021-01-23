package task

import "go-rest-echo/entity"

type usecase struct {
	repository Repository
}

// NewUsecase is
func NewUsecase(r Repository) Usecase {
	return &usecase{repository: r}
}

func (u *usecase) Fetch() ([]*entity.Task, error) {
	return u.repository.Fetch()
}

func (u *usecase) Get(ID string) (*entity.Task, error) {
	return u.repository.Get(ID)
}

func (u *usecase) Create(t *entity.Task) error {
	return u.repository.Create(t)
}

func (u *usecase) Update(t *entity.Task, ID string) error {
	return u.repository.Update(t, ID)
}

func (u *usecase) Delete(ID string) error {
	return u.repository.Delete(ID)
}
