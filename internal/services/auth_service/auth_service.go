package auth

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"test-task-auth-service-api/internal/config"
	"test-task-auth-service-api/internal/models"
	"test-task-auth-service-api/internal/store"
	"test-task-auth-service-api/pkg/hash"
	"test-task-auth-service-api/pkg/http/request"
	"test-task-auth-service-api/pkg/http/response"
	"test-task-auth-service-api/pkg/tokens"
	"test-task-auth-service-api/pkg/validate"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type AuthService struct {
	store store.Store
}

type HandleRefreshTokensRequestBody struct {
	RefreshToken string `json:"refreshToken"`
}

type TokensPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func New(store store.Store) *AuthService {
	return &AuthService{store}
}

func (as *AuthService) RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/auth/{user_id}/tokens/sign", as.HandleSignTokens).Methods(http.MethodPost)
	router.HandleFunc("/auth/tokens/refresh", as.HandleRefreshTokens).Methods(http.MethodPost)
}

func (as *AuthService) HandleSignTokens(rw http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["user_id"]
	if err := validate.ValidateGuid(userId); err != nil {
		response.WriteBadRequest(rw, "param \"user_id\" is not valid GUID")
		return
	}

	tokensPair, err := as.signTokens(req.Context(), userId)
	if err != nil {
		response.WriteInternalServerError(rw, err.Error())
		return
	}

	response.WriteOK(rw, tokensPair)
}

func (as *AuthService) HandleRefreshTokens(rw http.ResponseWriter, req *http.Request) {
	accessTokenString := request.ParseBearerToken(req)
	if accessTokenString == "" {
		response.WriteUnauthorized(rw, "no Bearer token provided")
		return
	}

	// get refresh token from request body
	b, err := io.ReadAll(req.Body)
	if err != nil {
		response.WriteInternalServerError(rw, err.Error())
		return
	}
	defer req.Body.Close()

	requestBody := &HandleRefreshTokensRequestBody{}
	if err := json.Unmarshal(b, requestBody); err != nil {
		response.WriteBadRequest(rw, err.Error())
		return
	}
	if requestBody.RefreshToken == "" {
		response.WriteBadRequest(rw, "no \"refreshToken\" in request body")
		return
	}
	refreshTokenString := requestBody.RefreshToken

	// validate tokens pair
	userId, refreshTokenId, err := as.validateTokensPair(req.Context(), accessTokenString, refreshTokenString)
	if err != nil {
		response.WriteUnauthorized(rw, err.Error())
		return
	}

	// delete old refresh token
	if err := as.store.DeleteRefreshTokenById(req.Context(), refreshTokenId); err != nil {
		response.WriteInternalServerError(rw, "failed to delete old refresh token")
		return
	}

	// sign new tokens
	tokensPair, err := as.signTokens(req.Context(), userId)
	if err != nil {
		response.WriteInternalServerError(rw, "failed to sign new tokens")
		return
	}

	response.WriteOK(rw, tokensPair)
}

func (as *AuthService) signTokens(ctx context.Context, userId string) (*TokensPair, error) {
	// generate refresh token
	refreshTokenString, err := tokens.GenerateRefreshTokenString()
	if err != nil {
		return nil, err
	}

	// hash refresh token before save
	refreshTokenHash, err := hash.HashString(refreshTokenString)
	if err != nil {
		return nil, err
	}

	// save refresh token
	refreshToken := &models.RefreshToken{
		UserId: userId,
		Value:  refreshTokenHash,
	}
	if err := as.store.SaveRefreshToken(ctx, refreshToken); err != nil {
		return nil, err
	}

	// generate access token (bind to refresh token ID)
	accessTokenString, err := tokens.GenerateAccessTokenString(userId, refreshToken.Id, config.Envs.AccessTokenLifetimeInMinutes, config.Envs.AccessTokenSecret)
	if err != nil {
		return nil, err
	}

	return &TokensPair{accessTokenString, refreshTokenString}, nil
}

func (as *AuthService) validateTokensPair(ctx context.Context, accessTokenString string, refreshTokenString string) (string, string, error) {
	// decode access token
	accessToken, err := tokens.DecodeJwtToken(accessTokenString, config.Envs.AccessTokenSecret)
	if err != nil {
		return "", "", errors.New("access token is not valid")
	}

	// get access token claims
	accessTokenClaims := accessToken.Claims.(jwt.MapClaims)
	userId, ok := accessTokenClaims["userId"].(string)
	if !ok {
		return "", "", errors.New("no \"userId\" claim in access token")
	}
	refreshTokenId, ok := accessTokenClaims["refreshTokenId"].(string)
	if !ok {
		return "", "", errors.New("no \"refreshTokenId\" claim in access token")
	}

	// find refresh token (access token stores refresh token ID)
	refreshToken, err := as.store.FindRefreshTokenById(ctx, refreshTokenId)
	if err != nil {
		return "", "", errors.New("no binded refresh token found")
	}
	refreshTokenHash := refreshToken.Value

	// compare refresh token from access token with the one from the database
	if err := hash.CompareHashAndString(refreshTokenHash, refreshTokenString); err != nil {
		return "", "", errors.New("refresh token is not valid")
	}

	return userId, refreshTokenId, nil
}
