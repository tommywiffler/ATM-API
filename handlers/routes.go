package handlers

import (
	"atm-api/middlewares"
)

func (s *Server) InitializeRoutes() {
	s.Router.HandleFunc("/api/atm/user", middlewares.MiddlewareJSON(UserLogin)).Methods("POST")
	s.Router.HandleFunc("/api/atm/balance/{id}", middlewares.MiddlewareJSON(GetBalance)).Methods("GET")
	s.Router.HandleFunc("/api/atm/deposit", middlewares.MiddlewareJSON(Deposit)).Methods("PATCH")
	s.Router.HandleFunc("/api/atm/withdraw", middlewares.MiddlewareJSON(Withdraw)).Methods("PATCH")
}
