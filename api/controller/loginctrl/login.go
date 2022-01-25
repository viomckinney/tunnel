package loginctrl

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"violet.wtf/tunnel/api/service/invitecodesvc"
	"violet.wtf/tunnel/api/service/sessionsvc"
	"violet.wtf/tunnel/api/service/usersvc"
	"violet.wtf/tunnel/api/tutil"
)

func LoginRoute(w http.ResponseWriter, r *http.Request) {
	var loginRequest loginRequestStruct

	if tutil.OhNo(w, json.NewDecoder(r.Body).Decode(&loginRequest)) {
		return
	}

	user, err := usersvc.GetUserByUsername(loginRequest.Username)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid username or password"))
			return
		}

		tutil.OhNo(w, err)
		return
	}

	correct, err := tutil.HashMatches(loginRequest.Password, user.PasswordHash)
	if tutil.OhNo(w, err) {
		return
	}

	if !correct {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid username or password"))
		return
	}

	token, err := tutil.GenerateRandomString(32)
	if tutil.OhNo(w, err) {
		return
	}

	fmt.Println("UID: " + user.ID.String())

	sessionsvc.Save(&sessionsvc.Session{
		ID:     uuid.Must(uuid.NewRandom()),
		UserID: user.ID,
		Token:  token,
	})

	w.Write([]byte(token))
}

func TokenExistsRoute(w http.ResponseWriter, r *http.Request) {
	tokenExists, err := sessionsvc.ExistsByToken(r.URL.Query().Get("token"))
	if tutil.OhNo(w, err) {
		return
	}

	w.Write(tutil.UnsafeMarshal(tokenExists))
}

func RegisterRoute(w http.ResponseWriter, r *http.Request) {
	var registerRequest registerRequestStruct

	if tutil.OhNo(w, json.NewDecoder(r.Body).Decode(&registerRequest)) {
		return
	}

	codeExists, err := invitecodesvc.UseInviteCodeAndReturnExists(
		registerRequest.InviteCode,
	)
	if tutil.OhNo(w, err) {
		return
	}

	if !codeExists {
		w.Write([]byte("Bad Code"))
		return
	}

	hashedPassword, err := tutil.Hash(registerRequest.Password)
	if tutil.OhNo(w, err) {
		return
	}

	usersvc.Save(&usersvc.User{
		ID:           uuid.Must(uuid.NewRandom()),
		Username:     registerRequest.Username,
		PasswordHash: hashedPassword,
	})

	w.Write([]byte("OK"))
}

type registerRequestStruct struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	InviteCode string `json:"inviteCode"`
}

type loginRequestStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
