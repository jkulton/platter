package auth

import (
	"crypto/subtle"
	"log"
	"net/http"
)

// SameStrings sonverts string inputs to byte slices then compares using crypto/subtle
func SameStrings(s1, s2 string) bool {
	b1 := []byte(s1)
	b2 := []byte(s2)
	return subtle.ConstantTimeCompare(b1, b2) == 1
}

// AuthMiddleware implements basic auth for HTTP routes
func AuthMiddleware(next http.Handler, auth string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth != "" {
			username, password, _ := r.BasicAuth()
			formattedAttempt := username + ":" + password

			if !SameStrings(auth, formattedAttempt) {
				log.Println("Unauthorized.")
				http.Error(w, "Unauthorized", 401)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
