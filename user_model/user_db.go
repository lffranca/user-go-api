package user_model

import (
	"database/sql"
	"time"
)

// UserModelDB UserModelDB
type UserModelDB struct {
	ID          sql.NullInt64
	Nome        sql.NullString
	Sobrenome   sql.NullString
	Username    sql.NullString
	Senha       sql.NullString
	Salt        sql.NullString
	Datacriacao sql.NullString
}

// GetObject GetObject
func (item *UserModelDB) GetObject() UserModel {
	t, _ := time.Parse(time.RFC3339, item.Datacriacao.String)

	return UserModel{
		ID:          int(item.ID.Int64),
		Nome:        item.Nome.String,
		Sobrenome:   item.Sobrenome.String,
		Username:    item.Username.String,
		Senha:       item.Senha.String,
		Salt:        item.Salt.String,
		Datacriacao: t,
	}
}
