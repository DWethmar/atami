package user

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

// RegisterUser handler
func RegisterUser(service auth.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// if err := r.ParseForm(); err != nil {
		// 	fmt.Fprintf(w, "ParseForm() err: %v", err)
		// 	response.SendServerError(w, r)
		// 	return
		// }

		var newUser = NewUser{}
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			fmt.Printf("Error while decoding entry: %v", err)
			response.SendBadRequestError(w, r, errors.New("Invalid input"))
			return
		}

		user, err := service.Register(auth.NewUser{
			Username: newUser.Username,
			Email:    newUser.Email,
			Password: newUser.Password,
		})

		if err != nil {
			fmt.Printf("Error while registering user: %v\n", err)
			response.SendBadRequestError(w, r, err)
			return
		}

		response.SendJSON(w, r, user, http.StatusCreated)
	})
}
