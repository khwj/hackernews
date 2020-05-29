package auth

import (
	"context"
	"net/http"
	"strconv"

	"github.com/khwj/hackernews/internal/users"
	"github.com/khwj/hackernews/pkg/jwt"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{name: "user"}

// Middleware ...
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			// Allow unauthenticated users in.
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			// Validate jwt token.
			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				// next.ServeHTTP(w, r)
				return
			}

			// Create user instance and check if user exists in db.
			user := users.User{Username: username}
			id, err := users.GetUserIDByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = strconv.Itoa(id)

			// Put user id in the context.
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// And call the next handler with our new context.
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to run.
func ForContext(ctx context.Context) *users.User {
	user, _ := ctx.Value(userCtxKey).(*users.User)
	return user
}
