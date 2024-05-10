package postgres

import (
	"context"

	"github.com/VanillaFox/system_architecture_lab/users/models"
)

func (r *Repository) Auth(ctx context.Context, creds *models.Creds) (string, error) {
	const sql = `
	SELECT pass
	FROM users
	WHERE username = $1;
	`

	var password string

	err := r.db.QueryRow(ctx, sql, creds.Username).Scan(&password)

	if err != nil {
		return password, err
	}

	return password, nil
}
