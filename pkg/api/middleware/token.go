package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/api/token"
	"github.com/dwethmar/atami/pkg/model"
)

type userUIDCtxKeyType string

const userUIDCtxKey userUIDCtxKeyType = "uuid"

// WithUserUID puts the request ID into the current context.
func WithUserUID(ctx context.Context, UID model.UserUID) context.Context {
	return context.WithValue(ctx, userUIDCtxKey, UID)
}

// UserUIDFromContext returns the user UID from the context.
func UserUIDFromContext(ctx context.Context) (model.UserUID, error) {
	v, ok := ctx.Value(userUIDCtxKey).(model.UserUID)
	if !ok {
		return model.UserUID(""), errors.New("No user UID found in context")
	}
	return v, nil
}

// Token handles auth requests
func Token(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")

		if len(splitToken) != 2 {
			response.SendBadRequestError(w, r, errors.New("Invalid authorization header"))
			return
		}

		tokenString := splitToken[1]

		if token, err := token.VerifyToken(tokenString); err == nil && token.Valid {
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				UID := claims["uid"].(model.UserUID)
				WithUserUID(r.Context(), UID)
			} else {
				response.SendBadRequestError(w, r, errors.New("Invalid JWT Token"))
				return
			}
		} else {
			fmt.Print(reflect.TypeOf(err))
			response.SendBadRequestError(w, r, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
