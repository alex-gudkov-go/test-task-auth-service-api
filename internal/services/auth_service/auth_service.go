package auth

import (
	"net/http"
	"test-task-auth-service-api/internal/config"
	"test-task-auth-service-api/internal/models"
	"test-task-auth-service-api/internal/store"
	"test-task-auth-service-api/pkg/hash"
	"test-task-auth-service-api/pkg/response"
	"test-task-auth-service-api/pkg/tokens"
	"test-task-auth-service-api/pkg/validate"

	"github.com/gorilla/mux"
)

type AuthService struct {
	store store.Store
}

func New(store store.Store) *AuthService {
	return &AuthService{store}
}

func (as *AuthService) RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/auth/{user_id}/tokens/sign", as.HandleSignTokens).Methods(http.MethodPost)
}

func (as *AuthService) HandleSignTokens(rw http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["user_id"]
	if err := validate.ValidateGuid(userId); err != nil {
		response.Write(rw, response.Error{Message: "param \"user_id\" is not valid GUID"}, http.StatusBadRequest)
		return
	}

	// generate refresh token
	refreshTokenString, err := tokens.GenerateRefreshTokenString()
	if err != nil {
		response.Write(rw, response.Error{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	// hash refresh token before save
	refreshTokenStringHash, err := hash.HashRefreshTokenString(refreshTokenString)
	if err != nil {
		response.Write(rw, response.Error{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	// save refresh token
	refreshToken := &models.RefreshToken{
		UserId: userId,
		Value:  refreshTokenStringHash,
	}
	if err := as.store.SaveRefreshToken(req.Context(), refreshToken); err != nil {
		response.Write(rw, response.Error{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	// generate access token (bind to refresh token ID)
	accessTokenString, err := tokens.GenerateAccessTokenString(userId, refreshToken.Id, config.Envs.AccessTokenLifetimeInMinutes, config.Envs.AccessTokenSecret)
	if err != nil {
		response.Write(rw, response.Error{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	responseData := &response.HandleSignTokensData{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}
	response.Write(rw, responseData, http.StatusOK)
}
