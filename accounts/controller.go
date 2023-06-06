package accounts

type Controller struct {
	uc *UseCase
}

func NewController(uc *UseCase) *Controller {
	return &Controller{
		uc: uc,
	}
}
