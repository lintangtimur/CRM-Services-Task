package accounts

type IAccount interface {
	Login(a *Actor, lr *LoginRequest) error
	CreateAdmin(a *Actor) error
}
type UseCase struct {
	repo *Repository
}

func (u UseCase) Login(a *Actor, lr *LoginRequest) error {
	return u.repo.Login(a, lr)
}

func (u UseCase) CreateAdmin(a *Actor) error {
	return u.repo.CreateAdmin(a)
}

func (u UseCase) GetApprovalList() ([]Actor, error) {
	return u.repo.FindAllApproval()
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{repo: repo}
}
