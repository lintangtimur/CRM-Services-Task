package customers

type ICustomer interface {
	CreateCustomer(c *Customer) error
}

type UseCase struct {
	repo *Repository
}

func (u UseCase) CreateCustomer(c *Customer) error {
	return u.repo.CreateCustomer(c)
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{repo: repo}
}
