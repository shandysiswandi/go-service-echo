package users

type usecase struct {
	userRepository UserRepository
}

// NewUsecase is
func NewUsecase(r UserRepository) UserUsecase {
	return &usecase{userRepository: r}
}

func (u *usecase) Fetch() (*[]User, error) {
	return u.userRepository.Fetch()
}

func (u *usecase) Get(ID string) (*User, error) {
	return u.userRepository.Get(ID)
}

func (u *usecase) Create(eu *User) error {
	eu.SetID()
	return u.userRepository.Create(eu)
}

func (u *usecase) Update(t *User, ID string) error {
	return u.userRepository.Update(t, ID)
}

func (u *usecase) Delete(ID string) error {
	return u.userRepository.Delete(ID)
}
