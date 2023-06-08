package dto

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
type RejectResponse struct {
	Message string `json:"message"`
}
type ActivateAdminRequest struct {
	AdminID uint `json:"admin_id" binding:"required"`
}
type ActivateAdminResponse struct {
	Message string `json:"message"`
}
type DeActivateAdminRequest struct {
	AdminID uint `json:"admin_id" binding:"required"`
}
type DeActivateAdminResponse struct {
	Message string `json:"message"`
}
type DeleteAdminRequest struct {
	AdminID int `json:"admin_id" binding:"required"`
}
type DeleteAdminResponse struct {
	Message string `json:"message"`
}
type DataActorResponse struct {
	Data []ActorResponse `json:"data"`
}
