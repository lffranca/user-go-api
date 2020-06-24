package login

// LoginRequest LoginRequest
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Senha    string `json:"senha" binding:"required"`
}
