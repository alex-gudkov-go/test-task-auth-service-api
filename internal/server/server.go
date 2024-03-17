package server

import (
	"log"
	"net/http"
	authService "test-task-auth-service-api/internal/services/auth_service"
	"test-task-auth-service-api/internal/store"

	"github.com/gorilla/mux"
)

type Server struct {
	address string
	store   store.Store
}

func New(address string, store store.Store) *Server {
	return &Server{
		address: address,
		store:   store,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()

	authService := authService.New(s.store)
	authService.RegisterHandlers(router)

	log.Println("Server listen at:", s.address)
	log.Fatalln(http.ListenAndServe(s.address, router))
}
