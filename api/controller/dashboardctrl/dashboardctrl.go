package dashboardctrl

import (
	"net/http"

	"violet.wtf/tunnel/api/service/authsvc"
	"violet.wtf/tunnel/api/service/sessionsvc"
	"violet.wtf/tunnel/api/tutil"
)

func DashboardRoute(w http.ResponseWriter, r *http.Request) {
	user, exists, err := authsvc.UserFromSessionToken(
		r.Header.Get("Authorization"),
	)
	if tutil.OhNo(w, err) {
		return
	}

	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Bad Authentication"))
		return
	}

	sessionCount, err := sessionsvc.CountSessionsByUserID(user.ID)
	if tutil.OhNo(w, err) {
		return
	}

	w.Write(tutil.UnsafeMarshal(dashboardResponseStruct{
		Username:     user.Username,
		IsAdmin:      user.Admin,
		Tunnels:      []string{"todo", "todo2"},
		SessionCount: sessionCount,
		APIKeyCount:  0,  // TODO
	}))
}

type dashboardResponseStruct struct {
	Username     string   `json:"username"`
	IsAdmin      bool     `json:"isAdmin"`
	Tunnels      []string `json:"tunnels"`
	SessionCount int      `json:"sessionCount"`
	APIKeyCount  int      `json:"apiKeyCount"`
}
