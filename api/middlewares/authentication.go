package middlewares

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/turgaysozen/littlejohn/dummy_data"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Extract the token from the Authorization header
		token := strings.TrimPrefix(authHeader, "Basic ")

		// Validate the username against the token
		if !isValidToken(token) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		usernameBytes, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		username := strings.TrimSuffix(string(usernameBytes), ":")
		ctx := context.WithValue(r.Context(), "username", username)

		// Replace the original request context with the new context
		r = r.WithContext(ctx)

		// If the token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

func isValidToken(token string) bool {
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return false
	}

	// Convert the decoded token to a string
	username := strings.TrimSuffix(string(decodedToken), ":")

	// Check if the username is present in the list of valid usernames
	for _, validUsername := range dummy_data.ValidUsernames {
		if username == validUsername {
			return true
		}
	}

	return false
}
