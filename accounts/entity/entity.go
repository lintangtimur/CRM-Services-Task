package entity

import "gorm.io/gorm"

type Actor struct {
	gorm.Model
	Username string             `json:"username" binding:"required"`
	Password string             `json:"password" binding:"required"`
	RoleID   uint               `json:"role_id" binding:"required"`
	Verified string             `json:"verified" gorm:"type:enum('true','false');default:'false'"`
	Active   string             `json:"active" gorm:"type:enum('true','false');default:'false'"`
	RA       []RegisterApproval `json:"RA" gorm:"foreignKey:AdminID;references:ID"`
}

type RegisterApproval struct {
	gorm.Model
	AdminID      uint   `json:"admin_id"`
	SuperAdminID uint   `json:"super_admin_id"`
	Status       string `json:"status"`
}
