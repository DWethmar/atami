package user

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/model"
	"github.com/dwethmar/atami/pkg/usecase/userusecase"
)

// User struct definition
// When updating this struct also update: TestMapUsers()
type User struct {
	UID      string `json:"uid"`
	Username string `json:"username"`
}

func toUsers(users []*model.User) []*User {
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
func ListUsers(usecase *userusecase.Usecase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result, err := usecase.List()

		if err == nil {
			users := toUsers(result)
			response.SendJSON(w, r, users, 200)
		} else {
			response.SendServerError(w, r)
		}
	})
}
