package hanlders

import (
	"encoding/json"
	"net/http"

	"pruebago.com/go/rest-ws/server"
)

type HomeResponse struct {
	Message string `json:"message"` //Se serializa para que al ser enviado a json lo lea de otra forma
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Se completo la redaccion de la peticion",
			Status:  true,
		})
	}
}
