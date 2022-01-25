package sessionsvc

import (
	"github.com/google/uuid"
	"violet.wtf/tunnel/api/appctx"
)

type Session struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	Token   string
}

func ExistsByToken(token string) (bool, error) {
	row := appctx.DB.QueryRow(
		"SELECT COUNT(1) FROM sessions WHERE token = $1",
		token,
	)

	if err := row.Err(); err != nil {
		return false, err
	}

	var count int
	row.Scan(&count)

	return count > 0, nil
}

func GetByToken(token string) (*Session, error) {
	row := appctx.DB.QueryRow(
		"SELECT id, user_id FROM sessions WHERE token = $1",
		token,
	)

	if err := row.Err(); err != nil {
		return nil, err
	}

	var idStr, userIDStr string
	if err := row.Scan(&idStr, &userIDStr); err != nil {
		return nil, err
	}

	return &Session{
		ID:     uuid.MustParse(idStr),
		UserID: uuid.MustParse(userIDStr),
		Token:  token,
	}, nil
}

func CountSessionsByUserID(userID uuid.UUID) (int, error) {
	row := appctx.DB.QueryRow(
		"SELECT COUNT(1) FROM sessions WHERE user_id = $1",
		userID,
	)

	if err := row.Err(); err != nil {
		return -1, err
	}

	var count int
	if err := row.Scan(&count); err != nil {
		return -1, err
	}

	return count, nil
}

func Save(session *Session) error {
	_, err := appctx.DB.Exec(
		"INSERT INTO sessions (id, user_id, token) VALUES ($1, $2, $3)",
		session.ID,
		session.UserID,
		session.Token,
	)

	return err
}
