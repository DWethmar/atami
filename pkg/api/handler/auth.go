package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/user"
)

type loginResponds struct {
	AccessToken string `json:"access_token"`
}

func newExpirationDate() int64 {
	return time.Now().Add(time.Minute * 15).Unix()
}

// Responds struct declaration
type Responds struct {
	UID       string    `json:"uid"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func toResponds(users []*user.User) []*Responds {
	r := make([]*Responds, len(users))
	for i, user := range users {
		r[i] = toRespond(user)
	}
	return r
}

func toRespond(u *user.User) *Responds {
	return &Responds{
		UID:       u.UID,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}
}

// ListUsers handler
func ListUsers(userService *user.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if users, err := userService.Find(); err == nil {
			response.SendJSON(w, r, toResponds(users), 200)
		} else {
			fmt.Printf("Error: %v \n", err)
			response.SendServerError(w, r)
		}
	})
}

// Register handler handles the request to create new user
func Register(authService *auth.Service) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")

		createUser := auth.CreateUser{
			Username: username,
			Email:    email,
			Password: password,
		}

		if err := authService.ValidateNewUser(createUser); err != nil {
			response.SendBadRequestError(w, r, err)
			return
		}

		user, err := authService.Register(createUser)

		if err != nil || user == nil {
			fmt.Printf("Error registering user: %v\n", err)
			response.SendBadRequestError(w, r, err)
			return
		}

		response.SendJSON(w, r, toRespond(user), http.StatusCreated)
	})
}

// Login handles login requests
func Login(authService *auth.Service, userService *user.Service) http.HandlerFunc {

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

		user, err := userService.FindByEmail(email, false)
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
