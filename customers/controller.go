package customers

type Controller struct {
	uc *UseCase
}

func (c Controller) CreateCustomer(cc *CreateCustomerRequest) (*CreateCustomerResponse, error) {
	cust := Customer{
		FirstName: cc.FirstName,
		LastName:  cc.LastName,
		Email:     cc.Email,
		Avatar:    cc.Avatar,
	}

	err := c.uc.CreateCustomer(&cust)

	if err != nil {
		return nil, err
	}

	res := &CreateCustomerResponse{
		Status: "sukses",
		Data: CustomerResponse{
			FirstName: cust.FirstName,
			LastName:  cust.LastName,
			Email:     cust.Email,
			Avatar:    cust.Avatar,
		},
	}

	return res, nil
}

func NewController(uc *UseCase) *Controller {
	return &Controller{
		uc: uc,
	}
}
