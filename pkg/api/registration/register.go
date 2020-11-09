package registration

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/auth"
)

// NewUser struct definition
type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register handler handles the request to create new user
func Register(service auth.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var newUser = NewUser{}
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			fmt.Printf("Error decoding entry: %v\n", err)
			response.SendBadRequestError(w, r, errors.New("Invalid input"))
			return
		}

		createUser := auth.CreateUser{
			Username: newUser.Username,
			Email:    newUser.Email,
			Password: newUser.Password,
		}

		if err := service.ValidateNewUser(createUser); err != nil {
			response.SendBadRequestError(w, r, err)
			return
		}

		user, err := service.Register(createUser)

		if err != nil || user == nil {
			fmt.Printf("Error registering user: %v\n", err)
			response.SendBadRequestError(w, r, err)
			return
		}

		response.SendJSON(w, r, toRespond(user), http.StatusCreated)
	})
}
