package middleware

import (
	"net/http"
	"strings"

	"pruebago.com/go/rest-ws/helpers"
	"pruebago.com/go/rest-ws/server"
)

var (
	NO_AUTH_NEEDED = []string{
		"login",
		"signup",
	}
)

// shouldCheckToken Helper - For non protected routes
func shouldCheckToken(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

// Definimos el Middleware - De verificacion de token para un usuario
func CheckAuthMiddleware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
			_, err := helpers.GetJWTAuthorizationInfo(s, w, r)
			if err != nil {
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
