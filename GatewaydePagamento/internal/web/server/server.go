package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vivianezzt/go-gateway.git/internal/service"
	"github.com/vivianezzt/go-gateway.git/internal/web/handlers"
)

type Server struct {
	// roteador
	router *chi.Mux
	server *http.Server
	accountService *service.AccountService
	port string
}

func NewServer(port string, accountService *service.AccountService) *Server {
	return &Server{
		router: chi.NewRouter(),
		port: port,
		accountService: accountService,
	}
}
func (s *Server) ConfigureRoutes(){
	accoutHandler := handlers.NewAccountHandler(s.accountService)
	s.router.Post("/accounts", accoutHandler.Create)
	s.router.Get("/accounts", accoutHandler.Get)
}
func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}