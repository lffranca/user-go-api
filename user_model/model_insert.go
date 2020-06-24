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
	senha string,
	salt string,
) error {
	query := `
		
	`

	if _, err := db.ExecContext(context.Background(), query,
		nome,
		sobrenome,
		username,
		senha,
		salt,
	): err != nil {
		return err
	}
	
	return nil
}
