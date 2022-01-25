package usersvc

import (
	"strings"

	"github.com/google/uuid"
	"violet.wtf/tunnel/api/appctx"
)

type User struct {
	ID           uuid.UUID
	Username     string
	PasswordHash []byte
	Admin        bool
}

func GetUserByID(id uuid.UUID) (*User, error) {
	row := appctx.DB.QueryRow(
		"SELECT username, password_hash, admin FROM users WHERE id = $1",
		id.String(),
	)

	var (
		username     string
		passwordHash []byte
		admin        bool
	)

	if err := row.Err(); err != nil {
		return nil, err
	}

	if err := row.Scan(&username, &passwordHash, &admin); err != nil {
		return nil, err
	}

	return &User{
		ID:           id,
		Username:     username,
		PasswordHash: passwordHash,
		Admin:        admin,
	}, nil
}

func GetUserByUsername(username string) (*User, error) {
	row := appctx.DB.QueryRow(
		"SELECT id, password_hash, admin FROM users WHERE username_lower = $1",
		strings.ToLower(username),
	)

	if err := row.Err(); err != nil {
		return nil, err
	}

	var (
		idStr        string
		passwordHash []byte
		admin        bool
	)

	if err := row.Scan(&idStr, &passwordHash, &admin); err != nil {
		return nil, err
	}

	return &User{
		ID:           uuid.MustParse(idStr),
		Username:     username,
		PasswordHash: passwordHash,
		Admin:        admin,
	}, nil
}

func Save(user *User) error {
	_, err := appctx.DB.Exec(
		"INSERT INTO users " + 
			"(id, username, username_lower, password_hash, admin) "+
			"VALUES ($1, $2, $3, $4, $5)",
		user.ID.String(),
		user.Username,
		strings.ToLower(user.Username),
		user.PasswordHash,
		user.Admin,
	)
	return err
}
