package customers

type ICustomer interface {
	CreateCustomer(c *Customer) error
	DeleteCustomer(c *Customer, d *DeleteCustomerRequest) error
}

type UseCase struct {
	repo *Repository
}

func (u UseCase) CreateCustomer(c *Customer) error {
	return u.repo.CreateCustomer(c)
}

func (u UseCase) DeleteCustomer(c *Customer, d *DeleteCustomerRequest) error {
	return u.repo.DeleteCustomer(c, d)
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{repo: repo}
}
