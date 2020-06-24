package user_model

import (
	"context"
	"database/sql"
)

// ModelInsert ModelInsert
func ModelInsert(db *sql.DB,
	nome string,
	sobrenome string,
	username string,
	hash string,
	salt string,
) error {
	query := `
		insert into usuario (nome, sobrenome, username, senha, salt, datacriacao)
		values ($1, $2, $3, $4, $5, current_timestamp)
		;
	`

	if _, err := db.ExecContext(context.Background(), query,
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
