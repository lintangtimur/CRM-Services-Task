package customers

type CreateCustomerRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Avatar    string `json:"avatar" binding:"required"`
}
type CustomerResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}
type CreateCustomerResponse struct {
	Status string `json:"status"`
	Data   CustomerResponse
}
