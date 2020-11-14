package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/response"
)

type loginResponds struct {
	AccessToken string `json:"access_token"`
}

func newExpirationDate() int64 {
	return time.Now().Add(time.Minute * 15).Unix()
}

// Login handles login requests
func Login(authService auth.Service) http.HandlerFunc {
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
		if ok, err := authService.Authenticate(auth.Credentials{
			Email:    email,
			Password: password,
		}); err == nil && ok {
			authenticated = true
		}

		if !authenticated {
			response.SendBadRequestError(w, r, errors.New("validation failure"))
			return
		}

		user, err := authService.FindByEmail(email)
		if err != nil || user == nil {
			fmt.Printf("error while retrieving user: %v\n", err)
			response.SendServerError(w, r)
			return
		}

		details, err := auth.CreateToken(user.UID, user.Username, newExpirationDate())
		if err != nil || details.AccessToken == "" {
			fmt.Printf("Error creating token: %v\n", err)
			response.SendServerError(w, r)
			return
		}

		response.SendJSON(w, r, loginResponds{
			AccessToken: details.AccessToken,
		}, http.StatusOK)
	})
}
