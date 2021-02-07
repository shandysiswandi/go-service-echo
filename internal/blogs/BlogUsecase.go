package blogs

type usecase struct {
	blogRepository Repository
}

// NewUsecase is
func NewUsecase(br Repository) Usecase {
	return &usecase{br}
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
	b := new(Blog)
	b.Title = payload.Title
	b.Body = payload.Body

	return u.blogRepository.Update(b, ID)
}

func (u *usecase) UpdateField(payload BlogPayloadPatch, ID string) error {
	b := new(Blog)

	if payload.Title == "" && payload.Body == "" {
		return nil
	}

	if payload.Title != "" {
		b.Title = payload.Title
	}
	if payload.Body != "" {
		b.Body = payload.Body
	}

	return u.blogRepository.UpdateField(b, ID)
}

func (u *usecase) Delete(ID string) error {
	return u.blogRepository.Delete(ID)
}
