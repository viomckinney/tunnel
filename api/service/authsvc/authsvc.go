package authsvc

import (
	"database/sql"
	"errors"

	"violet.wtf/tunnel/api/service/sessionsvc"
	"violet.wtf/tunnel/api/service/usersvc"
)

func UserFromSessionToken(
	token string,
) (user *usersvc.User, exists bool, err error) {
	session, err := sessionsvc.GetByToken(token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			exists = false
			err = nil
			return
		} else {
			return
		}
	}

	user, err = usersvc.GetUserByID(session.UserID)
	if err == nil {
		exists = true
	}
	return
}
