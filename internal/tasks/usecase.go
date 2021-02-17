package tasks

type usecase struct {
	taskRepository TaskRepository
}

// NewUsecase is
func NewUsecase(r TaskRepository) TaskUsecase {
	return &usecase{taskRepository: r}
}

func (u *usecase) Fetch() (*[]Task, error) {
	return u.taskRepository.Fetch()
}

func (u *usecase) Get(ID string) (*Task, error) {
	return u.taskRepository.Get(ID)
}

func (u *usecase) Create(t *Task) (*string, error) {
	return u.taskRepository.Create(t)
}

func (u *usecase) Update(t *Task, ID string) error {
	return u.taskRepository.Update(t, ID)
}

func (u *usecase) Delete(ID string) error {
	return u.taskRepository.Delete(ID)
}
