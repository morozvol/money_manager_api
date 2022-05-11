package server

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/morozvol/money_manager_api/internal/config"
	"github.com/morozvol/money_manager_api/internal/routes"
	"github.com/morozvol/money_manager_api/pkg/store"
	"github.com/sirupsen/logrus"
	"net/http"
)

//Server ...
type Server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
	config *config.Config
}

// Start Server
func (server *Server) Start() error {
	return http.ListenAndServe(server.config.BindAddr, server)
}

// New Server
func New(store store.Store, config *config.Config) *Server {
	s := &Server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
		config: config,
	}

	s.configureRouter()

	return s
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.router.ServeHTTP(w, r)
}

func (server *Server) configureRouter() {
	server.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	routes.OperationRoute(server.router, server.store)
	routes.CategoryRoute(server.router, server.store)
	routes.AccountRoute(server.router, server.store)
}

func (server *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	server.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (server *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (server *Server) handleLogin() func(http.ResponseWriter, *http.Request) {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			server.error(w, r, http.StatusBadRequest, err)
			return
		}

	}
}
