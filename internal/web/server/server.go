package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vivianezzt/go-gateway/internal/service"
	"github.com/vivianezzt/go-gateway/internal/web/handlers"
	"github.com/vivianezzt/go-gateway/internal/web/middleware"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	invoiceService *service.InvoiceService
	port           string
}

func NewServer(accountService *service.AccountService, invoiceService *service.InvoiceService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		invoiceService: invoiceService,
		port:           port,
	}
}

func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)
	invoiceHandler := handlers.NewInvoiceHandler(s.invoiceService)
	authMiddleware := middleware.NewAuthMiddleware(s.accountService)

	// 🔓 Rotas públicas (sem autenticação)
	s.router.Post("/accounts", accountHandler.Create) // Criação de conta
	s.router.Get("/accounts", accountHandler.Get)     // Consulta de conta

	// 🔐 Rotas protegidas por middleware (requer X-API-KEY)
	s.router.Group(func(r chi.Router) {
		r.Use(authMiddleware.Authenticate)

		// 🧾 Faturas
		r.Post("/invoice", invoiceHandler.Create)         // Cria nova fatura
		r.Get("/invoice/{id}", invoiceHandler.GetByID)    // Busca fatura por ID
		r.Get("/invoice", invoiceHandler.ListByAccount)   // Lista faturas da conta
	})
}


func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}
