package users

type userUsecase struct {
	userRepository UserRepository
}

// NewUserUsecase is
func NewUserUsecase(ur UserRepository) UserUsecase {
	return &userUsecase{userRepository: ur}
}

func (u *userUsecase) Fetch() (Users, error) {
	return u.userRepository.Fetch()
}

func (u *userUsecase) Get(ID string) (*User, error) {
	return u.userRepository.Get(ID)
}

func (u *userUsecase) GetByEmail(email string) (*User, error) {
	return u.userRepository.GetByEmail(email)
}

func (u *userUsecase) Create(eu *UserCreatePayload) error {
	eu.SetID()
	return u.userRepository.Create(eu)
}

func (u *userUsecase) Update(t *UserUpdatePayload, ID string) error {
	return u.userRepository.Update(t, ID)
}

func (u *userUsecase) Delete(ID string) error {
	return u.userRepository.Delete(ID)
}
