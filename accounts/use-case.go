package accounts

type IAccount interface {
	Login(a *Actor, lr *LoginRequest) error
	CreateAdmin(a *Actor) error
	GetApprovalList() ([]Actor, error)
	ApproveAdmin(a *Actor, ar *ApproveRequest, ra *RegisterApproval, update map[string]interface{}, ra2 map[string]interface{}) error
	ActivateAdmin(a *Actor, aar *ActivateAdminRequest, activeTrue map[string]interface{}) error
	DeleteAdmin(a *Actor, d *DeleteAdminRequest) error
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

func (u UseCase) ApproveAdmin(a *Actor, ar *ApproveRequest, ra *RegisterApproval, update map[string]interface{}, ra2 map[string]interface{}) error {
	err := u.repo.UpdateActor(a, ar, update)
	if err != nil {
		return err
	}
	err = u.repo.UpdateStatusRA(ra, ar, ra2)
	if err != nil {
		return err
	}
	return err
}

func (u UseCase) RejectAdmin(a *Actor, ar *ApproveRequest, ra *RegisterApproval, update map[string]interface{}, valueRa map[string]interface{}) error {
	err := u.repo.UpdateActor(a, ar, update)
	if err != nil {
		return err
	}
	err = u.repo.UpdateStatusRA(ra, ar, valueRa)
	if err != nil {
		return err
	}
	return err
}

func (u UseCase) ActivateAdmin(a *Actor, aar *ActivateAdminRequest, activeTrue map[string]interface{}) error {
	return u.repo.ActivateAdmin(a, aar, activeTrue)
}

func (u UseCase) DeactivateAdmin(a *Actor, d *DeActivateAdminRequest, val map[string]interface{}) error {
	return u.repo.DeactivateAdmin(a, d, val)
}

func (u UseCase) DeleteAdmin(a *Actor, d *DeleteAdminRequest) error {
	return u.repo.DeleteAdmin(a, d)
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{repo: repo}
}
