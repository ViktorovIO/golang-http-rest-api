package server

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"http-rest-api/internal/app/store"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Starting API server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/ping", s.handlePing())
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) handlePing() http.HandlerFunc {
	type request struct {
		name string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "Pong")
		if err != nil {
			s.logger.Error("Something is wrong")
		}
	}
}
