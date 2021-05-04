package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/domain"
	"github.com/dwethmar/atami/pkg/domain/user"
)

type userCTXKeyType string

const userCTXKey userCTXKeyType = "userCTXKey"

// WithUser puts the request ID into the current context.
func WithUser(ctx context.Context, user *user.User) context.Context {
	return context.WithValue(ctx, userCTXKey, user)
}

// GetUser returns the user UID from the context.
func GetUser(ctx context.Context) (*user.User, error) {
	value := ctx.Value(userCTXKey)
	if user, ok := value.(*user.User); ok {
		return user, nil
	}
	return nil, errors.New("No user found in context")
}

// Authenticated handles auth requests
func Authenticated(userStore *domain.UserStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			reqToken := r.Header.Get("Authorization")
			if reqToken == "" {
				response.UnauthorizedError(w, r, errors.New("unauthorized"))
				return
			}

			splitToken := strings.Split(reqToken, "Bearer ")

			if len(splitToken) != 2 {
				response.UnauthorizedError(w, r, errors.New("invalid authorization header"))
				return
			}

			tokenString := splitToken[1]
			var UID string

			if token, err := auth.VerifyAccessToken(tokenString); err == nil && token.Valid {
				if claims, ok := token.Claims.(*auth.CustomClaims); ok && token.Valid {
					UID = claims.Subject
				}
			} else {
				response.UnauthorizedError(w, r, errors.New("Invalid token"))
				return
			}

			if UID != "" {
				if user, err := userStore.FindByUID(UID); err == nil {
					ctx := WithUser(r.Context(), user)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					fmt.Print(err)
					response.ServerError(w, r)
					return
				}
			} else {
				response.UnauthorizedError(w, r, errors.New("Invalid token"))
			}
		}
		return http.HandlerFunc(fn)
	}
}
