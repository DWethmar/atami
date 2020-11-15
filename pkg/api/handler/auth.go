package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/model"
)

// NewUser struct definition
type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

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

func toResponds(users []*model.User) []*Responds {
	r := make([]*Responds, len(users))
	for i, user := range users {
		r[i] = toRespond(user)
	}
	return r
}

func toRespond(user *model.User) *Responds {
	return &Responds{
		UID:       user.UID.String(),
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
}

// ListUsers handler
func ListUsers(service auth.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if users, err := service.Find(); err == nil {
			response.SendJSON(w, r, toResponds(users), 200)
		} else {
			fmt.Printf("Error: %v \n", err)
			response.SendServerError(w, r)
		}
	})
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
