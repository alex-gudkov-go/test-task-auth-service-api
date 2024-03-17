package response

import (
	"encoding/json"
	"net/http"
)

func Write(rw http.ResponseWriter, data any, status int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(data)
}

type Error struct {
	Message string `json:"message"`
}

type HandleSignTokensData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
