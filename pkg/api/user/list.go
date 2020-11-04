package user

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/user"
)

// User struct definition
// When updating this struct also update: TestMapUsers()
type User struct {
	UID      user.UID `json:"uid"`
	Username string   `json:"username"`
}

func toUsers(users []*user.User) []*User {
	mapped := make([]*User, len(users))
	for i, user := range users {
		mapped[i] = &User{
			UID:      user.UID,
			Username: user.Username,
		}
	}
	return mapped
}

// ListUsers handler
func ListUsers(finder *user.Finder) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result, err := finder.FindAll()
		if err == nil {
			users := toUsers(result)
			response.SendJSON(w, r, users, 200)
		} else {
			response.SendServerError(w, r)
		}
	})
}