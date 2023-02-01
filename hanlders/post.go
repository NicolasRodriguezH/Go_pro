package hanlders

import (
	"encoding/json"
	"net/http"

	"github.com/segmentio/ksuid"
	"pruebago.com/go/rest-ws/helpers"
	"pruebago.com/go/rest-ws/models"
	"pruebago.com/go/rest-ws/repository"
	"pruebago.com/go/rest-ws/server"
)

type InsertPostRequest struct {
	Content string `json:"content"`
}

type PostResponse struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

// Se usa esquema/modelo de user.go (MeHandler) = Forma en la que traemos un token y lo descomponemos
func InsertPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := helpers.GetJWTAuthorizationInfo(s, w, r)

		if err != nil {
			return
		}

		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			var postRequest = InsertPostRequest{}
			if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			post := models.Post{
				Id:      id.String(),
				Content: postRequest.Content,
				UserId:  claims.UserId,
			}

			err = repository.InsertPost(r.Context(), &post)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(PostResponse{
				Id:      post.Id,
				Content: post.Content,
			})
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
