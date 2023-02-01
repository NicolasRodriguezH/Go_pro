package helpers

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"pruebago.com/go/rest-ws/models"
	"pruebago.com/go/rest-ws/server"
)

// Veo que solo recortó este codigo del achivo principal, y creo esta nueva carpeta donde lo agrego, luego desde los archivos donde se estaba haciendo rehuso del codigo lo llama por medio del paquete y el metodo, le pasa los parametros, comprueba que no haya errores y queda.
func GetJWTAuthorizationInfo(s server.Server, w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	return token, err
}
