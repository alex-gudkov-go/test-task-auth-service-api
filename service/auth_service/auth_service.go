package auth_service

import (
	"net/http"
	"test-task-auth-service-api/models"
	"test-task-auth-service-api/response"
	"test-task-auth-service-api/store"
	"test-task-auth-service-api/utils"

	"github.com/gorilla/mux"
)

type AuthService struct {
	store store.Store
}

func New(store store.Store) *AuthService {
	return &AuthService{store}
}

func (as *AuthService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/auth/{user_id}/tokens/sign", as.handleSignTokens).Methods(http.MethodPost)
}

func (as *AuthService) handleSignTokens(rw http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["user_id"]
	if err := utils.ValidateGuid(userId); err != nil {
		response.Write(rw, response.Error{Message: "param \"user_id\" is not valid GUID"}, http.StatusBadRequest)
		return
	}

	// generate refresh token
	refreshTokenString, err := utils.GenerateRefreshTokenString()
	if err != nil {
		response.Write(rw, response.Error{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	// hash refresh token before save
	refreshTokenStringHash, err := utils.HashRefreshTokenString(refreshTokenString)
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
	accessTokenString, err := utils.GenerateAccessTokenString(userId, refreshToken.Id)
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
