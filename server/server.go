package server

import (
	"log"
	"net/http"
	authService "test-task-auth-service-api/service/auth_service"
	"test-task-auth-service-api/store"

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

func (s *Server) Start() {
	r := mux.NewRouter()

	// setup services
	authService := authService.New(s.store)
	authService.RegisterRoutes(r)

	// start server
	log.Println("Start server at:", s.address)
	log.Fatalln(http.ListenAndServe(s.address, r))
}
