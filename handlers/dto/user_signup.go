package dto

type UserSignupJson struct {
	Username string `json:"username" binding:"required"`
	RoleID   int    `json:"role_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
