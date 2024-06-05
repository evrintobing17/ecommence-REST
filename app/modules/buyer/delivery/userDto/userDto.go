package userdto

import "time"

type Register struct {
	Address  string `json:"address" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,gte=6"`
}

type ResRegister struct {
	Address string `json:"address"`
	Email   string `json:"email"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
}

type ReqLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,gte=6"`
}

type ResLogin struct {
	Jwt string `json:"jwt"`
}

type ReqUpdate struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type RespUpdate struct {
	ID        int        `json:"id"`
	Email     string     `json:"email"`
	Name      string     `json:"name"`
	Age       string     `json:"age"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type DeleteResp struct {
	Message string `json:"message"`
}
