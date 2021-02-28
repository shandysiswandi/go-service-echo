package users

// UserUsecase is
type UserUsecase struct {
	userRepository UserRepository
}

// NewUserUsecase is
func NewUserUsecase(r UserRepository) *UserUsecase {
	return &UserUsecase{userRepository: r}
}

// Fetch is
func (u *UserUsecase) Fetch() (Users, error) {
	return u.userRepository.Fetch()
}

// Get is
func (u *UserUsecase) Get(ID string) (*User, error) {
	return u.userRepository.Get(ID)
}

// Create is
func (u *UserUsecase) Create(eu *User) error {
	eu.SetID()
	return u.userRepository.Create(eu)
}

// Update is
func (u *UserUsecase) Update(t *User, ID string) error {
	return u.userRepository.Update(t, ID)
}

// Delete is
func (u *UserUsecase) Delete(ID string) error {
	return u.userRepository.Delete(ID)
}
