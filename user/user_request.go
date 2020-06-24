package user

// UserRequest UserRequest
type UserRequest struct {
	Nome      string `json:"nome" binding:"required"`
	Sobrenome string `json:"sobrenome" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Senha     string `json:"senha" binding:"required"`
}
