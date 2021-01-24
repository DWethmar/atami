package beta

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
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
			response.JSON(w, r, toResponds(users), 200)
		} else {
			fmt.Printf("Error: %v \n", err)
			response.ServerError(w, r)
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
			response.BadRequestError(w, r, err)
			return
		}

		user, err := authService.Register(createUser)

		if err != nil || user == nil {
			fmt.Printf("Error registering user: %v\n", err)
			response.BadRequestError(w, r, err)
			return
		}

		response.JSON(w, r, toRespond(user), http.StatusCreated)
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
			response.BadRequestError(w, r, errors.New("email is empty"))
			return
		}

		if password == "" {
			response.BadRequestError(w, r, errors.New("password is empty"))
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
			response.BadRequestError(w, r, errors.New("validation failure"))
			return
		}

		user, err := userService.FindByEmail(email)
		if err != nil || user == nil {
			fmt.Printf("error while retrieving user: %v\n", err)
			response.ServerError(w, r)
			return
		}

		session := strconv.FormatInt(time.Now().UnixNano(), 10)

		accessTokenDuration := time.Minute * 60
		accessToken, err := auth.CreateAccessToken(user.UID, session, time.Now().Add(accessTokenDuration).Unix())
		if err != nil || accessToken == "" {
			fmt.Printf("Error creating access token: %v\n", err)
			response.ServerError(w, r)
			return
		}

		refreshTokenDuration := time.Hour * 730
		refreshToken, err := auth.CreateRefreshToken(user.UID, session, time.Now().Add(refreshTokenDuration).Unix())
		if err != nil || accessToken == "" {
			fmt.Printf("Error creating refresh token: %v\n", err)
			response.ServerError(w, r)
			return
		}

		cookie := createRefreshCookie(
			time.Now().Add(refreshTokenDuration),
			"localhost",
			refreshToken,
		)
		http.SetCookie(w, cookie)

		response.JSON(w, r, AccessDetails{
			AccessToken: accessToken,
		}, http.StatusOK)
	})
}

// Refresh handles refresh requests
func Refresh(authService *auth.Service, userService *user.Service) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("refresh_token")
		if err != nil {
			response.BadRequestError(w, r, errors.New(""))
			return
		}

		refreshToken := cookie.Value
		token, err := auth.VerifyRefreshToken(refreshToken)
		if err != nil {
			response.BadRequestError(w, r, err)
			return
		}

		claims, ok := token.Claims.(*auth.CustomClaims)
		if !ok {
			response.BadRequestError(w, r, errors.New("could not read token"))
			return
		}

		user, err := userService.FindByUID(claims.Subject)
		if err != nil || user == nil {
			fmt.Printf("error while retrieving user: %v\n", err)
			response.ServerError(w, r)
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
			response.ServerError(w, r)
			return
		}

		response.JSON(w, r, AccessDetails{
			AccessToken: accessToken,
		}, http.StatusOK)
	})
}

// NewAuthRouter returns the api routes handler
func NewAuthRouter(authService *auth.Service, userService *user.Service) http.Handler {
	r := chi.NewRouter()

	logger := httplog.NewLogger("auth", httplog.Options{})
	r.Use(httplog.RequestLogger(logger))
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		// AllowedOrigins: []string{"*"},
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			return true
		},
		AllowedMethods:   []string{"POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Post("/register", Register(authService))
	r.Post("/login", Login(authService, userService))
	r.Post("/refresh", Refresh(authService, userService))

	return r
}
