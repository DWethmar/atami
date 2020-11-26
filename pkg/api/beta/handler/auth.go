package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/user"
)

// AccessDetails contains details for accessing the api
type AccessDetails struct {
	AccessToken string `json:"access_token"`
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

func createRefreshCookie(expires time.Time, domain, token string) *http.Cookie {
	return &http.Cookie{
		Name:     "refresh_token",
		Value:    token,
		Domain:   domain,
		Path:     "/beta/auth/refresh",
		Expires:  expires,
		HttpOnly: true,
		MaxAge:   90000,
		Secure:   false,
	}
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

		user, err := userService.FindByEmail(email)
		if err != nil || user == nil {
			fmt.Printf("error while retrieving user: %v\n", err)
			response.SendServerError(w, r)
			return
		}

		session := strconv.FormatInt(time.Now().UnixNano(), 10)

		accessTokenDuration := time.Minute * 15
		accessToken, err := auth.CreateAccessToken(user.UID, session, time.Now().Add(accessTokenDuration).Unix())
		if err != nil || accessToken == "" {
			fmt.Printf("Error creating access token: %v\n", err)
			response.SendServerError(w, r)
			return
		}

		refreshTokenDuration := time.Hour * 730
		refreshToken, err := auth.CreateRefreshToken(user.UID, session, time.Now().Add(refreshTokenDuration).Unix())
		if err != nil || accessToken == "" {
			fmt.Printf("Error creating refresh token: %v\n", err)
			response.SendServerError(w, r)
			return
		}

		cookie := createRefreshCookie(
			time.Now().Add(refreshTokenDuration),
			"localhost",
			refreshToken,
		)
		http.SetCookie(w, cookie)

		response.SendJSON(w, r, AccessDetails{
			AccessToken: accessToken,
		}, http.StatusOK)
	})
}

// Refresh handles refresh requests
func Refresh(authService *auth.Service, userService *user.Service) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("refresh_token")
		if err != nil {
			response.SendBadRequestError(w, r, errors.New(""))
			return
		}

		refreshToken := cookie.Value
		token, err := auth.VerifyRefreshToken(refreshToken)
		if err != nil {
			response.SendBadRequestError(w, r, err)
			return
		}

		claims, ok := token.Claims.(*auth.CustomClaims)
		if !ok {
			response.SendBadRequestError(w, r, errors.New("could not read token"))
			return
		}

		user, err := userService.FindByUID(claims.Subject)
		if err != nil || user == nil {
			fmt.Printf("error while retrieving user: %v\n", err)
			response.SendServerError(w, r)
			return
		}

		accessTokenDuration := time.Minute * 15
		accessToken, err := auth.CreateAccessToken(
			user.UID,
			claims.SessionID,
			time.Now().Add(accessTokenDuration).Unix(),
		)
		if err != nil || accessToken == "" {
			fmt.Printf("Error creating token: %v\n", err)
			response.SendServerError(w, r)
			return
		}

		response.SendJSON(w, r, AccessDetails{
			AccessToken: accessToken,
		}, http.StatusOK)
	})
}
