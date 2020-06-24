package user_model

import (
	"context"
	"database/sql"
)

// ModelDelete ModelDelete
func ModelDelete(db *sql.DB, id int) error {
	query := `
		delete from usuario where id = $1;
	`

	if _, err := db.ExecContext(context.Background(), query, id); err != nil {
		return err
	}

	return nil
}
