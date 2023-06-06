package accounts

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type ActorResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	RoleID   uint   `json:"role_id" `
	Verified string `json:"verified" gorm:"type:enum('true','false');default:'false'"`
	Active   string `json:"active" gorm:"type:enum('true','false');default:'false'"`
}

type LoginResponse struct {
	User  ActorResponse
	Token string `json:"token"`
}
type AdminRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type AdminRegisterResponse struct {
	Status string        `json:"status"`
	Data   ActorResponse `json:"data"`
}
type ApprovalListResponse struct {
	AdminID []uint `json:"admin_id"`
}
type ApproveRequest struct {
	AdminID int `json:"admin_id" binding:"required"`
}
type ApproveResponse struct {
	Message string `json:"message"`
}
