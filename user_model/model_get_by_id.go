package user_model

import (
	"context"
	"database/sql"
)

// ModelGetByID ModelGetByID
func ModelGetByID(db *sql.DB, id int) (*UserModel, error) {
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
		where u.id = $1
		limit 1
		;
	`

	itemDB := UserModelDB{}
	if err := db.QueryRowContext(context.Background(), query, id).Scan(
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

	item := itemDB.GetObject()

	return &item, nil
}
