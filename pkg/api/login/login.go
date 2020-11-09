package login

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/api/token"
	"github.com/dwethmar/atami/pkg/auth"
)

type loginResponds struct {
	Token string `json:"token"`
}

// Login handles login requests
func Login(service auth.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "" {
			response.SendBadRequestError(w, r, errors.New("email is empty"))
			return
		}

		if password == "" {
			response.SendBadRequestError(w, r, errors.New("password is empty"))
			return
		}

		var authenticated = false

		if ok, err := service.Authenticate(auth.Credentials{
			Email:    email,
			Password: password,
		}); err == nil && ok {
			authenticated = true
		}

		if !authenticated {
			response.SendBadRequestError(w, r, errors.New("validation failure"))
			return
		}

		user, err := service.FindByEmail(email)
		if err != nil || user == nil {
			fmt.Printf("error while retrieving user\n")
			response.SendServerError(w, r)
			return
		}

		token, err := token.CreateToken(user.UID, user.Username)
		if err != nil || token == "" {
			fmt.Printf("Error creating token: %v\n", err)
			response.SendServerError(w, r)
			return
		}

		response.SendJSON(w, r, loginResponds{
			Token: token,
		}, http.StatusOK)
	})
}
