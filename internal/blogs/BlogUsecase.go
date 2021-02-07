package blogs

type usecase struct {
	blogRepository Repository
}

// NewUsecase is
func NewUsecase(br Repository) Usecase {
	return &usecase{
		blogRepository: br,
	}
}

func (u *usecase) Fetch() (*[]Blog, error) {
	return u.blogRepository.Fetch()
}

func (u *usecase) Get(ID string) (*Blog, error) {
	return u.blogRepository.Get(ID)
}

func (u *usecase) Create(payload BlogPayloadCreate) error {
	b := new(Blog)
	b.SetID()
	b.Title = payload.Title
	b.Body = payload.Body

	return u.blogRepository.Create(b)
}

func (u *usecase) Update(payload BlogPayloadPut, ID string) error {
	return u.blogRepository.Update(payload, ID)
}

func (u *usecase) UpdateField(payload BlogPayloadPatch, ID string) error {
	return u.blogRepository.UpdateField(payload, ID)
}

func (u *usecase) Delete(ID string) error {
	return u.blogRepository.Delete(ID)
}
