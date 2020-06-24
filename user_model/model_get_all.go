package user_model

import (
	"context"
	"database/sql"
	"log"
)

// ModelGetAll ModelGetAll
func ModelGetAll(db *sql.DB) ([]UserModel, error) {
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
		order by u.id desc
		;
	`

	rows, errRows := db.QueryContext(context.Background(), query)
	if errRows != nil {
		return nil, errRows
	}

	defer rows.Close()

	items := []UserModel{}
	for rows.Next() {
		itemDB := UserModelDB{}
		if err := rows.Scan(
			&itemDB.ID,
			&itemDB.Nome,
			&itemDB.Sobrenome,
			&itemDB.Username,
			&itemDB.Senha,
			&itemDB.Salt,
			&itemDB.Datacriacao,
		); err != nil {
			log.Println("USER MODEL GET ALL", err)
			continue
		}

		if itemDB.ID.Valid {
			items = append(items, itemDB.GetObject())
		}
	}

	return items, nil
}
