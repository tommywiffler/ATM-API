package handlers

import (
	"github.com/gorilla/mux"
)

type Server struct {
	Router         *mux.Router
	NetworkAddress string
}

func (s *Server) Init() {
	s.NetworkAddress = ":8080"
	s.Router = mux.NewRouter()
}
