package auth_service

import (
	"net/http"
	"test-task-auth-service-api/store"

	"github.com/gorilla/mux"
)

type AuthService struct {
	store store.Store
}

func New(store store.Store) *AuthService {
	return &AuthService{store}
}

func (as *AuthService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/auth/tokens", as.handleSignTokens).Methods(http.MethodPost)
	router.HandleFunc("/auth/tokens/refresh", as.handleRefreshTokens).Methods(http.MethodPost)
}

func (as *AuthService) handleSignTokens(rw http.ResponseWriter, req *http.Request) {
	// TODO: implement
}

func (as *AuthService) handleRefreshTokens(rw http.ResponseWriter, req *http.Request) {
	// TODO: implement
}
