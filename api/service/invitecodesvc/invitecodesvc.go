package invitecodesvc

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"violet.wtf/tunnel/api/appctx"
)

func UseInviteCodeAndReturnExists(inviteCode string) (bool, error) {
	row := appctx.DB.QueryRow(
		"DELETE FROM invite_codes WHERE code = $1 RETURNING id",
		inviteCode,
	)

	if err := row.Err(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else {
			return false, err
		}
	}

	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}
