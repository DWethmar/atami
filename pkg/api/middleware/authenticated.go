package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/api/token"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/model"
)

type userCTXKeyType string

const userCTXKey userCTXKeyType = "userCTXKey"

// WithUser puts the request ID into the current context.
func WithUser(ctx context.Context, user *model.User) context.Context {
	return context.WithValue(ctx, userCTXKey, user)
}

// UserFromContext returns the user UID from the context.
func UserFromContext(ctx context.Context) (*model.User, error) {
	value := ctx.Value(userCTXKey)
	if user, ok := value.(*model.User); ok {
		return user, nil
	}
	return nil, errors.New("No user found in context")
}

// Authenticated handles auth requests
func Authenticated(authService auth.Service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")

			if len(splitToken) != 2 {
				response.SendUnauthorizedError(w, r, errors.New("Invalid authorization header"))
				return
			}

			tokenString := splitToken[1]
			var UID model.UserUID

			if token, err := token.VerifyToken(tokenString); err == nil && token.Valid {
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					UIDString, ok := claims["uid"].(string)
					if ok {
						UID = model.UserUID(UIDString)
					} else {
						response.SendUnauthorizedError(w, r, errors.New("UID not set"))
						return
					}
				}
			} else {
				response.SendUnauthorizedError(w, r, err)
				return
			}

			if UID != "" {
				if user, err := authService.FindByUID(UID); err == nil {
					fmt.Printf("OK! %v \n", user)
					ctx := WithUser(r.Context(), user)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					response.SendServerError(w, r)
					fmt.Print("Not OK!\n")
					return
				}
			} else {
				response.SendUnauthorizedError(w, r, errors.New("Invalid JWT Token"))
			}
		}
		return http.HandlerFunc(fn)
	}
}
