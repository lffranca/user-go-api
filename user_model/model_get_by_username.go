package user_model

import (
	"context"
	"database/sql"
	"errors"
)

// ModelGetByUsername ModelGetByUsername
func ModelGetByUsername(db *sql.DB, username string) (*UserModel, error) {
	query := `
		select
			u.id,
			u.nome,
			u.sobrenome,
			u.username,
			u.senha,
			u.salt,
			u.datacriacao
		from usuario as u
		where u.username = $1
		order by u.id desc
		;
	`

	itemDB := UserModelDB{}
	if err := db.QueryRowContext(context.Background(), query, username).Scan(
		&itemDB.ID,
		&itemDB.Nome,
		&itemDB.Sobrenome,
		&itemDB.Username,
		&itemDB.Senha,
		&itemDB.Salt,
		&itemDB.Datacriacao,
	); err != nil {
		return nil, err
	}

	if !itemDB.ID.Valid {
		return nil, errors.New("Invalid user")
	}

	item := itemDB.GetObject()
	return &item, nil
}
