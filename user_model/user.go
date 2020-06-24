package user_model

import "time"

// UserModel UserModel
type UserModel struct {
	ID          int       `json:"id"`
	Nome        string    `json:"nome"`
	Sobrenome   string    `json:"sobrenome"`
	Username    string    `json:"username"`
	Senha       string    `json:"-"`
	Salt        string    `json:"-"`
	Datacriacao time.Time `json:"datacriacao"`
}
