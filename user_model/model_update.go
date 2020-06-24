package user_model

import (
	"context"
	"database/sql"
)

// ModelUpdate ModelUpdate
func ModelUpdate(db *sql.DB,
	id int,
	nome string,
	sobrenome string,
	username string,
	hash string,
	salt string,
) error {
	query := `
		update usuario set
			nome = $2,
			sobrenome = $3,
			username = $4,
			senha = $5,
			salt = $6
		where id  = $1
		;
	`

	if _, err := db.ExecContext(context.Background(), query,
		id,
		nome,
		sobrenome,
		username,
		hash,
		salt,
	); err != nil {
		return err
	}

	return nil
}
