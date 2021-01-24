package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/go-chi/chi"
)

type uuidCtxKeyType string

const uuidCtxKey uuidCtxKeyType = "uid"

// WithUID puts the request UID into the current context.
func WithUID(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, uuidCtxKey, uid)
}

// UIDFromContext returns the request ID from the context.
// A zero ID is returned if there are no identifers in the
// current context.
func UIDFromContext(ctx context.Context) (string, error) {
	v, ok := ctx.Value(uuidCtxKey).(string)
	if !ok {
		return "", errors.New("Could not receive UID")
	}
	return v, nil
}

// RequireUID requires that a uid is provided in the url.
func RequireUID(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid := chi.URLParam(r, "uid")
		if uid == "" {
			response.BadRequestError(w, r, errors.New("no UID found in URL"))
			return
		}
		ctx := WithUID(r.Context(), uid)
		next(w, r.WithContext(ctx))
	})
}
