package postgres

import (
	"context"
	"fmt"

	"github.com/VanillaFox/system_architecture_lab/users/models"
)

func (r *Repository) CreateUser(ctx context.Context, user *models.User) error {
	const sql = `
	INSERT INTO users(id, full_name, username, pass)
	VALUES (gen_random_uuid(), $1, $2, $3);
	`

	password, err := user.Password.Hash()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, sql, user.FullName, user.Username, password)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	const sql = `
	SELECT full_name, username
	FROM users
	WHERE username = $1;
	`

	var user models.User

	err := r.db.QueryRow(ctx, sql, username).Scan(&user.FullName, &user.Username)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetByFullNamePrefix(ctx context.Context, fullNamePrefix string) (*models.Users, error) {
	const sql = `
	SELECT full_name, username
	FROM users
	WHERE full_name LIKE $1;
	`

	rows, err := r.db.Query(ctx, sql, fmt.Sprint(fullNamePrefix, "%"))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users models.Users

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.FullName, &user.Username)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *Repository) GetUsers(ctx context.Context) (models.Users, error) {
	const sql = `
	SELECT full_name, username
	FROM users;
	`

	rows, err := r.db.Query(ctx, sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users models.Users

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.FullName, &user.Username)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) UpdateUser(ctx context.Context, username string, user *models.User) (*models.User, error) {
	const sql = `
	UPDATE users
	SET full_name=$1, username=$2, pass=$3
	WHERE username=$4;
	`

	password, err := user.Password.Hash()

	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(ctx, sql, user.FullName, user.Username, password, username)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) DeleteUser(ctx context.Context, username string) error {
	const sql = `
	DELETE from users
	WHERE username=$1;
	`

	_, err := r.db.Exec(ctx, sql, username)

	if err != nil {
		return err
	}

	return nil
}
