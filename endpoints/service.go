package endpoints

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Service struct {
	Cfg    Config
	logger *zap.Logger
}

func NewService(cfg Config, logger *zap.Logger) *Service {

	return &Service{
		Cfg:    cfg,
		logger: logger.Named("kafka_service"),
	}
}

func (s *Service) Propagate() {
	router := mux.NewRouter()

	//Pre-flight Endpoints
	router.PathPrefix("/").HandlerFunc(ok()).Methods("OPTIONS")
	router.PathPrefix("/").HandlerFunc(ok()).Methods("GET")

	// Unprotected Endpoints
	router.HandleFunc("/metrics", ok()).Methods("GET")
}
